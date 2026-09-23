# Content schema & style guide

Every file here is hand-checkable against a slide. If a claim cannot be traced to a
slide, it does not go in.

## Topic ids

| id | title | source deck |
|---|---|---|
| `value-chain` | Course Intro & Value Chain Fundamentals | new 01 + old 01 |
| `operations` | Operations Management | new 02 |
| `process-mapping` | Operations Design & Process Mapping (BPMN) | old 02 |
| `scm` | Supply Chain Management | new 03 |
| `logistics` | Supply Chain Strategy & Global Logistics | old 03 |
| `pm-foundations` | PM Foundations & PMBOK | new 04 + old 04 |
| `initiation` | Project Initiation & Scope Management | new 05 |
| `schedule-cost` | Project Planning: Schedule & Cost | new 06 |
| `risk-comm-quality` | Risk, Communication & Quality | new 07 |

## `lectures/<topic>.md`

```
---
id: schedule-cost
title: "Project Planning: Schedule & Cost"
deck: "new 06"
slides: 47
summary: "One sentence on what this lecture is for."
---

## <Section heading>            <!-- ~6-10 sections per lecture -->

Prose in short paragraphs. **Bold** every term the exam could ask you to define.
Tables are encouraged (GFM tables render) — the slides are full of them and a table
is easier to retrieve from than a paragraph.

> **Exam focus.** Use this callout when a slide maps onto one of the five named exam
> questions. Say what the examiner is likely to ask.

`[slide 27]` — end each section with the slide range it came from, in backticks.

?check id=sc-014                <!-- inline retrieval check: a question id from this
?check id=sc-015                     topic's question file. 2-3 per lecture, placed
                                     right after the section they test. -->
```

Write **condensed and exam-focused**, not a transcript. Diagram slides must be turned
into prose or a table — describe what the diagram *teaches*, not that it exists.
A slide that says only "Project Life Cycle" over a picture becomes the actual phases.
Aim for 900-1600 words per lecture.

## `questions/<topic>.json`

A JSON array. Question ids are `<prefix>-NNN` with a per-topic prefix
(`vc`, `ops`, `bpm`, `scm`, `log`, `pmf`, `ini`, `sc`, `rcq`).

```json
[
  {
    "id": "sc-014",
    "topic": "schedule-cost",
    "type": "mcq",
    "stem": "An activity has ES = 4 and LS = 9. What is its total float?",
    "choices": ["0 days", "4 days", "5 days", "9 days"],
    "answer": [2],
    "explanation": "Total float = LS − ES = 9 − 4 = 5. Only activities with float 0 are critical.",
    "source": { "deck": "new 06", "slide": 28 },
    "difficulty": 2,
    "tags": ["cpm", "float"],
    "examFocus": true
  }
]
```

Field rules:

- `type`: `"mcq"` (exactly one correct, 4 choices), `"tf"` (choices are exactly
  `["True", "False"]`), or `"multi"` (2-3 correct out of 5, stem must say "Select all").
- `answer`: array of **0-based** indices into `choices`. One element unless `multi`.
- `explanation`: one or two sentences saying *why*, and for a wrong-answer trap, why the
  tempting choice is wrong. This is the single most valuable field in the file — the
  reader has just got it wrong and this is what fixes the belief.
- `source.slide`: the PDF page number you read it from. Required, must be real.
- `difficulty`: 1 recall, 2 apply, 3 analyse/judge a scenario.
- `examFocus`: `true` only if it maps to one of the five named exam questions
  (critical path · risk type / knowledge area · BPM error or quality KPI ·
  case study · project definition & planning).

Quality bar:

- **Distractors must be plausible** — drawn from neighbouring concepts in the same
  deck, never filler. The classic good distractor is the adjacent term
  (effort vs duration, lead vs lag, QA vs QC, risk vs issue, verification vs validation).
- No "all of the above" / "none of the above".
- Spread the correct answer across positions; do not park it at index 1.
- Vary the stems: definition recall, scenario classification ("A supplier misses a
  delivery deadline. This is a…"), spot-the-error, and computation where the deck
  supports it.
- At least 40% `difficulty >= 2`. A bank of pure recall will not pass this exam.

## `cases/<name>.json`

```json
{
  "id": "freshmart",
  "title": "FreshMart goes omnichannel",
  "topics": ["value-chain", "operations", "pm-foundations"],
  "scenario": "300-500 words of narrative with concrete numbers and named people.",
  "questions": [ /* same question objects, ids prefixed `case-freshmart-01` */ ]
}
```

## `glossary.json`

```json
[{ "term": "Total float", "definition": "…", "topic": "schedule-cost", "source": { "deck": "new 06", "slide": 28 } }]
```
