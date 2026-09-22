import { useEffect, useState } from 'react'
import { Link, useParams } from 'react-router-dom'
import NotesBody from '../components/NotesBody'
import { api, type Notes as NotesData } from '../lib/api'

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
        <NotesBody notes={notes} />
      </div>
    </article>
  )
}
