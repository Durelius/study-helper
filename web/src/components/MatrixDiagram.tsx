import type { MatrixDiagram as Spec } from '../lib/diagram'

const ROW_H = 30
// Room for the hotspot badge, so a row number never sits on top of the row's name.
const LABEL_PAD = 26

type Box = { x: number; y: number; w: number; h: number }

export function matrixLayout(spec: Spec) {
  // The label column sizes itself to the longest row name: a fixed width either
  // wasted space or ran the names into the first column of data.
  const longest = Math.max(0, ...spec.rows.map((r) => r.label.length))
  const LABEL_W = Math.max(150, Math.min(330, longest * 5.4 + LABEL_PAD + 14))
  const colW = Math.max(62, Math.min(112, 660 / Math.max(1, spec.columns.length)))
  const boxes = new Map<string, Box>()
  spec.rows.forEach((row, r) => {
    row.cells.forEach((_, c) => {
      boxes.set(`${row.id}:${c}`, {
        x: LABEL_W + c * colW,
        y: ROW_H * (r + 1),
        w: colW,
        h: ROW_H,
      })
    })
    // A whole-row hotspot spans the row, so "click the row where nobody does the
    // work" highlights the row rather than just its label.
    boxes.set(row.id, { x: 0, y: ROW_H * (r + 1), w: LABEL_W + colW * spec.columns.length, h: ROW_H })
  })
  return {
    boxes,
    colW,
    labelW: LABEL_W,
    width: LABEL_W + colW * spec.columns.length,
    height: ROW_H * (spec.rows.length + 1),
  }
}

// A responsibility matrix, drawn the way the deck draws it: work packages down the
// side, roles across the top, one letter per cell.
export default function MatrixDiagram({ spec, layout }: { spec: Spec; layout: ReturnType<typeof matrixLayout> }) {
  const { colW, labelW: LABEL_W, width, height } = layout
  // A cell that can be clicked carries a number badge on its left, so its value is
  // nudged clear of it rather than being half hidden behind it.
  const spotted = new Set(spec.hotspots.map((h) => h.target))
  const colour = (v: string) =>
    ({ R: 'var(--focus)', A: 'var(--critical)', C: 'var(--warn)', I: 'var(--ink-faint)' })[v.trim().toUpperCase()] ??
    'var(--ink)'

  return (
    <g>
      <rect x={0} y={0} width={width} height={height} fill="var(--surface)" stroke="var(--line-strong)" strokeWidth={1.2} />
      <rect x={0} y={0} width={width} height={ROW_H} fill="var(--sunken)" />
      {spec.columns.map((c, i) => (
        <text
          key={i}
          x={LABEL_W + i * colW + colW / 2}
          y={ROW_H / 2}
          textAnchor="middle"
          dominantBaseline="middle"
          fontSize="9"
          fontWeight="600"
          fill="var(--ink-soft)"
        >
          {c}
        </text>
      ))}
      {spec.rows.map((row, r) => (
        <g key={row.id}>
          <text x={LABEL_PAD} y={ROW_H * (r + 1) + ROW_H / 2} dominantBaseline="middle" fontSize="9.5" fill="var(--ink)">
            {row.label}
          </text>
          {row.cells.map((v, c) => (
            <text
              key={c}
              x={LABEL_W + c * colW + colW / 2 + (spotted.has(`${row.id}:${c}`) ? 9 : 0)}
              y={ROW_H * (r + 1) + ROW_H / 2}
              textAnchor="middle"
              dominantBaseline="middle"
              fontSize="11"
              fontWeight="700"
              fill={colour(v)}
            >
              {v}
            </text>
          ))}
        </g>
      ))}
      {spec.rows.map((_, r) => (
        <line key={r} x1={0} y1={ROW_H * (r + 1)} x2={width} y2={ROW_H * (r + 1)} stroke="var(--line)" strokeWidth={1} />
      ))}
      {spec.columns.map((_, i) => (
        <line key={i} x1={LABEL_W + i * colW} y1={0} x2={LABEL_W + i * colW} y2={height} stroke="var(--line)" strokeWidth={1} />
      ))}
    </g>
  )
}
