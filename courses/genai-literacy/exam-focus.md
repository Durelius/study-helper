# 5602203 GAI Literacy & Applications — Midterm Focus

Chulalongkorn / CSII / BAScii, Autumn (Fall) 2026.
Lecturers: **Poomjai Nacaskul** (concepts, timelines, architectures) + **Marko Niinimaki** (labs, Colab/Mathematica) + Kongkan Kalakan (post-midterm workshops).

Compiled 2026-09-23 from `/Users/wilhelmdurelius/chula/ailiteracy/`. Revised the same day, after 14 further images arrived — see **§12**.
Tags used throughout: **[EVIDENCE]** = stated in a course document. **[INFERENCE]** = my reasoning. **[GUESS]** = weak. **[LOW-RES]** = read off a thumbnail where headings were legible but body text was not; structure only, nothing invented.

> **If you read one thing first, read §12.** The new images contain one finding that beats everything else in this document: **Nacaskul uploads full recordings of this year's classes to YouTube, including the pre-midterm ones.**

---

## 1. What the exam is

**[EVIDENCE — `Gen AI Lit 07.pdf`, p.3, Niinimaki's slide showing the actual answer sheet]**

- Header on the real paper: *"5602203 Generative AI Literacy and Applications — Fall 2026 Midterm Exam"*.
- **Closed book.**
- **10 questions.**
- **2 ruled lines per answer** (Q1 and Q2 show 2 lines, Q3 shows 3 — so roughly 2, occasionally 3). Instruction on the sheet: *"Your **brief** answer should be written on the lines of this sheet."*
- Explicit warning: *"We only read what you can write on the lines. And your handwriting must be readable."*
- Name + student ID required on the sheet.

**[EVIDENCE — syllabus `curriculum.pdf` = `_Autumn_2026_...Syllabus...pdf` (byte-identical, md5 f19689d2…)]**

- Midterm Examination: **20%** of the final grade. (Attendance 10, Participation 10, Quiz —, Midterm 20, Final 20, Project 40.)
- Marked "Midterm Examination (Yes): **Closed-book**" in the session table, sitting **after session 7**.
- Remark: *"Students must achieve a minimum of 50% in each grading criterion to pass."* So the midterm is a hurdle, not just 20% of a pot.
- Course Learning Outcome → Assessment mapping. Only **one** CLO is assessed by the exams:
  > **CLO 2: "Explain key GAI architectures using intuitive mental models, i.e. how Transformer and Diffusion Models 'work under the hood'." → Midterm/Final Exam**
  CLO 1 (history/philosophy) → Homework. CLO 3 (RAG/agentic workflows) → Term Project. CLO 4 (impact on society) → Class Discussion.
  **[INFERENCE]** That mapping is the single strongest syllabus-level signal: architecture explanation is the examined outcome. But see §7 — the lecturers' own midterm-topics slide is broader than the CLO map, so don't treat CLO 2 as a fence.

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

### Scope, in the lecturers' own words

**[EVIDENCE — `Gen AI Lit 07.pdf`, final slide, verbatim]**

> **Main topics for the midterm exam**
> Please check the Youtube playlist in MyCourseVille Supplementary Material!
> Some topics covered in the CNN course (assignment, Wolfram U), too.
> **Study the infographics sheets carefully!**
> Concepts and Principles (lecture 1)
> Neural networks and how they learn
> Architectures: Multi-Layer Perceptron, Convolutional NN, VAE, ..
> **No code, very little math.**

That slide is the highest-value single artifact in the folder. Note three things it does: it **includes** the YouTube playlist and the Wolfram CNN quizzes, it **elevates the infographics** to primary study material, and it **excludes code and maths**.

**[EVIDENCE — syllabus session table, sessions 1–7 = everything before the midterm]**

| # | Contents | Lecturer |
|---|---|---|
| 1 | The "I", "A", and "G" in "GAI" | Nacaskul |
| 2 | The "I", "A", and "G" in "GAI" | Nacaskul |
| 3 | The "Language AI" & "Vision AI" Timelines | Nacaskul |
| 4 | MLP, Neocognitron/CNN | Nacaskul |
| 5 | AE/VAE, Encoder-Decoder | Nacaskul |
| 6 | GOFAI, NLP, Transformer, LLM | Nacaskul |
| 7 | Prompt/RAG, Mix Architectures | Nacaskul |
| — | **Midterm Examination (Yes): Closed-book** | |

**[INFERENCE]** Delivery ran behind this plan. Your own notes for lecture 6 (`lec6/during.md`) are about ML vs AI, embodied AI, gradient descent, backprop and RMSE — syllabus session 4–5 material — and the Niinimaki decks 06/07 cover autoencoders and a Transformer lab. **Prompt/RAG (session 7) was pushed past the midterm**: `Gen AI Lit 07.pdf` says the *post*-midterm classes are the hands-on Agentic AI and RAG tutorials. So treat Prompt/RAG as low-probability for this exam (see §7).

---

## 2. Evidence from past papers

### 2025 midterm (`(excerpt) 2025 GAI Lit _amp_ Appl - midterm w soln.pdf`)

**Caveat: the excerpt starts at "[Page 2]", i.e. Q11. Questions 1–10 are not in the file.** The paper ran to at least Q25 plus two bonuses, so 2025 was a ~25-question paper — a different shape from this year's 10. The *style* still transfers: one-sentence prompts, 3–5 line boxed answers, heavy on "name the Nth element of X and briefly explain" and "what does ACRONYM stand for — explain letter L".

Grouped by topic, verbatim questions with the official answers condensed:

**Meta / programme concepts**
- **Q11.** *What does "and" mean within the context of your "Integrated Innovation" education?* → **And ⇔ Integrated**, as applied to BAScii's **multidisciplinary education** model, producing socially responsible **innovators**. It does *not* mean bolting desirable features together and calling it an innovation.
- **Q12.** *What does it mean to you to achieve the bare minimum level of "AI Literacy"?* → As a **global digital citizen**, appreciating **how AI got to where it is**, **how it is trained from data**, and the **capability as well as pitfalls** of AI solutions, so I don't blindly apply AI.
- **Q13.** *How does adding a chatbot to your e-commerce channel qualify as "AI Application"?* → AI Application **in the most basic form**, relying on the chatbot's ability to carry on "**human-like**" conversation, **without necessarily implying that its suggestions are intelligent solutions** to real problems.

**Architecture acronyms**
- **Q14.** *What does MLP stand for within the context of ANN? Explain the "P" in MLP.* → **Multilayer Perceptron**. The Perceptron is a **mathematical-computational model of a generic neuron**: it **sums weighted excitory and/or inhibitory incoming signals** and **fires if/when that exceeds an internal threshold**.
- **Q15.** *What does CNN stand for within the context of ANN? Explain the "C" in CNN.* → **Convolutional Neural Network**. Convolutional refers to the **convolutional layer**, a layer of **convolution matrices with trainable entries** which, applied to a **grid of pixel values**, **detect underlying visual patterns**.

**The five elements of "Embodied AI" / the Intelligent Being (Q16–Q20, one question each)**
1. **Corporal Form & Physical Embodiment** — the physical body, i.e. **robotics with some degree of autonomy** and an **array of sensors**, enabling the AI to interactively explore and learn from its operating environment.
2. **Valenced Feeling & Graded Sensation** — the ability to evaluate sensed input on a **preferential scale**, very/slightly negative → neutral → slightly/very positive, serving as **feedback to the learning algorithm**, hence conditioning future behaviours.
3. **Cognitive Faculty & Pattern Recognition** — the ability to **learn (form an internal model of) how reality operates**, form memory, and therefore identify/classify/predict likely outcomes from relevant inputs — **hence Predictive AI**.
4. **Mental Formation & Projective Generation** — the ability to **generatively imagine realistic/convincing scenarios, shapes, forms, ideas**, especially given that AI has learned to identify/classify/predict such items — **hence Generative AI**.
5. **Consciousness Marker & Program Register** — the momentary, sequential **awareness of the computation/inference step being performed**, from one moment to the next; central to the debate over whether AI can achieve genuine conscious awareness.

**Numerical optimisation (Q21–Q25)** — on f(x,y) = x² + y² − xy − 5x − 5y + 25, minimise: evaluate f(3,3)=4; derive ∇f = (2x−y−5, 2y−x−5)ᵀ; evaluate ∇f(3,3)=(−2,−2)ᵀ; step 0.1 against the gradient to (3.2,3.2); compare f values.
**[INFERENCE] This block is very unlikely to recur this year — "no code, very little math" rules it out, and 2 ruled lines cannot hold "show your work".** But the *concepts* (gradient = direction of steepest ascent; step against it to minimise loss) are fair game as a written definition.

**Bonus questions**
- **Bonus 1.** *1–3 key messages you got from watching "Is AI Slop Killing the Internet?" by Patrick Boyle* → old ad-revenue model paid creators for engaging content; AI search now delivers content without visits, so the incentive to produce high-quality trustworthy content evaporates; countermeasures exist but are limited, so consumers must stay vigilant about AI-generated content.
- **Bonus 2.** *Rank ChatGPT vs Gemini vs Claude Sonnet vs DeepSeek from your own experience* — on accuracy, use of RAG vs "pure" generation, linking to legitimate URLs, concise organised narratives. (Open opinion, left blank in the solution.)

**[EVIDENCE] That Patrick Boyle video is item #21 on the current YouTube playlist.** The lecturer sets bonus questions straight from his own playlist. See §4.

### 2024 "Applied AI" midterm (`(excerpt) 2024 Applied AI - midterm w soln.pdf`)

Opens with a **term bank** students must use and circle:
> {Artificial Intelligence (AI), Machine Learning (ML), Unsupervised Learning, Supervised Learning, Reinforcement Learning, Regression Analysis, Multilayer Perceptron (MLP), Cluster Analysis, Dimensionality Reduction, Predictive AI, Generative AI, Verifiable Ground Truths, Plausible Context Relevance, (Word2Vec) Word Embedding, Generative Pretrained Transformer (GPT), Large Language Model (LLM)}

- **Q1.** *What do we mean by "ML builds AI"?* → **AI cannot be programmed to respond using traditional "if-then" constructs**; instead AI is programmed to **calibrate its model parameters from data**, hence a **Machine Learning** algorithm. **Supervised, Unsupervised and Reinforcement Learning are the 3 principal ML paradigms.**
- **Q2.** *Elaborate on Predictive vs Generative AI. Give an example of each.* → **Predictive AI is evaluated on accuracy against verifiable ground truths**, plus quality of insight from real data; MLP is an example of supervised-learning predictive AI, cluster analysis an example of unsupervised. **Generative AI is evaluated on its ability to generate responses with a high degree of plausible context relevance**; GPT is the best-known example.

### What recurs across both years

| Recurring theme | 2024 | 2025 | Verdict |
|---|---|---|---|
| **Predictive vs Generative AI** | Q2 (core) | Q19/Q20 framing, whole Embodied-AI spine | Near-certain again |
| **Unpack an acronym and explain one letter** | term bank (MLP, GPT, LLM) | Q14 (MLP/"P"), Q15 (CNN/"C") | Near-certain |
| **Definition-by-contrast of two neighbours** | Predictive vs Generative | Generation vs Recognition, AI Literacy vs Artificial Literacy | Near-certain |
| **ML/AI relationship** | Q1 "ML builds AI" | — | Likely; also flagged in class this year |
| **BAScii / literacy / meta concepts** | — | Q11, Q12, Q13 | Possible but lecturer-specific |
| **A bonus drawn from a playlist video** | — | Bonus 1 (Boyle) | Likely |

---

## 3. Emphasis flagged *in class this year* (your own notes)

These are the strongest year-specific signals, because you wrote them down when the lecturer said them.

**[EVIDENCE — `lec1/during.md`, marked "Exam notice"/"Exam note" by you]**
- **"Define valenced (valenced feeling & graded sensation) in the context of an intelligent being and robotics"** — this is verbatim 2025's Q17.
- **"What are three forms (at least) of intelligence?"** — the three senses of the "I".
- Predictive AI came first; GANs were invented as a *device to improve discriminative training* (fake-Obama detector) and ended up being the star. Deepfakes made it famous.
- **"English words were vectorized. You can do word algebra with vectors"** — King − male + female = Queen.
- **"12288, ChatGPT vector size."**
- **"Diffusion model comes from transformer and is for visuals."**

**[EVIDENCE — `lec3.md`]**
- **"Exam note: Connectionist AI"**
- **"Exam notice: count parameters"** (with the worked toy: 2 inputs + 1 output → 3 parameters), repeated twice. **[INFERENCE]** Nacaskul flagged this, but Niinimaki wrote "very little math". A *one-line* parameter count (weights + biases for a tiny net) fits on 2 lines and is not really "math". Worth 10 minutes of prep, not more.

**[EVIDENCE — `lec5.md`, your note headed "Synopsis of class topic coverage (up to midterm = 'literacy' half of the course)"]**
1. Precise/pedantic meanings of **G, A, I**, and of **AI, GAI, AGI**
2. The **2 "AI task" timelines + 1 "neuroscience foundation" timeline**
3. **HIWUTH: the 6 foundational AI architectures + 2 foundational AI algorithms**

Plus: **"Exam for sure: 1958 cat line recognition experiment"** (Hubel & Wiesel — the slides date it 1959/1962, the notes say 1958; see §6).

**[EVIDENCE — `lec5/summarized_notes.md`, "What she says is on the exam"]**
1. The two timelines (vision and language)
2. HIWUTH: the 3 groups of AI and the 6 foundational architectures/algorithms — **"be able to tell them apart"**
3. **What a parameter, coefficient, weight and input are** (and that coefficient/parameter/weight are used interchangeably)

**[EVIDENCE — `lec6/during.md`]**
- **"Exam question: In what way is machine learning larger than artificial intelligence? And vice versa"**
- **"Exam sentence: BACKPROPAGATION IS A SPECIALIZATION OF GRADIENT DESCENT ALGORITHM, APPLIED TO MINIMIZE PREDICTION ERROR FOR MULTI LAYER PERCEPTRON (by tweaking weights and bias)"** — memorise this literally.
- RMSE = Root Mean Squared Error = the loss function L; GDA minimises it; ADAM = **Ada**ptive **M**oment Estimation.
- PAI mimics (re)cognition; GAI mimics thinking/generation.
- (AI =) ANN = f(parameters) = architecture.

---

## 4. The YouTube playlist, and what changed right before the exam

**[EVIDENCE]** The 2026.09.21 spreadsheet vs the 2026.09.15 one. Items are tagged 1=MustWatch, 2=QuickWin, 3=XtraInsight.

**The MustWatch list as of 21 Sep (this is your video revision list):**
| Topic | Video |
|---|---|
| Perceptron & ANN | 3Blue1Brown — *But what is a neural network?* (ch.1) |
| Gradient Descent / Backprop | 3Blue1Brown — *Gradient descent, how neural networks learn* (ch.2) |
| CNN | Nullpoint Error — *The real History of CNNs: From a cat's brain to ChatGPT's eyes* |
| VAE | Deepia — *Variational Autoencoders \| Generative AI Animated* |
| Transformer | 3Blue1Brown — *Attention in transformers, step-by-step* (ch.6) |
| Diffusion | Deepia — *Diffusion Models: DDPM* |
| **GAI Ethics & Post-Capitalism Economics** | **Thinkplicit — *The Rise of Collective Stupidity (End of Critical Thinking)*** |

Also QuickWin (short, cheap): 3B1B *Backpropagation, intuitively* (ch.3, 12 min), Infomity *CNN Animation* (8 min), Under The Hood *Transformer Architecture Explained* (20 min).

**What was added between 15 and 21 September (i.e. in the last week before the exam):**
1. A **whole new topic row: "GAI Ethics & Post-Capitalism Economics"** — three videos, with *The Rise of Collective Stupidity* tagged **MustWatch**. (The Patrick Boyle "AI Slop" video, which was a 2025 bonus question, is in this new row.)
2. The CNN history video ("from a cat's brain to ChatGPT's eyes") was **added and promoted to MustWatch**, displacing AlexNet down to XtraInsight. It is also the **first link in `links.md`**, so the lecturer pushed it directly to the class.

**[INFERENCE — moderately strong]** Last-minute additions are exam signals. Expect (a) a question or bonus touching **AI slop / cognitive offloading / the erosion of critical thinking / the economics of AI-generated content**, and (b) the **cat-brain → CNN narrative** to be the framing for the CNN question. The 2025 paper did exactly this with the Boyle video, and the syllabus reserves ~10% of the course for "How GAI Impacts My Generation".

---

## 5. Predicted topics, ranked

### HIGH confidence

**H1. The six foundational architectures — name them, place them, tell them apart.**
*Concepts:* **MLP, CNN, VAE, U-Net, Transformer, Diffusion** — plus the framing "predictive → representation bridge → generative".
*Why:* The infographic `From Prediction to Generation — 6 Key Architectures` is exactly these six; `lec5.md` and `lec5/summarized_notes.md` both record "6 foundational AI architectures … be able to tell them apart"; `Gen AI Lit 07` lists "Architectures: MLP, CNN, VAE, .."; syllabus sessions 4–6 are MLP/CNN, AE-VAE-Encoder-Decoder, Transformer/LLM; CLO 2 is the only exam-assessed outcome. Four independent sources.
**Confidence: high.** Likely 2–4 of the 10 questions.

**H2. The two foundational algorithms — gradient descent and backpropagation.**
*Concepts:* gradient = direction of steepest ascent; gradient descent takes finite steps *against* it to minimise a loss; backprop = the specialisation of GDA for an MLP, sending error backwards to tweak weights and biases; RMSE/loss; ADAM = Adaptive Moment Estimation.
*Why:* Your "Exam sentence" note in `lec6/during.md`; `Gen AI Lit 07` says "Neural networks and **how they learn**"; two MustWatch 3B1B videos; two infographics (`Gradient descent…` and `From Calculus to Backprop`); the 2025 paper devoted 5 questions to gradient descent numerically.
**Confidence: high** for a *verbal* question, **low** for a numerical one.

**H3. Predictive AI vs Generative AI.**
*Concepts:* Predictive is judged on **accuracy vs verifiable ground truths**; generative on **plausible context relevance**. Predictive learns x→y then you crank the input; generative learns z→x then you crank the PRNG. Discriminative p(y|x) vs generative p(x) or p(x|c). Predictive AI = element (iii) Cognitive Faculty & Pattern Recognition; Generative AI = element (iv) Mental Formation & Projective Generation.
*Why:* 2024 Q2 verbatim; the spine of the 2025 Embodied-AI block; dedicated infographic; Nacaskul's slides do "Predictive vs Generative AI, takes 1–4".
**Confidence: high.**

**H4. Unpacking G, A and I (and AI vs GAI vs AGI).**
*Concepts:*
- **I** — three senses: (1) "human-like", as in Turing's imitation game; (2) "advantageous information", as in Military/Business Intelligence (strategic/tactical value); (3) "solution to a problem", as in Computational/Digital Intelligence (optimal solution).
- **A** — three senses: (1) **Artifact/Artifice**, machine creation by humans; (2) **Artificial semblance**, when "seems about right" will do; (3) **Artful creation**, "truth manufactured" by whatever's trending.
- **G** — Generative ≠ Predictive (AI); Generation ≠ Recognition (Mind); Generate ⇔ {**Simulate** (given a distributional prior), **Connect** (graph/network adjacency), **Create** (integrated innovation)}.
- **GAI** = *Generative* AI; **AGI** = Artificial *General* Intelligence, explicitly out of scope for the course.
*Why:* Syllabus sessions 1–2 are literally this; `Gen AI Lit 07` names "Concepts and Principles (lecture 1)"; lecture-1 deck spends ~20 slides on it; XMind roadmap's first branch; your notes list it three times.
**Confidence: high.**

**H5. The two task timelines (+ the neuro/computational one).**
*Concepts:*
- **Language AI:** Turing 1950 *Computing Machinery and Intelligence* (imitation game) → symbolic/expert systems, retronymed **GOFAI**, AI winter → Mikolov 2013 **word2vec** (words as vectors, word algebra) → Vaswani 2017 **Attention Is All You Need** (Transformer, the "T" in GPT) → GPT-1/2/3 (0.117/1.5/175 bn params) → ChatGPT 2022.
- **Vision AI:** Hubel & Wiesel, cat's striate cortex — the cat's neurons fired for **oriented lines**, not for mice; lines compose into curves, curves into shapes → Fukushima 1980 **Neocognitron** → LeCun 1998 **LeNet/CNN** → Krizhevsky/Sutskever/Hinton 2012 **AlexNet** (GPU, ImageNet, error 25% → single digit) → phones telling cats from dogs.
- **Third timeline:** the **neurocomputational / pure-connectionist** one — McCulloch & Pitts 1943 → Rosenblatt 1958 Perceptron → Rumelhart/Hinton/Williams 1986 MLP+backprop. *The brain has no central processor; it is just neurons connected to neurons.* Your notes mark "Connectionist AI" as an exam note and "pure connectionist AI — VIKTIGT".
*Why:* Syllabus session 3; `lec5.md` item 2; `lec5/summarized_notes.md` item 1; the XMind roadmap's whole "Timelines" branch; the newly-promoted MustWatch CNN-history video.
**Confidence: high.**

### MEDIUM confidence

**M1. Perceptron / MLP mechanics and vocabulary.**
Perceptron = mathematical model of a neuron: weighted sum of inputs + bias, through an activation, fires past a threshold. **Bias = decision threshold** (`Gen AI Lit 03.pdf` says this in so many words). Single perceptron only separates **linearly separable** data — hence **XOR** as the classic counterexample, hence layers. MLP = feedforward layers stacked; no recurrence, no skipping layers. Distinguish **parameter = coefficient = weight** (learned) from **input/variable** (given) and **hyperparameter**.
*Why:* 2025 Q14; `lec5/summarized_notes.md` explicitly lists "what a parameter, coefficient, weight and input are" as examinable; `Gen AI Lit 06` quiz slides drill XOR, ReLU, sigmoid, and backprop code.
**Confidence: medium-high.**

**M2. CNN — why convolution, and the biological motivation.**
Convolutional layer = a layer of **convolution matrices (filters) with trainable entries** slid over a grid of pixel values to detect visual patterns. Weight sharing gives **translational covariance**; multiple scales handle objects at different distances; a **hierarchy of features** (lines → curves → shapes → objects) mirrors Hubel & Wiesel's cat. An MLP fed raw pixels fails because it cannot handle an object moving or changing colour.
*Why:* 2025 Q15; syllabus session 4; Wolfram quiz 1 is almost entirely this; the newly-promoted MustWatch video.
**Confidence: medium-high.**

**M3. AE → VAE, latent space, encoder–decoder.**
**Autoencoder:** x → compressed code → x̂; the **bottleneck** forces dimensionality reduction. **VAE:** encodes to a *probabilistic* latent **distribution** (≈ Gaussian) rather than a point, so you can **sample z ~ N(0,I)** and decode to get a *new* example — that's the predictive→generative transition. **Encoder–decoder** generally: compress to a latent, then expand to the target modality. **U-Net** = encoder–decoder CNN with **skip connections** (hence the U) preserving fine spatial detail; used for image segmentation and medical imaging.
*Why:* Syllabus session 5; `Gen AI Lit 06` lab was an autoencoder and asked "Latent Space?"; `Gen AI Lit 07` names VAE in the exam topics; VAE is a MustWatch video; two dedicated infographics.
**Confidence: medium-high.**

**M4. Transformer / attention / embedding.**
Words are not numbers, so **tokenize → embed** into vectors (GPT's embedding dimension ≈ **12288**, per your lecture-1 note). Embeddings give **global-context meaning** and support word algebra (king − man + woman ≈ queen). **Attention** gives **local-context interpretation**: it weighs which other tokens matter to each token, via **Q, K, V** matrices. **Positional encoding** supplies word order. GPT = **Generative Pre-trained Transformer**, decoder-only; "pre-trained" = trained on internet-scale text beforehand. An LLM generates by repeatedly predicting a contextually plausible next token.
*Why:* Syllabus session 6; CLO 2 names Transformer explicitly; `Gen AI Lit 07`'s lab was a Transformer with the prompts "How does the transformer decide which word it generates next? What are the main parts of this transformer?"; MustWatch 3B1B ch.6; *Attention Is All You Need* is in the seminal-papers zip and is syllabus reading [10].
**Confidence: medium-high.** (Slightly below CNN/VAE only because Niinimaki's topic slide trails off at "VAE, ..".)

**M5. ML vs AI — which contains which.**
Machine learning is **larger** than AI in that ML methods are used far outside AI (statistics, optimisation, any data-driven calibration); AI is **larger** than ML in that AI also includes rule-based/symbolic GOFAI, search, planning and robotics that involve no learning. Modern AI is built by ML because AI **cannot be pre-programmed with if-then constructs** — it must calibrate parameters from data. Three ML paradigms: **supervised, unsupervised, reinforcement** (Nacaskul's glossary adds **representation learning** as a fourth modality).
*Why:* Your note "Exam question" verbatim in `lec6/during.md`; 2024 Q1 is the same idea from the other side.
**Confidence: medium-high.**

**M6. GAN and the "adversarial" origin of the G.**
Two networks in a minimax game: a **generator** makes fakes, a **discriminator** judges real vs fake; both improve. **The irony Nacaskul stresses:** the adversarial network was built to *improve the predictive/discriminative* model's training, and ended up being the real star. Goodfellow et al. 2014. This is the *other* G in GAI — the one that gave "generative" its modern meaning.
*Why:* Flagged "Exam note" in `lec1/during.md`; slide "7) The Road to Generative AI (mid 2010s)" spells out the irony; `lec6/during.md` opens with GAN = Generative Adversarial Network.
**Confidence: medium.**

**M7. Diffusion models.**
A **fixed forward process** adds Gaussian noise to data step by step; the model **learns to reverse it** (denoise). Generation = start from pure noise and iteratively denoise. Text-to-image (Stable Diffusion, DALL·E) conditions this on a text embedding. Your note: "Diffusion model comes from transformer and is for visuals."
*Why:* CLO 2 names diffusion models alongside Transformers as the "under the hood" targets; MustWatch DDPM video; DDPM paper in the seminal zip; two infographics.
**[INFERENCE]** Tension: CLO 2 names it, but Niinimaki's topic slide stops at "VAE, ..", and diffusion belongs to session 5–6 material that may have run thin. Prepare a two-line answer; don't over-invest.
**Confidence: medium.**

**M8. The Embodied AI / Intelligent Being five elements.**
Verbatim from §2. Note **Valenced Feeling & Graded Sensation** is double-flagged: 2025 Q17 *and* your lecture-1 "Exam notice".
*Why:* Five of 2025's twenty visible questions; lecture-1 deck states it; recapped again in the architectures deck; XMind puts it under Embodied AI. Counter-signal: it is Nacaskul's personal framework and appears nowhere on Niinimaki's topic slide.
**Confidence: medium** overall, **high** specifically for "valenced".

**M9. Ethics / impact: AI slop, cognitive offloading, collective stupidity.**
Old web economics paid creators via clicks and ads; AI search answers without a visit, so the incentive to produce quality, trustworthy content collapses, degrading the content supply that trained the models. Related: **cognitive offloading and dependency crisis**; "Artificial Literacy" = leaning on ChatGPT to write your essay, pitch your project, read your papers — the failure mode of AI Literacy.
*Why:* Whole new playlist row added 21 Sep with a MustWatch; 2025 had exactly this as Bonus 1; syllabus reserves ~10% of the course for it; lecture-1 deck has "Artificial Literacy" and "Artificial Application" slides.
**Confidence: medium** for a regular question, **higher** for a bonus/final question if bonuses recur.

### LOW confidence

**L1. The Five Principles of Modern Deep Learning** — **Representation → Relation → Parameterization → Optimization → Architecture.** ("Not ever more complicated primitive mathematics, but increasingly ingenious orchestration of representations, relationships, parameters and learning.")
*Why medium-low:* It is the **newest infographic (15 Sep)** and Niinimaki said to study the infographics carefully — but no lecture note of yours mentions it, so it may not have been taught. **Cheap to memorise as five words; do that and move on.**

**L2. GOFAI / expert systems** — rule-based formal logic, if-then trees, early success in **medical diagnosis** (MYCIN 1975, ~65% accuracy), LISP and Prolog, AI winter, retronymed **GOFAI** ("Good Old-Fashioned AI"). Know what a **retronym** is. Postscript worth knowing: rule-based ML (decision trees, random forests, gradient-boosted trees) is still powerful; and today's LLMs, though pure ANN, are increasingly asked to do logical reasoning.
*Why low-medium:* Syllabus session 6, present in both lecturers' decks, but no "exam note" attaches to it. One question at most.

**L3. Wolfram CNN course specifics** — Niinimaki says "Some topics covered in the CNN course (assignment, Wolfram U), too." But those quizzes were multiple-choice about Wolfram Language layers (`NetEncoder`, `NetDecoder`, `FlattenLayer`, `SoftmaxLayer`, `CrossEntropyLossLayer`, `ElementwiseLayer[Ramp]`). "No code" argues against asking for API names. What *can* be asked in prose: **overfitting** (training error keeps falling while validation error plateaus or rises), **data augmentation**, **softmax** (turns raw activations into a probability distribution over classes), **ReLU/ramp** (returns the argument if > 0 else 0; introduces nonlinearity), **pooling** (reduces the size of the representation), **translational covariance**, **hierarchy of features**, **colour correlation length**.
**Confidence: low for code specifics, medium for the prose concepts.**

**L4. Prompt engineering / RAG / agentic workflows.** Syllabus session 7, but `Gen AI Lit 07` says these are the *post*-midterm hands-on classes. **[INFERENCE] Probably not on this exam.** Know one line each just in case: RAG = retrieve relevant documents first, then generate grounded on them, for factual verifiability.

---

## 6. The 10–15 most likely questions, with model two-line answers

Written as the examiner would phrase them (Nacaskul's house style: short prompt, often "…and briefly explain").

**1. What does MLP stand for, and what does the "P" refer to?**
> Multilayer Perceptron. The Perceptron is a mathematical-computational model of a neuron: it sums weighted excitory/inhibitory input signals plus a bias and fires when that sum exceeds an internal threshold.

**2. What does CNN stand for, and what does the "C" refer to?**
> Convolutional Neural Network. The convolutional layer is a layer of convolution matrices (filters) with trainable entries which, slid across a grid of pixel values, detect underlying visual patterns.

**3. Distinguish Predictive AI from Generative AI, with one example of each.**
> Predictive AI maps input to output and is judged on accuracy against verifiable ground truths — e.g. an MLP classifier. Generative AI models the data distribution and is judged on plausible context relevance — e.g. GPT.

**4. Name the six foundational architectures covered in this course, in order from predictive to generative.**
> MLP, CNN, VAE, U-Net, Transformer, Diffusion. MLP and CNN predict; VAE and U-Net are the representation bridge; Transformer and Diffusion generate.

**5. State, in one sentence, the relationship between the gradient descent algorithm and backpropagation.**
> Backpropagation is a specialisation of the gradient descent algorithm, applied to minimise prediction error in a multilayer perceptron by tweaking its weights and biases. The gradient gives the direction of steepest change; GDA steps against it in finite steps.

**6. What did Hubel and Wiesel observe in the cat's visual cortex, and why does it matter for AI?**
> Neurons in the cat's striate cortex fired for oriented lines, not for whole objects; line detectors combine into curve detectors and upward into shapes. That hierarchy of features is the biological blueprint for the Neocognitron and the CNN.

**7. Give the three distinct senses of the "I" in "AI".**
> (1) "Human-like", as in Turing's imitation game; (2) "advantageous information", as in Military or Business Intelligence; (3) "solution path to a problem", as in Computational or Digital Intelligence.

**8. Give the three distinct senses of the "A" in "AI".**
> (1) Artifact/Artifice — machine creation by humans; (2) Artificial semblance — when "seems about right" will do; (3) Artful creation — "truth manufactured" by whatever is trending.

**9. "Generate" can mean three things. Name them.**
> Simulate (given a distributional prior), Connect (graph/network adjacency), and Create (integrated innovation). Generative is opposed to Predictive in AI, and to Recognition in the mind.

**10. In what way is machine learning larger than artificial intelligence, and in what way is AI larger than ML?**
> ML is larger in that its methods calibrate parameters from data far outside AI — statistics, optimisation, forecasting. AI is larger in that it also includes rule-based GOFAI, search, planning and robotics that involve no learning at all.

**11. Explain what makes a Variational Autoencoder generative, when an ordinary autoencoder is not.**
> A plain autoencoder compresses x to a fixed code and reconstructs it. A VAE encodes to a probabilistic latent distribution (≈ Gaussian), so you can sample a new z and decode it into a new, previously unseen but realistic example.

**12. What is the attention mechanism, and why was it the breakthrough behind LLMs?**
> Attention weighs how much each token in a sequence should inform every other token, via query–key–value matrices, giving words local-context interpretation on top of their embedding's global-context meaning. It removed recurrence and made training parallel and scalable — Vaswani et al., 2017, the "T" in GPT.

**13. Name the 2nd element/component of an Intelligent Being ("Embodied AI") and briefly explain.**
> Valenced Feeling & Graded Sensation: the ability to evaluate sensed input on a preferential scale from very negative through neutral to very positive, which feeds back into the learning algorithm and thereby conditions future behaviour.

**14. What is a GAN, and what is ironic about its role in the history of Generative AI?**
> A Generative Adversarial Network pits a generator making fakes against a discriminator judging them, improving both. The irony: the generator was built as a device to sharpen a *predictive* discriminator, and ended up being the star of generative AI.

**15. How does a diffusion model generate an image?**
> A fixed forward process gradually corrupts data with Gaussian noise; the network learns to reverse it. Generation starts from pure noise and iteratively denoises, optionally conditioned on a text embedding.

**16. What is overfitting, and how would you detect it?**
> The model fits the training data ever better while failing to generalise. Symptom: training error keeps falling while validation error plateaus or starts rising.

**17. What is meant by "AI slop", and why is it an economic as well as a quality problem?**
> Low-value AI-generated content flooding the web. The old model paid creators through clicks; AI search now answers without a visit, so the incentive to produce high-quality trustworthy content evaporates and the overall reliability of internet content declines.

**18. What does it mean to achieve the bare minimum of "AI Literacy"?**
> As a global digital citizen, to appreciate how AI technology got to where it is, how it is trained from data, and both the capability and the pitfalls of AI solutions — so that I do not blindly apply AI.

---

## 7. Definitions and distinctions worth memorising

Short-answer papers live on these confusable pairs. Two lines each.

| Pair | The distinction in one line each |
|---|---|
| **GAI vs AGI** | GAI = *Generative* Artificial Intelligence (creates new content). AGI = Artificial *General* Intelligence (human-level breadth) — explicitly beyond this course. |
| **Predictive vs Generative AI** | Predictive: accuracy against **verifiable ground truths**; learn x→y, then crank the input. Generative: **plausible context relevance**; learn z→x, then crank the PRNG. |
| **Generation vs Recognition** | Recognition is an **interpolative** map within known territory (ground truth, validation). Generation is an **extrapolative** map into unexplored territory (imagination). |
| **BNN vs ANN** | Biological Neural Network = the actual brain. Artificial Neural Network = a mathematical model of it; every modern AI is built from ANNs. |
| **Perceptron vs MLP** | Perceptron = one artificial neuron, only separates linearly separable data. MLP = feedforward layers of perceptrons stacked, each layer's output the next layer's input. |
| **FFN vs RNN** | Feedforward: signal travels one way only. Recurrent: the signal loops back into the network. MLP is feedforward with no recurrence and no layer-skipping. |
| **Weight / coefficient / parameter vs input vs hyperparameter** | Weights, coefficients and parameters are the same thing — the numbers *learned* from data. Inputs are the x variables you supply. Hyperparameters are set before training. |
| **Gradient descent vs backpropagation** | GDA is the general "step downhill" optimiser. Backprop is GDA specialised to an MLP, sending error backwards through layers to update weights and biases. |
| **Loss function vs RMSE** | Loss L is the general "how wrong am I" number. RMSE (Root Mean Squared Error) is one specific loss. |
| **AE vs VAE** | AE: deterministic code, reconstruction only. VAE: probabilistic latent distribution you can *sample* from, hence generation. |
| **VAE vs Diffusion** | VAE *learns* both the encoder (to→simple) and decoder (fro→data). Diffusion *prescribes* the forward noising and only *learns* the reverse denoising. |
| **Encoder–decoder vs U-Net** | Encoder–decoder compresses then expands. U-Net is an encoder–decoder CNN with **skip connections** (the "U") that carry fine spatial detail across the bottleneck. |
| **Embedding vs attention** | Embedding gives a word its **global-context** meaning as a vector. Attention gives it a **local-context** interpretation relative to the other words in *this* sentence. |
| **GOFAI vs modern AI** | GOFAI = rule-based symbolic logic/expert systems, hand-written if-then. Modern AI = ANN whose parameters are calibrated from data by ML. "GOFAI" is a **retronym**. |
| **Retronym** | A new name for an older thing, coined because a newer thing took over the original name (expert systems → "GOFAI"; AI → "predictive AI"). |
| **ML vs AI** | Neither contains the other. ML reaches outside AI (any data-driven calibration); AI reaches outside ML (symbolic, search, planning). |
| **AI vs BI vs CI** | AI = approximating interactive/behavioural intelligence of a species. BI = extracted information of tactical value. CI = optimal solution/method for a complex problem. |
| **Apply/Applied vs Application/App** | "Applied" means working under real-world constraints, not simplified. "Application" ⇔ {Modification/Optimisation, Addition/Enhancement/Extension, Solution/Direction/Utilisation}. |
| **AI Literacy vs "Artificial Literacy"** | AI Literacy = can code, can safeguard, can explore/exploit. Artificial Literacy = outsourcing your essay, pitch and reading to ChatGPT — the failure mode. |
| **Sigmoid vs ReLU (ramp)** | Sigmoid squashes any real input to (0,1), interpretable as a probability. ReLU returns the argument if positive, else 0 — cheap nonlinearity. |
| **Softmax vs sigmoid** | Softmax turns a vector of raw activations into a probability distribution *across classes* (sums to 1). Sigmoid squashes a single value. |
| **Bias vs weight** | Weights scale each input. The bias shifts the decision threshold — `Gen AI Lit 03`: "Bias = decision threshold". |

**Dates and names worth having ready (the lecturer loves attribution):**
Turing 1950 (imitation game) · McCulloch & Pitts 1943 · Rosenblatt 1958 (Perceptron / Mark I) · Hubel & Wiesel 1959/1962 (cat's striate cortex) · Fukushima 1980 (Neocognitron) · Rumelhart, Hinton & Williams 1986 (backprop/MLP) · LeCun 1998 (LeNet/CNN) · Hinton & Salakhutdinov 2006 (autoencoder / dimensionality reduction) · Krizhevsky, Sutskever & Hinton 2012 (AlexNet) · Mikolov 2013 (word2vec) · Kingma & Welling 2013 (VAE) · Goodfellow 2014 (GAN) · Ronneberger 2015 (U-Net) · Vaswani 2017 (Transformer) · Ho, Jain & Abbeel 2020 (DDPM). Also Fei-Fei Li (ImageNet) and Geoffrey Hinton, both named in your lecture-2 notes.

**[Note on the cat date]** `lec5.md` records "1958 cat line recognition experiment"; Nacaskul's slides say **1959** in one place ("Receptive Fields of Single Neurones in the Cat's Striate Cortex", 1959) and **1962** in another. If asked for a year, "late 1950s / around 1959" is safest; the *finding* matters more than the year.

**The six papers the lecturer distributed** (`(seminal_foundational_papers)….zip`) — a compact map of what he considers canonical:
Bengio, LeCun & Hinton (2021) *Deep Learning for AI*; LeCun, Bengio & Hinton (2015) *Deep Learning*; Krizhevsky et al. (2012) *AlexNet*; Vaswani et al. (2017) *Attention Is All You Need*; Kingma & Welling (2019) *An Introduction to VAE*; Ho, Jain & Abbeel (2020) *DDPM*.
**[INFERENCE]** That set = CNN + Transformer + VAE + Diffusion + two deep-learning overviews. It corroborates H1/M3/M4/M7 and says nothing about RAG or agents.

---

## 8. What is probably NOT examinable

- **Numerical gradient descent / partial derivatives / computing ∇f by hand.** 2025 spent five questions on this, but Niinimaki's slide says **"No code, very little math"** and the answer sheet gives 2 lines with no room to "show your work". **[INFERENCE — confident]** Skip the arithmetic; keep the verbal definition.
- **Wolfram Language / Mathematica API names** (`NetEncoder`, `NetDecoder`, `FlattenLayer`, `Manipulate`, `NetPortGradient`). "No code." Keep the underlying concepts, drop the function names.
- **Python / Keras / PyTorch code** from the lab decks (`layers.Conv2D`, `nn.Transformer`, `d_model=128`, `<UNK>` tokens). "No code."
- **Prompt engineering, RAG, agentic workflows, knowledge graphs.** Syllabus session 7, but `Gen AI Lit 07` states these are the post-midterm hands-on classes and they map to the **Term Project**, not the exam. One-line awareness only.
- **Flow Matching, ResNet, LSTM, Vision Transformer, CLIP, Stable Diffusion internals, normalizing flows.** Present in infographics and the XMind, but never in a lecture note of yours and not on the topic slide. Recognise the names; do not study them.
- **Anything from sessions 8–15** (the Darkside lecture, tools/platforms workshops, project presentations) — after the midterm.
- **Deep quantum-computing / finance content** from the lecturer's other books shown in the lecture-1 prologue. Decoration.
- **[LOWER CONFIDENCE]** The BAScii/"Integrated Innovation"/"and" meta-questions (2025 Q11–Q13). They are Nacaskul's signature, so they *could* appear — but with only 10 slots and a topic slide focused on architectures, they compete badly. Memorise the three one-liners in §6 (Q18 and the "and ⇔ integrated" line) — that's 5 minutes of insurance, not a study topic.

---

## 9. A revision plan for the time available

**If you have one evening:**
1. §7 table of distinctions — read it twice, then cover the right column and recite. (60 min; this is where the marks are.)
2. The six architectures, one line each, in order (§5 H1). (20 min)
3. The "exam sentence" on backprop, verbatim (§3). (5 min)
4. The three meanings of I, of A, and of Generate (§5 H4). (20 min)
5. The two timelines as a list of names+years (§5 H5). (20 min)
6. The five Embodied-AI elements, especially **valenced** (§2). (15 min)
7. Skim the four infographics in §10 — they *are* the lecturer's answer key. (30 min)

**If you have two:** add the MustWatch videos you haven't seen — the CNN-history one and *The Rise of Collective Stupidity* first, since both were added days before the exam.

**Exam-room technique:** the sheet gives 2 lines and says only what fits on them gets read. Lead with the **term**, then the **mechanism**. Nacaskul's own model answers are written that way — bolded key term first, then the how. Don't write an intro clause.

---

## 10. Study these four infographics directly

Niinimaki said "Study the infographics sheets carefully!" — these are the four that map onto predicted topics. They are inside `(GAI-infographics) GAI Literacy _amp_ Appl.zip`:

1. **`cc_(2026.09.11_ChatGPT) From Prediction to Generation - 6 Key Architectures - v3.png`** — the six architectures with a key-idea line each. This *is* H1's answer key.
2. **`cc_(2026.09.09_ChatGPT) From Predictive AI to Generative AI - v2.png`** — "Learn x→y then crank the input" vs "Learn z→x then crank the PRNG".
3. **`cc_(2026.09.15_ChatGPT) Five Principles in Modern Deep Learning - v2.png`** — the newest sheet; Represent → Relate → Parameterize → Learn → Orchestrate, with a dated paper list.
4. **`cc_(2026.09.09_ChatGPT) From VAE to Diffusion Models to Flow Matching - v2.png`** — VAE vs diffusion side by side (skip the Flow Matching column and all the equations).

Also useful, lower priority: `The Encoder-Decoder Architecture - A Timeline & Genealogy`, `Master Timeline of AI - v3`, `From Calculus to Backprop`, `Generative AI - Literacy Roadmap - v2`, and the two 3Blue1Brown chapter summaries.

The **`(folded) Generative AI - Literacy Roadmap=160.xmind`** is Nacaskul's own map of the whole course (dated 2026.09.14) — four branches: *Words Matter* (G/A/I, Literacy, Roadmap) → *Timelines* (task-equivalence, computational-model, integrated-autonomy) → *HIWUTH* (Predictive: MLP/CNN; Transition: VAE/U-Net; Generative: Transformer/Diffusion) → *Application* (Chatbot "talks with you", RAG "searches for you", Agent "does things for you", Wiki "builds knowledge with you"). **That last four-way application split is a neat, memorable, quotable line and is cheap to learn.**

---

## 11. Open questions — things to check or get for me

1. **The exam date and time.** Nothing in the folder states it. My inference is Friday **25 September 2026** (the session-8 slot, before the 2 Oct online class). **Please confirm on MyCourseVille — this is the one thing that changes everything.**
2. **Questions 1–10 of the 2025 midterm are missing.** The PDF excerpt begins at "[Page 2]" with Q11. Page 1 would show ten more questions *and* would tell us what the lecturer puts at the *start* of a paper — most likely the G/A/I definitional questions. **If you can get page 1, it is the single highest-value item left.**
3. **Who writes the paper?** The format slide (10 questions, 2 lines) is Niinimaki's; the two past papers are Nacaskul's 25-question style. If it is jointly set, expect a mix — some conceptual (Nacaskul) and some "how neural networks learn" (Niinimaki). I have weighted for both. If you know who sets it, tell me and I'll re-rank.
4. **Are bonus questions still a thing?** 2025 had two. With only 10 questions this year, I don't know. If they are, the ethics/playlist-video bonus is highly likely.
5. **Links I did not fetch.** Per your instruction I did not chase the Google Slides link — confirmed to be lecture 6 = `Gen AI Lit 06.pdf`, which I read. I did **not** open the YouTube videos, the Cornell VAE demo (`cs.cornell.edu/courses/cs4782/2026sp/demos/vae/vae_viewer.html`), the OpenAI "how our models are developed" page, or the Wolfram U CNN course. I inferred their content only from titles, URLs and your quiz notes, and have said so where it matters. The two I'd most want: **the Wolfram U CNN course outline** (Niinimaki says it's examinable) and **the MyCourseVille supplementary-material page** (in case it holds a study guide the folder lacks).
6. **`6901009256_hwk3.nb`** is the Mathematica `Manipulate` sliders exercise on linear regression (y = mx + b and its relation to a simple neuron) — no exam content beyond "a linear model is a perceptron without the activation". Confirmed by reading it as text; nothing else in it.
7. **Whether "count the parameters" survives.** Nacaskul flagged it twice in lecture 3; Niinimaki says "very little math". A single-line count (e.g. "2 inputs + 1 output ⇒ 2 weights + 1 bias = 3 parameters") is my best guess at the most it could be. **[GUESS]**
