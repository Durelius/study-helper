// The name is asked for once and then reused everywhere, because typing it before
// every quiz is friction that stops people quizzing.

import { normaliseName } from './name'

const KEY = 'vc_player'

export function storedName(): string {
  try {
    const raw = localStorage.getItem(KEY) ?? ''
    if (!raw) return ''
    // A name saved before the rules existed may not be canonical; tidy it on the way
    // out so the board and the header agree.
    const check = normaliseName(raw)
    return check.ok ? check.name : ''
  } catch {
    // Private windows and blocked site data throw here; a nameless session still works,
    // it just asks again next time.
    return ''
  }
}

export function rememberName(name: string) {
  try {
    localStorage.setItem(KEY, name.trim())
  } catch {
    /* nothing to do: the name lives in memory for this visit */
  }
}
