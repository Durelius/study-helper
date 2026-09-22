export function percent(n: number): string {
  return `${Math.round(n * 100)}%`
}

export function accuracy(correct: number, seen: number): number {
  return seen > 0 ? correct / seen : 0
}

// duration renders a stopwatch figure: 1:04, or 12:31 for a long exam simulation.
export function duration(ms: number): string {
  const total = Math.max(0, Math.round(ms / 1000))
  const m = Math.floor(total / 60)
  const s = total % 60
  return `${m}:${String(s).padStart(2, '0')}`
}

// countdown breaks the remaining float into the units people actually think in.
export function countdown(seconds: number) {
  const clamped = Math.max(0, seconds)
  return {
    hours: Math.floor(clamped / 3600),
    minutes: Math.floor((clamped % 3600) / 60),
    seconds: clamped % 60,
    past: seconds <= 0,
  }
}

export function ago(msSinceEpoch: number): string {
  if (!msSinceEpoch) return 'not yet'
  const mins = Math.round((Date.now() - msSinceEpoch) / 60000)
  if (mins < 1) return 'just now'
  if (mins < 60) return `${mins}m ago`
  const hours = Math.round(mins / 60)
  if (hours < 24) return `${hours}h ago`
  return `${Math.round(hours / 24)}d ago`
}

// clock renders a media position: 4:07, or 1:12:30 once past an hour.
export function clock(seconds: number): string {
  const total = Math.max(0, Math.round(seconds))
  const h = Math.floor(total / 3600)
  const m = Math.floor((total % 3600) / 60)
  const s = total % 60
  if (h > 0) return `${h}:${String(m).padStart(2, '0')}:${String(s).padStart(2, '0')}`
  return `${m}:${String(s).padStart(2, '0')}`
}

export function megabytes(bytes: number): string {
  return `${(bytes / 1024 / 1024).toFixed(1)} MB`
}
