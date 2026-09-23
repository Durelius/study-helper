---
id: concepts
title: "The I, the A and the G"
deck: "lec 1"
slides: 81
summary: "What the three letters of GAI actually mean, how predictive AI differs from generative AI, and where machine learning sits relative to AI — the vocabulary the whole paper is written in."
---

## What the exam will ask from this lecture

The examiner's blueprint line is **"Concepts and Principles (lecture 1)"**, and this is the
best-corroborated material in the course. Expect two or three of the ten questions from
here. The shapes to prepare are: *give the three senses of the I (or the A)*, *explain the
difference between predictive and generative AI with an example of each*, and *in what way
is machine learning larger than AI, and in what way is AI larger than ML?*

Two ruled lines is about 25–35 words. Lead with the term, then the distinguishing clause.

## The three senses of "I"

**Intelligence** is used in three unrelated ways, and the course wants all three:

| Sense | What it means | Where you meet it |
|---|---|---|
| **Human-like** | Quality of interaction so human-like you cannot tell machine from person | **Turing's imitation game**; the "I" in AI |
| **Advantageous information** | Extracted information of tactical or strategic value, hard-won | **Military Intelligence**, **Business Intelligence** |
| **Solution path** | An optimal or clever method for a hard problem | **Computational Intelligence**, **Digital Intelligence** |

Note what the first sense does *not* require: being right. A chatbot passes Turing's test
by conversing convincingly, not by knowing anything. That is why the course keeps saying a
system can be "intelligent" in this sense and useless in every other.

> **Exam focus.** The three senses of "I" and the three senses of "A" are each a clean
> two-line answer, and they sit directly under the blueprint's "Concepts and Principles"
> line. Learn them as lists of three with one example word attached to each.

`[slides 15-22]`

?check id=con-001

> **Exam focus.** *"What are three forms (at least) of intelligence?"* was flagged in class as an
> exam notice, in those words. The answer is the three senses above: **human-like** (Turing's
> imitation game), **advantageous information** (military, business), and **an optimal solution
> path** (computational, digital). The trap is giving only the first — it is the one everybody
> thinks of, and on its own it is a third of the answer.

?check id=con-041

## The three senses of "A"

1. **Artifact / artifice** — a machine creation made by humans, as opposed to something
   found in nature. True of every AI system, so it never distinguishes one case from another.
2. **Artificial semblance** — the sense in which "seems about right" will do. This is the
   standard a generated image is held to.
3. **Artful creation** — *"truth manufactured"*, shaped by whatever is trending.

`[slides 19-21]`

## The "G", and what "generate" means

**GAI = Generative AI**. **AGI = Artificial General Intelligence** — human-level breadth
across tasks, and explicitly **out of scope for this course**. The G is the letter that
changed meaning: before 2014 "generative" was a technical term in statistics; the
**Generative Adversarial Network** made it popular, and GPT then inherited it.

"Generate" is unpacked three ways:

- **Simulate** — draw from a distributional prior.
- **Connect** — graph or network adjacency.
- **Create** — integrated innovation.

And it is opposed to two different things depending on the context: **generative vs
predictive** in AI, and **generation vs recognition** in the mind. Memorise this pair
verbatim: *recognition is an **interpolative** map within known territory; generation is an
**extrapolative** map into unexplored territory.*

`[slides 23-32]`

?check id=con-011

## Predictive AI vs Generative AI

The single best-corroborated topic in the course. Two framings, both worth having:

| | Predictive AI | Generative AI |
|---|---|---|
| **Judged on** | **Accuracy against verifiable ground truths** | **Plausible context relevance** |
| **Asks** | "What **will** be?" | "What **could** be?" |
| **Mechanics** | Learn **x → y**, then crank the input | Learn **z → x**, then crank the **PRNG** |
| **Probabilistically** | Discriminative, **p(y\|x)** | Generative, **p(x)** or **p(x\|c)** |
| **Example** | An MLP classifier (supervised); cluster analysis (unsupervised) | GPT |

The second framing comes with its own images: *from many throws, abstract the law of motion
and predict where the ball lands* (predictive); *from many faces, abstract the essence of
"faceness" and imagine a new face* (generative). Both rest on the same core operation —
**abstraction**: particulars in, generality out, applied to situations never seen.

A sentence worth writing down exactly as it stands:
> **"AI did not invent abstraction; it industrialized it — turning learned generalities into
> algorithmic capability at scale."**

