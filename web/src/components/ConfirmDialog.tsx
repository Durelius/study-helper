import { useEffect, useRef } from 'react'

// A confirm step for an action worth thinking about for a second — downloading the
// whole audiobook onto a phone, for instance. It states the cost in the words that
// matter (how many files, how many megabytes) rather than asking "are you sure?".
export default function ConfirmDialog({
  title,
  body,
  confirmLabel,
  onConfirm,
  onCancel,
}: {
  title: string
  body: React.ReactNode
  confirmLabel: string
  onConfirm: () => void
  onCancel: () => void
}) {
  const confirm = useRef<HTMLButtonElement | null>(null)
  const panel = useRef<HTMLDivElement | null>(null)

  useEffect(() => {
    confirm.current?.focus()
    function onKey(e: KeyboardEvent) {
      if (e.key === 'Escape') onCancel()
      if (e.key !== 'Tab') return
      // Keep focus inside the dialog: tabbing out of a modal and pressing enter on
      // something behind it is how people cancel things they meant to confirm.
      const focusable = panel.current?.querySelectorAll<HTMLElement>('button')
      if (!focusable || focusable.length === 0) return
      const first = focusable[0]
      const last = focusable[focusable.length - 1]
      if (e.shiftKey && document.activeElement === first) {
        e.preventDefault()
        last.focus()
      } else if (!e.shiftKey && document.activeElement === last) {
        e.preventDefault()
        first.focus()
      }
    }
    window.addEventListener('keydown', onKey)
    const previous = document.body.style.overflow
    document.body.style.overflow = 'hidden'
    return () => {
      window.removeEventListener('keydown', onKey)
      document.body.style.overflow = previous
    }
  }, [onCancel])

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4" role="dialog" aria-modal="true" aria-label={title}>
      <button className="absolute inset-0 bg-[rgba(0,0,0,0.55)]" tabIndex={-1} aria-label="Cancel" onClick={onCancel} />
      <div ref={panel} className="relative w-full max-w-md border border-line-strong bg-surface shadow-2xl">
        <h2 className="border-b border-line px-4 py-2.5 text-[0.95rem] font-medium">{title}</h2>
        <div className="px-4 py-3.5 text-[0.92rem] leading-relaxed text-ink-soft">{body}</div>
        <div className="flex justify-end gap-2 border-t border-line px-4 py-3">
          <button onClick={onCancel} className="btn btn-quiet">
            Cancel
          </button>
          <button ref={confirm} onClick={onConfirm} className="btn">
            {confirmLabel}
          </button>
        </div>
      </div>
    </div>
  )
}
