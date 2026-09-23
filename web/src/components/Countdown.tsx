import { countdown } from '../lib/format'

// The hero is the one thing this app is really about: how much float is left before the
// exam. It is drawn as a schedule bar running into a milestone diamond, which is
// how the course itself draws a deadline — and the bar turns critical as the slack
// runs out, for the same reason an activity does.
const WINDOW_HOURS = 72

export default function Countdown({ seconds, exam }: { seconds: number; exam?: string }) {
  // The day and time come from the course, not from whichever course was built first.
  const when = exam
    ? new Date(exam).toLocaleString(undefined, { weekday: 'long', hour: '2-digit', minute: '2-digit', hour12: false })
    : ''

  const { hours, minutes, past } = countdown(seconds)
  const remaining = Math.max(0, Math.min(1, seconds / (WINDOW_HOURS * 3600)))
  const spent = 1 - remaining
  const critical = seconds < 12 * 3600

  return (
    <section aria-label="Time until the midterm">
      <p className="text-[0.8rem] text-ink-soft">{when ? `Midterm · ${when}` : 'Midterm'}</p>
      <p
        className="mt-1 font-semibold tracking-tight"
        style={{ fontSize: 'clamp(2.4rem, 7vw, 4rem)', lineHeight: 1, color: critical ? 'var(--critical)' : 'var(--ink)' }}
      >
        {past ? 'Good luck' : `${hours}h ${String(minutes).padStart(2, '0')}m`}
      </p>
      <p className="mt-1 text-[0.95rem] text-ink-soft">{past ? 'the paper is now' : 'of float remaining'}</p>

      <div className="mt-5 flex items-center gap-2" aria-hidden="true">
        <div className="relative h-6 flex-1 border border-line-strong bg-sunken">
          <div
            className="absolute inset-y-0 left-0"
            style={{
              width: `${spent * 100}%`,
              background: critical ? 'var(--critical)' : 'var(--slack)',
              transition: 'width 1s linear',
            }}
          />
        </div>
        <svg width="18" height="18" viewBox="0 0 18 18" className="shrink-0">
          <path d="M9 1 L17 9 L9 17 L1 9 Z" fill={critical ? 'var(--critical)' : 'var(--ink)'} />
        </svg>
      </div>
      <div className="mt-1 flex justify-between text-[0.72rem] text-ink-faint">
        <span>72 hours out</span>
        <span>exam</span>
      </div>
    </section>
  )
}
