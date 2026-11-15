# Content schema & style guide

Every file here is hand-checkable against a slide. If a claim cannot be traced to a
slide, it does not go in. The general formats (lecture front matter, question fields,
`text` questions, glossary) are the same as in `courses/vitsup/content/SCHEMA.md`; this file
covers what is specific to this course.

## The paper (`examformat.md`)

| Part | Format | Points | Topic |
|---|---|---|---|
| 1 Wine label | questions about a label | 5 | `labels` |
| 2 Aromas + Introduction | true / false | 10 | `intro-aromas` |
| 3 French wine | match region ↔ grape | 5 | `france` |
| 4 Italian wine | match designation ↔ region | 5 | `italy` |

## Topic ids

| id | prefix | source |
|---|---|---|
| `exam-core` | `core` | all four parts |
| `labels` | `lab` | Intro 51-55, France 18-19/36, Italy 16-18 |
| `intro-aromas` | `aro` | Intro deck; Analysis (Week 2) 4-33 |
| `france` | `fra` | France deck |
| `italy` | `ita` | Italy deck |

Decks are cited as `Intro`, `Analysis`, `France`, `Italy` (slide = PDF page), the names in
`../materials.json`.

## Match questions (`type: "match"`)

```json
{
  "type": "match",
  "stem": "Match each designation to its region.",
  "items":   ["Barolo DOCG", "Taurasi DOCG", "Franciacorta DOCG"],
  "choices": ["Campania", "Lombardia", "Piemonte", "Puglia", "Sicilia"],
  "answer":  [2, 0, 1]
}
```

`answer[i]` is the index into `choices` for `items[i]`. Two items may share a choice.
The pool must hold **at least two spare choices** that no item pairs with, so the last
pair can never be found by elimination. Spare choices are real regions or grapes from the same deck, and
the pool is alphabetical, so its order says nothing. A match is graded all-or-nothing;
the runner marks each row so you can see which pairs you had.

## No format tells

The answer must never be guessable from how it is written. `TestWineedHasNoFormatTells`
enforces the measurable part:

- **MCQ length.** The right choice is the strictly longest in 8–30% of MCQs, and the
  strictly shortest in 8–30%. Chance is about 25%, minus ties. A key that is almost never
  the longest choice is a tell too. Write every choice to the same
  grammatical shape and length. The distractor is the neighbouring region, grape or term.
- **MCQ position.** Each of the four positions holds the answer in at least 15% of MCQs.
- **True/false.**
  - Between 40% and 60% of items are True.
  - A false statement is a true one with **one fact swapped for its neighbour**: malic for
    tartaric, Côte de Nuits for Côte de Beaune. Never make an item false by adding an
    absolute.
  - Absolutes (*always, never, only, all, must*) and hedges (*usually, often, may*) appear at
    about the same rate in true and false items. The two rates may differ by at most 15 points.
- **Both directions.** A pair drilled in a match also appears the other way round, as an
  MCQ or a one-word `text` question, so it is recalled and not just recognised.

## Label cards (`cases/*.json`)

Part 1 of the paper is one label and questions about it, so each card is a case with the
label as its scenario and exactly five questions. `lbl-*` cards carry a real label photo
in `image`, served from `courses/wineed/labels/`. That folder holds third-party photos, so
it is kept out of git and shipped by `deploy.sh`. Every question must be answerable from
what the photo shows plus the slides.

Inline `?check` lines in the lectures must be `mcq` or `tf` with no markdown in the stem:
the inline card renders the stem as plain text and has no input for typing or matching.
