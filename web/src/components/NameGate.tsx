import { useState } from 'react'
import { normaliseName } from '../lib/name'
import { rememberName } from '../lib/player'

// Asked once, then never again — the name is what puts a result on the leaderboard,
// so it is worth a screen of its own rather than a field buried in a form.
export default function NameGate({ onName }: { onName: (name: string) => void }) {
  const [value, setValue] = useState('')
  const [touched, setTouched] = useState(false)
  const check = normaliseName(value)
  // Only complain once they have actually typed something wrong.
  const problem = touched && value.trim() !== '' && !check.ok ? check.reason : ''

  function submit(e: React.FormEvent) {
    e.preventDefault()
    if (!check.ok) {
      setTouched(true)
      return
    }
    rememberName(check.name)
    onName(check.name)
  }

  return (
    <div className="mx-auto max-w-md px-4 py-16">
      <h1 className="text-2xl">What should the leaderboard call you?</h1>
      <p className="mt-2 text-[0.95rem] text-ink-soft">
        Your results are published on the front page so everyone revising can see where they stand.
        This is saved in this browser, so you only type it once.
      </p>
      <form onSubmit={submit} className="mt-6 flex gap-2">
        <input
          autoFocus
          value={value}
          onChange={(e) => setValue(e.target.value)}
          onBlur={() => setTouched(true)}
          maxLength={40}
          placeholder="Your name"
          aria-label="Your name"
          aria-invalid={!!problem}
          className="flex-1 border bg-surface px-3 py-2"
          style={{ borderColor: problem ? 'var(--critical)' : 'var(--line-strong)' }}
        />
        <button type="submit" disabled={!check.ok} className="btn">
          Start
        </button>
      </form>

      <p className="mt-2 min-h-[1.25rem] text-[0.85rem]">
        {problem ? (
          <span style={{ color: 'var(--critical)' }}>{problem}</span>
        ) : check.ok && check.name !== value ? (
          // Show the tidied form, so "WILHELM" does not look like it was ignored.
          <span className="text-ink-soft">
            You will appear as <strong className="text-ink">{check.name}</strong>
          </span>
        ) : null}
      </p>

      <p className="mt-4 text-[0.82rem] leading-relaxed text-ink-faint">
        Capitals do not matter — type it however you like and you will be the same person on the
        board each time.
      </p>
    </div>
  )
}
