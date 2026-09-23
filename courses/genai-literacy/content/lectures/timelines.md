---
id: timelines
title: "The Language and Vision Timelines"
deck: "lec 1 + lec 5"
slides: 34
summary: "Two AI task timelines plus the neurocomputational thread that made both possible, with the names and years a two-line answer can carry."
---

## What the exam will ask from this lecture

Flagged in class twice and written down in two separate sets of notes: **the two timelines
are on the exam**. The syllabus gives them a whole session. What a two-line answer can hold
is **a name, a year and what it changed** — so learn the milestones as triples, not as prose.

The other thing to be ready for is the *shape* question: **how many timelines are there and
what are they?** Answer: **two AI-task timelines — Language and Vision — plus a
neurocomputational (connectionist) thread underneath both**, with **hardware and compute** as
the enabler that let the third thread finally pay off.

## The Language AI timeline

Getting a computer to use language like a person.

| Year | Milestone | Why it matters |
|---|---|---|
| **1950** | **Turing**, *Computing Machinery and Intelligence* | The **imitation game**: if you cannot tell machine from person, stop asking whether it thinks |
| **1956** | **Dartmouth College summer workshop** | The field gets its name and its founding agenda |
| 1965– | **Expert systems**; **MYCIN 1975** (~65% accuracy on bacterial infections) | Rule-based AI's high-water mark, written in **LISP** and **Prolog** |
| **1981** | **METEO**, Canada | Machine translation of English↔French weather reports — NLP that actually shipped |
| **1997** | **Deep Blue** beats the world chess champion | Search plus hand-built evaluation, not learning |
| **2013** | **Mikolov**, **word2vec** | Words become **vectors**, so a network can take them as input |
| **2017** | **Vaswani**, *Attention Is All You Need* | The **Transformer** — the **T** in GPT |
| **2018–** | **GPT** | Generative Pre-trained Transformer; GPT-3 in 2020, then ChatGPT |

The long flat stretch in the middle is the point. Between Turing and word2vec, language AI
was **symbolic**: hand-written if-then rules over symbols, no learning from data. It stalled,
brought on the **AI winter**, and was **retronymed GOFAI** — Good Old-Fashioned AI — once the
neural approach took the name "AI" for itself.

> **Exam focus.** "What is symbolic AI (not neural networks)? Expert systems?" is one of the
> examiner's own questions. Two lines: hand-written if-then rules over symbols with no
> learning, e.g. MYCIN in LISP or Prolog; retronymed GOFAI once neural AI took the name.

`[slides 58-70]`

?check id=tim-003

## Why word2vec had to come before the Transformer

Neural networks take numbers. Words are not numbers. **Word2vec vectorised English**, and
once words were vectors, the geometry carried meaning: **king − male + female ≈ queen**. The
training task was simply to guess a **missing word** from its neighbours.

Attention can then weigh those vectors against one another. The ordering is not an accident
of history — **vectorising the words is what made attending to them possible**.

`[slides 68-70]`

## The Vision AI timeline

Getting a computer to recognise what it is looking at.

| Year | Milestone | Why it matters |
|---|---|---|
| **1959/1962** | **Hubel & Wiesel**, cat's **striate cortex** | Neurons fired for **oriented lines**, not for whole objects or for a mouse |
| **1980** | **Fukushima**, **Neocognitron** | The first convolutional architecture, built directly on that finding |
| **1998** | **LeCun**, **LeNet** | CNNs crack handwritten character recognition |
| **2010–** | **ImageNet** (**Fei-Fei Li**) | A large **labelled** dataset — the missing ingredient |
| **2012** | **AlexNet** (Krizhevsky, Sutskever, Hinton) | Error falls from **~25% to single digits** |
| ~2014 | Phones tell a cat from a dog | The capability reaches ordinary hardware |

The cat experiment is the one to be able to describe cold. **The cat's neurons did not fire
when it saw a mouse; they fired when it saw a line.** Two line detectors combine into a
curve detector, curves into shapes, shapes into objects. That **hierarchy of features** is
the blueprint the Neocognitron and every CNN since have copied.

> **Exam focus.** Flagged in class as "exam for sure". Expect either *what did Hubel and
> Wiesel observe and why does it matter for AI?* or the examiner's own version, *why did
> mobile phones start recognising images in the 2010s?* The second answer is a **convergence,
> not a breakthrough**: large labelled datasets (ImageNet), **GPU compute** for the matrix
> multiplications, and **deep CNN architectures** — all three arriving together.

`[slides 71-78]`

?check id=tim-008

## The third thread: connectionism

Underneath both task timelines runs the **neurocomputational** story — the one about
*mechanism* rather than about tasks.

| Year | Milestone |
|---|---|
| **1943** | **McCulloch & Pitts** — a mathematical model of a neuron |
| **1958** | **Rosenblatt** — the **Perceptron**, implemented in hardware as the Mark I |
| **1969** | **Minsky & Papert** — a single perceptron cannot learn XOR; funding collapses |
| **1980** | Multi-layer networks |
| **1986** | **Rumelhart, Hinton & Williams** — **backpropagation** trains an MLP |

The claim behind the label **"Connectionist AI"**, flagged in class as an exam note: **the
brain has no central processor. It is just neurons connected to neurons.** So to build
intelligence you replicate the *mechanism* of connection, not the outward *behaviour* — which
is exactly where symbolic AI went wrong. **BNN** is the biological neural network, the actual
brain; **ANN** is the mathematical model of it, and every modern AI is built from ANNs.

`[slides 61-67]`

?check id=tim-011

## Where the two timelines meet

They do not stay separate. **GANs** (2014) made "generative" a popular word on the vision
side; the **Transformer** (2017) did the generative turn on the language side; and **diffusion
models** carried the generative machinery back to images, using a **text** encoder to condition
what gets drawn. Today's text-to-image systems are the two timelines fused.

The lecturer's own poster runs several parallel threads down the decades — **Language &
Knowledge · Vision & Perception · Neuroscience Inspiration · Algorithms & Learning · Hardware
& Compute** — with **predictive** down one side and **generative** down the other, and arrows
carrying the story across. If asked how many timelines: **two task timelines plus the
neurocomputational one, with hardware and compute as the enabler.** AlexNet's GPUs are the
concrete hook.

## Names and years worth ten minutes

Turing 1950 · Dartmouth 1956 · McCulloch & Pitts 1943 · Rosenblatt 1958 · Hubel & Wiesel
~1959 · expert systems 1965, **MYCIN 1975** · Fukushima 1980 · **METEO 1981** · Rumelhart,
Hinton & Williams 1986 · **Deep Blue 1997** · LeCun 1998 · Hinton & Salakhutdinov 2006 ·
ImageNet 2010– · AlexNet 2012 · **word2vec 2013** · VAE 2013 (Kingma & Welling) · **GAN 2014**
(Goodfellow) · U-Net 2015 (Ronneberger) · **Transformer 2017** (Vaswani) · GPT 2018– · DDPM
2020 (Ho, Jain & Abbeel).

On the cat's date the sources disagree — 1958, 1959 and 1962 all appear. Say **"around 1959"**
and spend the words on the finding instead.

`[slides 58-80]`
