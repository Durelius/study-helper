import { useEffect, useState } from 'react'
import { Link } from 'react-router-dom'
import { api, type Topic } from '../lib/api'
import { accuracy, ago, percent } from '../lib/format'

export default function Read({ name }: { name: string }) {
  const [topics, setTopics] = useState<Topic[]>([])

  useEffect(() => {
    api.topics(name).then(setTopics)
  }, [name])

  return (
    <div>
      <h1 className="text-xl">The nine lectures, condensed</h1>
      <p className="mt-1 max-w-[62ch] text-[0.92rem] leading-relaxed text-ink-soft">
        Each page is the whole deck boiled down to what an exam can ask, with the slide number beside
        every section so you can go back to the original. Questions are dropped into the text as you
        read — answer them rather than skipping them, because that is the part that makes it stick.
      </p>

      <ul className="mt-6 divide-y divide-line border-y border-line">
        {topics.map((t) => (
          <li key={t.id}>
            <Link to={`/read/${t.id}`} className="flex items-center gap-4 py-3.5">
              <span className="min-w-0 flex-1">
                <span className="block">{t.title}</span>
                <span className="text-[0.78rem] text-ink-faint">
                  {t.questions} questions · {t.seen === 0 ? 'not started' : `last studied ${ago(t.lastSeen)}`}
                </span>
              </span>
              {t.seen > 0 && (
                <span className="text-[0.82rem] text-ink-soft">{percent(accuracy(t.correct, t.seen))}</span>
              )}
              <span aria-hidden="true" className="text-ink-faint">
                ›
              </span>
            </Link>
          </li>
        ))}
      </ul>
    </div>
  )
}
