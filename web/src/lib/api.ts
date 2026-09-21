// The browser never holds an answer key until it has submitted an answer, so these
// types deliberately have no `answer` field on a question.

export type Source = { deck: string; slide: number }

export type Question = {
  id: string
  topic: string
  topicTitle: string
  type: 'mcq' | 'tf' | 'multi'
  stem: string
  choices: string[]
  difficulty: number
  examFocus: boolean
  data?: unknown
}

export type Topic = {
  id: string
  title: string
  questions: number
  hasNotes: boolean
  seen: number
  correct: number
  lastSeen: number
}

export type Mode = {
  id: string
  title: string
  blurb: string
  why: string
  count: number
  timeLimitSec: number
  needsTopic: boolean
  available: number
}

export type Session = {
  id: string
  player: string
  mode: string
  topic?: string
  total: number
  score: number
  durationMs: number
  startedAt: number
  finishedAt?: number
}

export type Verdict = {
  correct: boolean
  answer: number[]
  explanation: string
  source: Source
}

export type Review = { question: Question; answer: number[]; explanation: string; source: Source }

export type Step = { action: 'read' | 'quiz'; mode?: string; topic?: string; label: string; why: string }

export type Plan = {
  exam: string
  secondsRemaining: number
  steps: Step[]
  seen: number
  bank: number
  accuracy: number
  coverage: number
  confidentlyWrong: number
}

export type Miss = {
  questionId: string
  topic: string
  wrong: number
  attempts: number
  confident: number
  lastSeen: number
  question: Question
  explanation: string
  source: Source
}

export type Notes = {
  id: string
  title: string
  deck: string
  slides: number
  summary: string
  body: string
  checks: Record<string, Question>
}

// Activity is one box in a generated critical-path network, carrying the six figures
// the slides' box notation shows.
export type Activity = {
  name: string
  duration: number
  predecessors: string[]
  es: number
  ef: number
  ls: number
  lf: number
  float: number
  critical: boolean
  layer: number
}

export type Network = { activities: Activity[]; duration: number; criticalPath: string[] }

async function request<T>(path: string, init?: RequestInit): Promise<T> {
  const res = await fetch(path, {
    ...init,
    headers: init?.body ? { 'Content-Type': 'application/json' } : undefined,
  })
  if (!res.ok) {
    const body = (await res.json().catch(() => null)) as { error?: string } | null
    throw new Error(body?.error ?? `Request failed (${res.status})`)
  }
  return res.json() as Promise<T>
}

export const api = {
  topics: (player: string) => request<Topic[]>(`/api/topics?player=${encodeURIComponent(player)}`),
  modes: () => request<Mode[]>('/api/modes'),
  notes: (topic: string) => request<Notes>(`/api/notes/${topic}`),
  leaderboard: (mode = '') =>
    request<{ top: Session[]; recent: Session[] }>(`/api/leaderboard?mode=${encodeURIComponent(mode)}`),
  plan: (player: string) => request<Plan>(`/api/plan?player=${encodeURIComponent(player)}`),
  stats: (player: string) =>
    request<{ topics: Topic[]; misses: Miss[] }>(`/api/stats?player=${encodeURIComponent(player)}`),
  start: (player: string, mode: string, topic = '', count = 0) =>
    request<{ id: string; player: string; mode: Mode; questions: Question[] }>('/api/quiz', {
      method: 'POST',
      body: JSON.stringify({ player, mode, topic, count }),
    }),
  answer: (id: string, questionId: string, given: number[], confident: boolean, ms: number) =>
    request<Verdict>(`/api/quiz/${id}/answer`, {
      method: 'POST',
      body: JSON.stringify({ questionId, given, confident, ms }),
    }),
  finish: (id: string) =>
    request<{ session: Session; review: Review[]; breakdown: { topic: string; title: string; seen: number; correct: number }[] }>(
      `/api/quiz/${id}/finish`,
      { method: 'POST', body: '{}' },
    ),
}
