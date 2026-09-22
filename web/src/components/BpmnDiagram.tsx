import {
  EVENT_R,
  GATE,
  LABEL_W,
  LANE_H,
  TASK_H,
  TASK_W,
  type BpmnDiagram as Spec,
  type BpmnNode,
} from '../lib/diagram'

type Box = { x: number; y: number; w: number; h: number }

// Where each element sits, so the flows, the labels and the clickable overlays all
// agree without recomputing the layout three times.
export function bpmnLayout(spec: Spec) {
  const laneOrder: { id: string; label: string; pool: string; top: number; h: number; blackBox: boolean }[] = []
  let y = 0
  const poolBands: { label: string; top: number; height: number; blackBox: boolean }[] = []
  for (const pool of spec.pools) {
    const top = y
    const lanes = pool.lanes.length > 0 ? pool.lanes : [{ id: `${pool.label}-empty`, label: '' }]
    for (const lane of lanes) {
      const h = lane.h ?? (pool.blackBox ? LANE_H * 0.5 : LANE_H)
      laneOrder.push({ id: lane.id, label: lane.label, pool: pool.label, top: y, h, blackBox: !!pool.blackBox })
      y += h
    }
    poolBands.push({ label: pool.label, top, height: y - top, blackBox: !!pool.blackBox })
  }

  const contentX = LABEL_W * 2
  const boxes = new Map<string, Box>()
  for (const n of spec.nodes) {
    const lane = laneOrder.find((l) => l.id === n.lane)
    if (!lane) continue
    const cy = lane.top + lane.h / 2 + (n.dy ?? 0) + (n.straddle ? lane.h / 2 : 0)
    const cx = contentX + n.x
    if (n.type === 'task' || n.type === 'subprocess') {
      boxes.set(n.id, { x: cx - TASK_W / 2, y: cy - TASK_H / 2, w: TASK_W, h: TASK_H })
    } else if (n.type === 'xor' || n.type === 'and' || n.type === 'or') {
      boxes.set(n.id, { x: cx - GATE, y: cy - GATE, w: GATE * 2, h: GATE * 2 })
    } else if (n.type === 'anchor') {
      boxes.set(n.id, { x: cx - 3, y: cy - 3, w: 6, h: 6 })
    } else {
      boxes.set(n.id, { x: cx - EVENT_R, y: cy - EVENT_R, w: EVENT_R * 2, h: EVENT_R * 2 })
    }
  }
  // The authored width is a minimum. The real width comes from where the elements
  // actually ended up, so a node added at the right-hand end cannot silently fall off
  // the edge of the diagram.
  let right = 0
  for (const n of spec.nodes) {
    const b = boxes.get(n.id)
    if (!b) continue
    // An event's caption is centred underneath it and is wider than the circle.
    const captionOverhang = n.type.startsWith('start') || n.type.startsWith('end') || n.type === 'intermediate' ? 46 : 0
    right = Math.max(right, b.x + b.w + captionOverhang)
  }
  const width = Math.max(spec.width, right + 16)

  return { laneOrder, poolBands, boxes, height: y, contentX, width }
}

const centre = (b: Box) => ({ x: b.x + b.w / 2, y: b.y + b.h / 2 })

// Flows leave the right edge and arrive at the left edge where that reads cleanly, and
// drop vertically when they change lane — which is how the deck draws them.
function routeFlow(a: Box, b: Box): string {
  const ca = centre(a)
  const cb = centre(b)
  const sameRow = Math.abs(ca.y - cb.y) < 6

  if (sameRow) {
    const x1 = cb.x > ca.x ? a.x + a.w : a.x
    const x2 = cb.x > ca.x ? b.x : b.x + b.w
    return `M${x1} ${ca.y} L${x2} ${cb.y}`
  }
  // Going forwards and changing row: out the right edge, along to a midpoint, across
  // to the target's row, then into its left edge. Entering horizontally is how the
  // deck draws a split into parallel branches, and it keeps arrowheads off the
  // corners of the boxes.
  if (b.x > a.x + a.w) {
    const x1 = a.x + a.w
    const x2 = b.x
    const midX = x1 + (x2 - x1) / 2
    return `M${x1} ${ca.y} L${midX} ${ca.y} L${midX} ${cb.y} L${x2} ${cb.y}`
  }
  // Doubling back: down out of the bottom, back along, and into the target's edge.
  const y1 = ca.y > cb.y ? a.y : a.y + a.h
  const midY = ca.y > cb.y ? a.y - 22 : a.y + a.h + 22
  const x2 = cb.x
  const y2 = cb.y > ca.y ? b.y : b.y + b.h
  return `M${ca.x} ${y1} L${ca.x} ${midY} L${x2} ${midY} L${x2} ${y2}`
}

