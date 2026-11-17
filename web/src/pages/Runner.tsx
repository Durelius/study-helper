import { useCallback, useEffect, useMemo, useRef, useState } from 'react'
import { useLocation, useNavigate, useParams } from 'react-router-dom'
import Markdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import Diagram from '../components/Diagram'
import NetworkDiagram from '../components/NetworkDiagram'
import PairingChart, { type PairingSpec } from '../components/PairingChart'
import type { DiagramSpec } from '../lib/diagram'
import { api, type Mode, type Network, type Question, type Verdict } from '../lib/api'
import { duration } from '../lib/format'
import { toggleChoice } from '../lib/choices'
import { SourceLink } from '../components/SourceViewer'

type RunState = { id: string; player: string; mode: Mode; questions: Question[] }

export default function Runner() {
  const { id = '' } = useParams()
  const nav = useNavigate()
  const location = useLocation()
  const passed = location.state as RunState | null

  const [run, setRun] = useState<RunState | null>(passed)
  const [index, setIndex] = useState(0)
  const [chosen, setChosen] = useState<number[]>([])
  const [typed, setTyped] = useState('')
  const [verdict, setVerdict] = useState<Verdict | null>(null)
  const [score, setScore] = useState(0)
  const [error, setError] = useState('')
  const [remaining, setRemaining] = useState(0)
  const startedAt = useRef(Date.now())

  // A refresh loses the questions held in router state, so they are fetched back.
  // The server still withholds every answer key, so this is safe to re-serve.
  useEffect(() => {
    if (run) return
    fetch(`/api/quiz/${id}`)
      .then((r) => (r.ok ? r.json() : Promise.reject(new Error('That quiz has expired.'))))
      .then((data: RunState & { answered: string[] }) => {
        setRun(data)
        setIndex(Math.min(data.answered.length, data.questions.length - 1))
      })
      .catch((e: Error) => setError(e.message))
  }, [id, run])

  const finish = useCallback(async () => {
    try {
      await api.finish(id)
      nav(`/results/${id}`, { replace: true })
    } catch (e) {
      setError((e as Error).message)
    }
  }, [id, nav])

  const limit = run?.mode.timeLimitSec ?? 0
  useEffect(() => {
    if (!limit) return
    const started = Date.now()
    const tick = setInterval(() => {
      const left = limit - Math.floor((Date.now() - started) / 1000)
      setRemaining(left)
      if (left <= 0) {
        clearInterval(tick)
        void finish()
      }
    }, 500)
    return () => clearInterval(tick)
  }, [limit, finish])

  const question = run?.questions[index]

  const network = useMemo(() => {
    if (!question?.data) return null
    const d = question.data as Partial<Network>
    return Array.isArray(d.activities) ? (question.data as Network) : null
  }, [question])

  // A hotspot question carries a diagram whose clickable regions are its choices.
  const diagram = useMemo(() => {
    if (!question?.data) return null
    const d = question.data as { kind?: string }
    return d.kind === 'bpmn' || d.kind === 'wbs' || d.kind === 'matrix' ? (question.data as DiagramSpec) : null
  }, [question])

  const pairing = useMemo(() => {
    if (!question?.data) return null
    const d = question.data as { kind?: string }
    return d.kind === 'pairing' ? (question.data as PairingSpec) : null
  }, [question])

  const scenario = useMemo(() => {
    if (!question?.data) return null
    const d = question.data as { scenario?: string; title?: string; image?: string }
    return typeof d.scenario === 'string' ? d : null
  }, [question])

  const toggle = useCallback(
    (i: number) => {
      if (verdict || !question) return
      setChosen((prev) => toggleChoice(question.type, prev, i))
    },
    [question, verdict],
  )

  // Number keys pick, Enter submits: the whole run should be doable without a mouse,
  // because speed is half of what the leaderboard measures.
  useEffect(() => {
    function onKey(e: KeyboardEvent) {
      if (!question) return
      // A slide open over the question owns the keyboard, or Enter on its Close
      // button would also skip to the next question.
      if (document.querySelector('[role=dialog]')) return
      // On a free-text question the number keys are part of the answer, not a choice,
      // and a match question is answered in its dropdowns.
      if (question.type !== 'text' && question.type !== 'num' && question.type !== 'match') {
        const n = Number(e.key)
        if (n >= 1 && n <= question.choices.length) toggle(n - 1)
      }
      if (e.key === 'Enter' && verdict) next()
    }
    window.addEventListener('keydown', onKey)
    return () => window.removeEventListener('keydown', onKey)
  })

  if (error) {
    return (
      <div>
        <p className="text-critical">{error}</p>
        <button onClick={() => nav('/quiz')} className="mt-3 underline underline-offset-4">
          Pick another run
        </button>
      </div>
    )
  }
  if (!run || !question) return <p className="text-ink-soft">Loading…</p>

  // A match question holds one pick per item, in item order; -1 is a row not yet paired.
  const isMatch = question.type === 'match'
  const pairs = isMatch ? (question.items ?? []).map((_, i) => chosen[i] ?? -1) : []
  const typedAnswer = question.type === 'text' || question.type === 'num'
  const ready = typedAnswer ? !!typed.trim() : isMatch ? pairs.every((p) => p >= 0) : chosen.length > 0

  async function submit(confident: boolean) {
    if (verdict || !question || !run) return
    if (!ready) return
    try {
      const res = await api.answer(id, question.id, isMatch ? pairs : chosen, confident, Date.now() - startedAt.current, typed)
      setVerdict(res)
      if (res.correct) setScore((s) => s + 1)
    } catch (e) {
      setError((e as Error).message)
    }
  }

  function next() {
    if (!run) return
    if (index + 1 >= run.questions.length) {
      void finish()
      return
    }
    setIndex((i) => i + 1)
    setChosen([])
    setTyped('')
    setVerdict(null)
    startedAt.current = Date.now()
    window.scrollTo({ top: 0 })
  }

  const progress = (index + (verdict ? 1 : 0)) / run.questions.length

  return (
    <div className="mx-auto max-w-2xl">
      <div className="flex items-baseline gap-3 text-[0.82rem] text-ink-soft">
        <span>
          {index + 1} of {run.questions.length}
        </span>
        {question.part ? (
          <span className="font-semibold text-ink">{question.part}</span>
        ) : (
          <span className="text-ink-faint">{question.topicTitle}</span>
        )}
        <span className="ml-auto tabular-nums">
          {limit > 0 ? (
            <span style={{ color: remaining < 120 ? 'var(--critical)' : undefined }}>
              {duration(Math.max(0, remaining) * 1000)} left
            </span>
          ) : (
            `${score} right`
          )}
        </span>
      </div>
      <div className="mt-1.5 h-1 bg-sunken" aria-hidden="true">
        <div className="h-full bg-ink" style={{ width: `${progress * 100}%`, transition: 'width .25s' }} />
      </div>

      {scenario && (
        <section className="mt-6 border-l-2 border-line-strong bg-surface px-4 py-3">
          <h2 className="text-[0.95rem] font-medium">{scenario.title}</h2>
          {scenario.image && (
            <img
              src={scenario.image}
              alt={scenario.title ?? 'Wine label'}
              className="mx-auto mt-3 max-h-[50vh] w-auto max-w-full"
            />
          )}
          <div className="prose mt-2 text-[0.92rem]">
            <Markdown remarkPlugins={[remarkGfm]}>{scenario.scenario ?? ''}</Markdown>
          </div>
        </section>
      )}

      <div className="prose mt-7 text-[1.02rem]" style={{ fontFamily: 'Archivo, sans-serif', lineHeight: 1.45 }}>
        <Markdown remarkPlugins={[remarkGfm]}>{question.stem}</Markdown>
      </div>

      {pairing && (
        <div className="mt-4 border border-line bg-surface p-3">
          <PairingChart spec={pairing} />
        </div>
      )}

      {network && (
        <div className="mt-4 border border-line bg-surface p-3">
          <NetworkDiagram net={network} reveal={!!verdict} />
        </div>
      )}

      {diagram && (
        <div className="mt-4 border border-line bg-surface p-3">
          <Diagram
            spec={diagram}
            selected={chosen}
            correct={verdict?.answer ?? []}
            revealed={!!verdict}
            onPick={toggle}
          />
          {/* Only a hotspot question has anything to click; other questions just
              carry a table or a network for reference. */}
          {question.type === 'hotspot' && (
            <p className="mt-1 text-[0.72rem] text-ink-faint">
              {verdict
                ? 'The correct element is outlined in green.'
                : 'Click the part of the diagram you think is wrong, or pick from the list below.'}
            </p>
          )}
        </div>
      )}

      {question.type === 'multi' && !verdict && (
        <p className="mt-3 text-[0.82rem] text-ink-faint">Pick every answer that applies.</p>
      )}

      {isMatch ? (
        <MatchRows
          items={question.items ?? []}
          choices={question.choices}
          pairs={pairs}
          answer={verdict?.answer}
          onPair={(row, choice) =>
            setChosen(pairs.map((p, i) => (i === row ? choice : p)))
          }
        />
      ) : typedAnswer ? (
        <div className="mt-4">
          <label htmlFor="answer" className="block text-[0.82rem] text-ink-soft">
            {question.type === 'num' ? 'Your answer' : 'Type your answer'}
          </label>
          <div className="mt-1.5 flex items-center gap-2">
          <input
            id="answer"
            autoFocus
            autoComplete="off"
            autoCapitalize="off"
            spellCheck={false}
            value={typed}
            disabled={!!verdict}
            onChange={(e) => setTyped(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === 'Enter' && !verdict && typed.trim()) submit(true)
            }}
            inputMode={question.type === 'num' ? 'decimal' : undefined}
            placeholder={question.type === 'num' ? 'a number' : 'one word is enough'}
            className="w-full border px-3 py-2.5 text-[1.05rem]"
            style={{
              borderColor: verdict
                ? verdict.correct
                  ? 'var(--slack)'
                  : 'var(--critical)'
                : 'var(--line-strong)',
              background: verdict
                ? verdict.correct
                  ? 'var(--slack-soft)'
                  : 'var(--critical-soft)'
                : 'var(--surface)',
              borderWidth: verdict ? 2 : 1,
            }}
          />
          {question.unit && <span className="shrink-0 text-[0.95rem] text-ink-soft">{question.unit}</span>}
          </div>
          <p className="mt-1 text-[0.76rem] text-ink-faint">
            {question.type === 'num'
              ? 'Work it out on paper first. Fractions like 1/2, commas like 4,096 and two decimal places are all fine.'
              : 'Spelling is marked generously — capitals, punctuation and a slipped letter are fine.'}
          </p>
          {verdict && !verdict.correct && verdict.accept && verdict.accept.length > 0 && (
            <p className="mt-2 text-[0.92rem]">
              <span className="text-ink-soft">The answer was </span>
              <strong style={{ color: 'var(--slack)' }}>{verdict.accept[0]}</strong>
              {verdict.accept.length > 1 && (
                <span className="text-ink-faint">
                  {' '}
                  (also accepted: {verdict.accept.slice(1).join(', ')})
                </span>
              )}
            </p>
          )}
        </div>
      ) : (
      <ul className="mt-4 space-y-1.5">
        {question.choices.map((c, i) => {
          const picked = chosen.includes(i)
          const isAnswer = verdict?.answer.includes(i)
          const wrongPick = verdict && picked && !isAnswer
          return (
            <li key={i}>
              <button
                onClick={() => toggle(i)}
                disabled={!!verdict}
                className="flex w-full items-start gap-3 border px-3 py-2.5 text-left"
                style={{
                  borderColor: isAnswer
                    ? 'var(--slack)'
                    : wrongPick
                      ? 'var(--critical)'
                      : picked
                        ? 'var(--ink)'
                        : 'var(--line)',
                  background: isAnswer
                    ? 'var(--slack-soft)'
                    : wrongPick
                      ? 'var(--critical-soft)'
                      : picked
                        ? 'var(--sunken)'
                        : 'var(--surface)',
                  borderWidth: picked || isAnswer || wrongPick ? 2 : 1,
                }}
              >
                <span
                  className="w-4 shrink-0 text-[0.85rem]"
                  // The marker has to sit on the verdict tint once the answer is in,
                  // so it takes the matching ink rather than the neutral grey.
                  style={{
                    color: isAnswer
                      ? 'var(--slack-ink)'
                      : wrongPick
                        ? 'var(--critical-ink)'
                        : 'var(--ink-faint)',
                  }}
                >
                  {i + 1}
                </span>
                <span className="text-[0.96rem]">{c}</span>
              </button>
            </li>
          )
        })}
      </ul>
      )}

      {!verdict ? (
        // Two buttons rather than a checkbox: committing to "sure" or "not sure" is
        // one decision, and it is the one that makes the weak-spot list worth having.
        <div className="mt-5 flex gap-2">
          <button
            onClick={() => submit(true)}
            disabled={!ready}
            className="btn flex-1"
          >
            I'm sure
          </button>
          <button
            onClick={() => submit(false)}
            disabled={!ready}
            className="btn btn-quiet flex-1"
          >
            Not sure
          </button>
        </div>
      ) : (
        <div className="mt-5">
          <p
            className="font-medium"
            style={{ color: verdict.correct ? 'var(--slack)' : 'var(--critical)' }}
          >
            {verdict.correct ? 'Right' : 'Not this time'}
          </p>
          <p className="mt-1.5 text-[0.95rem] leading-relaxed text-ink-soft">{verdict.explanation}</p>
          {verdict.source.slide > 0 && (
            <p className="mt-2 text-[0.78rem] text-ink-faint">
              Source: <SourceLink source={verdict.source} />
            </p>
          )}
          <button
            onClick={next}
            autoFocus
            className="btn mt-4 w-full"
          >
            {index + 1 >= run.questions.length ? 'See how you did' : 'Next question'}
          </button>
        </div>
      )}
    </div>
  )
}

