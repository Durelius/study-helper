import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { api, type Mode, type Topic } from '../lib/api'

export default function QuizPicker({ name }: { name: string }) {
  const nav = useNavigate()
  const [modes, setModes] = useState<Mode[]>([])
  const [topics, setTopics] = useState<Topic[]>([])
  const [picking, setPicking] = useState<Mode | null>(null)
  const [busy, setBusy] = useState('')
  const [error, setError] = useState('')

  useEffect(() => {
    api.modes().then(setModes).catch((e: Error) => setError(e.message))
    api.topics(name).then(setTopics)
  }, [name])

  async function start(mode: Mode, topic = '') {
    setBusy(mode.id)
    setError('')
    try {
      const run = await api.start(name, mode.id, topic)
      nav(`/quiz/${run.id}`, { state: run })
    } catch (e) {
      setError((e as Error).message)
      setBusy('')
    }
  }

  if (picking) {
    return (
      <div>
        <button onClick={() => setPicking(null)} className="text-[0.82rem] text-ink-soft underline underline-offset-4">
          All modes
        </button>
        <h1 className="mt-3 text-xl">{picking.title} — which lecture?</h1>
        <ul className="mt-5 divide-y divide-line border-y border-line">
          {topics
            .filter((t) => t.questions > 0)
            .map((t) => (
              <li key={t.id}>
                <button onClick={() => start(picking, t.id)} className="flex w-full items-center gap-3 py-3 text-left">
                  <span className="flex-1">{t.title}</span>
                  <span className="text-[0.8rem] text-ink-faint">{t.questions}</span>
                </button>
              </li>
            ))}
        </ul>
      </div>
    )
  }

  return (
    <div>
      <h1 className="text-xl">What do you want to practise?</h1>
      <p className="mt-1 max-w-[62ch] text-[0.92rem] leading-relaxed text-ink-soft">
        Every answer is marked on the server and the explanation comes back with it, so a wrong answer
        is worth more than a right one. Finished runs go on the leaderboard.
      </p>
      {error && <p className="mt-4 text-critical">{error}</p>}

      <ul className="mt-6 grid gap-3 sm:grid-cols-2">
        {modes.map((m) => {
          const empty = m.available === 0
          return (
            <li key={m.id}>
              <button
                disabled={empty || busy !== ''}
                onClick={() => (m.needsTopic ? setPicking(m) : start(m))}
                className="flex h-full w-full flex-col border border-line bg-surface p-4 text-left disabled:opacity-45"
              >
                <span className="flex items-baseline gap-2">
                  <span className="font-medium">{m.title}</span>
                  <span className="text-[0.76rem] text-ink-faint">
                    {empty ? 'not written yet' : m.count > 0 ? `${m.count} questions` : 'a full case'}
                    {m.timeLimitSec > 0 && ` · ${Math.round(m.timeLimitSec / 60)} min`}
                  </span>
                </span>
                <span className="mt-1 text-[0.9rem] text-ink-soft">{m.blurb}</span>
                <span className="mt-2.5 border-t border-line pt-2 text-[0.8rem] leading-relaxed text-ink-faint">
                  {m.why}
                </span>
                {busy === m.id && <span className="mt-2 text-[0.8rem] text-ink-soft">Building the run…</span>}
              </button>
            </li>
          )
        })}
      </ul>
    </div>
  )
}
