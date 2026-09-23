# 5602203 GAI Literacy & Applications — Midterm Focus

Chulalongkorn / CSII / BAScii, Autumn (Fall) 2026.
**Poomjai Nacaskul gives the lectures. Marko Niinimaki writes the exam.** (Kongkan Kalakan runs post-midterm workshops.)

Compiled 2026-09-23 from `/Users/wilhelmdurelius/chula/ailiteracy/`. Revised twice the same day: once after 14 further images arrived (**§12**), once after the student confirmed **who sets the paper** — which changed the ranking materially.

Tags: **[EVIDENCE]** = stated in a course document. **[MARKO]** = written by the examiner himself, the strongest class of evidence here. **[INFERENCE]** = my reasoning. **[GUESS]** = weak. **[LOW-RES]** = *(no longer used — every image has since been read at full resolution; see §11 and §12)*.

> **Two things to read first:** §1.1, the examiner's own blueprint for this paper, and §12, which contains the one finding that beats everything else in this document — **Nacaskul uploads full recordings of this year's classes to YouTube, including the pre-midterm ones**.

---

## 1. What the exam is

### 1.1 Who writes it, and why that changes everything

**[EVIDENCE]** The deck `Gen AI Lit 07.pdf` is **Marko Niinimaki's**. His decks in this folder are `Gen AI Lit 02 / 03 / 06 / 07.pdf`; the 02 title slide reads *"Generative AI Literacy and Applications 02 — Chulalongkorn University, School of Integrated Innovation, Fall 2026 — Marko Niinimaki"*, and 03/06/07 continue the same series and numbering.

That matters because **all three exam-defining statements in this course come from that deck** — which means they come from the man setting the paper:

1. **[MARKO]** The **answer sheet mock-up** (p.3): 10 questions, two ruled lines each, closed book, *"We only read what you can write on the lines. And your handwriting must be readable."*
2. **[MARKO]** **"No code, very little math."**
3. **[MARKO]** The final slide, **"Main topics for the midterm exam"**.

**[INFERENCE — the pivot of this whole document]** Those are not a colleague's summary of someone else's exam. They are the examiner describing his own paper. **§1.3 below should be read as the blueprint, and everything else ranked against it.** Where Nacaskul's lecture emphasis and Marko's blueprint agree, confidence is high. Where they diverge, Marko wins on *style and question shape*; Nacaskul still wins on *content*, because Marko examines the course Nacaskul teaches.

### 1.2 The mechanics

### Date

**Friday 25 September 2026, 13:00** (Bangkok). Confirmed by the student on 23 September —
this supersedes the inference below, which was drawn from the lecture-7 sequence before
the date was known.

Closed book. **10 questions, two ruled lines of answer each**, per the answer sheet on
p.3 of `Gen AI Lit 07.pdf`, which carries the warning: *"We only read what you can write
on the lines. And your handwriting must be readable."* Worth **20%** of the grade, with a
**50% hurdle in every grading criterion** (syllabus).

Two lines is roughly 25–35 words. That is one definition plus one distinguishing clause,
or one mechanism plus one consequence — nothing longer will be read.

**[EVIDENCE — syllabus `curriculum.pdf` = `_Autumn_2026_...Syllabus...pdf`, byte-identical, md5 f19689d2…]** Assessment: Attendance 10, Participation 10, Midterm **20**, Final 20, Project 40. The session table marks the midterm "Closed-book", sitting after session 7.

**[EVIDENCE — syllabus CLO → assessment map]** Only one learning outcome is assessed by the exams:
> **CLO 2: "Explain key GAI architectures using intuitive mental models, i.e. how Transformer and Diffusion Models 'work under the hood'." → Midterm/Final Exam**

CLO 1 (history/philosophy) → Homework. CLO 3 (RAG/agentic workflows) → Term Project. CLO 4 (impact on society) → Class Discussion.
**[INFERENCE]** Architecture explanation is the formally examined outcome — but Marko's topic slide is broader than the CLO map, so treat CLO 2 as a floor, not a fence.

### 1.3 Scope, in the examiner's own words

**[MARKO — `Gen AI Lit 07.pdf`, final slide, verbatim]**

> **Main topics for the midterm exam**
> Please check the Youtube playlist in MyCourseVille Supplementary Material!
> Some topics covered in the CNN course (assignment, Wolfram U), too.
> **Study the infographics sheets carefully!**
> Concepts and Principles (lecture 1)
> Neural networks and how they learn
> Architectures: Multi-Layer Perceptron, Convolutional NN, VAE, ..
> **No code, very little math.**

Six instructions, each now a first-party statement about this paper:

| Blueprint line | What it means for revision |
|---|---|
| YouTube playlist (MyCourseVille) | The MustWatch list in §5 is examinable material, not optional enrichment. |
| **Wolfram U CNN course** | **Upgraded.** Marko runs the labs and set that assignment; the two quiz files are his chosen concepts. See M2b. |
| "Study the infographics sheets **carefully**" | The infographics are the answer key. §11. |
| Concepts and Principles (lecture 1) | The G/A/I unpacking, predictive vs generative, the timelines. H3–H5. |
| Neural networks and **how they learn** | Perceptron, MLP, weights/bias, gradient descent, backprop. H2, M1. |
| Architectures: MLP, CNN, VAE, .. | H1. The "…" is doing work; Transformer and diffusion sit under it. |
| **No code, very little math** | No Python, no Wolfram function names, no hand-computed derivatives. §9. |

### 1.4 Syllabus sessions before the midterm

| # | Contents | Lecturer |
|---|---|---|
| 1–2 | The "I", "A", and "G" in "GAI" | Nacaskul |
| 3 | The "Language AI" & "Vision AI" Timelines | Nacaskul |
| 4 | MLP, Neocognitron/CNN | Nacaskul |
| 5 | AE/VAE, Encoder-Decoder | Nacaskul |
| 6 | GOFAI, NLP, Transformer, LLM | Nacaskul |
| 7 | Prompt/RAG, Mix Architectures | Nacaskul |

**[INFERENCE]** Delivery ran behind this plan. Your lecture-6 notes are about ML vs AI, embodied AI, gradient descent and RMSE — session 4–5 material — and Marko's decks 06/07 cover an autoencoder lab and a Transformer lab. **Prompt/RAG (session 7) was pushed past the midterm**: `Gen AI Lit 07.pdf` says the *post*-midterm classes are the hands-on Agentic AI and RAG tutorials. Treat Prompt/RAG as low probability (§9).

---

## 2. Evidence from past papers — and why it is weaker than it looks

### 2.0 The attribution caveat (read before using §2)

**[INFERENCE — important]** Both past papers are in **Nacaskul's** style: 25-ish questions, his signature "name the Nth element of X and briefly explain", his BAScii "Integrated Innovation" vocabulary, his five-element Intelligent Being framework, and a five-question block of hand-computed gradient descent. **Marko is setting this paper.** So:

- **Topic evidence still holds.** Same course, same syllabus, same lecture content. What was examinable in 2024 and 2025 is still roughly what is examinable now.
- **Style evidence largely does not.** "Unpack this acronym and explain one letter", "name element (ii) of the Intelligent Being", and the BAScii meta-questions are *Nacaskul's habits*, not necessarily Marko's. Every prediction resting on them is downgraded below and marked.
- Marko's own question style is a separate and better guide — see §3.

### 2.1 A second caveat: the 2025 paper is a biased sample

**[EVIDENCE]** The file is named `(excerpt) …` and begins at "[Page 2]" with **Q11**. **Questions 1–10 are not obtainable** — the student does not have them, and both past papers are filenamed "(excerpt)", so the first page appears to have been withheld deliberately.

**[INFERENCE]** Two consequences. First, the visible sample is **biased toward the later questions**, so do not read my topic frequencies as the true frequencies. Second — and I am applying this as a deliberate weighting, not a fact — **material withheld from students is usually withheld because it gets reused**, and a paper's *opening* questions are almost always the foundational ones. Since a 10-question short-answer exam is mostly made of basics anyway, I have **nudged the foundational topics (H3, H4, H5) up rather than down**, even though they are under-represented in the visible portion. I will not guess at the specific missing questions.

### 2.2 The 2025 midterm, Q11 onward

**Meta / programme concepts** *(Nacaskul-specific — downgraded, see D1)*
- **Q11.** *What does "and" mean within the context of your "Integrated Innovation" education?* → **And ⇔ Integrated**, as applied to BAScii's **multidisciplinary education** model, producing socially responsible **innovators** — *not* bolting desirable features together and calling it an innovation.
- **Q12.** *What does it mean to achieve the bare minimum level of "AI Literacy"?* → As a **global digital citizen**, appreciating **how AI got where it is**, **how it is trained from data**, and the **capability as well as pitfalls** of AI solutions, so as not to apply AI blindly.
- **Q13.** *How does adding a chatbot to your e-commerce channel qualify as "AI Application"?* → AI Application **in the most basic form**, relying on the chatbot's **"human-like"** conversation, **without implying its suggestions are intelligent solutions** to real problems.