The abstraction ladder behind it: **animal** — abstraction becomes *adaptive* (implicit,
embodied); **human** — *deliberate* (explicit, systematic, taught as disciplines); **AI** —
*scalable* (algorithmized, extracted from massive data).

> **Exam focus.** 2024's paper asked this outright — *"Elaborate on Predictive vs Generative
> AI. Give an example of each."* Write the evaluation standard first, the example second.
> The marks are in "verifiable ground truths" against "plausible context relevance".

`[slides 29-33, 43]`

?check id=con-007

## ML and AI: neither contains the other

An exam question flagged in class, in both directions:

- **ML is larger than AI** because its methods — calibrating parameters from data — reach far
  outside anything anyone calls AI: statistics, optimisation, forecasting, any data-driven
  model fitting.
- **AI is larger than ML** because it also includes **rule-based symbolic systems, search,
  planning and robotics** that involve no learning at all.

And the reason they overlap so heavily now: **"ML builds AI."** You **cannot pre-program
intelligence with if-then constructs**; you program the system to **calibrate its model
parameters from data**, and that calibration is machine learning. The three principal
paradigms are **supervised** (labelled targets, error correction), **unsupervised**
(structure only — clustering), and **reinforcement** (policy optimised against reward), with
**representation learning** added as a fourth modality.

`[slides 59, 76]`

## The Intelligent Being, in five elements

Nacaskul's framework, and the spine of the 2025 paper. Learn all five; expect at most one.

1. **Corporal Form & Physical Embodiment** — a body with sensors and some autonomy, so the
   agent can interactively explore and learn from its environment. *Without a body you
   cannot learn to swim.*
2. **Valenced Feeling & Graded Sensation** — grading what is sensed on a **preferential
   scale** from very negative through neutral to very positive, which acts as **feedback to
   the learning algorithm** and conditions future behaviour.
3. **Cognitive Faculty & Pattern Recognition** — learning an internal model of reality,
   forming memory, classifying and predicting — **hence Predictive AI**.
4. **Mental Formation & Projective Generation** — imagining realistic scenarios, shapes and
   ideas — **hence Generative AI**.
5. **Consciousness Marker & Program Register** — momentary, sequential awareness of the
   computation step being performed; the hook for the machine-consciousness debate.

> **Exam focus.** "Define **valenced** feeling and graded sensation" was flagged in class and
> appeared verbatim on the 2025 paper. Two lines: the preferential scale, then the fact that
> it is the feedback signal shaping later behaviour.

`[slides 16-18, 54]`

?check id=con-018

## The five principles of modern deep learning

The course's own compression of the whole field, and five words that cost thirty seconds:

> **Representation → Relation → Parameterization → Optimization → Architecture**
>
> Represent → Relate → Parameterize → Learn → **Orchestrate**.

| # | Principle | In one line | Example |
|---|---|---|---|
| 1 | **Representation** | Information → **numerical** representation: encode what matters as vectors or tensors so it becomes computable | Tokens and **embeddings**; image patches |
| 2 | **Relation** | Relationships → **operations on** representations: express abstract relationships algebraically or geometrically | **Attention** — query, key, value |
| 3 | **Parameterization** | Complexity → **learned parameters + composition**: simple operations, massively parameterised, express complex functions | A deep network's weights |
| 4 | **Optimization** | Parameters → **gradient-based learning**: make the system differentiable so parameters can be learned from data | **Gradient descent**, Adam |
| 5 | **Architecture** | Intelligence → **clever orchestration of 1–4**: design how representations, relations, parameterised operations and learning interact | The Transformer block |

And the closing claim, which is the quotable sentence:

> **"Not ever more complicated primitive mathematics, but increasingly ingenious orchestration
> of representations, relationships, parameters and learning."**

The primitives stay simple — a weighted sum, a nonlinearity, a derivative. What changes is how
ingeniously they are arranged. A useful test of whether you have the scheme: **Adam serves
optimization alone**, because it changes only how parameters are updated; **the Transformer
serves all five at once**, because its contribution *is* the orchestration.

> **Exam focus.** The examiner said to study the infographic sheets carefully, and this is the
> newest of them. Five words in order, plus the closing sentence, is a complete two-line
> answer — and each principle is also a ready-made "which principle is this?" question.

`[slides 1-5]`

?check id=con-028
?check id=con-030

## The literacy roadmap: four stages, three timeline families

