import type { Episode } from '../lib/api'
import { clock } from '../lib/format'

// The transport controls, drawn wherever the player happens to be: pinned to the
// bottom of the Audio tab, or along the top of the reading panel.
//
// The <audio> element itself is NOT in here. It stays mounted on the page, because
// moving it between parents would remount it and stop playback — which is precisely
// what someone reading along while listening does not want.
export default function PlayerControls({
  episode,
  playing,
  position,
  rate,
  onToggle,
  onSkip,
  onSeek,
  onRate,
  compact = false,
}: {
  episode: Episode
  playing: boolean
  position: number
  rate: number
  onToggle: () => void
  onSkip: (seconds: number) => void
  onSeek: (seconds: number) => void
  onRate: (rate: number) => void
  compact?: boolean
}) {
  return (
    <div className="flex items-center gap-2 sm:gap-3">
      <button onClick={onToggle} className="btn h-10 w-10 shrink-0 !px-0" aria-label={playing ? 'Pause' : 'Play'}>
        {playing ? <PauseIcon /> : <PlayIcon />}
      </button>
      <button onClick={() => onSkip(-15)} className="btn btn-quiet shrink-0 !px-2.5 text-[0.78rem]">
        −15s
      </button>
      <button onClick={() => onSkip(30)} className="btn btn-quiet shrink-0 !px-2.5 text-[0.78rem]">
        +30s
      </button>

      <div className="min-w-0 flex-1">
        {!compact && (
          <div className="flex items-baseline justify-between gap-2 text-[0.76rem] text-ink-faint">
            <span className="truncate">{episode.title}</span>
            <span className="shrink-0">
              {clock(position)} / {clock(episode.seconds)}
            </span>
          </div>
        )}
        <input
          type="range"
          min={0}
          max={episode.seconds}
          step={1}
          value={position}
          onChange={(e) => onSeek(Number(e.target.value))}
          aria-label="Seek"
          className={`w-full accent-[var(--ink)] ${compact ? '' : 'mt-1'}`}
        />
        {compact && (
          <div className="mt-0.5 text-right text-[0.72rem] text-ink-faint">
            {clock(position)} / {clock(episode.seconds)}
          </div>
        )}
      </div>

      <select
        value={rate}
        onChange={(e) => onRate(Number(e.target.value))}
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
  )
}

export function PlayIcon() {
  return (
    <svg width="14" height="14" viewBox="0 0 14 14" fill="currentColor" aria-hidden="true">
      <path d="M2 1l11 6-11 6z" />
    </svg>
  )
}

export function PauseIcon() {
  return (
    <svg width="14" height="14" viewBox="0 0 14 14" fill="currentColor" aria-hidden="true">
      <path d="M2 1h4v12H2zM8 1h4v12H8z" />
    </svg>
  )
}
