import { useEffect, useState } from 'react'
import { api, type Mode, type Session } from '../lib/api'
import { accuracy, ago, duration, percent } from '../lib/format'

export default function Leaderboard({ name }: { name: string }) {
  const [mode, setMode] = useState('')
  const [modes, setModes] = useState<Mode[]>([])
  const [top, setTop] = useState<Session[]>([])
  const [recent, setRecent] = useState<Session[]>([])

  useEffect(() => {
    api.modes().then(setModes).catch(() => setModes([]))
  }, [])

  useEffect(() => {
    api.leaderboard(mode).then((l) => {
      setTop(l.top)
      setRecent(l.recent)
    })
  }, [mode])

  return (
    <div className="grid gap-10 md:grid-cols-[1.4fr_1fr]">
      <div>
        <h1 className="text-xl">Leaderboard</h1>
        <p className="mt-1 max-w-[56ch] text-[0.9rem] text-ink-soft">
          Accuracy first, then the clock. Grading happens on the server, so the only way up the board is
          to know the material.
        </p>

        <div className="mt-4 flex flex-wrap gap-1.5">
          <Chip active={mode === ''} onClick={() => setMode('')}>
            Every mode
          </Chip>
          {modes.map((m) => (
            <Chip key={m.id} active={mode === m.id} onClick={() => setMode(m.id)}>
              {m.title}
            </Chip>
          ))}
        </div>

        {top.length === 0 ? (
          <p className="mt-6 border border-dashed border-line-strong p-4 text-[0.9rem] text-ink-soft">
            No finished runs in this mode yet.
          </p>
        ) : (
          <ol className="mt-5 divide-y divide-line border-y border-line">
            {top.map((s, i) => (
              <li key={s.id} className="flex items-center gap-3 py-3">
                <span className="w-6 text-[0.85rem] text-ink-faint">{i + 1}</span>
                <span className="min-w-0 flex-1">
                  <span className={`block truncate ${s.player === name ? 'font-semibold' : ''}`}>{s.player}</span>
                  <span className="text-[0.76rem] text-ink-faint">
                    {s.mode}
                    {s.topic ? ` · ${s.topic}` : ''} · {duration(s.durationMs)} · {ago(s.finishedAt ?? 0)}
                  </span>
                </span>
                <span className="relative h-5 w-24 shrink-0 bg-sunken">
                  <span
                    className="absolute inset-y-0 left-0"
                    style={{ width: `${accuracy(s.score, s.total) * 100}%`, background: 'var(--slack)' }}
                  />
                </span>
                <span className="w-20 text-right text-[0.9rem] font-medium">
                  {percent(accuracy(s.score, s.total))}
                  <span className="ml-1.5 text-[0.76rem] font-normal text-ink-faint">
                    {s.score}/{s.total}
                  </span>
                </span>
              </li>
            ))}
          </ol>
        )}
      </div>

      <div>
        <h2 className="text-[0.95rem]">Just finished</h2>
        <ul className="mt-3 divide-y divide-line border-y border-line">
          {recent.map((s) => (
            <li key={s.id} className="flex items-baseline gap-2 py-2 text-[0.88rem]">
              <span className="min-w-0 flex-1 truncate">{s.player}</span>
              <span className="text-ink-faint">{s.mode}</span>
              <span className="font-medium">
                {s.score}/{s.total}
              </span>
            </li>
          ))}
          {recent.length === 0 && <li className="py-2 text-[0.88rem] text-ink-soft">Nothing yet.</li>}
        </ul>
      </div>
    </div>
  )
}

function Chip({ active, onClick, children }: { active: boolean; onClick: () => void; children: React.ReactNode }) {
  return (
    <button
      onClick={onClick}
      className={`border px-2.5 py-1 text-[0.82rem] ${
        active ? 'border-ink bg-ink text-paper' : 'border-line-strong text-ink-soft'
      }`}
    >
      {children}
    </button>
  )
}
