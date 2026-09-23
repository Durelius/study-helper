import { useState } from 'react'
import Markdown from 'react-markdown'
import remarkGfm from 'remark-gfm'
import type { Notes, Question } from '../lib/api'
import { toggleChoice } from '../lib/choices'

// NotesBody renders a lecture's markdown with its inline retrieval checks in place.
// It is shared by the Read page and the reading panel on the Audio tab, so a question
// answered while listening behaves exactly as it does while reading.
export default function NotesBody({ notes }: { notes: Notes }) {
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
      }
      continue
    }
    buffer.push(line)
  }
  blocks.push({ kind: 'md', text: buffer.join('\n') })

  return (
    <>
      {blocks.map((b, i) =>
        b.kind === 'md' ? (
          <div key={i} className="prose">
            <Markdown remarkPlugins={[remarkGfm]}>{b.text}</Markdown>
          </div>
        ) : (
          <Check key={i} question={b.question} />
        ),
      )}
    </>
  )
}

// Check is a question dropped into the middle of the reading. It grades on the server
// like everything else, and it is not recorded — the point is to interrupt the reading
// with a moment of recall, not to score it.
export function Check({ question }: { question: Question }) {
  const [chosen, setChosen] = useState<number[]>([])
  const [verdict, setVerdict] = useState<{ correct: boolean; answer: number[]; explanation: string } | null>(null)
  const multi = question.type === 'multi'

  async function grade(given: number[]) {
    const res = await fetch('/api/check', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ questionId: question.id, given }),
    })
    setVerdict(await res.json())
  }

  // A single-answer check grades on the click, as it always has. A select-all check
  // collects picks the way the quiz runner does and waits for an explicit Check, so
  // the first pick does not give away the rest of the answer.
  function pick(i: number) {
    if (verdict) return
    const next = toggleChoice(question.type, chosen, i)
    setChosen(next)
    if (!multi) grade(next)
  }

  return (
    <aside className="my-8 max-w-[68ch] border border-line-strong bg-surface">
      <p className="border-b border-line px-4 py-1.5 text-[0.74rem] text-ink-faint">Check yourself</p>
      <div className="px-4 py-3">
        <p className="text-[0.95rem] leading-snug">{question.stem}</p>
        <ul className="mt-3 space-y-1">
          {question.choices.map((c, i) => {
            const picked = chosen.includes(i)
            const isAnswer = verdict?.answer.includes(i)
            const isWrongPick = verdict && picked && !isAnswer
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
                        : picked
                          ? 'var(--ink)'
                          : 'var(--line)',
                    background: isAnswer
                      ? 'var(--slack-soft)'
                      : isWrongPick
                        ? 'var(--critical-soft)'
                        : picked
                          ? 'var(--sunken)'
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
        {multi && !verdict && (
          <div className="mt-3 flex items-center justify-between gap-3">
            <p className="text-[0.8rem] text-ink-faint">Pick every answer that applies.</p>
            <button onClick={() => grade(chosen)} disabled={!chosen.length} className="btn">
              Check
            </button>
          </div>
        )}
        {verdict && (
          <p
            className="mt-3 border-l-2 pl-3 text-[0.88rem] leading-relaxed text-ink-soft"
            style={{ borderColor: verdict.correct ? 'var(--slack)' : 'var(--critical)' }}
          >
            {verdict.explanation}
          </p>
        )}
      </div>
    </aside>
  )
}
