import type { WbsDiagram as Spec } from '../lib/diagram'

const BOX_W = 118
const BOX_H = 46
const GAP_X = 12
const GAP_Y = 34

type Box = { x: number; y: number; w: number; h: number }

// Lays the tree out bottom-up: a parent sits centred over its children, which is how
// the deck draws a WBS and what makes a missing or misplaced branch visible.
export function wbsLayout(spec: Spec) {
  const children = new Map<string, string[]>()
  const roots: string[] = []
  for (const n of spec.nodes) {
    if (n.parent) children.set(n.parent, [...(children.get(n.parent) ?? []), n.id])
    else roots.push(n.id)
  }
  const byId = new Map(spec.nodes.map((n) => [n.id, n]))
  const boxes = new Map<string, Box>()
  let cursor = 0

  function place(id: string): number {
    const node = byId.get(id)!
    const kids = children.get(id) ?? []
    const y = node.level * (BOX_H + GAP_Y)
    if (kids.length === 0) {
      const x = cursor
      cursor += BOX_W + GAP_X
      boxes.set(id, { x, y, w: BOX_W, h: BOX_H })
      return x + BOX_W / 2
    }
    const centres = kids.map(place)
    const mid = (centres[0] + centres[centres.length - 1]) / 2
    boxes.set(id, { x: mid - BOX_W / 2, y, w: BOX_W, h: BOX_H })
    return mid
  }
  roots.forEach(place)

  const maxY = Math.max(...[...boxes.values()].map((b) => b.y + b.h))
  return { boxes, children, byId, width: Math.max(cursor - GAP_X, 0), height: maxY }
}

export default function WbsDiagram({ spec, layout }: { spec: Spec; layout: ReturnType<typeof wbsLayout> }) {
  const { boxes, children } = layout
  const shade = ['var(--ink)', 'var(--focus)', 'var(--surface)']

  return (
    <g>
      {[...children.entries()].flatMap(([parentId, kids]) => {
        const p = boxes.get(parentId)
        if (!p) return []
        const stem = p.y + p.h + GAP_Y / 2
        return [
          <line
            key={`${parentId}-stem`}
            x1={p.x + p.w / 2}
            y1={p.y + p.h}
            x2={p.x + p.w / 2}
            y2={stem}
            stroke="var(--line-strong)"
            strokeWidth={1.2}
          />,
          ...kids.map((kid) => {
            const c = boxes.get(kid)!
            return (
              <g key={`${parentId}-${kid}`}>
                <line x1={p.x + p.w / 2} y1={stem} x2={c.x + c.w / 2} y2={stem} stroke="var(--line-strong)" strokeWidth={1.2} />
                <line x1={c.x + c.w / 2} y1={stem} x2={c.x + c.w / 2} y2={c.y} stroke="var(--line-strong)" strokeWidth={1.2} />
              </g>
            )
          }),
        ]
      })}

      {spec.nodes.map((n) => {
        const b = boxes.get(n.id)
        if (!b) return null
        const level = Math.min(n.level, 2)
        const filled = level < 2
        return (
          <g key={n.id}>
            <rect
              x={b.x}
              y={b.y}
              width={b.w}
              height={b.h}
              rx={3}
              fill={filled ? shade[level] : 'var(--surface)'}
              stroke={filled ? shade[level] : 'var(--line-strong)'}
              strokeWidth={1.3}
            />
            <Label
              code={n.code}
              text={n.label}
              x={b.x + b.w / 2}
              y={b.y + b.h / 2}
              width={b.w - 10}
              colour={filled ? 'var(--paper)' : 'var(--ink)'}
            />
          </g>
        )
      })}
      {/* The deck labels its levels, and the level is often what the question turns on. */}
      {[0, 1, 2].map((lvl) => {
        const any = spec.nodes.find((n) => n.level === lvl)
        if (!any) return null
        const b = boxes.get(any.id)!
        return (
          <text key={lvl} x={-8} y={b.y + b.h / 2} textAnchor="end" dominantBaseline="middle" fontSize="9" fill="var(--ink-faint)">
            {['Level 1', 'Level 2', 'Level 3'][lvl]}
          </text>
        )
      })}
    </g>
  )
}

function Label({
  code,
  text,
  x,
  y,
  width,
  colour,
}: {
  code?: string
  text: string
  x: number
  y: number
  width: number
  colour: string
}) {
  const perLine = Math.max(10, Math.floor(width / 5))
  const words = text.split(' ')
  const lines: string[] = []
  let line = ''
  for (const w of words) {
    if ((line + ' ' + w).trim().length > perLine && line) {
      lines.push(line)
      line = w
    } else line = (line + ' ' + w).trim()
  }
  if (line) lines.push(line)
  const all = code ? [code, ...lines] : lines
  const start = y - ((all.length - 1) * 10.5) / 2
  return (
    <text textAnchor="middle" fontSize="9.5" fill={colour} fontFamily="Archivo, sans-serif">
      {all.map((l, i) => (
        <tspan key={i} x={x} y={start + i * 10.5} fontWeight={code && i === 0 ? 700 : 400}>
          {l}
        </tspan>
      ))}
    </text>
  )
}
