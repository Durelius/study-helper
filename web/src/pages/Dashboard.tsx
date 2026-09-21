import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import Countdown from '../components/Countdown'
import { api, type Plan, type Session, type Topic } from '../lib/api'
import { accuracy, ago, duration, percent } from '../lib/format'

export default function Dashboard({ name }: { name: string }) {
  const nav = useNavigate()
  const [plan, setPlan] = useState<Plan | null>(null)
  const [topics, setTopics] = useState<Topic[]>([])
  const [board, setBoard] = useState<Session[]>([])
  const [error, setError] = useState('')
  const [tick, setTick] = useState(0)

  useEffect(() => {
    Promise.all([api.plan(name), api.topics(name), api.leaderboard()])
      .then(([p, t, l]) => {
        setPlan(p)
        setTopics(t)
        setBoard(l.top.slice(0, 6))
      })
      .catch((e: Error) => setError(e.message))
  }, [name])

  // One interval drives the countdown; the rest of the page is static until reload.
  useEffect(() => {
    const id = setInterval(() => setTick((n) => n + 1), 1000)
    return () => clearInterval(id)
  }, [])

  if (error) return <p className="text-critical">{error}</p>
  if (!plan) return <p className="text-ink-soft">Loading…</p>

  const seconds = plan.secondsRemaining - tick
  const next = plan.steps[0]

  return (
    <div className="grid gap-10 md:grid-cols-[1.15fr_1fr]">
      <div>
        <Countdown seconds={seconds} />

        {next && (
          <section className="mt-8 border-l-2 border-critical bg-surface p-4">
            <p className="text-[0.8rem] text-ink-soft">Do this next</p>
            <h2 className="mt-1 text-lg">{next.label}</h2>
            <p className="mt-1.5 max-w-[52ch] text-[0.92rem] leading-relaxed text-ink-soft">{next.why}</p>
            <button
              onClick={() => (next.action === 'read' ? nav(`/read/${next.topic}`) : start(next.mode!, next.topic ?? ''))}
              className="mt-3 border border-ink bg-ink px-4 py-2 text-[0.9rem] font-medium text-paper"
            >
              {next.action === 'read' ? 'Open the notes' : 'Start'}
            </button>
          </section>
        )}

        {plan.steps.length > 1 && (
          <section className="mt-6">
            <h3 className="text-[0.82rem] text-ink-soft">Then</h3>
            <ul className="mt-2 divide-y divide-line border-y border-line">
              {plan.steps.slice(1).map((s, i) => (
                <li key={i}>
                  <button
                    onClick={() => (s.action === 'read' ? nav(`/read/${s.topic}`) : start(s.mode!, s.topic ?? ''))}
                    className="w-full py-2.5 text-left"
                  >
                    <span className="text-[0.92rem]">{s.label}</span>
                    <span className="mt-0.5 block max-w-[52ch] text-[0.82rem] text-ink-faint">{s.why}</span>
                  </button>
                </li>
              ))}
            </ul>
          </section>
        )}

        <section className="mt-8">
          <h3 className="text-[0.82rem] text-ink-soft">
            Your coverage · {plan.seen} of {plan.bank} questions seen · {percent(plan.accuracy)} right
          </h3>
          <ul className="mt-3 grid gap-2 sm:grid-cols-2">
            {topics.map((t) => (
              <TopicTile key={t.id} topic={t} onDrill={() => start('topic', t.id)} />
            ))}
          </ul>
        </section>
      </div>

      <div>
        <h2 className="text-lg">Leaderboard</h2>
        <p className="mt-1 text-[0.85rem] text-ink-soft">
          Ranked by accuracy, then by time. Runs of five questions or more.
        </p>
        {board.length === 0 ? (
          <p className="mt-4 border border-dashed border-line-strong p-4 text-[0.9rem] text-ink-soft">
            Nobody has finished a quiz yet. Be the first and the board is yours.
          </p>
        ) : (
          <ol className="mt-4 divide-y divide-line border-y border-line">
            {board.map((s, i) => (
              <li key={s.id} className="flex items-center gap-3 py-2.5">
                <span className="w-5 text-[0.85rem] text-ink-faint">{i + 1}</span>
                <span className="min-w-0 flex-1">
                  <span className={`block truncate text-[0.94rem] ${s.player === name ? 'font-semibold' : ''}`}>
                    {s.player}
                  </span>
                  <span className="text-[0.76rem] text-ink-faint">
                    {s.mode} · {duration(s.durationMs)}
                  </span>
                </span>
                <span className="relative h-5 w-20 shrink-0 bg-sunken">
                  <span
                    className="absolute inset-y-0 left-0"
                    style={{ width: `${accuracy(s.score, s.total) * 100}%`, background: 'var(--slack)' }}
                  />
                </span>
                <span className="w-14 text-right text-[0.88rem] font-medium">
                  {s.score}/{s.total}
                </span>
              </li>
            ))}
          </ol>
        )}
        <button onClick={() => nav('/leaderboard')} className="mt-3 text-[0.85rem] underline underline-offset-4">
          Full leaderboard
        </button>
      </div>
    </div>
  )

  async function start(mode: string, topic: string) {
    try {
      const run = await api.start(name, mode, topic)
      nav(`/quiz/${run.id}`, { state: run })
    } catch (e) {
      setError((e as Error).message)
    }
  }
}

// The topic tile borrows the shape of the activity box from the critical-path slides:
// figures along the top, the name in a band through the middle, a bar underneath.
function TopicTile({ topic, onDrill }: { topic: Topic; onDrill: () => void }) {
  const acc = accuracy(topic.correct, topic.seen)
  const untouched = topic.seen === 0
  const shaky = !untouched && acc < 0.7
  return (
    <li>
      <button
        onClick={onDrill}
        disabled={topic.questions === 0}
        className="w-full border border-line bg-surface text-left disabled:opacity-50"
      >
        <span className="flex justify-between border-b border-line px-2.5 py-1 text-[0.72rem] text-ink-faint">
          <span>{topic.questions} questions</span>
          <span>{untouched ? 'not started' : ago(topic.lastSeen)}</span>
        </span>
        <span className="block px-2.5 py-1.5 text-[0.92rem] leading-tight">{topic.title}</span>
        <span className="flex items-center gap-2 border-t border-line px-2.5 py-1">
          <span className="relative h-2 flex-1 bg-sunken">
            <span
              className="absolute inset-y-0 left-0"
              style={{
                width: `${acc * 100}%`,
                background: shaky ? 'var(--warn)' : 'var(--slack)',
              }}
            />
          </span>
          <span className="w-16 text-right text-[0.74rem] text-ink-soft">
            {untouched ? '—' : `${percent(acc)} of ${topic.seen}`}
          </span>
        </span>
      </button>
    </li>
  )
}