// MatchRows is a matching question: each item on the left gets one choice from the
// pool. The pool is longer than the list, so the last pair is never a gimme. Once the
// answer is in, each row says whether it was right and, if not, what it should have been.
function MatchRows({
  items,
  choices,
  pairs,
  answer,
  onPair,
}: {
  items: string[]
  choices: string[]
  pairs: number[]
  answer?: number[]
  onPair: (row: number, choice: number) => void
}) {
  return (
    <ul className="mt-4 space-y-1.5">
      {items.map((item, i) => {
        const right = answer ? pairs[i] === answer[i] : null
        return (
          <li
            key={i}
            className="flex flex-wrap items-center gap-x-3 gap-y-1.5 border px-3 py-2.5"
            style={{
              borderColor: right === null ? 'var(--line)' : right ? 'var(--slack)' : 'var(--critical)',
              background: right === null ? 'var(--surface)' : right ? 'var(--slack-soft)' : 'var(--critical-soft)',
            }}
          >
            <span className="min-w-[9rem] flex-1 text-[0.96rem]">{item}</span>
            <select
              value={pairs[i]}
              disabled={!!answer}
              onChange={(e) => onPair(i, Number(e.target.value))}
              aria-label={`Match for ${item}`}
              className="min-w-[11rem] flex-1 border border-line-strong bg-surface px-2 py-1.5 text-[0.95rem]"
            >
              <option value={-1}>Choose…</option>
              {choices.map((c, ci) => (
                <option key={ci} value={ci}>
                  {c}
                </option>
              ))}
            </select>
            {answer && !right && (
              <span className="w-full text-[0.85rem]" style={{ color: 'var(--slack-ink)' }}>
                → {choices[answer[i]]}
              </span>
            )}
          </li>
        )
      })}
    </ul>
  )
}
