import { useEffect, useRef, useState } from 'react'
import { Link } from 'react-router-dom'
import NotesBody from './NotesBody'
import { api, type Notes } from '../lib/api'

// ReadingView shows a lecture's notes in place of the episode list, so you can follow
// the text while it is read to you.
//
// It is deliberately not a modal. Staying on the Audio route means the <audio> element
// never unmounts and playback continues untouched, and the player pinned to the bottom
// of the page keeps working exactly as it did — no overlay, no scroll trap, and the
// notes get the full width they were designed for.
export default function ReadingView({
  topic,
  onBack,
  isPlaying,
  onPlay,
}: {
  topic: string
  onBack: () => void
  // Whether the episode currently loaded is the one these notes belong to. Reading one
  // lecture while listening to another is allowed, but it should be obvious.
  isPlaying: boolean
  onPlay: () => void
}) {
  const [notes, setNotes] = useState<Notes | null>(null)
  const [error, setError] = useState('')
  const top = useRef<HTMLDivElement | null>(null)

  useEffect(() => {
    setNotes(null)
    setError('')
    api.notes(topic).then(setNotes).catch((e: Error) => setError(e.message))
    window.scrollTo({ top: 0 })
  }, [topic])

  return (
    <div ref={top}>
      <div className="flex items-baseline gap-3">
        <button onClick={onBack} className="text-[0.85rem] text-ink-soft underline underline-offset-4">
          Episodes
        </button>
        {!isPlaying && (
          <button onClick={onPlay} className="ml-auto text-[0.85rem] underline underline-offset-4">
            Listen to this one
          </button>
        )}
        <Link
          to={`/read/${topic}`}
          className={`${isPlaying ? 'ml-auto' : ''} text-[0.8rem] text-ink-soft underline underline-offset-4`}
        >
          Open on its own page
        </Link>
      </div>

      {error && <p className="mt-4 text-critical">{error}</p>}
      {!notes && !error && <p className="mt-4 text-ink-soft">Loading the notes…</p>}

      {notes && (
        <article className="mt-3">
          <h1 className="text-2xl">{notes.title}</h1>
          <p className="mt-1 text-[0.8rem] text-ink-faint">
            Deck {notes.deck} · {notes.slides} slides
            {isPlaying && <span style={{ color: 'var(--slack)' }}> · playing now</span>}
          </p>
          <p className="mt-3 max-w-[62ch] text-[0.98rem] leading-relaxed text-ink-soft">{notes.summary}</p>
          <div className="mt-8">
            <NotesBody notes={notes} />
          </div>
        </article>
      )}
    </div>
  )
}
