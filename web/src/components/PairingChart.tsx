// The AIS food-and-wine pairing chart, drawn the way the class sheet lays it out:
// ten axes from a common centre, rings 0-10, the dish in red and the wine in blue.
// Each wine axis faces the food axes it has to answer, so a good pairing gives two
// polygons of the same shape and size.

export type FoodProfile = {
  sweetTendency: number
  fattiness: number
  juiciness: number
  oiliness: number
  hardness: number
  persistence: number
}

export type WineProfile = {
  softness: number
  intensity: number
  alcTannin: number
  acidity: number
}

export type PairingSpec = { kind: 'pairing'; food?: FoodProfile; wine?: WineProfile }

type Axis = { angle: number; label: string[]; side: 'food' | 'wine' }

// Angles in degrees, counter-clockwise from 3 o'clock, as on the sheet.
const FOOD: Record<keyof FoodProfile, Axis> = {
  oiliness: { angle: 22, label: ['Oiliness'], side: 'food' },
  juiciness: { angle: 52, label: ['Juiciness'], side: 'food' },
  sweetTendency: { angle: 128, label: ['Sweet tendency'], side: 'food' },
  fattiness: { angle: 158, label: ['Fattiness'], side: 'food' },
  hardness: { angle: 263, label: ['Saltiness', 'Bitter tendency', 'Sour tendency', 'Sweetness'], side: 'food' },
  persistence: { angle: 277, label: ['T-O persistence', 'Spiciness', 'Aromaticity'], side: 'food' },
}

const WINE: Record<keyof WineProfile, Axis> = {
  intensity: { angle: 84, label: ['T-O intensity', 'Persistence'], side: 'wine' },
  softness: { angle: 96, label: ['Sweetness', 'Roundness'], side: 'wine' },
  alcTannin: { angle: 205, label: ['Alcoholicity', 'Tannicity'], side: 'wine' },
  acidity: { angle: 335, label: ['Acidity', 'Effervescence', 'Saltiness'], side: 'wine' },
}

const R = 150 // radius of the 10 ring
const FOOD_COLOUR = '#d64545'
const WINE_COLOUR = '#3b7bd6'

function point(angle: number, score: number) {
  const rad = (angle * Math.PI) / 180
  return { x: (Math.cos(rad) * R * score) / 10, y: (-Math.sin(rad) * R * score) / 10 }
}

function polygon<K extends string>(axes: Record<K, Axis>, values: Record<K, number>) {
  return (Object.keys(axes) as K[])
    .map((k) => ({ a: axes[k].angle, v: values[k] }))
    .sort((p, q) => p.a - q.a)
    .map(({ a, v }) => point(a, v))
}

export default function PairingChart({ spec }: { spec: PairingSpec }) {
  const axes = [...Object.values(FOOD), ...Object.values(WINE)]
  const food = spec.food ? polygon(FOOD, spec.food) : null
  const wine = spec.wine ? polygon(WINE, spec.wine) : null
  const path = (pts: { x: number; y: number }[]) => pts.map((p) => `${p.x.toFixed(1)},${p.y.toFixed(1)}`).join(' ')

  return (
    <div>
      <svg viewBox="-270 -215 540 430" style={{ width: '100%', maxWidth: 560, height: 'auto', display: 'block', margin: '0 auto' }}
        role="img" aria-label={describe(spec)}>
        {Array.from({ length: 10 }, (_, i) => (
          <circle key={i} r={(R * (i + 1)) / 10} fill="none" stroke="var(--line)" strokeWidth={i === 9 ? 1.2 : 0.7} />
        ))}
        {axes.map((ax) => {
          const end = point(ax.angle, 10.6)
          const lab = point(ax.angle, 11.4)
          const anchor = Math.abs(lab.x) < 12 ? (ax.angle > 90 && ax.angle < 270 ? 'end' : 'start') : lab.x > 0 ? 'start' : 'end'
          const vertical = ax.angle > 80 && ax.angle < 100 ? -1 : ax.angle > 260 && ax.angle < 280 ? 1 : 0
          const colour = ax.side === 'food' ? FOOD_COLOUR : WINE_COLOUR
          return (
            <g key={ax.angle}>
              <line x1={0} y1={0} x2={end.x} y2={end.y} stroke="var(--ink-faint)" strokeWidth={0.9} />
              <text x={lab.x} y={lab.y + (vertical > 0 ? 6 : vertical < 0 ? -6 - 11 * (ax.label.length - 1) : -5.5 * (ax.label.length - 1))}
                textAnchor={anchor} fontSize="10.5" fill={colour} fontWeight={600}>
                {ax.label.map((l, i) => (
                  <tspan key={l} x={lab.x} dy={i === 0 ? 0 : 11.5}>{l}</tspan>
                ))}
              </text>
            </g>
          )
        })}
        {[2, 4, 6, 8, 10].map((n) => {
          const p = point(90, n)
          return <text key={n} x={p.x + 3} y={p.y + 3} fontSize="8" fill="var(--ink-faint)">{n}</text>
        })}
        {food && <polygon points={path(food)} fill={FOOD_COLOUR} fillOpacity={0.12} stroke={FOOD_COLOUR} strokeWidth={2.2} strokeLinejoin="round" />}
        {wine && <polygon points={path(wine)} fill={WINE_COLOUR} fillOpacity={0.12} stroke={WINE_COLOUR} strokeWidth={2.2} strokeLinejoin="round" />}
        {food && food.map((p, i) => <circle key={`f${i}`} cx={p.x} cy={p.y} r={3} fill={FOOD_COLOUR} />)}
        {wine && wine.map((p, i) => <circle key={`w${i}`} cx={p.x} cy={p.y} r={3} fill={WINE_COLOUR} />)}
      </svg>
      <p className="mt-1 text-center text-[0.75rem] text-ink-faint">
        <span style={{ color: FOOD_COLOUR, fontWeight: 700 }}>Red: dish</span>
        {' · '}
        <span style={{ color: WINE_COLOUR, fontWeight: 700 }}>Blue: wine</span>
        {' · '}scores 0-10 from the centre
      </p>
      <p className="mt-1 text-center text-[0.72rem] text-ink-faint">{describe(spec)}</p>
    </div>
  )
}

// The same scores in words, for screen readers and for reading values exactly.
function describe(spec: PairingSpec) {
  const parts: string[] = []
  if (spec.food) {
    const f = spec.food
    parts.push(`Dish: sweet tendency ${f.sweetTendency}, fattiness ${f.fattiness}, juiciness ${f.juiciness}, oiliness ${f.oiliness}, salt/bitter/sour ${f.hardness}, persistence/spice/aroma ${f.persistence}.`)
  }
  if (spec.wine) {
    const w = spec.wine
    parts.push(`Wine: softness ${w.softness}, intensity ${w.intensity}, alcohol/tannin ${w.alcTannin}, acidity/effervescence ${w.acidity}.`)
  }
  return parts.join(' ')
}
