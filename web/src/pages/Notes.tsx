import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import Markdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import { api, type Notes as NotesData, type Question } from '../lib/api'

export default function Notes() {
  const { topic = '' } = useParams()
  const [notes, setNotes] = useState<NotesData | null>(null)
  const [error, setError] = useState('')

  useEffect(() => {
    setNotes(null)
    api.notes(topic).then(setNotes).catch((e: Error) => setError(e.message))
  }, [topic])

  if (error) return <p className="text-critical">{error}</p>
  if (!notes) return <p className="text-ink-soft">Loading…</p>

  // The notes carry "?check id=…" markers. Splitting on them turns the page into an
  // alternation of reading and retrieval, which is the whole point of putting them there.
  const blocks: Array<{ kind: 'md'; text: string } | { kind: 'check'; question: Question }> = []
  let buffer: string[] = []
  for (const line of notes.body.split('\n')) {
    const marker = line.trim().match(/^\?check id=(\S+)$/)
    if (marker) {
      const question = notes.checks[marker[1]]
      if (question) {
        blocks.push({ kind: 'md', text: buffer.join('\n') })
        buffer = []
        blocks.push({ kind: 'check', question })
        continue
      }
      continue
    }
    buffer.push(line)
  }
  blocks.push({ kind: 'md', text: buffer.join('\n') })

  return (
    <article>
      <Link to="/read" className="text-[0.82rem] text-ink-soft underline underline-offset-4">
        All lectures
      </Link>
      <h1 className="mt-3 text-2xl">{notes.title}</h1>
      <p className="mt-1 text-[0.8rem] text-ink-faint">
        Deck {notes.deck} · {notes.slides} slides
      </p>
      <p className="mt-3 max-w-[62ch] text-[0.98rem] leading-relaxed text-ink-soft">{notes.summary}</p>

      <div className="mt-8">
        {blocks.map((b, i) =>
          b.kind === 'md' ? (
            <div key={i} className="prose">
              <Markdown remarkPlugins={[remarkGfm]}>{b.text}</Markdown>
            </div>
          ) : (
            <Check key={i} question={b.question} />
          ),
        )}
      </div>
    </article>
  )
}

// Check is a question dropped into the middle of the reading. It grades on the server
// like everything else, and it is not recorded — the point is to interrupt the reading,
// not to score it.
function Check({ question }: { question: Question }) {
  const [chosen, setChosen] = useState<number | null>(null)
  const [verdict, setVerdict] = useState<{ correct: boolean; answer: number[]; explanation: string } | null>(null)

  async function pick(i: number) {
    if (verdict) return
    setChosen(i)
    const res = await fetch('/api/check', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ questionId: question.id, given: [i] }),
    })
    setVerdict(await res.json())
  }

  return (
    <aside className="my-8 max-w-[68ch] border border-line-strong bg-surface">
      <p className="border-b border-line px-4 py-1.5 text-[0.74rem] text-ink-faint">Check yourself</p>
      <div className="px-4 py-3">
        <p className="text-[0.95rem] leading-snug">{question.stem}</p>
        <ul className="mt-3 space-y-1">
          {question.choices.map((c, i) => {
            const isAnswer = verdict?.answer.includes(i)
            const isWrongPick = verdict && chosen === i && !verdict.correct
            return (
              <li key={i}>
                <button
                  onClick={() => pick(i)}
                  disabled={!!verdict}
                  className="flex w-full items-start gap-2 border px-2.5 py-1.5 text-left text-[0.9rem]"
                  style={{
                    borderColor: isAnswer
                      ? 'var(--slack)'
                      : isWrongPick
                        ? 'var(--critical)'
                        : 'var(--line)',
                    background: isAnswer
                      ? 'var(--slack-soft)'
                      : isWrongPick
                        ? 'var(--critical-soft)'
                        : 'transparent',
                  }}
                >
                  <span
                    style={{
                      color: isAnswer
                        ? 'var(--slack-ink)'
                        : isWrongPick
                          ? 'var(--critical-ink)'
                          : 'var(--ink-faint)',
                    }}
                  >
                    {'ABCD'[i]}
                  </span>
                  <span>{c}</span>
                </button>
              </li>
            )
          })}
        </ul>
        {verdict && (
          <p className="mt-3 border-l-2 pl-3 text-[0.88rem] leading-relaxed text-ink-soft"
             style={{ borderColor: verdict.correct ? 'var(--slack)' : 'var(--critical)' }}>
            {verdict.explanation}
          </p>
        )}
      </div>
    </aside>
  )
}