export default function BpmnDiagram({
  spec,
  layout,
}: {
  spec: Spec
  layout: ReturnType<typeof bpmnLayout>
}) {
  const { laneOrder, poolBands, boxes, height, contentX, width } = layout

  return (
    <g>
      {/* Pools and their lanes. */}
      {poolBands.map((p, i) => (
        <g key={i}>
          <rect x={0} y={p.top} width={width} height={p.height} fill="none" stroke="var(--line-strong)" strokeWidth={1.4} />
          <line x1={LABEL_W} y1={p.top} x2={LABEL_W} y2={p.top + p.height} stroke="var(--line-strong)" strokeWidth={1.4} />
          <text
            x={LABEL_W / 2}
            y={p.top + p.height / 2}
            transform={`rotate(-90 ${LABEL_W / 2} ${p.top + p.height / 2})`}
            textAnchor="middle"
            dominantBaseline="middle"
            fontSize="9.5"
            fontWeight="600"
            fill="var(--ink-soft)"
          >
            {p.label}
          </text>
        </g>
      ))}
      {laneOrder.map((l, i) => (
        <g key={i}>
          {i > 0 && (
            <line x1={LABEL_W} y1={l.top} x2={width} y2={l.top} stroke="var(--line)" strokeWidth={1} />
          )}
          {l.label && (
            <text
              x={LABEL_W + LABEL_W / 2}
              y={l.top + l.h / 2}
              transform={`rotate(-90 ${LABEL_W + LABEL_W / 2} ${l.top + l.h / 2})`}
              textAnchor="middle"
              dominantBaseline="middle"
              fontSize="9"
              fill="var(--ink-soft)"
            >
              {l.label}
            </text>
          )}
          {i > 0 && <line x1={LABEL_W} y1={l.top} x2={LABEL_W * 2} y2={l.top} stroke="var(--line)" strokeWidth={1} />}
        </g>
      ))}
      <line x1={contentX} y1={0} x2={contentX} y2={height} stroke="var(--line)" strokeWidth={1} />

      {/* Flows first, so nodes sit on top of the arrowheads. */}
      {spec.flows.map((f) => {
        const a = boxes.get(f.from)
        const b = boxes.get(f.to)
        if (!a || !b) return null
        const message = f.kind === 'message'
        const d = routeFlow(a, b)
        const mid = centre(b)
        return (
          <g key={`${f.from}>${f.to}`}>
            <path
              d={d}
              fill="none"
              stroke="var(--ink-soft)"
              strokeWidth={1.3}
              strokeDasharray={message ? '5 4' : undefined}
              markerEnd={message ? 'url(#bpmn-open)' : 'url(#bpmn-arrow)'}
            />
            {f.label && (
              <text
                x={(centre(a).x + mid.x) / 2}
                y={Math.min(centre(a).y, mid.y) - 5}
                textAnchor="middle"
                fontSize="9"
                fill="var(--ink-faint)"
              >
                {f.label}
              </text>
            )}
          </g>
        )
      })}

      {spec.nodes.map((n) => {
        const b = boxes.get(n.id)
        if (!b) return null
        return <Node key={n.id} node={n} box={b} />
      })}
    </g>
  )
}

