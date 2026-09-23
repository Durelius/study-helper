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

One survival worth a clause: rule-based reasoning is **still used in heuristic medicine**,
because a doctor's differential diagnosis genuinely *is* an if-then tree — rule this out, then
that. Everywhere else it lost to systems that calibrate their parameters from data.

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

The fullest version of the course's own history runs **five parallel threads** down the
decades — **Language & Knowledge · Vision & Perception · Neuroscience Inspiration · Algorithms
& Learning · Hardware & Data Enablers** — with **predictive AI** down one side and
**generative AI** down the other, and arrows carrying the story across. If asked how many
timelines: **two task timelines plus the neurocomputational one, with hardware and compute as
the enabler.** AlexNet's GPUs are the concrete hook.

`[slides 58-80]`

## What runs on the other three threads

The two task timelines above are the ones flagged for the exam, but the remaining threads
carry names worth a line each.

**Neuroscience inspiration.** **Hebb 1949**, *The Organization of Behavior* — **neurons that
fire together wire together**, the first account of learning as a change in connection
strength, and the reason unsupervised learning is described in this course as Hebbian.
**Hodgkin & Huxley 1952** — a mathematical model of the **action potential**. **Hubel &
Wiesel 1962-68** — the visual cortex hierarchy. **1970s** cognitive science and brain theory;
**2000s** brain imaging (fMRI, DTI) mapping functional connectivity; **2010s** deep brain
networks and neural coding.

**Algorithms & learning.** **Wiener 1948**, *Cybernetics: Or Control and Communication in the
Animal and the Machine* — feedback, control and information flow treated as one subject shared
by animals and machines; it predates the founding of AI by eight years and supplies its
vocabulary. **Dartmouth 1956.** **Minsky & Papert 1969.** **Expert systems** — **DENDRAL
(1965)** then **MYCIN (1970s)**, rule-based reasoning of the form *IF fever AND rash THEN
measles*, with a **confidence** attached. **Backpropagation 1986.** **LSTM 1997** (Hochreiter
& Schmidhuber). **Unsupervised and probabilistic models 2006-09** — RBMs, sparse coding,
graphical models. Then the **2013-2020s generative revolution**: VAE, GAN, diffusion, large
models.

**Hardware & data enablers.** **Electronic computers** in the 1940s-50s (ENIAC 1945, EDVAC
1949, the stored-program architecture) → **transistors 1947** → **integrated circuits** in the
1960s → **the microprocessor 1971** → **workstations and graphics acceleration** in the 1980s
→ the **internet, more data and distributed computing** in the 1990s → **clusters, multi-core
CPUs and early GPUs** in the 2000s. This is the thread that decides *when* an idea becomes
practical rather than *whether* it is correct.

**Language, the parts not in the table above.** **ELIZA (Weizenbaum, 1966)** — the first
well-known chatbot, matching patterns and reflecting them back as a psychotherapist would,
with no understanding at all. **Statistical NLP** from the 1980s (n-grams, HMMs).
**Bengio et al. 2003** — a neural probabilistic language model with **learned word vectors**,
a decade before word2vec made them cheap. **Seq2Seq 2014**, **Attention 2015**, **BERT 2018**.

**Vision, the parts not in the table above.** **1960s** edge, line and shape detection and 3-D
reconstruction. **1990s-2000s** the pre-deep-learning state of the art: **support vector
machines** fed **hand-designed local features** — **SIFT (1999)**, SURF, HOG. A human designed
the features and a classifier did the rest; AlexNet's real break was **learning the features
too**.

`[slides 58-80]`

?check id=tim-035
?check id=tim-038

## The seminal papers, as a list of triples

The course's own reading of deep learning is a list of papers, each pinned to one or more of
the **five principles** (*representation, relation, parameterization, optimization,
architecture* — see the concepts lecture). Learn them as **author, year, one line**; that is
exactly the shape a two-line answer can carry.

