import { useMemo } from 'react'
import BpmnDiagram, { bpmnLayout } from './BpmnDiagram'
import WbsDiagram, { wbsLayout } from './WbsDiagram'
import MatrixDiagram, { matrixLayout } from './MatrixDiagram'
import type { DiagramSpec } from '../lib/diagram'

// Draws a diagram and makes its hotspots clickable, so a question about a model is
// answered by pointing at the offending element — the app's version of circling it.
//
// Each hotspot maps to one of the question's choices by index, so the click and the
// list of options underneath are the same answer by two routes.
export default function Diagram({
  spec,
  selected,
  correct,
  revealed,
  onPick,
}: {
  spec: DiagramSpec
  /** Indices the reader has chosen. */
  selected: number[]
  /** The answer key, once it may be shown. */
  correct: number[]
  revealed: boolean
  onPick: (index: number) => void
}) {
  const laid = useMemo(() => {
    if (spec.kind === 'bpmn') return { kind: 'bpmn' as const, layout: bpmnLayout(spec) }
    if (spec.kind === 'wbs') return { kind: 'wbs' as const, layout: wbsLayout(spec) }
    return { kind: 'matrix' as const, layout: matrixLayout(spec) }
  }, [spec])

  const pad = 18
  const width = (laid.layout as { width: number }).width
  const height = laid.layout.height
  const left = spec.kind === 'wbs' ? -54 : spec.kind === 'bpmn' ? 0 : -2

  return (
    <div className="overflow-x-auto">
      <svg
        viewBox={`${left - pad / 2} ${-pad} ${width - left + pad * 1.5} ${height + pad * 2.2}`}
        width={width - left + pad * 1.5}
        height={height + pad * 2.2}
        style={{ maxWidth: '100%', height: 'auto' }}
        role="img"
        aria-label="Diagram. The answer options below describe each part you can pick."
      >
        <defs>
          <marker id="bpmn-arrow" viewBox="0 0 8 8" refX="7" refY="4" markerWidth="6.5" markerHeight="6.5" orient="auto">
            <path d="M0 0 L8 4 L0 8 z" fill="var(--ink-soft)" />
          </marker>
          <marker id="bpmn-open" viewBox="0 0 8 8" refX="7" refY="4" markerWidth="7" markerHeight="7" orient="auto">
            <path d="M0 0 L8 4 L0 8" fill="none" stroke="var(--ink-soft)" strokeWidth="1.2" />
          </marker>
        </defs>

        {laid.kind === 'bpmn' && <BpmnDiagram spec={spec as never} layout={laid.layout as never} />}
        {laid.kind === 'wbs' && <WbsDiagram spec={spec as never} layout={laid.layout as never} />}
        {laid.kind === 'matrix' && <MatrixDiagram spec={spec as never} layout={laid.layout as never} />}

        {spec.hotspots.map((h, i) => {
          const box = hotspotBox(h.target, laid)
          if (!box) return null
          // Matrix cells tile with no gap between them, so an outset highlight would
          // spill into the neighbours and the badge would straddle a row boundary.
          const pad = spec.kind === 'matrix' ? 0 : 6
          const badge =
            spec.kind === 'matrix'
              ? { x: box.x + 11, y: box.y + box.h / 2 }
              : { x: box.x - pad, y: box.y - pad }
          const isCorrect = revealed && correct.includes(i)
          const isWrongPick = revealed && selected.includes(i) && !correct.includes(i)
          const isPicked = !revealed && selected.includes(i)
          const stroke = isCorrect
            ? 'var(--slack)'
            : isWrongPick
              ? 'var(--critical)'
              : isPicked
                ? 'var(--ink)'
                : 'transparent'
          return (
            <g key={i} className="vc-hotspot" onClick={() => !revealed && onPick(i)}>
              <rect
                x={box.x - pad}
                y={box.y - pad}
                width={box.w + pad * 2}
                height={box.h + pad * 2}
                rx={4}
                fill="transparent"
                stroke={stroke}
                strokeWidth={2.2}
                style={{ cursor: revealed ? 'default' : 'pointer' }}
              />
              {/* A number in the corner ties the clickable region to its option below. */}
              <circle
                cx={badge.x}
                cy={badge.y}
                r={spec.kind === 'matrix' ? 7 : 8}
                fill={stroke === 'transparent' ? 'var(--sunken)' : stroke}
                stroke="var(--line-strong)"
                strokeWidth={0.8}
              />
              <text
                x={badge.x}
                y={badge.y}
                textAnchor="middle"
                dominantBaseline="central"
                fontSize="9"
                fontWeight="700"
                fill={stroke === 'transparent' ? 'var(--ink-soft)' : 'var(--paper)'}
                style={{ pointerEvents: 'none' }}
              >
                {i + 1}
              </text>
            </g>
          )
        })}
      </svg>
    </div>
  )
}

type Laid =
  | { kind: 'bpmn'; layout: ReturnType<typeof bpmnLayout> }
  | { kind: 'wbs'; layout: ReturnType<typeof wbsLayout> }
  | { kind: 'matrix'; layout: ReturnType<typeof matrixLayout> }

// A hotspot points at a node, a matrix cell, or a flow written "from>to".
function hotspotBox(target: string, laid: Laid) {
  if (target.includes('>') && laid.kind === 'bpmn') {
    const [from, to] = target.split('>')
    const a = laid.layout.boxes.get(from)
    const b = laid.layout.boxes.get(to)
    if (!a || !b) return null
    const x1 = Math.min(a.x + a.w, b.x)
    const x2 = Math.max(a.x + a.w, b.x)
    const y1 = Math.min(a.y + a.h / 2, b.y + b.h / 2)
    const y2 = Math.max(a.y + a.h / 2, b.y + b.h / 2)
    return { x: x1, y: y1 - 8, w: Math.max(26, x2 - x1), h: Math.max(16, y2 - y1) }
  }
  return laid.layout.boxes.get(target) ?? null
}
