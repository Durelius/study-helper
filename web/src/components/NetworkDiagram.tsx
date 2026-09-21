import type { Network } from '../lib/api'

const BOX_W = 96
const BOX_H = 54
const COL_GAP = 62
const ROW_GAP = 26

// NetworkDiagram draws an activity-on-node network in the notation the slides use:
// a three-column box with ES | Duration | EF across the top, the activity name in a
// band through the middle, and LS | Float | LF along the bottom.
//
// While the question is live the box is blank except for the name and the duration —
// exactly like the deck's worked exercise, whose boxes students fill in themselves.
// Printing the schedule figures before the answer is in would give the whole thing away.
export default function NetworkDiagram({ net, reveal }: { net: Network; reveal: boolean }) {
  const layers = new Map<number, typeof net.activities>()
  for (const a of net.activities) {
    const bucket = layers.get(a.layer) ?? []
    bucket.push(a)
    layers.set(a.layer, bucket)
  }
  const columns = [...layers.keys()].sort((a, b) => a - b)
  const tallest = Math.max(...columns.map((c) => layers.get(c)!.length))

  const width = columns.length * (BOX_W + COL_GAP) - COL_GAP
  const height = tallest * (BOX_H + ROW_GAP) - ROW_GAP

  const place = new Map<string, { x: number; y: number }>()
  columns.forEach((col, ci) => {
    const bucket = layers.get(col)!
    const blockHeight = bucket.length * (BOX_H + ROW_GAP) - ROW_GAP
    const top = (height - blockHeight) / 2
    bucket.forEach((a, ri) => {
      place.set(a.name, { x: ci * (BOX_W + COL_GAP), y: top + ri * (BOX_H + ROW_GAP) })
    })
  })

  return (
    <div className="overflow-x-auto py-2">
      <svg
        viewBox={`-8 -8 ${width + 16} ${height + 16}`}
        width={width + 16}
        height={height + 16}
        role="img"
        aria-label={`Activity network with ${net.activities.length} activities`}
        style={{ maxWidth: '100%', height: 'auto' }}
      >
        <defs>
          <marker id="arrow" viewBox="0 0 8 8" refX="7" refY="4" markerWidth="7" markerHeight="7" orient="auto">
            <path d="M0 0 L8 4 L0 8 z" fill="var(--line-strong)" />
          </marker>
          <marker id="arrow-critical" viewBox="0 0 8 8" refX="7" refY="4" markerWidth="7" markerHeight="7" orient="auto">
            <path d="M0 0 L8 4 L0 8 z" fill="var(--critical)" />
          </marker>
        </defs>

        {net.activities.flatMap((a) =>
          (a.predecessors ?? []).map((p) => {
            const from = place.get(p)!
            const to = place.get(a.name)!
            const x1 = from.x + BOX_W
            const y1 = from.y + BOX_H / 2
            const x2 = to.x
            const y2 = to.y + BOX_H / 2
            const mid = (x1 + x2) / 2
            const critical = net.criticalPath ?? []
            const onPath = reveal && critical.includes(p) && critical.includes(a.name)
            return (
              <path
                key={`${p}-${a.name}`}
                d={`M${x1} ${y1} C${mid} ${y1} ${mid} ${y2} ${x2 - 3} ${y2}`}
                fill="none"
                stroke={onPath ? 'var(--critical)' : 'var(--line-strong)'}
                strokeWidth={onPath ? 2 : 1.25}
                markerEnd={onPath ? 'url(#arrow-critical)' : 'url(#arrow)'}
              />
            )
          }),
        )}

        {net.activities.map((a) => {
          const { x, y } = place.get(a.name)!
          const critical = reveal && a.critical
          const third = BOX_W / 3
          const band = BOX_H / 3
          return (
            <g key={a.name}>
              <rect
                x={x}
                y={y}
                width={BOX_W}
                height={BOX_H}
                rx="2"
                fill={critical ? 'var(--critical-soft)' : 'var(--surface)'}
                stroke={critical ? 'var(--critical)' : 'var(--line-strong)'}
                strokeWidth={critical ? 2 : 1}
              />
              <line x1={x} y1={y + band} x2={x + BOX_W} y2={y + band} stroke="var(--line)" strokeWidth="1" />
              <line x1={x} y1={y + band * 2} x2={x + BOX_W} y2={y + band * 2} stroke="var(--line)" strokeWidth="1" />
              <line x1={x + third} y1={y} x2={x + third} y2={y + band} stroke="var(--line)" strokeWidth="1" />
              <line x1={x + third * 2} y1={y} x2={x + third * 2} y2={y + band} stroke="var(--line)" strokeWidth="1" />
              <line x1={x + third} y1={y + band * 2} x2={x + third} y2={y + BOX_H} stroke="var(--line)" strokeWidth="1" />
              <line x1={x + third * 2} y1={y + band * 2} x2={x + third * 2} y2={y + BOX_H} stroke="var(--line)" strokeWidth="1" />

              <text
                x={x + BOX_W / 2}
                y={y + band + band / 2 + 4}
                textAnchor="middle"
                fontSize="14"
                fontWeight="600"
                fill={critical ? 'var(--critical-ink)' : 'var(--ink)'}
                fontFamily="Archivo, sans-serif"
              >
                {a.name}
              </text>

              {/* The duration always shows: it is given in the question, not derived. */}
              <text
                x={x + BOX_W / 2}
                y={y + band - 5}
                textAnchor="middle"
                fontSize="11"
                fontWeight="600"
                fill="var(--ink)"
                fontFamily="Archivo, sans-serif"
              >
                {reveal ? a.duration : a.duration}
              </text>

              {reveal && (
                <>
                  <text x={x + third / 2} y={y + band - 5} textAnchor="middle" fontSize="10" fill="var(--ink-soft)" fontFamily="Archivo, sans-serif">
                    {a.es}
                  </text>
                  <text x={x + third * 2.5} y={y + band - 5} textAnchor="middle" fontSize="10" fill="var(--ink-soft)" fontFamily="Archivo, sans-serif">
                    {a.ef}
                  </text>
                  <text x={x + third / 2} y={y + BOX_H - 5} textAnchor="middle" fontSize="10" fill="var(--ink-soft)" fontFamily="Archivo, sans-serif">
                    {a.ls}
                  </text>
                  <text
                    x={x + BOX_W / 2}
                    y={y + BOX_H - 5}
                    textAnchor="middle"
                    fontSize="10"
                    fontWeight="600"
                    fill={a.float === 0 ? 'var(--critical)' : 'var(--ink-soft)'}
                    fontFamily="Archivo, sans-serif"
                  >
                    {a.float}
                  </text>
                  <text x={x + third * 2.5} y={y + BOX_H - 5} textAnchor="middle" fontSize="10" fill="var(--ink-soft)" fontFamily="Archivo, sans-serif">
                    {a.lf}
                  </text>
                </>
              )}
            </g>
          )
        })}
      </svg>
      <p className="mt-1 text-[0.72rem] text-ink-faint">
        {reveal
          ? 'ES · duration · EF across the top, LS · float · LF underneath. Float 0 marks the critical path.'
          : 'The number above each activity is its duration in days.'}
      </p>
    </div>
  )
}