| Year | Author(s) | Contribution |
|---|---|---|
| **1847** | **Cauchy** | The **gradient method** — steepest descent, formalised about **140 years** before anyone had a network to apply it to |
| **1986** | **Rumelhart, Hinton & Williams** | **Backpropagation**: adjust weights to minimise output error; hidden units learn useful features |
| **1989** | **Cybenko** | **Universal approximation** — a **single** hidden layer with a sigmoid nonlinearity can approximate any continuous function |
| **1997** | **Hochreiter & Schmidhuber** | **LSTM** — gates that stop the error signal decaying, so long-term dependencies become learnable |
| **1998** | **LeCun et al.** | **LeNet / CNN**: locality, convolution, **weight sharing**, end-to-end learning for document recognition |
| **2006** | **Hinton & Salakhutdinov** | The **deep autoencoder** — high-dimensional data into low-dimensional learned codes |
| **2012** | **Krizhevsky, Sutskever & Hinton** | **AlexNet** — 60M parameters, GPU training; **scale works** |
| **2013** | **Mikolov et al.** | **word2vec** — continuous vector representations capturing syntactic and semantic similarity |
| **2013/14** | **Kingma & Welling** | **VAE** — continuous latent variables and the reparameterisation that makes them trainable by gradient descent |
| **2014** | **Sutskever, Vinyals & Le** | **Seq2Seq** — an LSTM maps a sequence to one fixed-length vector and decodes the target sequence |
| **2014** | **Bahdanau, Cho & Bengio** | **Attention** — soft-search the relevant parts of the source for each output word, three years before the Transformer |
| **2014** | **Goodfellow et al.** | **GAN** — the adversarial generator/discriminator minimax game |
| **2014/15** | **Kingma & Ba** | **Adam** — a first-order optimiser using **adaptive moment estimates** |
| **2015** | **Ronneberger, Fischer & Brox** | **U-Net** — a contracting path for context and a symmetric expanding path for localisation |
| **2015/16** | **He, Zhang, Ren & Sun** | **ResNet** — residual functions and identity shortcuts ease the optimisation of much deeper networks |
| **2017** | **Vaswani et al.** | **Transformer** — attention alone, **dispensing with recurrence** |
| **2020/21** | **Dosovitskiy et al.** | **Vision Transformer** — a pure Transformer applied directly to sequences of **image patches** |
| **2020** | **Ho, Jain & Abbeel** | **DDPM** — a generative diffusion model trained with a **denoising** objective |

Two observations the course draws from that list. First, **Cauchy is 140 years early**: the
mathematics was finished long before there was anything to apply it to. Second, the later the
paper, the **more** of the five principles it touches — Adam serves optimization alone, while
the Transformer serves all five, because its contribution *is* the orchestration.

`[slides 58-80]`

?check id=tim-026
?check id=tim-030

## How the ideas built on one another

A second reading of the same history runs one row per **concept lineage** — cybernetics,
artificial intelligence, machine learning, artificial neural network, perceptron, feedforward
network, MLP, deep learning, CNN, ResNet — rather than per task. What it adds:

- **AI was envisioned in 1950 and formally launched in 1956.** Turing proposed; Dartmouth
  founded.
- **Machine learning** was named by **Arthur Samuel**, whose **checkers program learned from
  self-play**; **Mitchell (1986)** defined and unified the field as a discipline.
- **Widrow & Hoff (1965)** introduced the **ADALINE**, the step between the perceptron and the
  trainable feedforward network.
- **Hornik, Stinchcombe & White (1989)** proved that **multilayer feedforward networks are
  universal approximators** — the companion result to Cybenko's.
- **LeCun et al. (1989)** applied backpropagation to **handwritten zip-code recognition**,
  nine years before LeNet-5.
- **Hinton, Osindero & Teh (2006)**, *A Fast Learning Algorithm for Deep Belief Nets* —
  **greedy layer-wise pretraining** reopened a line of work stalled since the 1980s. This is
  where "deep learning" restarts.

The dependency chain, which is the most compressible form of the whole history:

> **Perceptron (1957-58) → feedforward net + backprop (1986) → MLP as universal approximator
> (1989-90s) → deep learning (2006-2012) → CNN at scale (2012) → ResNet (2015)**, with
> **functional-link networks (1995)** as the alternative branch, and **cybernetics (1948)** and
> **AI (1950/1956)** as the context all of it sits in.

Each step removed the obstacle the previous one exposed: XOR defeated the perceptron, so
layers; layers could not be trained, so backprop; depth degraded, so residual connections.

`[slides 58-80]`

?check id=tim-042

## Names and years worth ten minutes

Cauchy 1847 (gradient method) · McCulloch & Pitts 1943 · Wiener 1948 (cybernetics) · Hebb
1949 · Turing 1950 · Hodgkin & Huxley 1952 · Dartmouth 1956 · Rosenblatt 1958 · Hubel & Wiesel
~1959 · DENDRAL 1965, expert systems, **MYCIN 1975** · **ELIZA 1966** (Weizenbaum) · Minsky &
Papert 1969 · Fukushima 1980 · **METEO 1981** · Rumelhart, Hinton & Williams 1986 · **Cybenko
1989** (universal approximation) · Pao 1995 (functional-link nets) · **LSTM 1997** (Hochreiter
& Schmidhuber) · **Deep Blue 1997** · LeCun 1998 · SIFT 1999 · Bengio 2003 (neural language
model) · Hinton & Salakhutdinov 2006; Hinton, Osindero & Teh 2006 · ImageNet 2010– · AlexNet
2012 · **word2vec 2013** · VAE 2013 (Kingma & Welling) · **GAN 2014** (Goodfellow) · Seq2Seq
2014 (Sutskever, Vinyals & Le) · Attention 2014 (Bahdanau, Cho & Bengio) · Adam 2014/15
(Kingma & Ba) · U-Net 2015 (Ronneberger) · **ResNet 2015** (He, Zhang, Ren & Sun) ·
**Transformer 2017** (Vaswani) · GPT and BERT 2018– · Vision Transformer 2020/21
(Dosovitskiy) · DDPM 2020 (Ho, Jain & Abbeel).

On the cat's date the sources disagree — 1958, 1959 and 1962 all appear. Say **"around 1959"**
and spend the words on the finding instead.

`[slides 58-80]`
