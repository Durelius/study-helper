import { useRef, useState } from 'react'
import { Link } from 'react-router-dom'
import ConfirmDialog from '../components/ConfirmDialog'
import { useAudio } from '../components/AudioPlayer'
import type { Episode } from '../lib/api'
import { clock, megabytes } from '../lib/format'
import { forgetDownload, loadDownloads, rememberDownload, type DownloadRecord } from '../lib/downloads'

// Just the episode list: the player itself lives in the app shell, so this page can be
// navigated away from without interrupting anything.
export default function AudioPage() {
  const { episodes, minutes, current, playing, listened, play } = useAudio()
  const [downloads, setDownloads] = useState<Record<string, DownloadRecord>>(loadDownloads)
  const [saving, setSaving] = useState<Record<string, number>>({})
  const [confirming, setConfirming] = useState(false)
  const [bulk, setBulk] = useState<{ done: number; total: number } | null>(null)
  const [cancelBulk, setCancelBulk] = useState(false)
  // A ref, not state: the download loop reads this between episodes and would never
  // see a state update captured in its own closure.
  const stopRequested = useRef(false)
  const [error, setError] = useState('')

  // Fetching the file rather than linking it means the progress is real and the
  // "downloaded" mark records something that actually happened.
  async function download(ep: Episode) {
    setSaving((s) => ({ ...s, [ep.topic]: 0 }))
    try {
      const res = await fetch(ep.url)
      if (!res.ok || !res.body) throw new Error(`Download failed (${res.status})`)
      const total = Number(res.headers.get('Content-Length')) || ep.bytes
      const reader = res.body.getReader()
      const parts: Uint8Array[] = []
      let received = 0
      for (;;) {
        const { done, value } = await reader.read()
        if (done) break
        parts.push(value)
        received += value.length
        setSaving((s) => ({ ...s, [ep.topic]: received / total }))
      }
      const url = URL.createObjectURL(new Blob(parts as BlobPart[], { type: 'audio/mp4' }))
      const a = document.createElement('a')
      a.href = url
      a.download = ep.file
      a.click()
      URL.revokeObjectURL(url)
      setDownloads(rememberDownload(ep.topic, received))
    } catch (e) {
      setError((e as Error).message)
    } finally {
      setSaving((s) => {
        const next = { ...s }
        delete next[ep.topic]
        return next
      })
    }
  }

  // Downloading the whole book means nine separate saves, so it runs one at a time and
  // can be stopped part-way. Anything already on disk is skipped rather than refetched.
  async function downloadAll() {
    setConfirming(false)
    setCancelBulk(false)
    stopRequested.current = false
    const pending = episodes.filter((e) => !downloads[e.topic])
    setBulk({ done: 0, total: pending.length })
    for (const [i, ep] of pending.entries()) {
      if (stopRequested.current) break
      await download(ep)
      setBulk({ done: i + 1, total: pending.length })
      // A breath between saves: browsers throttle a burst of downloads, and some ask
      // permission once before allowing the rest.
      await new Promise((r) => setTimeout(r, 400))
    }
    setBulk(null)
    setCancelBulk(false)
  }

  const totalListened = Object.values(listened).reduce((a, b) => a + b, 0)
  const notSaved = episodes.filter((e) => !downloads[e.topic])
  const notSavedBytes = notSaved.reduce((sum, e) => sum + e.bytes, 0)

  if (episodes.length === 0) {
    return (
      <div>
        <h1 className="text-xl">Audio</h1>
        <p className="mt-2 max-w-[58ch] text-[0.95rem] leading-relaxed text-ink-soft">
          No audiobook has been rendered yet. Run <code>tools/audiobook.py</code> and{' '}
          <code>tools/synth.py</code>, then point the server at the output with <code>-audio</code>.
        </p>
      </div>
    )
  }

  return (
    <div>
      <h1 className="text-xl">The lectures, read aloud</h1>
      <p className="mt-1 max-w-[62ch] text-[0.92rem] leading-relaxed text-ink-soft">
        All {episodes.length} lectures spoken, {Math.round(minutes)} minutes in total. The questions in
        the notes are read out too, with a pause to answer before you hear why. Open the notes while an
        episode plays and it keeps playing.
      </p>
      <p className="mt-2 text-[0.85rem] text-ink-soft">
        You have listened for <strong>{clock(totalListened)}</strong>
        {totalListened > 0 && ` · ${Math.round((totalListened / (minutes * 60)) * 100)}% of the book`}
      </p>

      <div className="mt-5 flex flex-wrap items-center gap-3">
        {notSaved.length > 0 ? (
          <button onClick={() => setConfirming(true)} disabled={!!bulk} className="btn btn-quiet">
            {bulk ? `Saving ${bulk.done} of ${bulk.total}…` : `Download all (${megabytes(notSavedBytes)})`}
          </button>
        ) : (
          <span className="text-[0.85rem]" style={{ color: 'var(--slack)' }}>
            Every episode is saved on this device.
          </span>
        )}
        {bulk && (
          <button
            onClick={() => {
              setCancelBulk(true)
              stopRequested.current = true
            }}
            className="text-[0.82rem] text-ink-soft underline underline-offset-4"
          >
            {cancelBulk ? 'Stopping…' : 'Stop after this one'}
          </button>
        )}
      </div>

      <ul className="mt-4 space-y-1.5">
        {episodes.map((ep) => {
          const heard = listened[ep.topic] ?? ep.listenedSec
          const done = heard / ep.seconds
          const isCurrent = current?.topic === ep.topic
          const progress = saving[ep.topic]
          const saved = downloads[ep.topic]
          return (
            <li key={ep.topic} className="border border-line bg-surface">
              <div className="flex items-start gap-3 p-3">
                <button
                  onClick={() => play(ep)}
                  aria-label={isCurrent && playing ? `Pause ${ep.title}` : `Play ${ep.title}`}
                  className="mt-0.5 flex h-9 w-9 shrink-0 items-center justify-center border"
                  style={{
                    borderColor: isCurrent ? 'var(--ink)' : 'var(--line-strong)',
                    background: isCurrent ? 'var(--ink)' : 'transparent',
                    color: isCurrent ? 'var(--paper)' : 'var(--ink)',
                  }}
                >
                  {isCurrent && playing ? <PauseIcon /> : <PlayIcon />}
                </button>

                <div className="min-w-0 flex-1">
                  <div className="flex items-baseline gap-2">
                    <span className="text-[0.72rem] text-ink-faint">{ep.index}</span>
                    <span className="truncate text-[0.95rem]">{ep.title}</span>
                  </div>
                  <div className="mt-0.5 flex flex-wrap items-center gap-x-2 text-[0.76rem] text-ink-faint">
                    <span>{clock(ep.seconds)}</span>
                    <span>·</span>
                    <span>{megabytes(ep.bytes)}</span>
                    {heard > 0 && (
                      <>
                        <span>·</span>
                        <span style={{ color: done >= 0.95 ? 'var(--slack)' : undefined }}>
                          {done >= 0.95 ? 'finished' : `${clock(heard)} listened`}
                        </span>
                      </>
                    )}
                    {saved && (
                      <>
                        <span>·</span>
                        <span style={{ color: 'var(--slack)' }}>saved</span>
                      </>
                    )}
                    <span>·</span>
                    <Link to={`/read/${ep.topic}`} className="underline underline-offset-2">
                      read
                    </Link>
                  </div>
                  {heard > 0 && (
                    <div className="mt-2 h-1.5 bg-sunken" aria-hidden="true">
                      <div
                        className="h-full"
                        style={{
                          width: `${Math.min(100, done * 100)}%`,
                          background: done >= 0.95 ? 'var(--slack)' : 'var(--warn)',
                        }}
                      />
                    </div>
                  )}
                </div>

                {progress !== undefined ? (
                  <span className="w-24 shrink-0 self-center text-right text-[0.78rem] text-ink-soft">
                    {Math.round(progress * 100)}%
                  </span>
                ) : saved ? (
                  <button
                    onClick={() => setDownloads(forgetDownload(ep.topic))}
                    className="action-go shrink-0 self-center"
                    title="Forget that this was downloaded"
                  >
                    Saved
                  </button>
                ) : (
                  <button onClick={() => download(ep)} className="action-go shrink-0 self-center">
                    Download
                  </button>
                )}
              </div>
            </li>
          )
        })}
      </ul>

      {confirming && (
        <ConfirmDialog
          title="Download the whole audiobook?"
          confirmLabel={`Download ${notSaved.length}`}
          onCancel={() => setConfirming(false)}
          onConfirm={downloadAll}
          body={
            <>
              <p>
                {notSaved.length} {notSaved.length === 1 ? 'episode' : 'episodes'},{' '}
                <strong>{megabytes(notSavedBytes)}</strong> in total, saved one at a time to this
                device. Worth doing on wifi rather than mobile data.
              </p>
              {episodes.length - notSaved.length > 0 && (
                <p className="mt-2">
                  The {episodes.length - notSaved.length} you already have will be skipped.
                </p>
              )}
              <p className="mt-2">
                Your browser may ask once whether to allow multiple downloads. Say yes, or only the
                first will be saved.
              </p>
            </>
          }
        />
      )}

      {error && <p className="mt-4 text-[0.85rem] text-critical">{error}</p>}
    </div>
  )
}

function PlayIcon() {
  return (
    <svg width="14" height="14" viewBox="0 0 14 14" fill="currentColor" aria-hidden="true">
      <path d="M2 1l11 6-11 6z" />
    </svg>
  )
}

function PauseIcon() {
  return (
    <svg width="14" height="14" viewBox="0 0 14 14" fill="currentColor" aria-hidden="true">
      <path d="M2 1h4v12H2zM8 1h4v12H8z" />
    </svg>
  )
}
