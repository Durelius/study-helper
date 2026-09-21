import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import { api, type Review, type Session } from '../lib/api'
import { accuracy, duration, percent } from '../lib/format'

type Breakdown = { topic: string; title: string; seen: number; correct: number }

export default function Results() {
  const { id = '' } = useParams()
  const [data, setData] = useState<{ session: Session; review: Review[]; breakdown: Breakdown[] } | null>(null)
  const [error, setError] = useState('')

  useEffect(() => {
    api.finish(id).then(setData).catch((e: Error) => setError(e.message))
  }, [id])

  if (error) return <p className="text-critical">{error}</p>
  if (!data) return <p className="text-ink-soft">Marking…</p>

  const { session, review } = data
  const acc = accuracy(session.score, session.total)
  // The misses are the reason to be on this page, so they come first and in full.
  const missed = review.filter((_, i) => i < review.length)

  return (
    <div className="mx-auto max-w-2xl">
      <p className="text-[0.82rem] text-ink-soft">{session.mode}</p>
      <h1 className="mt-1 flex items-baseline gap-3">
        <span className="text-4xl font-semibold tracking-tight">
          {session.score}
          <span className="text-ink-faint">/{session.total}</span>
        </span>
        <span className="text-lg text-ink-soft">{percent(acc)}</span>
        <span className="ml-auto text-[0.85rem] text-ink-faint">{duration(session.durationMs)}</span>
      </h1>

      <div className="mt-3 h-2 bg-sunken">
        <div className="h-full" style={{ width: `${acc * 100}%`, background: 'var(--slack)' }} />
      </div>

      <p className="mt-4 text-[0.92rem] leading-relaxed text-ink-soft">
        This run is on the leaderboard. Everything below is the whole paper with the reasoning — read the
        ones you got wrong now, while you still remember what you were thinking.
      </p>

      <div className="mt-5 flex gap-2">
        <Link to="/quiz" className="btn">
          Another run
        </Link>
        <Link to="/" className="btn btn-quiet">
          Back to the plan
        </Link>
      </div>

      <h2 className="mt-10 text-[0.95rem]">Every question, with the answer</h2>
      <ol className="mt-3 space-y-5">
        {missed.map((r, i) => (
          <li key={r.question.id} className="border-t border-line pt-4">
            <p className="text-[0.76rem] text-ink-faint">
              {i + 1} · {r.question.topicTitle}
              {r.source.slide > 0 && ` · deck ${r.source.deck}, slide ${r.source.slide}`}
            </p>
            <p className="mt-1.5 text-[0.98rem] leading-snug">{r.question.stem}</p>
            <ul className="mt-2 space-y-1">
              {r.question.choices.map((c, ci) => {
                const isAnswer = r.answer.includes(ci)
                return (
                  <li
                    key={ci}
                    className="border px-2.5 py-1.5 text-[0.9rem]"
                    style={{
                      borderColor: isAnswer ? 'var(--slack)' : 'var(--line)',
                      background: isAnswer ? 'var(--slack-soft)' : 'transparent',
                    }}
                  >
                    {c}
                  </li>
                )
              })}
            </ul>
            <p className="mt-2 border-l-2 border-line-strong pl-3 text-[0.88rem] leading-relaxed text-ink-soft">
              {r.explanation}
            </p>
          </li>
        ))}
      </ol>
    </div>
  )
}
