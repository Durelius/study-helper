import { useState } from 'react'
import { rememberName } from '../lib/player'

// Asked once, then never again — the name is what puts a result on the leaderboard,
// so it is worth a screen of its own rather than a field buried in a form.
export default function NameGate({ onName }: { onName: (name: string) => void }) {
  const [value, setValue] = useState('')

  function submit(e: React.FormEvent) {
    e.preventDefault()
    const name = value.trim()
    if (!name) return
    rememberName(name)
    onName(name)
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
          maxLength={24}
          placeholder="Your name"
          aria-label="Your name"
          className="flex-1 border border-line-strong bg-surface px-3 py-2"
        />
        <button
          type="submit"
          disabled={!value.trim()}
          className="btn"
        >
          Start
        </button>
      </form>
    </div>
  )
}
