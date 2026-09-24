import { useEffect, useState } from 'react'
import { NavLink, Route, Routes } from 'react-router-dom'
import NameGate from './components/NameGate'
import Dashboard from './pages/Dashboard'
import Audio from './pages/Audio'
import { AudioProvider, useAudio } from './components/AudioPlayer'
import { SourceProvider } from './components/SourceViewer'
import Leaderboard from './pages/Leaderboard'
import Read from './pages/Read'
import Notes from './pages/Notes'
import QuizPicker from './pages/QuizPicker'
import Runner from './pages/Runner'
import Results from './pages/Results'
import { api, type CourseInfo } from './lib/api'
import { normaliseName } from './lib/name'
import { rememberName, storedName } from './lib/player'

// Audio only appears for a course that has an audiobook rendered, so a course without
// one does not advertise an empty tab.
const tabs = [
  { to: '/', label: 'Plan', end: true, always: true },
  { to: '/read', label: 'Read', end: false, always: true },
  { to: '/quiz', label: 'Quiz', end: false, always: true },
  { to: '/audio', label: 'Audio', end: false, always: false },
  { to: '/leaderboard', label: 'Leaderboard', end: false, always: true },
]

export default function App() {
  const [name, setName] = useState(storedName)

  if (!name) return <NameGate onName={setName} />

  // The provider sits above the router so the audio element survives navigation: that
  // is what lets "read" be an ordinary link to the notes instead of a second copy of
  // them rendered inside the Audio tab.
  return (
    <AudioProvider name={name}>
      <SourceProvider>
        <Shell name={name} setName={setName} />
      </SourceProvider>
    </AudioProvider>
  )
}

function Shell({ name, setName }: { name: string; setName: (n: string) => void }) {
  const { current, episodes } = useAudio()
  // The course names itself, so one build serves every course.
  const [course, setCourse] = useState<CourseInfo | null>(null)
  useEffect(() => {
    api.course().then(setCourse).catch(() => setCourse(null))
  }, [])

  return (
    <div className="min-h-full">
      <header className="border-b border-line bg-surface">
        <div className="mx-auto flex max-w-5xl items-baseline gap-4 px-4 pt-4">
          <span className="font-semibold tracking-tight">{course?.shortName ?? '…'}</span>
          <span className="text-[0.8rem] text-ink-faint">
            {course?.code ? `${course.code} midterm` : ''}
          </span>
          <button
            onClick={() => {
              const typed = window.prompt('Leaderboard name', name)
              if (typed === null) return
              const check = normaliseName(typed)
              if (!check.ok) {
                window.alert(check.reason)
                return
              }
              rememberName(check.name)
              setName(check.name)
            }}
            className="ml-auto text-[0.82rem] text-ink-soft underline decoration-line underline-offset-4"
          >
            {name}
          </button>
        </div>
        {/* The tabs are drawn as lanes, which is the structural motif the whole app
            borrows from the course's own swimlane diagrams. */}
        <nav className="mx-auto flex max-w-5xl gap-1 px-4">
          {tabs.filter((t) => t.always || episodes.length > 0).map((t) => (
            <NavLink
              key={t.to}
              to={t.to}
              end={t.end}
              className={({ isActive }) =>
                `border-b-2 px-3 py-2.5 text-[0.9rem] ${
                  isActive ? 'border-ink font-medium text-ink' : 'border-transparent text-ink-soft'
                }`
              }
            >
              {t.label}
            </NavLink>
          ))}
        </nav>
      </header>

      {/* Room for the player pinned to the bottom, when there is one. */}
      <main className={`mx-auto max-w-5xl px-4 py-8 ${current ? 'pb-32' : ''}`}>
        <Routes>
          <Route path="/" element={<Dashboard name={name} />} />
          <Route path="/read" element={<Read name={name} />} />
          <Route path="/read/:topic" element={<Notes />} />
          <Route path="/quiz" element={<QuizPicker name={name} />} />
          <Route path="/quiz/:id" element={<Runner />} />
          <Route path="/results/:id" element={<Results />} />
          <Route path="/audio" element={<Audio />} />
          <Route path="/leaderboard" element={<Leaderboard name={name} />} />
          <Route
            path="*"
            element={
              <div className="py-16 text-center">
                <h1 className="text-xl">That page does not exist.</h1>
                <NavLink to="/" className="mt-3 inline-block underline underline-offset-4">
                  Back to the plan
                </NavLink>
              </div>
            }
          />
        </Routes>
      </main>
    </div>
  )
}
