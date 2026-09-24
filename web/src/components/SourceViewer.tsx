import { createContext, useCallback, useContext, useEffect, useRef, useState } from 'react'
import type { Source } from '../lib/api'
import { loadMaterials, pageNumbers, parseCitation, type Materials } from '../lib/sources'

// Every citation in the app — under a quiz answer, in the results, at the end of a
// section of notes — opens the slide it came from. The point of a slide number is to go
// and read the slide, and a modal does that without losing your place in the quiz.

type Open = (deck: string, cited: number[]) => void
const Ctx = createContext<{ materials: Materials; open: Open }>({ materials: {}, open: () => {} })

export function SourceProvider({ children }: { children: React.ReactNode }) {
  const [materials, setMaterials] = useState<Materials>({})
  const [view, setView] = useState<{ deck: string; cited: number[]; page: number } | null>(null)
  useEffect(() => {
    loadMaterials().then(setMaterials)
  }, [])
  const open = useCallback<Open>((deck, cited) => setView({ deck, cited, page: cited[0] }), [])

  return (
    <Ctx.Provider value={{ materials, open }}>
      {children}
      {view && materials[view.deck] && (
        <Viewer
          materials={materials}
          {...view}
          onPage={(page) => setView({ ...view, page })}
          onClose={() => setView(null)}
        />
      )}
    </Ctx.Provider>
  )
}

// SourceLink is a question's single {deck, slide} source.
export function SourceLink({ source, className }: { source: Source; className?: string }) {
  const { materials, open } = useContext(Ctx)
  const deck = materials[source.deck]
  const label = `${source.deck}, ${deck?.label ?? 'slide'} ${source.slide}`
  if (!deck?.pages[String(source.slide)]) return <span className={className}>{label}</span>
  return (
    <button onClick={() => open(source.deck, [source.slide])} className={`cite ${className ?? ''}`}>
      {label}
    </button>
  )
}

// Citation is a lecture reference like `[lec 1, slides 32, 59; infographics 3]`, with
// each deck in it opening separately.
export function Citation({ text, lectureDeck }: { text: string; lectureDeck?: string }) {
  const { materials, open } = useContext(Ctx)
  // A course with no rendered slides keeps its references exactly as written.
  if (!Object.keys(materials).length) return <code>[{text}]</code>
  const parts = parseCitation(text, materials, lectureDeck)
  if (!parts.length) return null
  return (
    <span className="text-[0.8rem] text-ink-faint">
      [
      {parts.map((p, i) =>
        p.deck && p.pages ? (
          <button key={i} onClick={() => open(p.deck!, p.pages!)} className="cite">
            {p.text}
          </button>
        ) : (
          <span key={i}>{p.text}</span>
        ),
      )}
      ]
    </span>
  )
}

function Viewer({
  materials,
  deck,
  cited,
  page,
  onPage,
  onClose,
}: {
  materials: Materials
  deck: string
  cited: number[]
  page: number
  onPage: (n: number) => void
  onClose: () => void
}) {
  const d = materials[deck]
  const all = pageNumbers(d)
  const at = all.indexOf(page)
  const prev = at > 0 ? all[at - 1] : null
  const next = at < all.length - 1 ? all[at + 1] : null
  // Fit shows the whole slide at once; full width is for the dense infographics,
  // which are taller than any screen and need to be scrolled and read.
  const [fit, setFit] = useState(d.label !== 'sheet')
  const close = useRef<HTMLButtonElement | null>(null)
  const body = useRef<HTMLDivElement | null>(null)

  useEffect(() => {
    close.current?.focus()
    const previous = document.body.style.overflow
    document.body.style.overflow = 'hidden'
    return () => {
      document.body.style.overflow = previous
    }
  }, [])

  useEffect(() => {
    body.current?.scrollTo(0, 0)
    function onKey(e: KeyboardEvent) {
      if (e.key === 'Escape') onClose()
      if (e.key === 'ArrowLeft' && prev !== null) onPage(prev)
      if (e.key === 'ArrowRight' && next !== null) onPage(next)
    }
    window.addEventListener('keydown', onKey)
    return () => window.removeEventListener('keydown', onKey)
  }, [prev, next, onPage, onClose])

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-2 sm:p-4" role="dialog" aria-modal="true" aria-label={d.title}>
      <button className="absolute inset-0 bg-[rgba(0,0,0,0.7)]" tabIndex={-1} aria-label="Close" onClick={onClose} />
      <div className="relative flex max-h-full w-full max-w-6xl flex-col border border-line-strong bg-surface shadow-2xl">
        <header className="flex flex-wrap items-center gap-x-3 gap-y-1.5 border-b border-line px-3 py-2">
          <span className="text-[0.92rem] font-medium">{d.title}</span>
          <span className="text-[0.8rem] text-ink-faint">
            {d.label} {page}
            {all.length > 1 && ` of ${all[all.length - 1]}`}
          </span>
          {cited.length > 1 && (
            <span className="flex flex-wrap gap-1">
              {cited.map((n) => (
                <button
                  key={n}
                  onClick={() => onPage(n)}
                  className="border px-1.5 text-[0.76rem]"
                  style={{
                    borderColor: n === page ? 'var(--ink)' : 'var(--line)',
                    background: n === page ? 'var(--sunken)' : 'transparent',
                  }}
                >
                  {n}
                </button>
              ))}
            </span>
          )}
          <span className="ml-auto flex gap-1.5">
            <button onClick={() => setFit(!fit)} className="btn btn-quiet px-2 py-1 text-[0.8rem]">
              {fit ? 'Full width' : 'Fit'}
            </button>
            <button ref={close} onClick={onClose} className="btn btn-quiet px-2 py-1 text-[0.8rem]" aria-label="Close">
              Close
            </button>
          </span>
        </header>

        <div ref={body} className="min-h-0 flex-1 overflow-auto bg-sunken">
          {d.pages[String(page)].map((src) => (
            <img
              key={src}
              src={`/materials/${src}`}
              alt={`${d.title}, ${d.label} ${page}`}
              onClick={() => setFit(!fit)}
              className={fit ? 'mx-auto max-h-[calc(100vh-7rem)] w-auto cursor-zoom-in' : 'w-full cursor-zoom-out'}
            />
          ))}
        </div>

        {all.length > 1 && (
          <footer className="flex items-center justify-between border-t border-line px-3 py-1.5">
            <button onClick={() => prev !== null && onPage(prev)} disabled={prev === null} className="btn btn-quiet px-2 py-1 text-[0.8rem]">
              ← {d.label} {prev ?? ''}
            </button>
            <span className="hidden text-[0.74rem] text-ink-faint sm:inline">Arrow keys turn the page · Esc closes</span>
            <button onClick={() => next !== null && onPage(next)} disabled={next === null} className="btn btn-quiet px-2 py-1 text-[0.8rem]">
              {d.label} {next ?? ''} →
            </button>
          </footer>
        )}
      </div>
    </div>
  )
}