**Architecture acronyms** *(style is Nacaskul's; the content is core for anyone)*
- **Q14.** *What does MLP stand for? Explain the "P".* → **Multilayer Perceptron**. The Perceptron is a **mathematical-computational model of a generic neuron**: it **sums weighted excitory and/or inhibitory incoming signals** and **fires if/when that exceeds an internal threshold**.
- **Q15.** *What does CNN stand for? Explain the "C".* → **Convolutional Neural Network**. The **convolutional layer** is a layer of **convolution matrices with trainable entries** which, applied to a **grid of pixel values**, **detect underlying visual patterns**.

**The five elements of the Intelligent Being / "Embodied AI" (Q16–Q20, one per question)** *(framework downgraded as a question pattern — but learn the five, they are short)*
1. **Corporal Form & Physical Embodiment** — a physical body: robotics with some autonomy and an array of sensors, letting the AI interactively explore and learn from its environment.
2. **Valenced Feeling & Graded Sensation** — evaluating sensed input on a **preferential scale** (very/slightly negative → neutral → slightly/very positive), serving as **feedback to the learning algorithm** and thereby conditioning future behaviour.
3. **Cognitive Faculty & Pattern Recognition** — learning an internal model of how reality operates, forming memory, identifying/classifying/predicting outcomes — **hence Predictive AI**.
4. **Mental Formation & Projective Generation** — generatively imagining realistic scenarios, shapes, forms and ideas — **hence Generative AI**.
5. **Consciousness Marker & Program Register** — momentary, sequential awareness of the computation/inference step being performed; the hook for the machine-consciousness debate.

**Numerical optimisation (Q21–Q25)** — on f(x,y) = x² + y² − xy − 5x − 5y + 25: evaluate f(3,3)=4; derive ∇f = (2x−y−5, 2y−x−5)ᵀ; evaluate ∇f(3,3)=(−2,−2)ᵀ; step 0.1 against it to (3.2,3.2); compare.
**[INFERENCE — confident] Will not recur.** Marko's own words are "very little math", and two ruled lines cannot hold working. Keep the verbal definition only.

**Bonus questions**
- **Bonus 1.** *1–3 key messages from "Is AI Slop Killing the Internet?" (Patrick Boyle)* → the old ad model paid creators for engaging content; AI search now answers without a visit, so the incentive to produce high-quality trustworthy content evaporates; countermeasures are limited, so consumers must stay vigilant.
- **Bonus 2.** *Rank ChatGPT / Gemini / Claude Sonnet / DeepSeek from your own experience* — accuracy, RAG vs pure generation, legitimate URLs, concise organised narrative. (Open opinion.)

**[EVIDENCE]** That Boyle video is item #21 on the **current** playlist. Bonus questions have been drawn straight from the playlist before.

### 2.3 The 2024 "Applied AI" midterm

Opens with a **term bank** to use and circle: *{Artificial Intelligence (AI), Machine Learning (ML), Unsupervised Learning, Supervised Learning, Reinforcement Learning, Regression Analysis, Multilayer Perceptron (MLP), Cluster Analysis, Dimensionality Reduction, Predictive AI, Generative AI, Verifiable Ground Truths, Plausible Context Relevance, (Word2Vec) Word Embedding, Generative Pretrained Transformer (GPT), Large Language Model (LLM)}*

- **Q1.** *What do we mean by "ML builds AI"?* → **AI cannot be programmed with traditional if-then constructs**; instead it is programmed to **calibrate its model parameters from data**, hence a Machine Learning algorithm. **Supervised, Unsupervised and Reinforcement Learning are the 3 principal ML paradigms.**
- **Q2.** *Elaborate on Predictive vs Generative AI. Give an example of each.* → **Predictive AI is evaluated on accuracy against verifiable ground truths**; MLP is a supervised example, cluster analysis an unsupervised one. **Generative AI is evaluated on its ability to generate responses with a high degree of plausible context relevance**; GPT is the best-known example.

### 2.4 What recurs — with the attribution discount applied

| Theme | 2024 | 2025 | Verdict now |
|---|---|---|---|
| **Predictive vs Generative AI** | Q2 (core) | whole Embodied-AI spine | **Still near-certain** — the course's spine, not a stylistic tic |
| **ML/AI relationship** | Q1 | — | **Still high** — also flagged in class this year |
| Unpack an acronym, explain one letter | term bank | Q14, Q15 | **Downgraded to medium** — Nacaskul's format. Know the content; don't expect the wording |
| Five-elements-one-per-question | — | Q16–20 | **Downgraded to low-medium** as a *pattern*; "valenced" survives separately |
| BAScii / Integrated Innovation / literacy meta | — | Q11–13 | **Downgraded to low** — Nacaskul's personal vocabulary, absent from the blueprint |
| A bonus drawn from a playlist video | — | Bonus 1 | **Held at medium** — the blueprint also points at the playlist |

---

## 3. The examiner's own question style

**[MARKO]** Marko's decks contain questions he wrote and put to the class. These are the best available model of how *this* paper will be phrased — short, plain English, mechanism-seeking, no acronym games.

From **`Gen AI Lit 02.pdf`** ("Questions to start with" / "Bonus"):
> *What were the applications? Some product/person/team names?*
> ***What is symbolic AI (not neural networks)? Expert systems?***
> ***Can we have generative AI without neural networks?***
> ***What was the "hard part" about applied mathematics?***
> ***Why did mobile phones start recognizing images in the 2010s?***

From **`Gen AI Lit 06/07.pdf`** (lab prompts):
> *What is a latent space?*
> ***How does the transformer decide which word it generates next? What are the main parts of this transformer?***

**[INFERENCE]** Note the shape: *"Why did X happen?"*, *"Can we have X without Y?"*, *"What are the main parts of X?"*, *"How does X decide Y?"*. These are **causal and compositional**, not definitional recall. Prepare to **explain a mechanism in one sentence and name its consequence in the second**, rather than to recite a numbered list. Model answers for the five bolded ones:

- **Symbolic AI / expert systems** — rule-based formal logic: hand-written if-then rules over symbols, e.g. MYCIN (1975) diagnosing bacterial infections at ~65% accuracy, written in LISP or Prolog. No learning from data; retronymed **GOFAI** once neural AI took the name.
- **Generative AI without neural networks?** — Yes in principle: Markov chains (1906) generate text from next-token statistics, and any sampler from a fitted distribution is generative. But without a learned representation it has no context, so output is locally plausible and globally incoherent. (Your lecture-2 note: *"possible, but it's going to be random and not have context."*)
- **The "hard part" about applied mathematics** — not that the maths is simplified, but that the conclusion must **conform to real-world constraints**; applied is harder than theoretical for exactly that reason.
- **Why phones started recognising images in the 2010s** — three things arrived together: **large labelled datasets** (ImageNet), **GPU compute** for the matrix multiplications, and the **deep CNN** (AlexNet 2012, error 25% → single digit). Not one breakthrough, a convergence.
- **How a transformer decides the next word** — it embeds and positionally encodes the tokens, uses **attention (Q/K/V)** to weigh how much every other token should inform each one, and outputs a probability distribution over the vocabulary via softmax, from which the next token is sampled.

---

## 4. Emphasis flagged *in class this year* (your own notes)

Nacaskul's lecture emphasis. **[INFERENCE]** These now carry weight as *content* signals rather than as guarantees of wording, since he is not setting the paper — but he tells the class what matters, and Marko examines his course.

**[EVIDENCE — `lec1/during.md`, your "Exam notice"/"Exam note" markers]**
- **"Define valenced (valenced feeling & graded sensation) in the context of an intelligent being and robotics"** — verbatim 2025 Q17.
- **"What are three forms (at least) of intelligence?"**
- GANs were invented as a device to *improve discriminative training* and ended up the star; deepfakes made them famous.
- **"English words were vectorized. You can do word algebra"** — King − male + female = Queen.
- **"12288, ChatGPT vector size."**
- **"Diffusion model comes from transformer and is for visuals."**

**[EVIDENCE — `lec3.md`]** "Exam note: **Connectionist AI**"; "Exam notice: **count parameters**" (2 inputs + 1 output → 3 parameters), noted twice.

**[EVIDENCE — `lec5.md`, headed "Synopsis of class topic coverage (up to midterm)"]**
1. Precise meanings of **G, A, I**, and of **AI, GAI, AGI**
2. The **2 "AI task" timelines + 1 "neuroscience foundation" timeline**
3. **HIWUTH: the 6 foundational architectures + 2 foundational algorithms**
Plus **"Exam for sure: 1958 cat line recognition experiment"**.

**[EVIDENCE — `lec5/summarized_notes.md`, "What she says is on the exam"]** The two timelines; the 3 groups and 6 architectures, **"be able to tell them apart"**; **what a parameter, coefficient, weight and input are**.

**[EVIDENCE — `lec6/during.md`]**
- **"Exam question: In what way is machine learning larger than artificial intelligence? And vice versa"**
- **"Exam sentence: BACKPROPAGATION IS A SPECIALIZATION OF GRADIENT DESCENT ALGORITHM, APPLIED TO MINIMIZE PREDICTION ERROR FOR MULTI LAYER PERCEPTRON (by tweaking weights and bias)"** — memorise literally.
- RMSE = the loss L; GDA minimises it; **ADAM = Adaptive Moment Estimation**.
- PAI mimics (re)cognition; GAI mimics thinking/generation. (AI =) ANN = f(parameters) = architecture.

---

## 5. The YouTube playlist, and what changed right before the exam

**[EVIDENCE]** `(2026.09.21)` vs `(2026.09.15)` spreadsheets. Tags: 1=MustWatch, 2=QuickWin, 3=XtraInsight. The blueprint points at this playlist explicitly, so treat MustWatch as assigned.

| Topic | MustWatch video |
|---|---|
| Perceptron & ANN | 3Blue1Brown — *But what is a neural network?* (ch.1) |
| Gradient descent / backprop | 3Blue1Brown — *Gradient descent, how neural networks learn* (ch.2) |
| CNN | Nullpoint Error — *The real History of CNNs: From a cat's brain to ChatGPT's eyes* |
| VAE | Deepia — *Variational Autoencoders \| Generative AI Animated* |
| Transformer | 3Blue1Brown — *Attention in transformers, step-by-step* (ch.6) |
| Diffusion | Deepia — *Diffusion Models: DDPM* |
| **GAI Ethics & Post-Capitalism Economics** | **Thinkplicit — *The Rise of Collective Stupidity (End of Critical Thinking)*** |

QuickWins if short of time: 3B1B *Backpropagation, intuitively* (12 min), Infomity *CNN Animation* (8 min), Under The Hood *Transformer Architecture Explained* (20 min).

**Added in the last week before the exam (15 → 21 Sept):**
1. A **whole new topic row, "GAI Ethics & Post-Capitalism Economics"**, with *The Rise of Collective Stupidity* as **MustWatch**. (Boyle's "AI Slop" — a 2025 bonus — sits in this new row.)
2. The **CNN history video added and promoted to MustWatch**, displacing AlexNet to XtraInsight. It is also the **first link in `links.md`**, so it was pushed directly to the class.

**[INFERENCE — moderately strong]** Expect (a) a question or bonus on **AI slop / cognitive offloading / erosion of critical thinking / the economics of AI content**, and (b) the **cat-brain → CNN narrative** as the framing for the CNN question.

---

## 6. Predicted topics, ranked against the examiner's blueprint

Ranking rule: a topic scores high if it appears in **Marko's blueprint** *and* has independent corroboration from Nacaskul's lectures, the infographics or the past papers.

### HIGH confidence

**H1. The six foundational architectures — name them, place them, tell them apart.**
*Concepts:* **MLP, CNN, VAE, U-Net, Transformer, Diffusion**, arranged predictive → representation bridge → generative.
*Why:* **Blueprint line "Architectures: Multi-Layer Perceptron, Convolutional NN, VAE, .."**; the infographic *From Prediction to Generation — 6 Key Architectures* is exactly these six; `lec5.md` and `lec5/summarized_notes.md` both record "6 foundational architectures … be able to tell them apart"; syllabus sessions 4–6; CLO 2. Five independent sources, one of them the examiner.
**Confidence: high.** Likely 2–4 of the 10 questions.

**[MARKO — strengthened]** The Wolfram U CNN course *he assigned* has a whole design section covering **U-Net (lesson 17)** and **Transformers (lesson 19)**, plus **Cycle GAN (lesson 23)**. So four of the six architectures — CNN, U-Net, Transformer, and GAN alongside them — are backed by the examiner's own assigned material, not only by Nacaskul's lectures. **U-Net in particular is no longer a weak link in the list of six; treat it as properly examinable.**

**H2. How neural networks learn — gradient descent and backpropagation.**
*Concepts:* loss/cost function; **gradient = direction of steepest ascent**, so step against it; gradient descent is iterative and stepwise; **backprop = the specialisation of GDA for an MLP**, sending error backwards to tweak weights and biases; RMSE; ADAM = Adaptive Moment Estimation.
*Why:* **Blueprint line "Neural networks and how they learn"** — the examiner's own phrase; your "Exam sentence" note; two MustWatch 3B1B videos; two infographics.
**Confidence: high** for a *verbal* question, **very low** for a numerical one.

*The 3B1B chapter-2 summary sheet (`Screenshot … 13.46.59`) supplies the phrasing:* the **cost function** is the sum of squared differences between actual output activations and desired targets, **averaged over tens of thousands of examples**; the **gradient vector gives the direction of steepest downhill** once negated; the **ball rolling down a surface** is the intended mental model; trained nets reach **96–98% accuracy** but settle in *a* local minimum rather than a human-like abstraction.
*Corroborated by the Mathematica demo (`Screenshot … 13.49.12`, and your `6901009256_hwk3.nb`):* sliders fit a curve to scattered data with an **"RMSE Gauge — match blue to red"** and an **"RMSE Surface — aim to descend"**. That surface *is* the loss landscape; descending it by hand *is* gradient descent.

**H3. Predictive AI vs Generative AI.**
*Concepts:* Predictive is judged on **accuracy against verifiable ground truths**; generative on **plausible context relevance**. Predictive learns **x → y, then you crank the input**; generative learns **z → x, then you crank the PRNG**. Discriminative p(y|x) vs generative p(x) or p(x|c).
*Why:* 2024 Q2 verbatim; the spine of the 2025 paper; **two** dedicated infographics; blueprint line "Concepts and Principles (lecture 1)"; and §2.1's weighting pushes it further up.
**Confidence: high — the best-corroborated topic in the course.**

*Added from `S__79970354_0.jpg`, read at full resolution: "From Instinct to Imagination to Intelligence at Scale" (Nacaskul 2026).* A second, more quotable framing:
- **Predictive intelligence asks "What WILL be?"** — abstract regularities from observed examples, predict new outcomes. *From many throws, abstract the law of motion, predict where the ball will land.* Learns a mapping **x → y**.
- **Generative intelligence asks "What COULD be?"** — abstract the underlying structure or distribution of a domain, generate plausible instances never explicitly observed. *From many faces, abstract the essence of "faceness", and imagine a new face.* Learns a distribution **p(x)** or a transformation **x = G(z)**.
- **The common core, across animals, humans and AI:** specific experiences (particulars) → **abstraction** (find generality) → new situations (beyond what was seen).
- **A three-rung ladder:** **animal** — abstraction becomes *adaptive* (implicit, embodied, learn from experience to survive); **human** — *deliberate* (explicit, systematic, cumulative; we invent disciplines and teach them); **AI** — *scalable* (algorithmize the learning disciplines; extract generality from massive data).
- Its own takeaway, already exam-shaped: ***"AI did not invent abstraction; it industrialized it — turning learned generalities into algorithmic capability at scale."***

**H4. Unpacking G, A and I (and AI vs GAI vs AGI).**
*Concepts:*
- **I** — three senses: (1) "human-like", Turing's imitation game; (2) "advantageous information", as in Military/Business Intelligence; (3) "solution path to a problem", as in Computational/Digital Intelligence.
- **A** — (1) **Artifact/Artifice**, machine creation by humans; (2) **Artificial semblance**, when "seems about right" will do; (3) **Artful creation**, "truth manufactured".
- **G** — Generative ≠ Predictive (AI); Generation ≠ Recognition (Mind); Generate ⇔ {**Simulate** (given a distributional prior), **Connect** (graph/network adjacency), **Create** (integrated innovation)}.
- **GAI** = *Generative* AI; **AGI** = Artificial *General* Intelligence, explicitly out of scope.
*Why:* **Blueprint line "Concepts and Principles (lecture 1)"**; syllabus sessions 1–2; ~20 slides of the lecture-1 deck; the XMind roadmap's first branch; three separate notes of yours; plus §2.1's weighting.
**Confidence: high.**

**H5. The timelines.**
*Concepts:*
- **Language AI:** Turing 1950 *Computing Machinery and Intelligence* (imitation game) → symbolic/expert systems, retronymed **GOFAI**, AI winter → Mikolov 2013 **word2vec** → Vaswani 2017 **Attention Is All You Need** (Transformer, the "T" in GPT) → GPT-3 2020 → ChatGPT.
- **Vision AI:** Hubel & Wiesel, cat's striate cortex — neurons fired for **oriented lines**, not mice; lines → curves → shapes → objects → Fukushima 1980 **Neocognitron** → LeCun 1998 **LeNet/CNN** → **AlexNet 2012** (GPU + ImageNet, error 25% → single digit) → phones telling cats from dogs.
- **Third thread:** the **neurocomputational / pure-connectionist** one — McCulloch & Pitts 1943 → Rosenblatt 1958 Perceptron → Rumelhart/Hinton/Williams 1986 MLP + backprop. *The brain has no central processor; it is just neurons connected to neurons* ("Connectionist AI", flagged as an exam note).
*Why:* Syllabus session 3; `lec5.md` item 2; `lec5/summarized_notes.md` item 1; the XMind; the newly-promoted MustWatch CNN-history video; **two dedicated timeline infographics**; and Marko's own milestone slides (H5-extra).
**Confidence: high.**

**[EVIDENCE — both posters now read at full resolution] A correction to "two timelines".** Two of the infographics are large timeline posters, and both are legible in the zip:

- **`Master Timeline of AI - v3`** (sheet 7) — *"A Multi-Threaded Journey of Ideas, Models, Data and Hardware"*. A decade spine runs 1940s → 2020s, with **five parallel threads: Language & Knowledge · Vision & Perception · Neuroscience Inspiration · Algorithms & Learning (Models & Paradigms) · Hardware & Data Enablers.** A **PREDICTIVE AI** column runs down the left (models p(y|x); goal predict outcomes; classification, regression, ranking, forecasting), **GENERATIVE AI** down the right (models p(x); goal generate realistic, novel and useful content), with arrows carrying the story left to right. Entries now usable: **Wiener 1948** (cybernetics) · **Hebb 1949** · **Hodgkin–Huxley 1952** · **Rosenblatt 1956/58** · **ELIZA 1966 (Weizenbaum)** · **Hubel & Wiesel 1962–68** · **Minsky & Papert 1969** · **DENDRAL 1965 / MYCIN 1970s** (*IF fever AND rash THEN measles*, confidence 0.8) · **backprop 1986** · **LSTM 1997** · **SIFT 1999, SURF, HOG with SVMs** · **Bengio et al. 2003** word embeddings · **RBMs 2006** · **AlexNet 2012** · **Seq2Seq 2014, Attention 2015, BERT 2018** · and the hardware chain **transistors 1947 → ICs 1960s → microprocessor 1971 → workstations and GPUs 1980s → clusters and multi-core 2000s**. It also carries two side panels: **predictive vs generative as two complementary paradigms**, and **three complementary ways to build functional complexity** (enrich the links · many paths · depth).
- **`Fr Idea to Intelligence - v4`** (sheet 9) — one row per **concept lineage**: **Cybernetics · Artificial Intelligence · Machine Learning · Artificial Neural Network · Perceptron · Feedforward NN · Multi-Layer Perceptron · Deep Learning · Convolutional NN · Residual Network.** New names it supplies: **Samuel** (coined "machine learning", checkers from self-play) · **Widrow & Hoff 1965, ADALINE** · **Mitchell 1986**, ML as a discipline · **Hornik, Stinchcombe & White 1989**, multilayer feedforward networks are universal approximators · **LeCun et al. 1989**, backprop on handwritten zip codes · **Hinton, Osindero & Teh 2006**, deep belief nets and greedy layer-wise pretraining · **He et al. 2015**, ResNet. Its own takeaway: **"AI was envisioned in 1950 and formally launched in 1956."**

**[EVIDENCE]** So the lecturer's framing is richer than "two timelines": **two task timelines (Language, Vision) + a neuroscience-inspiration thread + an algorithms-and-learning thread + a hardware-and-data thread**, all feeding a left-to-right predictive → generative flow. **If asked how many timelines and what they are: two task timelines plus the neurocomputational one, with hardware and compute as the enabler** — AlexNet's GPUs are the concrete hook. If there is room, name all five threads; the five-thread version is what the lecturer drew.

**H5-extra. [MARKO] Marko's own milestone list** (`Gen AI Lit 02.pdf`, "Notes 2" and "Notes 3") — his selection, in his words: Turing 1950, the Turing test · **Dartmouth College summer workshop 1956** · Bernstein chess program (IBM) 1957; **Deep Blue beats the top human player 1997** · NLP: **METEO, Canada, 1981**, translating EN↔FR weather reports · Neural networks: **Perceptron 1943, implemented 1958, multi-layer 1980** · **Expert systems 1965, MYCIN 1975 (~65% accuracy), LISP and Prolog** · "The new AI" 2000–: **ImageNet 2010–, AlexNet, error rates 25% → single digit** · **Encoder–decoder 1997 and "deep learning" 2014, used by Google Translate** · **Generative AI 2018–, GPT**.
**[INFERENCE]** Dartmouth 1956, Deep Blue 1997, MYCIN and METEO appear *only* in Marko's deck, not in Nacaskul's. Since Marko sets the paper, **these are worth ten minutes** even though nothing else in the folder mentions them.

### MEDIUM confidence

**M1. Perceptron / MLP mechanics and vocabulary.**
Perceptron = mathematical model of a neuron: **weighted sum of inputs + bias → activation → fires past a threshold**. **Bias = decision threshold** (`Gen AI Lit 03.pdf` says exactly this). A single perceptron separates only **linearly separable** data — hence **XOR** as the classic counterexample, hence layers. MLP = feedforward layers stacked; no recurrence, no layer-skipping. Distinguish **parameter = coefficient = weight** (learned) from **input/variable** (given) and **hyperparameter** (set beforehand).
*Why:* Blueprint's "Neural networks and how they learn" and "MLP"; 2025 Q14; `lec5/summarized_notes.md` lists the parameter/weight/input vocabulary as examinable; **[MARKO]** `Gen AI Lit 03/06` drill XOR, AND/OR, sigmoid, ReLU and backprop in his own lab slides.
**Confidence: high-medium** (raised — squarely Marko's lab territory).

*Reinforced by the 3B1B chapter-1 summary sheet (`Screenshot … 13.47.12`):* the canonical worked example is a **28×28 pixel grid → 784 input neurons → two hidden layers of 16 → 10 output neurons ≈ 13,000 weights and biases** ("knobs and dials"); the progression **layer 1 pixels → layer 2 edges → layer 3 patterns → layer 4 identifies the digit** (the same hierarchical abstraction as Hubel & Wiesel); the neuron as **weighted sum + bias → activation**; **sigmoid squashes to (0,1)**; a layer step is **a matrix–vector multiplication**; **ReLU replaced sigmoid in modern nets because sigmoid suffers from vanishing gradients**.
**[INFERENCE]** This keeps Nacaskul's twice-flagged "count the parameters" alive in minimal form — e.g. *"how many parameters does a 2-input, 1-output perceptron have?"* → 2 weights + 1 bias = 3. That is arithmetic, not mathematics, so it survives "very little math". **Ten minutes, not an hour.**

**M2. CNN — why convolution, and the biological motivation.**
A convolutional layer is a set of **filters (convolution matrices) with trainable entries** slid across a grid of pixel values to detect visual patterns. **Weight sharing across positions gives translational covariance**; multiple scales handle objects at different distances; a **hierarchy of features** (lines → curves → shapes → objects) mirrors Hubel & Wiesel's cat. An MLP fed raw pixels fails because it cannot cope with the object moving or changing colour.
*Why:* Blueprint names "Convolutional NN"; 2025 Q15; syllabus session 4; the newly-promoted MustWatch video; and **[MARKO]** his own question *"Why did mobile phones start recognizing images in the 2010s?"*
**Confidence: high-medium.**

**M2b. [UPGRADED TWICE] The Wolfram U CNN course — concepts, not code.**
**[MARKO]** The blueprint says *"Some topics covered in the CNN course (assignment, Wolfram U), too."* He set that assignment and wrote the quizzes, so **the concepts he chose are the concepts he thinks matter**.

**[EVIDENCE — the course's own lesson list, read off the player, not inferred]** *Convolutional Neural Networks for Computer Vision*, Wolfram U, 2h30, Intermediate, four sections:

| § | Lessons |
|---|---|
| **1. Basics of Image Classification** (65 min) | 1 Overview and **Human Visual System** · 2 **Properties of Image Data** · 3 **Network Encoders** · 4 CIFAR-10; **Arrays and Tensors** · 5 **Convolution Layers** · 6 **ElementwiseLayer and ReLU** · 7 **Pooling and Softmax** · 8 **Loss Layer and Cross Entropy** · Quiz 1 |
| **2. Training and Analysis** (34 min) | 9 **Detecting and Avoiding Overfitting** · 10 **Dropout Layers** · 11 **Optical Illusions** · 12 **Receptive Fields and Sensitivity Maps** · 13 Network Gradient Analysis · 14 Image Feature Visualization · Quiz 2 |
| **3. Neural Net Design** (31 min) | 15 Inception Module · 16 ResNets · **17 U-Nets** · 18 EfficientNet · **19 Transformers** · Quiz 3 |
| **4. Image Applications** (16 min) | 20 Image Retrieval · 21 Object Detection · 22 Image Reconstruction · **23 Cycle GAN** · 24 Image Segmentation · Quiz 4 |

Course blurb, verbatim on what it stresses: *"the key building blocks of neural networks, discussing network layers, methods to enhance training and techniques to analyze networks through receptive fields, sensitivity maps and feature visualizations… innovative architectures like Inception, ResNets and U-Nets, as well as emerging designs such as visual transformers… applications in image classification, segmentation, reconstruction and generative modeling."*

**[INFERENCE]** Three things follow. (a) **Section 1 is a glossary of exactly the two-line definitions this format rewards** — encoder, tensor, convolution layer, ReLU, pooling, softmax, cross-entropy. It overlaps almost exactly with Marko's own lecture-6 recap quizzes on XOR, ReLU, sigmoid and backprop, which is strong convergent evidence. (b) **Lesson 1 "Human Visual System" and lesson 11 "Optical Illusions" tie the course straight into the Hubel & Wiesel vision thread** — so that thread now has support from *both* lecturers, and a question framed "what does the CNN borrow from biological vision?" is well-supported. (c) **Dropout and receptive fields** were absent from my earlier list and are now in it. The course is Mathematica-based, but "no code, very little math" still applies: **learn the concepts, not the Wolfram Language syntax.**

The quiz files `hk4.2/quiz1.md` and `quiz2.md` add the mechanism-level detail, stripped of the function names "no code" excludes:

*The guiding principles of natural vision that CNN design copies* — **translational and scale covariance · rotation invariance · symmetry transformations** (sometimes symmetry matters: "bp" flipped becomes "qd") · **inherent scale of image features · hierarchy of features** (leaves → trees → forest) · **colour correlation length** (neighbouring pixels have similar colours, and the correlation falls off with distance) · **motion detection** · **feature binding** (the brain binds spatially aligned features into objects).

*Mechanisms:*
- **Translational covariance** — the network processes every region with the **same learned filters**, so features are detected consistently wherever they appear.
- **Why multiple scales** — objects appear at different sizes depending on distance from the camera.
- **Flattening before the fully-connected layers** — converts the 3-D tensor from the convolutional stack into a 1-D vector.
- **Softmax** — converts raw activations into a **probability distribution over classes**.
- **Cross-entropy loss** — the distance between predicted class probabilities and the true (one-hot encoded) label.
- **ReLU / ramp after conv and linear layers** — introduces **nonlinearity**, letting the network learn complex decision boundaries.
- **Pooling** — reduces the size of the representation.
- **One-hot encoding** — exactly one slot on, the rest off.
- **Overfitting** — training error keeps falling while **validation error plateaus or rises**.
- **Data augmentation** (flips, small translations) — synthesises plausible variation, improving generalisation.
- **Input gradients** — the direction in pixel space that most changes the loss for a given class.

Two more from the lesson list, not covered by the quizzes:
- **Dropout** — randomly switching off a fraction of units during training so the network cannot rely on any single path; a regulariser against overfitting.
- **Receptive field** — the region of the input image that a given neuron can actually "see"; it widens with depth, which is how early layers detect edges and deep layers detect objects. **Sensitivity maps** show which input pixels most affect a given output.
- **Encoder** (in this course's sense) — the component that converts raw images into the numeric arrays/**tensors** a network can consume.

**Confidence: high-medium** for all of these in prose — raised again now that the syllabus of the assigned course is known; **very low** for `NetEncoder`, `NetDecoder`, `FlattenLayer`, `SoftmaxLayer` by name.

**M3. AE → VAE, latent space, encoder–decoder.**
**Autoencoder:** x → compressed code → x̂; the **bottleneck** forces dimensionality reduction. **VAE:** encodes to a *probabilistic* latent **distribution** (≈ Gaussian) rather than a point, so you can **sample z ~ N(0,I)** and decode to a *new* example — the predictive→generative transition. **Encoder–decoder** generally: compress to a latent, expand to the target. **U-Net** = encoder–decoder CNN with **skip connections** (hence the U) preserving fine spatial detail; image segmentation, medical imaging.
*Why:* **Blueprint names VAE explicitly**; **[MARKO]** his lecture-6 lab *was* an autoencoder and asked *"What is a latent space?"*; syllabus session 5; MustWatch video; dedicated infographic; the Cornell VAE interactive demo in `links.md`.
**Confidence: high-medium** (raised — Marko's own lab).

**M4. Transformer / attention / embedding.**
Words are not numbers, so **tokenize → embed** into vectors (GPT's embedding dimension ≈ **12288**). Embeddings carry **global-context meaning** and support word algebra (king − man + woman ≈ queen). **Attention** supplies **local-context interpretation**: it weighs which other tokens matter to each token, via **Q, K, V** matrices. **Positional encoding** supplies order. **GPT = Generative Pre-trained Transformer**, decoder-only; "pre-trained" = trained on internet-scale text beforehand. Generation = repeatedly predicting a contextually plausible next token.
*Why:* Syllabus session 6; CLO 2 names Transformer; **[MARKO]** his lecture-7 lab asked *"How does the transformer decide which word it generates next? What are the main parts of this transformer?"*; MustWatch 3B1B ch.6; *Attention Is All You Need* in the seminal-papers zip.
**Confidence: medium-high.**

**M5. ML vs AI — which contains which.**
ML is **larger** than AI in that its methods reach far outside AI (statistics, optimisation, any data-driven calibration). AI is **larger** than ML in that it also includes rule-based/symbolic GOFAI, search, planning and robotics involving no learning. Modern AI is *built by* ML because AI **cannot be pre-programmed with if-then constructs** — it must calibrate parameters from data. Three paradigms: **supervised, unsupervised, reinforcement** (Nacaskul's glossary adds **representation learning**).
*Why:* Your "Exam question" note verbatim; 2024 Q1 from the other side; and it is exactly the compositional kind of question Marko asks (§3).
**Confidence: medium-high.**

**M6. GOFAI / symbolic AI / expert systems.** *(raised from low)*
Rule-based formal logic and if-then trees; early success in **medical diagnosis** (**MYCIN 1975, ~65% accuracy**); **LISP and Prolog**; the AI winter; retronymed **GOFAI**. Know what a **retronym** is. Two postscripts from Nacaskul's slides: rule-based ML (decision trees, random forests, gradient-boosted trees) remains powerful and popular; and today's LLMs, though pure ANN, are increasingly tasked with logical *reasoning*.
*Why raised:* **[MARKO]** this is one of his own bonus questions verbatim (*"What is symbolic AI (not neural networks)? Expert systems?"*), and MYCIN/LISP/Prolog appear in **his** Notes-3 slide. Syllabus session 6.
**Confidence: medium-high.**

**M7. GAN and the adversarial origin of the "G".**
Two networks in a minimax game: a **generator** makes fakes, a **discriminator** judges real vs fake; both improve. **The irony:** the adversarial network was built to *sharpen a predictive/discriminative* model, and ended up being the star. Goodfellow et al. 2014. This is the *other* G in GAI — the one that gave "generative" its modern meaning.
*Why:* Flagged "Exam note" in `lec1/during.md`; the "Road to Generative AI (mid 2010s)" slide spells out the irony; **[MARKO]** his deck-02 recap lists *"Adversarial networks → deep fakes"*, right beside his *"Can we have generative AI without neural networks?"* bonus.
**Confidence: medium.**

**M8. Diffusion models.**
A **fixed forward process** adds Gaussian noise step by step; the model **learns to reverse it** (denoise). Generation starts from pure noise and iteratively denoises, optionally conditioned on a text embedding (Stable Diffusion, DALL·E).
*Why:* CLO 2 names diffusion alongside Transformers; MustWatch DDPM video; DDPM paper in the seminal zip; two infographics. **Counter-signal:** the blueprint stops at "VAE, ..".
**Confidence: medium.** Prepare two lines; don't over-invest.

**M9. Ethics / impact: AI slop, cognitive offloading, collective stupidity.**
Old web economics paid creators via clicks and ads; AI search answers without a visit, so the incentive to produce quality, trustworthy content collapses — degrading the very content supply that trained the models. Related: **cognitive offloading** and dependency; **"Artificial Literacy"** = leaning on ChatGPT to write your essay, pitch your project and read your papers, i.e. the failure mode of AI Literacy.
*Why:* A whole new playlist row added 21 Sept with a MustWatch; 2025 Bonus 1 was exactly this; ~10% of the syllabus is "How GAI Impacts My Generation"; the blueprint points at the playlist.
**Confidence: medium** as a regular question, **higher** as a bonus.

### LOWER confidence — and what got demoted

**D1. [DEMOTED] BAScii meta-questions: "Integrated Innovation", the "and", "bare minimum AI Literacy", "chatbot as AI Application".**
Three of the twenty visible 2025 questions, but this is **Nacaskul's personal vocabulary and it appears nowhere in Marko's blueprint**. Previously "lower confidence insurance"; now **low**.
*Counterweight, and the reason not to drop it entirely:* one of the new thumbnails (`Screenshot … 13.46.41`) is an infographic Nacaskul made about the **Krueger 100XR** counter-drone interceptor, subtitled ***"Integrated Innovation: Lessons from the Krueger 100XR"***, whose headline reads — legibly — ***"Integrated Innovation ≠ the integration of impressive features. Integrated Innovation = designing an innovation to work within an integrated reality."*** Its six lesson titles are legible (bodies not): *1 Not a Feature Bin, but a System Architecture · 2 Technology Layering Creates Real Capability · 3 Startup Situational Awareness is Essential · 4 Modular Design & Future-Proofing · 5 OODA at Two Levels · 6 Engineered Interoperability Multiplies Impact*.
**→ Learn the one-line definition above. Thirty seconds, and it directly answers 2025 Q11. Learn nothing else about drones.**

**D2. [DEMOTED as a pattern] The five elements of the Intelligent Being.**
Five of twenty visible 2025 questions, but it is Nacaskul's framework and the one-per-question format is his. **Low-medium as a pattern.** However **"valenced feeling and graded sensation" stays medium-high on its own merits**, being double-flagged: 2025 Q17 *and* an explicit "Exam notice" in your lecture-1 notes. Learn all five (they are short); expect at most one.

**D3. [DEMOTED] "Unpack the acronym, explain the letter."** Know what every acronym stands for — free marks in any format — but do not expect Marko to phrase a question that way. **Content stays high; the format expectation drops to medium.**

**M10. [PROMOTED from L1 — the earlier reasoning here was wrong] The Five Principles of Modern Deep Learning**, and the seminal-papers list that comes with them.
**Representation → Relation → Parameterization → Optimization → Architecture** — represent, relate, parameterize, learn, **orchestrate**. Closing line: *"Not ever more complicated primitive mathematics, but increasingly ingenious **orchestration** of representations, relationships, parameters and learning."*

**The correction.** This sheet was previously ranked lowest on the grounds that *"no lecture note of yours mentions it, so it may not have been taught"*. **That reasoning was wrong and is withdrawn.** The examiner's own blueprint slide says ***"Study the infographics sheets carefully!"*** — so a dedicated sheet from the lecturer is *first-party evidence of what he considers examinable*, and the absence of a lecture note is not evidence against it. Being the **newest** sheet (15 Sept, ten days before the exam) is a point in its favour, not against.

**Why it now sits at medium.** Three things:
1. The five words are a complete, quotable answer to *"what does modern deep learning actually consist of?"* — the exact shape a two-line paper rewards.
2. Each principle is a ready-made "which principle is this?" question, and the answers are single words.
3. The sheet carries a **timeline of 18 seminal papers with author, year and a one-line contribution each**, colour-coded to the five principles. That is the single densest block of names-and-dates material in the whole course, and it is **the lecturer's own selection**. It is now in `content/lectures/timelines.md` as a table.

**A useful self-test the sheet supports:** *Adam serves optimization alone; the Transformer serves all five, because its contribution* is *the orchestration.*
**Confidence: medium.** Half an hour, not five minutes.

**M11. [NEW] The degradation problem and the residual connection.**
By **early 2015** deeper had stopped meaning better: beyond **20–30 layers** accuracy stalled and then reversed — a **30-layer model at 16.59% error against a 14-layer model's 13.34%**, *on the training set*. So **not overfitting**: an **optimisation** failure. The confounding part is that the solution provably existed — add 16 **identity** pass-through layers to a trained 14-layer model and it must do at least as well — and gradient descent could not find it. **The fix: learn the residual F(x) = H(x) − x and add the input back, so the block outputs H(x) = F(x) + x.** Refine rather than re-create; identity is available by construction. Depth then scaled to **152 layers and beyond**, and the residual block became the **backbone** under AlphaGo, AlphaFold and the LLMs. He, Zhang, Ren & Sun, Microsoft Research, Dec 2015 — **the most cited paper of the 21st century**.
*Why:* a **whole dedicated infographic** (sheet 10) plus a second one covering it from the functional-complexity angle (sheet 12), and **[MARKO]** **ResNets are lesson 16 of the Wolfram U CNN course he assigned**. Two independent sources, one of them the examiner's own assigned material.
**Confidence: medium.** The two-line answer: *deeper networks fitted worse, which was an optimisation failure rather than overfitting; adding the input back so each block learns a correction made arbitrary depth trainable.*

**M12. [NEW] Universal approximation, stated properly — and why depth exists anyway.**
**One hidden layer** with finitely many neurons approximates any continuous function on a closed bounded region to any tolerance (**Cybenko 1989**, sigmoid; Funahashi 1989; Hornik 1991). The examinable consequence is the one-liner: ***expressivity is abundant, trainability is scarce*** — the obstacle is never the function class, it is **optimisation and inductive bias**. Depth is not needed for power; it is needed for **parameter efficiency and optimisability**.
*Why:* sheet 12 is devoted to it, sheet 8 lists Cybenko among the 18 seminal papers, and it sharpens an answer the bank already had.
**Confidence: medium.** It also gives the cleanest available answer to *"why go deep if one layer is enough?"*

**L2. Flow matching.** *(previously listed as not examinable — softened)*
Defines a **probability path** from a simple distribution to the data distribution and learns the **vector field** that transports samples along it; generation = start from noise and **solve the learned ODE**. It is the third column of the *VAE → Diffusion → Flow Matching* infographic, which Marko told you to study, so it earns exactly one line: ***VAE, diffusion and flow matching are three different learning games with the same goal — a learnable generator mapping a simple distribution to the data distribution.*** That sentence covers it. **Do not study it further.**

**L3. Prompt engineering / RAG / agentic workflows.** Syllabus session 7, but `Gen AI Lit 07` states these are the *post*-midterm hands-on classes, mapped to the **Term Project**. **[INFERENCE] Probably not on this paper.** One line of insurance: *RAG = retrieve relevant documents first, then generate grounded on them, for factual verifiability.* The XMind's four-way split is a cheap bonus: **Chatbot "talks with you" · RAG "searches for you" · Agent "does things for you" · Wiki "builds knowledge with you".**

---

## 7. Definitions and distinctions worth memorising

Short-answer papers live on these. Two lines each. **This table is the highest-yield hour in this document.**

| Pair | The distinction |
|---|---|
| **GAI vs AGI** | GAI = *Generative* AI (creates new content). AGI = Artificial *General* Intelligence (human-level breadth) — explicitly beyond this course. |
| **Predictive vs Generative AI** | Predictive: **accuracy against verifiable ground truths**; asks "what WILL be?"; learn x→y, then crank the input. Generative: **plausible context relevance**; asks "what COULD be?"; learn z→x, then crank the PRNG. |
| **Generation vs Recognition** | Recognition is an **interpolative** map within known territory. Generation is an **extrapolative** map into unexplored territory. |
| **BNN vs ANN** | Biological Neural Network = the actual brain. Artificial Neural Network = a mathematical model of it; every modern AI is built from ANNs. |
| **Perceptron vs MLP** | Perceptron = one artificial neuron; separates only **linearly separable** data (XOR defeats it). MLP = feedforward layers of perceptrons stacked, each layer's output the next layer's input. |
| **FFN vs RNN** | Feedforward: signal travels one way only. Recurrent: the signal loops back. MLP is feedforward, no recurrence, no layer-skipping. |
| **Weight / coefficient / parameter vs input vs hyperparameter** | Weights, coefficients and parameters are the same thing — the numbers **learned** from data. Inputs are the x variables you supply. Hyperparameters are set before training. |
| **Gradient descent vs backpropagation** | GDA is the general "step downhill" optimiser. Backprop is GDA **specialised to an MLP**, sending error backwards through layers to update weights and biases. |
| **Loss function vs RMSE** | Loss L is the general "how wrong am I" number. RMSE (Root Mean Squared Error) is one specific loss. |
| **Variable selection vs feature engineering** | Variable selection = choosing which x's to use. Feature engineering = **building a new input from existing ones** by a logical function (e.g. x₁/x₃). |
| **AE vs VAE** | AE: deterministic code, reconstruction only. VAE: a probabilistic latent **distribution** you can **sample** from — hence generation. |
| **VAE vs Diffusion** | VAE **learns both** directions (encoder to the simple distribution, decoder back). Diffusion **prescribes** the forward noising and only **learns** the reverse denoising. |
| **Encoder–decoder vs U-Net** | Encoder–decoder compresses then expands. U-Net is an encoder–decoder CNN with **skip connections** (the "U") carrying fine spatial detail across the bottleneck. |
| **Embedding vs attention** | Embedding gives a word its **global-context** meaning as a vector. Attention gives it a **local-context** interpretation relative to the other words in *this* sentence. |
| **GOFAI vs modern AI** | GOFAI = rule-based symbolic logic / expert systems, hand-written if-then, no learning. Modern AI = ANN whose parameters are calibrated from data by ML. "GOFAI" is a **retronym**. |
| **Retronym** | A new name for an older thing, coined because a newer thing took the original name (expert systems → "GOFAI"; AI → "predictive AI"). |
| **ML vs AI** | Neither contains the other. ML reaches outside AI (any data-driven calibration); AI reaches outside ML (symbolic, search, planning). |
| **AI vs BI vs CI** | AI = approximating the interactive/behavioural intelligence of a species. BI = extracted information of tactical value. CI = optimal solution/method for a complex problem. |
| **Supervised vs unsupervised vs reinforcement** | Supervised: labelled targets, error correction. Unsupervised: structure only, Hebbian/clustering. Reinforcement: policy optimisation against reward. |
| **Apply/Applied vs Application/App** | "Applied" = working under **real-world constraints**, not simplified. "Application" ⇔ {Modification/Optimisation, Addition/Enhancement/Extension, Solution/Direction/Utilisation}. |
| **AI Literacy vs "Artificial Literacy"** | AI Literacy = can code, can safeguard, can explore/exploit. Artificial Literacy = outsourcing your essay, pitch and reading to ChatGPT — the failure mode. |
| **Sigmoid vs ReLU (ramp)** | Sigmoid squashes any real input into (0,1); suffers from **vanishing gradients** in deep nets. ReLU returns the argument if positive, else 0 — cheap nonlinearity, now the default. |
| **Softmax vs sigmoid** | Softmax turns a vector of raw activations into a **probability distribution across classes** (sums to 1). Sigmoid squashes a single value. |
| **Bias vs weight** | Weights scale each input. The **bias shifts the decision threshold** (`Gen AI Lit 03`: "Bias = decision threshold"). |
| **Covariance vs invariance (vision)** | **Translational covariance**: the same filters are applied everywhere, so a feature is detected wherever it appears. **Rotation invariance**: the label does not change when the input is rotated — though **symmetry sometimes matters** ("bp" flipped becomes "qd"). |
| **Overfitting vs underfitting** | Overfitting: training error keeps falling while **validation error plateaus or rises**. Underfitting: both stay high — the model lacks capacity. |
| **Dropout vs data augmentation** | Both fight overfitting. Dropout randomly switches off units during training so no single path is relied on. Augmentation (flips, small translations) manufactures extra plausible training variation. |
| **Receptive field vs sensitivity map** | The receptive field is the region of the input a given neuron can see; it widens with depth, so early layers catch edges and deep layers catch objects. A sensitivity map shows which input pixels most affect a given output. |
| **Encoder vs embedding vs tensor** | An encoder converts raw data (pixels, words) into numbers a network can consume. An embedding is the learned vector meaning of an item. A tensor is just the multi-dimensional array those numbers live in. |
| **Autoencoder vs general encoder–decoder** | AE: the target **is** the input (y = x), usually with a **bottleneck**. General encoder–decoder: **y ≠ x**, and **no bottleneck is required** — translation, summarisation, segmentation, captioning. |
| **Fixed bottleneck vs attention** | Before attention the decoder read **one fixed-length vector** for the whole input, and long inputs lost detail. Attention lets it **weigh a whole sequence of encoder states per output element**. |
| **Degradation vs overfitting** | Overfitting: training error falls while validation error rises. Degradation: the **deeper** model is worse **on the training set too** — an optimisation failure, not a capacity one. |
| **Plain depth vs residual depth** | A plain deep net must push signal and gradient through every layer. A **residual** block adds the input back, so it learns a **correction** and identity is free — which changes **trainability**, not the function class. |
| **Expressivity vs trainability** | Universal approximation makes expressivity **abundant** (one hidden layer suffices). What is **scarce** is finding the parameters — hence architecture and optimisation choices. |
| **Attention vs the feed-forward block** | Attention **moves information between tokens**. The FFN is a per-token MLP acting as a **key–value memory** that adds stored facts — and holds about **two thirds** of the parameters. |
| **TO vs FRO mapping** | TO maps **data → simple distribution**; FRO maps **simple → data**. **Generation always uses FRO**; TO exists only to create a supervised task with known targets. *TO sets the homework; FRO learns to solve it backwards.* |
| **Prompt vs PRNG in conditional generation** | The **prompt is WHERE to navigate** — it steers the generator to a region of the possibility space. The **PRNG is WHICH possibility** — it picks one sample there, which is why the same prompt gives different images. |
| **Backpropagation vs the gradient-descent variants** | Batch, SGD, mini-batch, momentum, RMSProp and Adam are variants of the **optimiser**. Backprop is not one of them — it is **how the gradient they all consume is computed**. |
| **Werbos 1974 vs Rumelhart 1986** | Werbos **published** backpropagation in a 1974 dissertation; the 1986 Nature paper is where it **became known**. 1986 is the popularisation, not the invention. |
| **False positive vs false negative** | Both are misclassifications, with different costs. A **civil court** minimises **false positives** (better to acquit the guilty than convict the innocent); a **military command** minimises **false negatives**. Moving the threshold trades one for the other; no line removes both. |
| **Step vs sigmoid** | A step fires or does not, and has **no useful derivative** — so nothing can be learned by gradient descent. The sigmoid smooths the edge, making the model **differentiable** and its outputs graded rather than rigid. |
| **What counts as a parameter** | **Weights and biases** count — training changes them. **Inputs** are data you supply; **hyperparameters** are set before training. Neither counts. *Two inputs, one output → three parameters.* |
| **Embedding = vocabulary, attention = grammar** | The embedding gives each word a stable meaning (*cat is closer to dog than to refrigerator*). Attention supplies which word bears on which **in this sentence** — the harder half, because a sentence is not recoverable from distances between isolated words. |
| **AI vs robot vs embodied AI** | AI is intelligent but cannot act. A robot can act but is not intelligent. **Embodied AI** is the join: a physical form able to experience and act in its environment. |
| **Predictive vs generative, the one-liner** | *"Approximating reality is predictive AI. Taking certain things and running with it is generative AI. All fiction is realistic but not real."* |

**Names and dates worth having ready** (the course is attribution-heavy on both sides):
**Cauchy 1847** (gradient method, ~140 years early) · Turing 1950 (imitation game) · **Dartmouth workshop 1956** · **Wiener 1948** (cybernetics) · **Hebb 1949** · **Hodgkin & Huxley 1952** · **ELIZA 1966** (Weizenbaum) · McCulloch & Pitts 1943 · Rosenblatt 1958 (Perceptron / Mark I) · Hubel & Wiesel 1959/1962 (cat's striate cortex) · **Expert systems 1965, MYCIN 1975** · Fukushima 1980 (Neocognitron) · **METEO 1981** · Rumelhart, Hinton & Williams 1986 (backprop / MLP) · **Deep Blue 1997** · LeCun 1998 (LeNet/CNN) · Hinton & Salakhutdinov 2006 (autoencoder / dimensionality reduction) · ImageNet 2010– · Krizhevsky, Sutskever & Hinton 2012 (AlexNet) · Mikolov 2013 (word2vec) · Kingma & Welling 2013 (VAE) · Goodfellow 2014 (GAN) · Ronneberger 2015 (U-Net) · Vaswani 2017 (Transformer) · GPT and **BERT** 2018– · Ho, Jain & Abbeel 2020 (DDPM). Add from the seminal-papers sheet: **Cybenko 1989** (universal approximation) · **Pao 1995** (functional-link nets) · **LSTM 1997** (Hochreiter & Schmidhuber) · **Seq2Seq 2014** (Sutskever, Vinyals & Le) · **Attention 2014** (Bahdanau, Cho & Bengio) · **Adam 2014/15** (Kingma & Ba) · **ResNet 2015** (He, Zhang, Ren & Sun) · **ViT 2020/21** (Dosovitskiy). Also **Fei-Fei Li** (ImageNet) and **Geoffrey Hinton**.
*(The bolded ones come from **Marko's** slides specifically.)*

**[Note on the cat's date]** `lec5.md` records "1958"; Nacaskul's slides say **1959** in one place and **1962** in another. Say "around 1959 / late 1950s" — the finding matters far more than the year.

**The six papers distributed** (`(seminal_foundational_papers)….zip`): Bengio, LeCun & Hinton (2021) *Deep Learning for AI*; LeCun, Bengio & Hinton (2015) *Deep Learning*; Krizhevsky et al. (2012) *AlexNet*; Vaswani et al. (2017) *Attention Is All You Need*; Kingma & Welling (2019) *An Introduction to VAE*; Ho, Jain & Abbeel (2020) *DDPM*.
**[INFERENCE]** CNN + Transformer + VAE + Diffusion + two overviews. Corroborates H1/M3/M4/M8; says nothing about RAG or agents.

---

## 8. The most likely questions, with model two-line answers

Phrased the way **Marko** asks things where I can (mechanism-seeking, plain English), rather than in Nacaskul's acronym-unpacking style.

**1. Explain the difference between Predictive AI and Generative AI, and give one example of each.**
> Predictive AI maps input to output and is judged on accuracy against verifiable ground truths — e.g. an MLP classifier. Generative AI models the data distribution and is judged on plausible context relevance — e.g. GPT.

**2. What are the main parts of a multilayer perceptron, and what does each neuron actually compute?**
> Input layer, one or more hidden layers, output layer, connected feedforward with no recurrence. Each neuron takes a weighted sum of its inputs plus a bias, passes it through an activation function, and outputs the result to the next layer.

**3. Why can a convolutional network recognise an object wherever it appears in the image, when a plain MLP cannot?**
> A CNN slides the same learned filters over every region (weight sharing), so a feature is detected consistently regardless of position — translational covariance. An MLP gives every pixel position its own weight, so a shifted object looks entirely different to it.

**4. Name the six foundational architectures covered in this course, from predictive to generative.**
> MLP, CNN, VAE, U-Net, Transformer, Diffusion. MLP and CNN predict; VAE and U-Net form the representation bridge; Transformer and Diffusion generate.

**5. State the relationship between gradient descent and backpropagation.**
> Backpropagation is a specialisation of the gradient descent algorithm, applied to minimise prediction error in a multilayer perceptron by tweaking its weights and biases. The gradient gives the direction of steepest ascent; the algorithm steps against it, repeatedly.

**6. What is a loss function, and what does the network do with it?**
> It returns a single number measuring how far the network's output is from the ground truth — e.g. RMSE, or cross-entropy for classification. Training minimises it by gradient descent, adjusting weights and biases step by step.

**7. What did Hubel and Wiesel observe in the cat's visual cortex, and why does it matter for AI?**
> Neurons in the cat's striate cortex fired for oriented lines, not for whole objects; line detectors combine into curve detectors and upward into shapes. That hierarchy of features is the blueprint for the Neocognitron and the CNN.

**8. Why did mobile phones start recognising images in the 2010s?**
> Three things converged: large labelled datasets (ImageNet), GPUs fast enough for the matrix multiplications, and deep convolutional architectures. AlexNet in 2012 cut ImageNet error from about 25% to single digits.

**9. Give the three distinct senses of the "I" in "AI".**
> (1) "Human-like", as in Turing's imitation game; (2) "advantageous information", as in Military or Business Intelligence; (3) "solution path to a problem", as in Computational or Digital Intelligence.

**10. Give the three distinct senses of the "A" in "AI".**
> (1) Artifact/Artifice — machine creation by humans; (2) Artificial semblance — when "seems about right" will do; (3) Artful creation — "truth manufactured" by whatever is trending.

**11. "Generate" can mean three things in this course. Name them.**
> Simulate (given a distributional prior), Connect (graph/network adjacency), and Create (integrated innovation). Generative is opposed to Predictive in AI, and to Recognition in the mind.

**12. In what way is machine learning larger than AI, and in what way is AI larger than ML?**
> ML is larger in that its methods calibrate parameters from data far outside AI — statistics, optimisation, forecasting. AI is larger in that it also includes rule-based GOFAI, search and planning that involve no learning at all.

**13. What is symbolic AI, and what happened to it?**
> Rule-based formal logic over symbols — hand-written if-then expert systems such as MYCIN (1975), written in LISP or Prolog, with no learning from data. It stalled, brought on the AI winter, and was retronymed GOFAI once neural approaches took the name.

**14. What is a latent space, and what makes a VAE generative when a plain autoencoder is not?**
> The latent space is the compressed internal representation at the autoencoder's bottleneck. A plain AE maps each input to one fixed code; a VAE encodes to a probability distribution, so you can sample a new point and decode it into a new, realistic, previously unseen example.

**15. How does a transformer decide which word to generate next?**
> It embeds and positionally encodes the tokens, then uses attention (query, key, value) to weigh how much every other token should inform each one. The final layer produces a probability distribution over the vocabulary, and the next token is sampled from it.

**16. How does a diffusion model generate an image?**
> A fixed forward process gradually corrupts data with Gaussian noise; the network learns to reverse it. Generation starts from pure noise and iteratively denoises, optionally conditioned on a text embedding.

**17. What is a GAN, and what is ironic about its place in the history of generative AI?**
> A Generative Adversarial Network pits a generator making fakes against a discriminator judging them, improving both. The irony: the generator was built to sharpen a *predictive* discriminator, and ended up being the star.

**18. What is overfitting, and how would you detect it?**
> The model fits the training data ever better while failing to generalise. Symptom: training error keeps falling while validation error plateaus or starts rising.

**19. What is a receptive field, and why does it widen with depth?**
> The receptive field is the region of the input image a given neuron can actually see. Each layer pools information from the one below, so it widens with depth — which is why early layers detect edges and deep layers detect whole objects.

**20. Name two ways of preventing overfitting and say how each works.**
> Dropout randomly switches off a fraction of units during training, so the network cannot rely on any single path. Data augmentation (flips, small translations) manufactures extra plausible training variation, forcing the model to generalise.

**21. Can we have generative AI without neural networks?**
> Yes in principle — a Markov chain generates text from next-token statistics, and any sampler from a fitted distribution is generative. But without learned representations it has no context, so the output is locally plausible and globally incoherent.

**22. What is meant by "AI slop", and why is it an economic problem and not just a quality one?**
> Low-value AI-generated content flooding the web. The old model paid creators through clicks; AI search now answers without a visit, so the incentive to produce high-quality trustworthy content evaporates and the reliability of the whole content supply declines.

---

## 9. What is probably NOT examinable

- **Numerical gradient descent, partial derivatives, computing ∇f by hand.** 2025 spent five questions on this, but **[MARKO]** "No code, very little math", and two ruled lines leave no room to show work. **Skip the arithmetic; keep the verbal definition.**
- **Python, Keras, PyTorch, Colab code** from the lab decks (`layers.Conv2D`, `d_model=128`, `<UNK>` tokens). "No code."
- **Wolfram Language / Mathematica function names** (`NetEncoder`, `NetDecoder`, `FlattenLayer`, `NetPortGradient`, `Manipulate`). The **concepts behind them are examinable** (M2b); the names are not.
- **Prompt engineering, RAG, agentic workflows, knowledge graphs.** Session 7, pushed past the midterm and mapped to the Term Project. One line of insurance only (L3).
- **ResNet, LSTM, Vision Transformer, CLIP, normalizing flows, Stable Diffusion internals.** Present in infographics and the XMind, never in a lecture note of yours, absent from the blueprint. Recognise the names; do not study them.
- **Flow matching beyond one sentence.** See L2.
- **Anything from sessions 8–15** (the Darkside lecture, tools/platform workshops, project presentations).
- **The Krueger 100XR drone specifics**, quantum computing, and the finance material in Nacaskul's prologue. Decoration — except the one-line "Integrated Innovation" definition in D1.
- **[NOW LOWER]** The BAScii meta-questions (2025 Q11–Q13) — D1. Memorise the three one-liners as insurance; do not treat them as a study topic.

---

## 10. Revision plan — Wednesday evening, Thursday, Friday morning

Exam **Friday 25 Sept, 13:00**. You have **Wednesday evening, all of Thursday, and Friday until about 11:30**. Ordered so that **if you run out of time you lose the least**: everything in Wednesday's block is high-confidence and high-yield.

### Wednesday evening (~2½ hours) — the irreducible core
1. **§7 table of distinctions.** Read it twice, then cover the right-hand column and recite. **(75 min. The single highest-yield block in this plan.)**
2. **The six architectures, one line each, in order** (H1). (20 min)
3. **The backprop "exam sentence", verbatim**, plus gradient-descent-in-one-sentence (H2). (10 min)
4. **Predictive vs Generative, both framings**: ground truth vs plausible relevance, and "what WILL be" vs "what COULD be" (H3). (15 min)
5. **The three meanings of I, of A, and of Generate** (H4). (20 min)
6. Stop. Do not start anything new.

### Thursday morning (~2½ hours) — mechanisms
7. **How neural networks learn, end to end**: perceptron → weighted sum + bias → activation → layers → loss → gradient descent → backprop. Say it aloud once without notes. (40 min)
8. **CNN**: the cat, the hierarchy of features, weight sharing and translational covariance, why an MLP fails on pixels (M2). (30 min)
9. **The Wolfram CNN concepts in prose** (M2b) — the vision-principles list, plus the Section-1 glossary the examiner assigned: **encoder, tensor, convolution layer, ReLU, pooling, softmax, cross-entropy**, then **overfitting, dropout, receptive fields, data augmentation**. These are the cheapest marks in the paper: each is a clean two-line definition. **(40 min — worth the extra ten.)**
10. **AE → VAE → latent space → encoder–decoder → U-Net** (M3). (30 min)

### Thursday afternoon (~2½ hours) — history and the remaining architectures
11. **The timelines** as lists of names + years: Language, Vision, neurocomputational — plus **Marko's own milestone list** (H5-extra: Dartmouth 1956, MYCIN 1975, METEO 1981, Deep Blue 1997). (45 min)
12. **Transformer / attention / embedding** (M4), then **GAN** (M7) and **diffusion** (M8), two lines each. (45 min)
13. **GOFAI / symbolic AI / expert systems / retronym** (M6) and **ML vs AI** (M5). (30 min)
14. **The five Embodied-AI elements**, especially **valenced** (D2). (15 min)

### Thursday evening (~1½ hours) — the assigned media
15. The **MustWatch videos you have not seen**, in this order: the **CNN history** one and **The Rise of Collective Stupidity** (both added days before the exam), then whichever of VAE / attention / DDPM you are least sure of. Watch at 1.5×. (60 min)
16. **§8 model answers** — read all 22 once, out loud. (20 min)

### Friday morning (~2 hours, finishing by 11:30) — consolidation only, no new material
17. **Re-read the core infographics** (§11). They are the lecturer's own answer key and they re-encode everything above in a form that survives exam nerves. (45 min)
18. **§7 table again, cover-and-recite** — only the rows you got wrong on Wednesday. (30 min)
19. **§8 model answers, cover-and-recite** — the ten you feel least secure about. (30 min)
20. **Stop by 11:30.** Eat. Open nothing new after that: nothing learned in the last ninety minutes will survive, and it will crowd out what you already know.

### Exam-room technique
Two ruled lines is **25–35 words**. Lead with the **term**, then the **mechanism** — that is how the model answers are written and it is what fits. One definition plus one distinguishing clause, or one mechanism plus one consequence. No introductory clause. If you know more than fits, **the extra will not be read**, so spend the words on the distinguishing detail rather than the setup. Write legibly; Marko says so twice on the sheet.

---

## 11. Study these infographics directly

**[MARKO]** *"Study the infographics sheets carefully!"* — his instruction, so these count as
assigned material, and a sheet on a topic no lecture note mentions is **still** evidence about
the paper. That is the standard applied here.

**[RESOLVED]** All sheets are available at **full resolution** inside
`(GAI-infographics) GAI Literacy _amp_ Appl.zip` — **15 content sheets at 1024×1536 or
1055×1491, fully legible**, plus 11 course-icon files that carry no content. Everything in
this document that was previously marked **[LOW-RES]** from a thumbnail has now been read from
the original. **No re-export of anything is needed**, and the earlier request for one is
withdrawn (§12).

All 15 have been read and folded into `content/`. Question sources cite them as
`{"deck": "infographics", "slide": N}` with **N as the sheet number below**, so a claim can be
traced back by hand:

| N | Sheet | What it gives you |
|---|---|---|
| 1 | *Generative AI* definition material | p(y\|x) vs p(x) |
| 2 | **From Predictive AI to Generative AI — v2** | "Learn x→y then crank the input" vs "learn z→x then crank the PRNG"; the common engine (universal approximation → gradient-based learning → learned mapping); the **TO/FRO** table; **conditional generation: the prompt is WHERE to navigate, the PRNG is WHICH possibility**; the manifold caveat |
| 3 | **From Prediction to Generation — 6 Key Architectures — v3** | The six, with key idea, typical tasks and a one-phrase summary each. **H1's answer key.** |
| 4 | **From VAE to Diffusion Models to Flow Matching — v2** | The side-by-side table; ELBO; which mapping each family fixes and which it learns |
| 5 | **From Instinct to Imagination to Intelligence at Scale** | "What WILL be?" vs "What COULD be?"; the animal → human → AI abstraction ladder; *"AI did not invent abstraction; it industrialized it."* |
| 6 | **The two 3Blue1Brown chapter summaries** | 784 → 16 → 16 → 10 ≈ **13,000** weights and biases; the **cost function** as the averaged sum of squared differences; the ball-on-a-surface picture; 96–98% accuracy and the local-minimum caveat |
| 7 | **Master Timeline of AI — v3** | **Five threads**: Language & Knowledge · Vision & Perception · Neuroscience Inspiration · Algorithms & Learning · **Hardware & Data Enablers**. ELIZA 1966, Hebb 1949, Hodgkin–Huxley 1952, Wiener 1948, DENDRAL 1965, the hardware chain, SVMs and SIFT, Bengio 2003, BERT 2018 |
| 8 | **Five Principles in Modern Deep Learning — v2** | The five principles **and the 18 seminal papers** with author, year and contribution. See **M10** |
| 9 | **From Idea to Intelligence — v4** | Ten **concept lineages** (cybernetics, AI, ML, ANN, perceptron, FNN, MLP, deep learning, CNN, ResNet) and the dependency chain showing how each step removed the previous obstacle |
| 10 | **The Story of ResNets — v2** | The **degradation problem**, why it is not overfitting, the identity-mapping argument, H(x) = F(x) + x, and the depth race 8 → 19/22 → 30 → 152 |
| 11 | **The Encoder–Decoder Architecture — A Timeline & Genealogy** | The one-line definition; **AE vs general encoder–decoder** (y = x and a bottleneck, versus neither); the **denoising** autoencoder; **attention as the end of the fixed bottleneck**; cross-attention |
| 12 | **Functional Complexity in NN** | The **UAT stated precisely — one hidden layer**; *"expressivity is abundant, trainability is scarce"*; the **three ways** to build functional complexity; ResNets as ensembles of short paths |
| 13 | **From Calculus to Backprop** | derivative → gradient → gradient descent → backpropagation; the learning rate η; forward vs backward pass; **Werbos 1974**; the non-convex loss surface |
| 14 | **Decoding the MLP/FFN Backbone (Gemini)** | The FFN inside a Transformer block; **FFNs as key–value memories storing facts**; layer specialisation; **superposition** and sparse autoencoders; FFN ≈ two thirds of the parameters |
| 15 | **Generative AI — Literacy Roadmap — v2** | The four stages; the **three timeline families**; the **integrated-autonomy ladder** (software → operating system → hardware); **chatbot / RAG / agent / wiki** |

**If you only read four:** 3, 2, 4 and 8. **If you have another twenty minutes:** 10 and 12,
because the degradation problem and "expressivity is abundant, trainability is scarce" are the
two best two-line answers on any of them that the rest of this document did not already give
you.

**Lower priority, but now covered in `content/`:** 13 (it mostly restates H2), 14 (advanced, and
the only sheet that goes beyond the syllabus), 15 (structure rather than mechanism).

**The XMind** (`(folded) Generative AI - Literacy Roadmap=160.xmind`, dated 2026.09.14) is the
same material as sheet 15 in outline form: *Words Matter* → *Timelines* → *HIWUTH* →
*Application*. **The four-way application split is quotable and costs 60 seconds.**

---

## 12. The 14 new images — what they are

**[RESOLVED 23 Sept]** Everything below that was marked **[LOW-RES]** has since been read from
the full-resolution originals in the infographics zip. **Nothing needs re-exporting.** The
findings are folded into §11 and into `content/`; the table is kept only as a record of where
each image came from.

### Read at full resolution (genuinely new content)

- **`S__79970354_0.jpg`** — **"From Instinct to Imagination to Intelligence at Scale"**, Nacaskul 2026. Genuinely useful; folded into **H3** and §11. The best new find among the images.
- **`S__80085001_0.jpg`** — **not course content**: a photo of a screen showing a **Zoom invitation from Nacaskul** (meeting ID 818 1068 7287, plus an agenda link). **[INFERENCE]** Possibly the 2 Oct online class, possibly a pre-exam review session. **Worth asking whether a review session is scheduled before Friday — that would outrank everything in this file.**
- **`(2026.09.09_ChatGPT) From Predictive AI to Generative AI_0.png`** and **`…From VAE to Diffusion Models to Flow Matching_0.png`** — **duplicates** of sheets already in the infographics zip (identical dimensions, 1054×1492). No new content; already covered in H3, M3, M8 and L2.

### Low-resolution thumbnails — what I could and could not read

| File | What it is | Readable? |
|---|---|---|
| `…13.46.41` | Krueger 100XR drone sheet + **"Integrated Innovation: Lessons from the Krueger 100XR"** | Headline and six lesson titles legible; bodies not. Covered in **D1**. |
| `…13.46.59` | 3Blue1Brown **chapter 2** (gradient descent) summary | Mostly legible; folded into **H2**. |
| `…13.47.12` | 3Blue1Brown **chapter 1** (what is a neural network) summary | Mostly legible; folded into **M1**. |
| `…13.48.18` | **MASTER TIMELINE OF AI** | Structure and thread names only. See **H5**. |
| `…13.48.27` | **FROM IDEA TO INTELLIGENCE** combined timeline | Row labels only, no entries. See **H5**. |
| `…13.48.41` | **"A4InfoExplain: Psychology of Intelligence"** — a study-technique sheet (retrieval practice, spacing, interleaving, desirable difficulty, Bjork citations) | Headings legible. **Not exam content** — advice on *how* to revise, not *what*. It happens to recommend exactly the cover-and-recite method in §10. |
| `…13.48.49` | Screenshot of Nacaskul's **Quantum Computing** YouTube playlist | Legible. **Irrelevant — different course.** |
| `…13.49.02` | Screenshot of Nacaskul's **"BAScii - Generative AI (Autumn 2026)"** YouTube playlist | **See below — this is the important one.** |
| `…13.49.12` | The Mathematica **RMSE Gauge / RMSE Surface** demo | Legible; folded into **H2**. |
| `messageImage_1789028631413_0.jpg` | MyCourseVille **Video Playlist** page for 5602203.01, with a Thai "file retention expired" overlay | Item titles partly legible: *Essence of Linear Algebra*, *…Neural Network*, *Deep Learning Animated*, *Generative AI Animated*, *Computational …*, ***The Wolfram Neural Net Framework: A Gentle Introduction***, *AI Mathematics*, ***BAScii - Generative AI (Autumn 2026)***. |

### ★ The most actionable finding in this document

**[EVIDENCE — `Screenshot … 13.49.02`]** Nacaskul runs a YouTube playlist called **"BAScii - Generative AI (Autumn 2026)"**, 4 videos, containing:

- **GAILit&Appl(2026) Class04b** — 1:13:47, uploaded ~2 hours before the screenshot
- **GAILit&Appl(2026) Class04a** — 31:51, same upload
- **GAILit&Appl(2026) Class03b** — 1:05:50, ~5 days earlier

**These are full recordings of this year's actual classes**, on the channel already in `links.md` (`youtube.com/@poomjainacaskulph.d.4869/playlists`) and listed on the MyCourseVille video page. The thumbnails show red/blue half-plane diagrams — **linear separability and perceptron decision boundaries**, consistent with class 3–4 content on MLP and CNN.

**[INFERENCE]** This beats any inference I can make from slides, because it is the lecturer saying the material in his own words, at the length he chose, two days before the exam. **If there is time for exactly one thing beyond §7, it is Class04a (32 min) then Class04b at 1.5×.** Check the playlist for a Class 05/06/07 that may have appeared since the screenshot.

### Which thumbnails were worth re-exporting

**None — closed.** The two that mattered, `13.48.18` (**MASTER TIMELINE OF AI**) and
`13.48.27` (**FROM IDEA TO INTELLIGENCE**), are sheets **7** and **9** in the infographics zip
and have been read at full resolution. Their threads, lineages and dated entries are now in
`content/lectures/timelines.md` and in the `tim-` questions. The remaining thumbnails were
either already understood (the 3B1B sheets, the RMSE demo, Krueger) or not exam content (the
study-psychology sheet, the quantum playlist).

---

## 13. Open questions for the student

1. **Is a Zoom review session scheduled before Friday?** `S__80085001_0.jpg` shows a Zoom invite from Nacaskul with an agenda document. If that is a pre-exam review, it outranks this entire document.
2. **Watch the class recordings.** "BAScii - Generative AI (Autumn 2026)" on `youtube.com/@poomjainacaskulph.d.4869/playlists` — Class03b, Class04a, Class04b, and anything newer. §12.
3. **Page 1 of the 2025 midterm is not obtainable** — closed. Note only that **questions 1–10 of that paper were withheld**, so the visible sample is partial and biased toward its later questions; §2.1 explains how I compensated.
4. ~~**The Wolfram U CNN course itself.**~~ **CLOSED.** The full lesson list has been read off the course player and folded into **M2b**; it also strengthened **H1** (U-Net, Transformer and GAN are all in the material Marko assigned). No further action needed.
5. ~~**Re-export the two timeline posters.**~~ **CLOSED.** All 15 content sheets were in the
   infographics zip at full resolution all along; every one has now been read and folded into
   `content/`. See §11 for the sheet-number index used by the question sources.
6. **The MyCourseVille supplementary-material page**, in case it holds a study guide the folder lacks. The screenshot shows the video list, but the page had expired.
6. **Confirm the blueprint slide is Marko's**, if you can. I inferred it from the deck series (`Gen AI Lit 02`'s title slide names him; 03/06/07 continue it), and the whole ranking in §6 leans on that. If "Main topics for the midterm exam" turns out to be Nacaskul's, tell me and I will re-rank again.
7. **Are bonus questions still a thing?** 2025 had two. With ten questions this year, unknown. If they are, the ethics/playlist-video bonus (M9) is the likely one.
8. **Links not fetched** (inferred from titles and your notes only, flagged where it matters): the YouTube videos themselves, the Cornell VAE interactive demo, the OpenAI "how our models are developed" page. The Google Slides link is lecture 6 = `Gen AI Lit 06.pdf`, which I read.
9. **`6901009256_hwk3.nb`** is the Mathematica `Manipulate` linear-regression exercise, matching the RMSE Gauge/Surface screenshot. No exam content beyond "a linear model is a perceptron without the activation, and the RMSE surface is the loss landscape".

---

## 14. A classmate's notes — an independent record of what was said in class

**[EVIDENCE — 21 pages of a classmate's contemporaneous notes, read 23 Sept]** A different
class of evidence from anything above: not a deck, not a past paper, but a record of **what was
actually said in the room**, by someone who was in it. Her covering message: *"He talked a lot
about being able to distinguish the concepts he's mentioned from each other for the exams."*
That independently confirms what `lec5/summarized_notes.md` already said — **A-vs-B
discrimination is the dominant question shape**, not free-standing definition.

Questions sourced from these pages cite `{"deck": "class notes", "slide": N}`, where N is the
page number 1–21 in filename order.

### What they confirm (first-hand, from the room)

- **The six architectures are the exam material.** A page headed **"6 architectures for
  midterms"**, dated **Sep 11**, lists them by name: **MLP · CNN · VAE · U-Net · Transformer ·
  Diffusion.** A second page, **Sep 4**, opens **"Understand 6 architectures of AI for exam"**.
  **H1 now has a first-hand witness as well as the examiner's blueprint.**
- **"What led the field of AI to transition from predictive to generative AI?"** — written down
  as a question in its own right. Now `gen-034`.
- **Parameter counting**, the "(AI =) ANN = f(θ) = architecture" identity, the
  backprop-as-specialisation-of-GDA sentence, the cat's line-detecting neurons, word algebra,
  ADAM, BNN vs ANN and the three senses of I and A — **all corroborated verbatim**.

### What they add that the decks only glanced at

| Finding | Where it went |
|---|---|
| **"Learn the vocabulary, then learn the grammar."** Embedding is the vocabulary (*cat is closer to dog than to refrigerator*); attention is the grammar. Grammar is harder because a sentence is not recoverable from distances between isolated words | `tra-031`, `tra-033`, `transformers.md` |
| **The dot product** is what attention actually computes — how congruent two vectors are | `tra-032` |
| **Why the hard edge is smoothed:** a step has no useful derivative, so the sigmoid is what makes the model *differentiable* and therefore learnable | `nn-050`, `neural-networks.md` |
| **A perceptron is a partition**, a hyperplane of **one dimension fewer** than the data; ML is adjusting the parameters of those partitions | `nn-052` |
| **False positives vs false negatives:** *civil courts minimise false positives; military commands minimise false negatives.* Both are errors; the threshold is a judgement about which one costs more | `nn-053` |
| **The GAN posed as two questions:** can a model judge a photo to be Obama? — then, can an engine make an image the first model judges real? | `gen-033` |
| **The road-and-houses picture of latent space**, and why sampling a plain autoencoder's latent space is risky | `ae-031`, `ae-032` |
| **The postal-code system that broke when the digits moved**, forcing the convolutional filter layer | `cnn-030` |
| **"AI is intelligent but cannot do things. Robots can do things but are not intelligent."** Embodied AI is the join | `con-043` |
| **"Being hungry is not an intelligent calculation — it is reading a negative signal coming from the body."** The clearest statement of valenced feeling there is | `con-044` |
| **"All fiction is realistic but not real."** Approximating reality is predictive; taking it and running with it is generative | `con-046` |
| **Computer scientist vs data scientist** on whether ML sits inside AI or AI is branding on top of ML | `con-045` |
| **GOFAI's one surviving niche:** heuristic medicine, because differential diagnosis genuinely is an if-then tree | `timelines.md` |

### Treated as colour, not as examinable

**[INFERENCE]** The **Buddhism and symbolism** material — the **five aggregates** (form,
feeling, perception, mental formations, consciousness) that the **five elements of the
Intelligent Being** are drawn from — appears **nowhere in any deck**. Marko sets the paper, so
**no question was written on it**. It is recorded in `concepts.md` as a *handle* on a list that
*is* examinable: if you can hold the five aggregates you can hold the five elements. Same for
*"integration is not a salad bowl"* and *"a boy can play catch without knowing Newtonian
mechanics"* — good sentences, not question material.

### Deliberately left out as uncorroborated

- **"Neme2Vec"**, described as a finer-grained word2vec embedding the smaller etymological parts
  of a word. No such algorithm appears in any deck or infographic; it is most likely a garbled
  note on **subword or morpheme tokenisation**, which is a real thing and is already covered by
  the tokenisation material. **No question written.**
- **"The culmination of visual processing was differential neural networks."** No deck says
  this, and the term does not exist as written. Probably *diffusion*. **Not used.**
- **"6 AI architectures: feedforward neural network, recurrent neural network"** on one page
  conflicts with the same student's own **"MLP · CNN · VAE · U-Net · Transformer · Diffusion"**
  on another, and with every deck. **The deck list wins**; the stray note is a transcription
  slip and was ignored.
- **The Mathematica homework page** (`Plot[...]`, `ControlPlacement`) is excluded by **"no
  code"**.
- **"1958 Cat Brain Test."** Her date agrees with `lec5.md`; Nacaskul's slides say 1959 in one
  place and 1962 in another. The existing advice stands: **say "around 1959"** and spend the
  words on the finding.

### Four gaps this pass closed

Items the lecturer marked "Exam notice" in the student's own notes that had **no question**:

1. **"What are three forms (at least) of intelligence?"** → `con-041` (select-all) and
   `con-042`, sourced to the lecture-1 deck's own list rather than to the notes.
2. **12288, the ChatGPT vector size.** **Was** already covered by `tra-013` — but written
   `12,288` with a comma, which is why a grep for `12288` missed it. `tra-013` is now marked
   `examFocus` and its explanation carries both spellings.
3. **Counting parameters** (flagged **twice** in `lec3.md`, the strongest signal in those
   notes) → `nn-046` to `nn-049`, drilling 3-input, single-layer and two-layer cases, plus a
   table in `neural-networks.md`. **The convention, stated in every explanation: one weight per
   incoming connection plus one bias per neuron; inputs and hyperparameters do not count.** The
   worked example dictated in class — *two inputs, one output, three parameters* — is row one.
4. **Buddhism and symbolism** → deliberately not examined; see above.
