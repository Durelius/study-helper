# Value Chain midterm trainer

A study tool for *5601203 Value Chain Excellence and Project Management*: the nine
lectures condensed into something readable, and 337 questions that make you retrieve
them instead of re-reading them.

One Go binary serves the React app, the API and the leaderboard.

## Why it is built this way

Two days of revision is not enough time to read 420 slides twice, and re-reading is the
weakest thing you can do with the time anyway. So the app is built around retrieval:

- **Grading happens on the server.** The browser never receives an answer key until it
  has submitted an answer. Without that the leaderboard means nothing and neither does
  your own accuracy.
- **Every answer comes back with the reasoning and the slide number**, so a miss sends
  you to the exact page of the exact deck rather than leaving you to hunt.
- **You mark each answer *sure* or *not sure*.** Being confidently wrong is the state
  that actually costs marks, so Weak Spots replays those first.
- **Questions are interleaved across topics** rather than blocked by lecture. Blocked
  practice feels easier and measurably teaches less.
- **The critical path drill is generated**, not authored. You cannot recognise your way
  past a network you have never seen, which is the point — exam question 1 is a critical
  path problem.
- **The dashboard tells you what to do next** and why, based on what you have not
  touched, what you keep getting wrong, and how many hours are left.

## The material

Everything under `content/` is authored from the lecture PDFs and traceable to a slide.
`content/SCHEMA.md` is the format and the quality bar.

| Topic | Questions | Source deck |
|---|---|---|
| Value Chain Fundamentals | 32 | new 01 + old 01 |
| Operations Management | 38 | new 02 |
| Process Mapping & BPMN | 35 | old 02 |
| Supply Chain Management | 28 | new 03 |
| Supply Chain Strategy & Global Logistics | 20 | old 03 |
| PM Foundations & PMBOK | 54 | new 04 + old 04 |
| Project Initiation & Scope | 35 | new 05 |
| Schedule & Cost Planning | 45 | new 06 |
| Risk, Communication & Quality | 50 | new 07 |

Plus four case studies (32 questions) and a 139-term glossary.

The server refuses to start if any answer index is out of range, any question has an
unknown type, or any id is duplicated — an authoring mistake is a boot failure, not a
wrong mark found halfway through someone's revision.

**Earned value is deliberately absent.** Deck 06's cost section stops at the baseline,
so PV/EV/AC/CPI/SPI/EAC are not examinable from this syllabus and drilling them would
waste the time.

## Running it

```sh
cd web && npm install && npm run build   # writes into internal/web/dist
cd .. && go build ./cmd/server && ./server
```

Then open http://localhost:8080.

### Development

```sh
go run ./cmd/server -dev -content ./content   # API only, on :8080, reading content from disk
cd web && npm run dev                         # UI on :5173, proxying /api to :8080
```

`-content ./content` reads the material from disk instead of the embedded copy, so
editing a question file only needs a server restart rather than a rebuild.

Other flags: `-db` (SQLite path, default `study.db`), `-exam` (the countdown target,
default `2026-09-23T09:00:00+07:00`), `-addr`, `-behind-proxy`.

## Tests

```sh
go test ./...
```

The critical-path solver is checked against the worked example on slide 30 of deck 06 —
the one whose boxes the deck leaves blank — plus a hand-computed fixture, and the
generator is checked over 200 seeds for a unique critical path and internally consistent
float. If the solver and the slides ever disagree, the tests fail rather than the app
quietly teaching the wrong method.

## Deploying

`deploy/INSTALL.md` has the one-time setup: DNS, the systemd unit, the Apache vhost and
certbot. After that, `./deploy/deploy.sh` runs the tests, builds the frontend,
cross-compiles for linux/amd64, ships the binary and restarts the service.

In production the app runs on `127.0.0.1:8093` behind Apache, with the database in
`/var/lib/chulavaluechain/study.db`. The binary is replaced on every deploy; the
database is not, so the leaderboard survives.
