import type { Question } from './api'

// toggleChoice is how a click changes the current selection. A multi-select question
// collects picks until the answer is submitted; every other kind holds exactly one.
// The quiz runner and the inline checks on the reading share it so they cannot drift.
export function toggleChoice(type: Question['type'], prev: number[], i: number): number[] {
  if (type !== 'multi') return [i]
  return prev.includes(i) ? prev.filter((x) => x !== i) : [...prev, i]
}
