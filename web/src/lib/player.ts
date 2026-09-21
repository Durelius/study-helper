// The name is asked for once and then reused everywhere, because typing it before
// every quiz is friction that stops people quizzing.

const KEY = 'vc_player'

export function storedName(): string {
  try {
    return localStorage.getItem(KEY) ?? ''
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