The lecturer's own map of the course, and a cheap source of structure answers.

**Stage 1 — Words Matter.** Build the language: (Machine) Intelligence, Artificial, Generative,
Literacy, Roadmap. Clear concepts first, because you cannot ask how a thing works until you can
name its parts.

**Stage 2 — Timelines.** **Three complementary families**, explicitly *not* a strict taxonomy:

| Family | What it tracks | Landmarks |
|---|---|---|
| **Task equivalence** | Reproducing human **capabilities** | Turing 1950 → symbolic AI and expert systems 1960s-80s → deep learning 2010s → LLMs 2020+ |
| **Computational model** | Competing **conceptions of intelligence** | McCulloch & Pitts 1943 → Dartmouth 1956 → connectionist models 1980s → deep learning 2010s+ |
| **Integrated autonomy** | From models to **real-world autonomy** | AI as **software** → AI as **operating system** (agentic: reasoning, planning, tool use) → AI as **hardware** (robotics, embodied AI) |

That last row ends in **Embodied AI**, described with five elements that map onto the
Intelligent Being: **physical form** (sensing, manipulation, motion), **valenced sensing**
(positive, negative, neutral), **cognition function** (predictive intelligence), **concept
formation** (generative intelligence) and **experience register** (event processing).

**Stage 3 — HIWUTH, "How It Works Under The Hood".** The six architectures in learning order:
**MLP (1986) → CNN (2012) → VAE (2013) → U-Net (2015) → Transformer (2017) → Diffusion
(2020→2022)**, running predictive → transition → generative. Note those are a **learning
sequence**, not a claim of historical or architectural descent.

**Stage 4 — Application.** Four modes, one verb each:

> **Chatbot "talks with you" · RAG "searches for you" · Agent "does things for you" · Wiki
> "builds your knowledge".**

`[slides 1-5]`

?check id=con-037
?check id=con-039

## How the lecturer frames these in class

Not examinable in themselves — Marko sets the paper and none of this is in his decks — but
these are the framings the lecturer uses, and they are often the clearest statement of a
distinction that *is* examinable. Learn the distinction; the framing is just the handle.

- **The five elements of the Intelligent Being are drawn loosely from the five aggregates of
  Buddhism**: *form* (the physical body — robotics), *feeling* (categorising sensations as
  pleasant or unpleasant), *perception* (recognising, labelling and identifying sensory inputs
  and mental objects), *mental formations*, and *consciousness*. If you can hold the five
  aggregates you can hold the five elements, because they are the same list.
- **On valenced feeling:** *"Being hungry is not an intelligent calculation — it is reading a
  negative signal coming from the body."* That is the whole element in one sentence.
- **On embodiment:** *"AI is intelligent but cannot do things. Robots can do things but are not
  intelligent."* **Embodied AI** is the join: a physical form with the ability to experience
  its environment.
- **On Integrated Innovation:** *"Integration is not simply the combination of features A + B +
  C + D. Not a salad bowl, putting in many things at your convenience."*
- **On generative AI:** *"Taking certain things and running with it is the premise of
  generative AI. Approximating reality is predictive AI. All fiction is realistic but not
  real."* And the warning that goes with it: **extrapolation beyond the basis of validation is
  why AI is so powerful and so dangerous** — there is no ground truth left to check against.
- **On knowing without understanding:** *"A boy can play catch without knowing Newtonian
  mechanics."* The same point as the dog that knows where the frisbee will land: abstraction
  can be entirely implicit.
- **On ML and AI:** to a **computer scientist**, machine learning sits *inside* AI research; to
  a **data scientist**, AI is the *packaging and branding* on top of machine learning. Two
  professions looking at one stack from opposite ends — which is exactly why neither field
  contains the other.

`[slides 15-25]`

?check id=con-044
?check id=con-046

## Three more one-liners to have ready

- **Applied vs application.** "Applied" means working under **real-world constraints**, not
  simplified — which is why applied mathematics is *harder* than theoretical, not easier.
  "Application" splits into modification/optimisation, addition/enhancement, and
  solution/utilisation.
- **Integrated Innovation.** *Not* the integration of impressive features; it is **designing
  an innovation to work within an integrated reality**.
- **AI Literacy vs "Artificial Literacy."** Literacy means you can code it, safeguard it and
  explore or exploit it. Artificial Literacy is the failure mode — outsourcing your essay,
  your pitch and your reading to a chatbot.

`[slides 36-42]`
