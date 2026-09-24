// Citations, and the rendered slides behind them.
//
// tools/materials.py renders each cited deck to one image per page and writes a
// manifest keyed by the deck names the content uses. A course without it gets a 404,
// and every citation stays plain text.

export type Deck = { title: string; label: string; pages: Record<string, string[]> }
export type Materials = Record<string, Deck>

let pending: Promise<Materials> | null = null

export function loadMaterials(): Promise<Materials> {
  pending ??= fetch('/materials/manifest.json')
    .then((r) => (r.ok ? r.json() : {}))
    .catch(() => ({}))
  return pending
}

export function pageNumbers(deck: Deck): number[] {
  return Object.keys(deck.pages).map(Number).sort((a, b) => a - b)
}

// A part of a citation: either plain text, or a run of pages in a deck we can show.
export type CitePart = { text: string; deck?: string; pages?: number[] }

// parseCitation splits a lecture reference such as
//   "lec 1, slides 32, 59; 2024 midterm p.1; infographics 3-4; Wolfram CNN lesson 12"
// into parts, one per deck. A segment naming no deck ("slides 5-6") belongs to the
// lecture's own deck, but only when that is a single deck: a composite lecture's bare
// numbers cannot be pinned to one, so they are dropped rather than shown as a page
// number nobody can look up. A segment naming a deck without slides stays as text.
export function parseCitation(text: string, materials: Materials, lectureDeck?: string): CitePart[] {
  const names = Object.keys(materials).sort((a, b) => b.length - a.length)
  const kept: CitePart[] = []
  for (const raw of text.split(';')) {
    const segment = raw.trim()
    const named = names.find((n) => segment.toLowerCase().startsWith(n.toLowerCase()))
    const bare = !named && /^(slides?|p\.|pages?)\s*\d/i.test(segment)
    const deck = named ?? (bare && lectureDeck && materials[lectureDeck] ? lectureDeck : undefined)
    const pages = deck ? readPages(named ? segment.slice(named.length) : segment, materials[deck]) : []
    if (deck && pages.length) kept.push({ text: segment, deck, pages })
    else if (!bare) kept.push({ text: segment })
  }
  return kept.flatMap((p, i) => (i === 0 ? [p] : [{ text: '; ' }, p]))
}

function readPages(rest: string, deck: Deck): number[] {
  const all = pageNumbers(deck)
  if (/final (slide|page)/i.test(rest)) return all.length ? [all[all.length - 1]] : []
  const list = rest.replace(/^[\s,]*(slides?|sheets?|pages?|p\.)?\s*/i, '')
  if (!/^\d/.test(list)) return []
  const out: number[] = []
  for (const piece of list.split(',')) {
    const m = piece.trim().match(/^(\d+)(?:\s*[-–]\s*(\d+))?$/)
    if (!m) break
    const from = Number(m[1])
    const to = m[2] ? Number(m[2]) : from
    for (let n = from; n <= to && n - from < 60; n++) out.push(n)
  }
  return out.filter((n) => deck.pages[String(n)])
}
