import { useCallback, useEffect, useRef, useState } from 'react'
import { Link } from 'react-router-dom'
import { api, type Episode } from '../lib/api'
import { clock, megabytes } from '../lib/format'
import { forgetDownload, loadDownloads, rememberDownload, type DownloadRecord } from '../lib/downloads'

// How often playing time is reported. Short enough that closing the tab loses almost
// nothing, long enough not to write to the database every second.
const REPORT_EVERY = 15

export default function AudioPage({ name }: { name: string }) {
  const [episodes, setEpisodes] = useState<Episode[]>([])
  const [minutes, setMinutes] = useState(0)
  const [current, setCurrent] = useState<Episode | null>(null)
  const [playing, setPlaying] = useState(false)
  const [position, setPosition] = useState(0)
  const [rate, setRate] = useState(1)
  const [downloads, setDownloads] = useState<Record<string, DownloadRecord>>(loadDownloads)
  const [saving, setSaving] = useState<Record<string, number>>({})
  const [error, setError] = useState('')

  const audio = useRef<HTMLAudioElement | null>(null)
  // Listening time is accumulated from playback, not from wall-clock time, so a paused
  // tab or a scrub through the episode adds nothing.
  const lastTime = useRef(0)
  const unreported = useRef(0)
  const listened = useRef<Record<string, number>>({})

  useEffect(() => {
    api
      .episodes(name)
      .then((d) => {
        setEpisodes(d.episodes)
        setMinutes(d.minutes)
        for (const e of d.episodes) listened.current[e.topic] = e.listenedSec
      })
      .catch((e: Error) => setError(e.message))
  }, [name])

  const report = useCallback(
    (episode: string, pos: number, delta: number) => {
      if (delta <= 0 && pos <= 0) return
      void api.listen(name, episode, pos, delta)
    },
    [name],
  )

  // Anything still unreported when the tab closes would otherwise be lost.
  useEffect(() => {
    function flush() {
      if (current && unreported.current > 0) {
        report(current.topic, position, unreported.current)
        unreported.current = 0
      }
    }
    window.addEventListener('pagehide', flush)
    return () => {
      flush()
      window.removeEventListener('pagehide', flush)
    }
  }, [current, position, report])

  function choose(ep: Episode) {
    if (current?.topic === ep.topic) {
      toggle()
      return
    }
    if (current && unreported.current > 0) {
      report(current.topic, position, unreported.current)
      unreported.current = 0
    }
    setCurrent(ep)
    setPosition(ep.positionSec)
    lastTime.current = ep.positionSec
    setPlaying(true)
    // The element needs the new source before it can be told where to start.
    requestAnimationFrame(() => {
      const el = audio.current
      if (!el) return
      el.playbackRate = rate
      // Resume where they stopped, unless they were basically at the end.
      if (ep.positionSec > 5 && ep.positionSec < ep.seconds - 10) el.currentTime = ep.positionSec
      void el.play().catch(() => setPlaying(false))
    })
  }

  function toggle() {
    const el = audio.current
    if (!el) return
    if (el.paused) void el.play().catch(() => setPlaying(false))
    else el.pause()
  }

  function onTimeUpdate() {
    const el = audio.current
    if (!el || !current) return
    const now = el.currentTime
    const delta = now - lastTime.current
    lastTime.current = now
    setPosition(now)
    // Normal playback advances by well under a second per tick. A larger jump is a
    // seek, and a negative one is a rewind; neither is time spent listening.
    if (delta > 0 && delta < 2) {
      unreported.current += delta
      listened.current[current.topic] = (listened.current[current.topic] ?? 0) + delta
      if (unreported.current >= REPORT_EVERY) {
        report(current.topic, now, unreported.current)
        unreported.current = 0
      }
    }
  }

  function skip(seconds: number) {
    const el = audio.current
    if (!el) return
    el.currentTime = Math.max(0, Math.min(el.duration || 0, el.currentTime + seconds))
    lastTime.current = el.currentTime
  }

  function changeRate(next: number) {
    setRate(next)
    if (audio.current) audio.current.playbackRate = next
  }

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

  const totalListened = Object.values(listened.current).reduce((a, b) => a + b, 0)

  if (error && episodes.length === 0) return <p className="text-critical">{error}</p>

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
    <div className="pb-28">
      <h1 className="text-xl">The lectures, read aloud</h1>
      <p className="mt-1 max-w-[62ch] text-[0.92rem] leading-relaxed text-ink-soft">
        All nine lectures spoken, {Math.round(minutes)} minutes in total. The questions in the notes are
        read out too, with a pause to answer before you hear why. Download an episode and it plays with
        no signal.
      </p>
      <p className="mt-2 text-[0.85rem] text-ink-soft">
        You have listened for <strong>{clock(totalListened)}</strong>
        {totalListened > 0 && ` · ${Math.round((totalListened / (minutes * 60)) * 100)}% of the book`}
      </p>

      <ul className="mt-6 space-y-1.5">
        {episodes.map((ep) => {
          const heard = listened.current[ep.topic] ?? ep.listenedSec
          const done = heard / ep.seconds
          const isCurrent = current?.topic === ep.topic
          const progress = saving[ep.topic]
          const saved = downloads[ep.topic]
          return (
            <li key={ep.topic} className="border border-line bg-surface">
              <div className="flex items-start gap-3 p-3">
                <button
                  onClick={() => choose(ep)}
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
                      notes
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

      {error && <p className="mt-4 text-[0.85rem] text-critical">{error}</p>}

      {current && (
        <div className="fixed inset-x-0 bottom-0 border-t border-line-strong bg-surface">
          <div className="mx-auto max-w-5xl px-4 py-2.5">
            <div className="flex items-center gap-3">
              <button onClick={toggle} className="btn h-10 w-10 shrink-0 !px-0" aria-label={playing ? 'Pause' : 'Play'}>
                {playing ? <PauseIcon /> : <PlayIcon />}
              </button>
              <button onClick={() => skip(-15)} className="btn btn-quiet shrink-0 !px-2.5 text-[0.78rem]">
                −15s
              </button>
              <button onClick={() => skip(30)} className="btn btn-quiet shrink-0 !px-2.5 text-[0.78rem]">
                +30s
              </button>

              <div className="min-w-0 flex-1">
                <div className="flex items-baseline justify-between gap-2 text-[0.76rem] text-ink-faint">
                  <span className="truncate">{current.title}</span>
                  <span className="shrink-0">
                    {clock(position)} / {clock(current.seconds)}
                  </span>
                </div>
                <input
                  type="range"
                  min={0}
                  max={current.seconds}
                  step={1}
                  value={position}
                  onChange={(e) => {
                    const el = audio.current
                    if (!el) return
                    el.currentTime = Number(e.target.value)
                    lastTime.current = el.currentTime
                    setPosition(el.currentTime)
                  }}
                  aria-label="Seek"
                  className="mt-1 w-full accent-[var(--ink)]"
                />
              </div>

              <select
                value={rate}
                onChange={(e) => changeRate(Number(e.target.value))}
                aria-label="Playback speed"
                className="shrink-0 border border-line-strong bg-surface px-1.5 py-1 text-[0.78rem]"
              >
                {[0.8, 1, 1.25, 1.5, 1.75, 2].map((r) => (
                  <option key={r} value={r}>
                    {r}×
                  </option>
                ))}
              </select>
            </div>
          </div>
        </div>
      )}

      <audio
        ref={audio}
        src={current?.url}
        onTimeUpdate={onTimeUpdate}
        onPlay={() => setPlaying(true)}
        onPause={() => {
          setPlaying(false)
          if (current && unreported.current > 0) {
            report(current.topic, position, unreported.current)
            unreported.current = 0
          }
        }}
        onEnded={() => {
          setPlaying(false)
          if (current) report(current.topic, current.seconds, unreported.current)
          unreported.current = 0
          // Roll straight into the next lecture, the way a podcast app would.
          const next = episodes.find((e) => e.index === (current?.index ?? 0) + 1)
          if (next) choose(next)
        }}
        preload="metadata"
      />
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