function Node({ node, box }: { node: BpmnNode; box: Box }) {
  const c = centre(box)
  const label = node.label ?? ''

  // A black-box pool has no flow objects in it, so its anchors draw nothing; they
  // exist only to give a message flow somewhere on the boundary to land.
  if (node.type === 'anchor') return null

  if (node.type === 'task' || node.type === 'subprocess') {
    return (
      <g>
        <rect x={box.x} y={box.y} width={box.w} height={box.h} rx={7} fill="var(--surface)" stroke="var(--focus)" strokeWidth={1.4} />
        <Wrapped text={label} x={c.x} y={c.y} width={box.w - 12} />
        {node.type === 'subprocess' && (
          <g>
            <rect x={c.x - 6} y={box.y + box.h - 14} width={12} height={12} fill="none" stroke="var(--focus)" strokeWidth={1.1} />
            <line x1={c.x - 3} y1={box.y + box.h - 8} x2={c.x + 3} y2={box.y + box.h - 8} stroke="var(--focus)" strokeWidth={1.1} />
            <line x1={c.x} y1={box.y + box.h - 11} x2={c.x} y2={box.y + box.h - 5} stroke="var(--focus)" strokeWidth={1.1} />
          </g>
        )}
      </g>
    )
  }

  if (node.type === 'xor' || node.type === 'and' || node.type === 'or') {
    const mark =
      node.type === 'xor' ? (
        <g stroke="var(--ink)" strokeWidth={1.8}>
          <line x1={c.x - 6} y1={c.y - 6} x2={c.x + 6} y2={c.y + 6} />
          <line x1={c.x + 6} y1={c.y - 6} x2={c.x - 6} y2={c.y + 6} />
        </g>
      ) : node.type === 'and' ? (
        <g stroke="var(--ink)" strokeWidth={1.8}>
          <line x1={c.x - 7} y1={c.y} x2={c.x + 7} y2={c.y} />
          <line x1={c.x} y1={c.y - 7} x2={c.x} y2={c.y + 7} />
        </g>
      ) : (
        <circle cx={c.x} cy={c.y} r={6.5} fill="none" stroke="var(--ink)" strokeWidth={1.8} />
      )
    return (
      <g>
        <path
          d={`M${c.x} ${box.y} L${box.x + box.w} ${c.y} L${c.x} ${box.y + box.h} L${box.x} ${c.y} Z`}
          fill="var(--warn-soft)"
          stroke="var(--warn)"
          strokeWidth={1.4}
        />
        {mark}
        {label && (
          <text x={c.x} y={box.y - 6} textAnchor="middle" fontSize="9" fill="var(--ink-faint)">
            {label}
          </text>
        )}
      </g>
    )
  }

  // Events.
  const isEnd = node.type === 'end' || node.type === 'endMessage'
  const stroke = isEnd ? 'var(--critical)' : 'var(--slack)'
  const width = isEnd ? 3 : node.type === 'intermediate' ? 1.4 : 1.6
  return (
    <g>
      <circle cx={c.x} cy={c.y} r={EVENT_R} fill="var(--surface)" stroke={stroke} strokeWidth={width} />
      {node.type === 'intermediate' && (
        <circle cx={c.x} cy={c.y} r={EVENT_R - 3.5} fill="none" stroke={stroke} strokeWidth={1.2} />
      )}
      {(node.type === 'startMessage' || node.type === 'endMessage') && (
        <g stroke={stroke} strokeWidth={1.1} fill="none">
          <rect x={c.x - 7} y={c.y - 5} width={14} height={10} />
          <path d={`M${c.x - 7} ${c.y - 5} L${c.x} ${c.y + 1} L${c.x + 7} ${c.y - 5}`} />
        </g>
      )}
      {label && (
        <Wrapped text={label} x={c.x} y={c.y + EVENT_R + 11} width={96} faint anchorTop />
      )}
    </g>
  )
}

// Text that breaks onto a second line rather than spilling out of its box.
function Wrapped({
  text,
  x,
  y,
  width,
  faint,
  anchorTop,
}: {
  text: string
  x: number
  y: number
  width: number
  faint?: boolean
  anchorTop?: boolean
}) {
  const perLine = Math.max(8, Math.floor(width / 5.4))
  const words = text.split(' ')
  const lines: string[] = []
  let line = ''
  for (const w of words) {
    if ((line + ' ' + w).trim().length > perLine && line) {
      lines.push(line)
      line = w
    } else {
      line = (line + ' ' + w).trim()
    }
  }
  if (line) lines.push(line)
  const start = anchorTop ? y : y - ((lines.length - 1) * 11) / 2
  return (
    <text
      textAnchor="middle"
      fontSize="9.5"
      fill={faint ? 'var(--ink-soft)' : 'var(--ink)'}
      fontFamily="Archivo, sans-serif"
    >
      {lines.map((l, i) => (
        <tspan key={i} x={x} y={start + i * 11}>
          {l}
        </tspan>
      ))}
    </text>
  )
}
