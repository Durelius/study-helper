import { createContext, useCallback, useContext, useEffect, useRef, useState } from 'react'
import { Link } from 'react-router-dom'
import PlayerControls from './PlayerControls'
import { api, type Episode } from '../lib/api'

// The player lives above the router, so navigating between the episode list and a
// lecture's notes never unmounts the <audio> element and playback simply continues.
// That is what makes "read along" an ordinary link rather than a second copy of the
// notes page rendered inside the Audio tab.

// How often playing time is reported. Short enough that closing the tab loses almost
// nothing, long enough not to write to the database every second.
const REPORT_EVERY = 15

type AudioState = {
  episodes: Episode[]
  minutes: number
  current: Episode | null
  playing: boolean
  position: number
  /** Seconds actually listened per episode, updated live. */
  listened: Record<string, number>
  play: (episode: Episode) => void
  toggle: () => void
}

const Ctx = createContext<AudioState | null>(null)

export function useAudio(): AudioState {
  const ctx = useContext(Ctx)
  if (!ctx) throw new Error('useAudio must be used inside AudioProvider')
  return ctx
}

export function AudioProvider({ name, children }: { name: string; children: React.ReactNode }) {
  const [episodes, setEpisodes] = useState<Episode[]>([])
  const [minutes, setMinutes] = useState(0)
  const [current, setCurrent] = useState<Episode | null>(null)
  const [playing, setPlaying] = useState(false)
  const [position, setPosition] = useState(0)
  const [rate, setRate] = useState(1)
  const [listened, setListened] = useState<Record<string, number>>({})

  const audio = useRef<HTMLAudioElement | null>(null)
  // Listening time is accumulated from playback, not from wall-clock time, so a paused
  // tab or a scrub through the episode adds nothing.
  const lastTime = useRef(0)
  const unreported = useRef(0)

  useEffect(() => {
    api
      .episodes(name)
      .then((d) => {
        setEpisodes(d.episodes)
        setMinutes(d.minutes)
        setListened(Object.fromEntries(d.episodes.map((e) => [e.topic, e.listenedSec])))
      })
      .catch(() => setEpisodes([]))
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

  const play = useCallback(
    (ep: Episode) => {
      if (current?.topic === ep.topic) {
        const el = audio.current
        if (el) el.paused ? void el.play().catch(() => setPlaying(false)) : el.pause()
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
      requestAnimationFrame(() => {
        const el = audio.current
        if (!el) return
        el.playbackRate = rate
        // Resume where they stopped, unless they were basically at the end.
        if (ep.positionSec > 5 && ep.positionSec < ep.seconds - 10) el.currentTime = ep.positionSec
        void el.play().catch(() => setPlaying(false))
      })
    },
    [current, position, rate, report],
  )

  const toggle = useCallback(() => {
    const el = audio.current
    if (!el) return
    if (el.paused) void el.play().catch(() => setPlaying(false))
    else el.pause()
  }, [])

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
      setListened((l) => ({ ...l, [current.topic]: (l[current.topic] ?? 0) + delta }))
      if (unreported.current >= REPORT_EVERY) {
        report(current.topic, now, unreported.current)
        unreported.current = 0
      }
    }
  }

  function seek(seconds: number) {
    const el = audio.current
    if (!el) return
    el.currentTime = Math.max(0, Math.min(el.duration || seconds, seconds))
    lastTime.current = el.currentTime
    setPosition(el.currentTime)
  }

  return (
    <Ctx.Provider value={{ episodes, minutes, current, playing, position, listened, play, toggle }}>
      {children}

      {current && (
        <div className="fixed inset-x-0 bottom-0 z-40 border-t border-line-strong bg-surface">
          <div className="mx-auto max-w-5xl px-4 py-2.5">
            <Link
              to={`/read/${current.topic}`}
              className="mb-1.5 inline-block text-[0.76rem] text-ink-soft underline underline-offset-4"
            >
              Read along with this episode
            </Link>
            <PlayerControls
              episode={current}
              playing={playing}
              position={position}
              rate={rate}
              onToggle={toggle}
              onSkip={(s) => seek((audio.current?.currentTime ?? 0) + s)}
              onSeek={seek}
              onRate={(r) => {
                setRate(r)
                if (audio.current) audio.current.playbackRate = r
              }}
            />
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
          if (next) play(next)
        }}
        preload="metadata"
      />
    </Ctx.Provider>
  )
}
