import { useState } from 'react'
import { NavLink, Route, Routes } from 'react-router-dom'
import NameGate from './components/NameGate'
import Dashboard from './pages/Dashboard'
import Audio from './pages/Audio'
import { AudioProvider, useAudio } from './components/AudioPlayer'
import Leaderboard from './pages/Leaderboard'
import Read from './pages/Read'
import Notes from './pages/Notes'
import QuizPicker from './pages/QuizPicker'
import Runner from './pages/Runner'
import Results from './pages/Results'
import { rememberName, storedName } from './lib/player'

const tabs = [
  { to: '/', label: 'Plan', end: true },
  { to: '/read', label: 'Read', end: false },
  { to: '/quiz', label: 'Quiz', end: false },
  { to: '/audio', label: 'Audio', end: false },
  { to: '/leaderboard', label: 'Leaderboard', end: false },
]

export default function App() {
  const [name, setName] = useState(storedName)

  if (!name) return <NameGate onName={setName} />

  // The provider sits above the router so the audio element survives navigation: that
  // is what lets "read" be an ordinary link to the notes instead of a second copy of
  // them rendered inside the Audio tab.
  return (
    <AudioProvider name={name}>
      <Shell name={name} setName={setName} />
    </AudioProvider>
  )
}

function Shell({ name, setName }: { name: string; setName: (n: string) => void }) {
  const { current } = useAudio()

  return (
    <div className="min-h-full">
      <header className="border-b border-line bg-surface">
        <div className="mx-auto flex max-w-5xl items-baseline gap-4 px-4 pt-4">
          <span className="font-semibold tracking-tight">Value Chain</span>
          <span className="text-[0.8rem] text-ink-faint">5601203 midterm</span>
          <button
            onClick={() => {
              const next = window.prompt('Leaderboard name', name)?.trim()
              if (next) {
                rememberName(next)
                setName(next)
              }
            }}
            className="ml-auto text-[0.82rem] text-ink-soft underline decoration-line underline-offset-4"
          >
            {name}
          </button>
        </div>
        {/* The tabs are drawn as lanes, which is the structural motif the whole app
            borrows from the course's own swimlane diagrams. */}
        <nav className="mx-auto flex max-w-5xl gap-1 px-4">
          {tabs.map((t) => (
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
