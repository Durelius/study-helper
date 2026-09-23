# GenAI Literacy — course skeleton

Nothing here yet but the shape. Fill it in and the app serves it; no Go changes.

## What to edit

**`course.json`** — the only file that needs thought.

- `topics` — replace the three placeholders with the real lectures, in teaching order.
  `id` becomes the filename under `content/`, so keep it short and kebab-case.
- `weight` — each topic's share of an exam simulation. Write them however suits: the
  marks each topic carries on the real paper, percentages, or all 1s for an even
  spread. They are normalised on load.
- `exam` — the countdown target, RFC 3339 with the offset (`+07:00` for Bangkok).
- `generators` — computed drills. `"cpm"` gives the critical-path lab; leave empty
  unless the course teaches it.
- `modes` — drills specific to this syllabus. Omit until the question bank exists:
  a mode with nothing behind it is greyed out, and the generic drills (Quick 10, exam
  simulation, weak spots, topic drill, glossary, case study, spot-the-error) are
  always there.

## Where the material goes

```
content/
  lectures/<topic-id>.md      notes, front matter as in SCHEMA.md
  questions/<topic-id>.json   the question bank for that topic
  visuals/*.json              diagram questions (click the error)
  cases/*.json                case studies
  glossary.json
audio/                        rendered audiobook + manifest.json (not in git)
```

`courses/valuechain/content/SCHEMA.md` is the authoring guide and applies unchanged.

## Running it

```sh
go run ./cmd/server -course courses/genai-literacy -addr :8081
```

The tests in `internal/content` check the valuechain material. Point them at this
course too once it has questions.
