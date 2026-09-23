// The server is the authority on names — see internal/store/name.go, which these rules
// mirror. Doing it here too means you see what the board will call you while you type,
// instead of finding out after you submit.

const ALLOWED = /^[A-Za-zÅÄÖåäö ]+$/

export type NameCheck = { ok: true; name: string } | { ok: false; reason: string }

/** Trimmed, single-spaced, each word capitalised: "  wilhelm   DURELIUS " → "Wilhelm Durelius". */
export function normaliseName(input: string): NameCheck {
  const trimmed = input.trim()
  if (!trimmed) return { ok: false, reason: 'A name is needed — it goes on the leaderboard.' }
  if (!ALLOWED.test(trimmed)) return { ok: false, reason: 'Letters and spaces only (Å, Ä and Ö are fine).' }
  const name = trimmed
    .split(/\s+/)
    .map((w) => w[0].toUpperCase() + w.slice(1).toLowerCase())
    .join(' ')
  if ([...name].length > 24) return { ok: false, reason: 'That name is too long for the board.' }
  return { ok: true, name }
}
