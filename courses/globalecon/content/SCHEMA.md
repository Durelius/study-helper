# Content schema & style guide

Every file here is hand-checkable against a slide. If a claim cannot be traced to a
slide, it does not go in.

## Topic ids

All material comes from one deck, `Quiz Prep` (the 89-slide Quiz Prep Packet).

| id | title | prefix | slides |
|---|---|---|---|
| `exam-core` | Exam Core: Formula Sheet & Traps | `core` | all |
| `market-for-ideas` | Market for Ideas: Mokyr & the Enlightenment | `mfi` | 5–18 |
| `oligopoly` | Oligopoly: Grab vs Uber | `olig` | 19–42 |
| `mixed-strategy` | Mixed-Strategy Nash Equilibrium | `mix` | 43–48 |
| `solow` | Solow Model: Capital, Saving & Steady State | `sol` | 49–72, 82–83 |
| `golden-rule` | Golden Rule, Population & Technology | `gr` | 72–89 |

Tags the modes draw on: `packet` (the packet's own numbers and pop-quiz questions),
`explain` (pick the right reasoning), `formula` (recall or apply one formula), `graph`
(the figure cases).

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
  `["True", "False"]`), `"multi"` (2-3 correct out of 5, stem must say "Select all"),
  `"text"` (a typed one-word answer — see below), or `"hotspot"` (click the diagram).
- `answer`: array of **0-based** indices into `choices`. One element unless `multi`.
- `explanation`: one or two sentences saying *why*, and for a wrong-answer trap, why the
  tempting choice is wrong. This is the single most valuable field in the file — the
  reader has just got it wrong and this is what fixes the belief.
- `source.slide`: the PDF page number you read it from. Required, must be real.
- `difficulty`: 1 recall, 2 apply, 3 analyse/judge a scenario.
- `examFocus`: `true` for anything the quiz can ask from the packet; `false` only for
  quiz logistics (seat, phone, timing).

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

### Free-text questions (`type: "text"`)

For a one-word recall answer, which is the closest thing the app has to a short-answer
exam question. There are no `choices`; `accept` lists the answers that count.

```json
{
  "id": "tra-004",
  "topic": "transformers",
  "type": "text",
  "stem": "What does the T in ChatGPT stand for?",
  "choices": [],
  "answer": [],
  "accept": ["transformer"],
  "explanation": "Transformer — the architecture underneath. GPT is Generative Pre-trained Transformer: generative because it produces new text, pre-trained because it learns from a corpus before any task, transformer because attention lets it weigh every other token when choosing the next one.",
  "source": { "deck": "new 06", "slide": 12 },
  "difficulty": 1,
  "tags": ["transformer", "gpt"],
  "examFocus": true
}
```

Grading is deliberately generous, because the point is whether the term is in your head,
not whether you can spell it under time pressure. Case, punctuation, surrounding
whitespace and a leading "the"/"a" are all ignored, and the answer is accepted within an
edit distance of 20% of its length, with a floor of one edit so short words survive a
slip. Transposition counts as one edit, since "bais" for "bias" is the commonest typo
there is. So `transformer` accepts *Transformer*, *transfomer*, *transformr*,
*transformers* and *the transformer.* — but not *transducer*.

Write `accept` with every form you would give a mark for in a real exam: alternate
spellings, the expansion as well as the acronym, and the singular where you wrote the
plural. Do not rely on the fuzz to cover a genuinely different word.

Use this type for terms worth recalling cold. Do not use it where the answer is a
phrase, a number, or anything with more than one reasonable wording — use `mcq` there,
since the reader cannot guess which wording you had in mind.

### Numeric questions (`type: "num"`)

For anything the booklet asks you to calculate. There are no `choices`; `value` is the key.

```json
{
  "id": "olig-031",
  "topic": "oligopoly",
  "type": "num",
  "stem": "Grab and Uber face P = 339 − Q and MC = 147. In the Cournot equilibrium, how much does Grab produce?",
  "choices": [],
  "answer": [],
  "value": 64,
  "unit": "rides/min",
  "explanation": "q_G = 96 − ½(96 − ½q_G) ⇒ ¾q_G = 48 ⇒ q_G = 64.",
  "source": { "deck": "Quiz Prep", "slide": 36 },
  "difficulty": 2,
  "tags": ["cournot", "packet"],
  "examFocus": true
}
```

The reader types a number. `4,096`, `4096 THB`, `1/2`, `0.5` and `50%` all parse. A
whole-number key must be hit exactly. Any other key accepts rounding to two decimal
places. Set `tolerance` (absolute) only when a looser answer should count. `unit` is shown
beside the input and is never part of the answer. Grading is never fuzzy: 65 is not 64.

Inline `?check` cards in the notes only render choice questions, so point them at an
`mcq`, `tf` or `multi`.

## `cases/<name>.json`

```json
{
  "id": "freshmart",
  "title": "FreshMart goes omnichannel",
  "topics": ["value-chain", "operations", "pm-foundations"],
  "scenario": "The setup the questions share.",
  "image": "/labels/cs-graph.webp",
  "questions": [ /* same question objects, ids prefixed `case-freshmart-01` */ ]
}
```

## `glossary.json`

```json
[{ "term": "Total float", "definition": "…", "topic": "schedule-cost", "source": { "deck": "new 06", "slide": 28 } }]
```
