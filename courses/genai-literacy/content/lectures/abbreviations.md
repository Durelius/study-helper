---
id: abbreviations
title: "Abbreviations and Hard Terms"
deck: "lec 1 + lec 5 + GAI Lit 03-07 + infographics"
slides: 30
summary: "A lookup-and-drill sheet: every acronym in the course expanded, the letter that carries the meaning, a one-line gloss, and the neighbouring term it gets swapped with."
---

## How to use this sheet

This is not another lecture. It is the table you check when a letter goes blank. Each row
gives four things: the **expansion**, the **letter that carries the meaning** (the one a
question will ask you to explain), a **one-line gloss**, and the **term it gets confused with**.

The 2025 paper asked *"What does MLP stand for? Explain the P"* and *"What does CNN stand for?
Explain the C"*. That format is Nacaskul's and may not recur under Marko, but the two-line
shape still works for any acronym: **line one, the expansion; line two, what the key letter
means and what it rules out.**

> **Exam focus.** Spelling out the acronym is worth the least. The mark is in the second line:
> what the letter is, plus the clause that tells it apart from its neighbour.

`[2025 midterm, slide 2]`

## The letters of "AI" itself

| Term | Expansion | The letter that matters | One-line gloss | Confused with |
|---|---|---|---|---|
| **AI** | Artificial Intelligence | **I** = human-like; **A** = artifact, semblance or artful creation | Approximating the interactive or behavioural intelligence of a species | ML, as if one contained the other |
| **GAI** | Generative AI | **G = Generative** | Captures a dataset's probability distribution so it can synthesise new samples that look like they came from it | **AGI** |
| **AGI** | Artificial General Intelligence | **G = General** | Human-level breadth across every task; **out of scope** for this course | **GAI** |
| **ML** | Machine Learning | — | Calibrating model parameters from data instead of writing if-then rules | AI: **neither contains the other** |
| **BI** | Business Intelligence | **I = advantageous information** | Extracted information of tactical or strategic value, as in *military* intelligence | The I in AI |
| **CI** | Computational Intelligence | **I = an optimal solution path** | A clever or optimal method for a hard problem (also "digital intelligence") | The I in AI |
| **GOFAI** | Good Old-Fashioned AI | "Old-Fashioned" marks it as a **retronym** | Symbolic, rule-based AI (expert systems, MYCIN), renamed once neural nets took the name "AI" | Modern AI, which learns its parameters from data |
| **BNN** | Biological Neural Network | **B = Biological** | The real brain: nerve cells passing stimulus to each other | **ANN** |
| **ANN** | Artificial Neural Network | **A = Artificial** | The mathematical model of a BNN; every modern AI is built from these | **BNN** |

**The three I's, as a list of three:** human-like (Turing's imitation game, the I in AI) ·
advantageous information (military, business) · optimal solution path (computational,
digital). **The three A's:** artifact / artifice (made by humans) · artificial semblance
("seems about right" will do) · artful creation ("truth manufactured").

**The G changed meaning.** Before 2014 "generative" was a term in statistics. The
**Generative Adversarial Network** made the word popular, and GPT then inherited it.

`[lec 1, slides 15-24, 53, 59, 63, 76]`

?check id=abb-002
?check id=abb-004
?check id=abb-006

## Network building blocks

| Term | Expansion | The letter that matters | One-line gloss | Confused with |
|---|---|---|---|---|
| **Perceptron** | (not an acronym) | — | A model of one generic neuron: **sums weighted excitatory or inhibitory inputs and fires past an internal threshold** | The MLP, which stacks many of them |
| **MLP** | **Multi-Layer Perceptron** | **P = Perceptron** | Feedforward layers of perceptrons, each layer's output the next one's input; **no recurrence, no skipping layers** | A single perceptron, which cannot learn XOR |
| **FFN** (1) | Feedforward Network | **FF = feedforward** | The signal travels **one way**, input to output. An MLP is one | **RNN** |
| **FFN** (2) | Feed-forward block | same letters, second meaning | The per-token MLP **inside each Transformer block** (up-projection, ReLU, down-projection), about **two thirds of the parameters** | Attention, which moves information *between* tokens |
| **RNN** | Recurrent Neural Network | **R = Recurrent** | The signal **loops back**, so earlier outputs feed later computation | FFN; also ResNet (the R there is *Residual*) |
| **LSTM** | Long Short-Term Memory | the **gates** | A gated recurrent network (Hochreiter & Schmidhuber, 1997) that stops the error signal decaying over long sequences | A plain RNN; later replaced by attention |
| **FLN** | Functional-Link Network | **FL = functional link** | Pao, 1995: inputs augmented with explicit nonlinear transforms, so it is expressive while staying shallow | FFN, which is the letters in a different order |

**XOR** is the logic function no single straight line can separate. That is why a single
perceptron fails on it, and why the M in MLP (the layers) is needed.

> **Exam focus.** *"Explain the P"* means explain the **perceptron**, not the parameters or the
> propagation. Model answer: *Multilayer Perceptron. The perceptron is a mathematical model of
> a generic neuron that sums weighted incoming signals and fires when the sum exceeds a
> threshold.*

`[2025 midterm, slide 2; lec 5, slide 1; GAI Lit 03, slide 9; infographics, slides 8, 12, 14]`

?check id=abb-009
?check id=abb-010
?check id=abb-013

## The architectures

The six the course names, in learning order: **MLP → CNN → VAE → U-Net → Transformer →
Diffusion.** The six run from predictive (MLP, CNN), through a representation bridge (VAE, U-Net),
to generative (Transformer, Diffusion). **The GAN is not one of the six.**

| Term | Expansion | The letter that matters | One-line gloss | Confused with |
|---|---|---|---|---|
| **CNN** | Convolutional Neural Network | **C = Convolutional** | Layers of **convolution matrices with trainable entries**, slid across a pixel grid to detect visual patterns; local filters and weight sharing | An MLP fed raw pixels |
| **AE** | Autoencoder | **auto** = the target is the input itself | Reproduce your own input through a narrow **bottleneck**; needs no labels | VAE |
| **DAE** | Denoising Autoencoder | **D = Denoising** | Reconstruct the **clean** input from a deliberately **corrupted** copy, so the features must be robust | Diffusion, which reuses the idea step by step |
| **VAE** | Variational Autoencoder | **V**: the encoder outputs a **probability distribution**, not a point | Sample a new latent point, decode it, and get a new example. The AE only reconstructs; the VAE generates | AE; also diffusion |
| **SAE** | Sparse Autoencoder | **S = Sparse** | A wide, mostly-inactive re-encoding of a model's activations, used to pull superposed concepts apart | VAE |
| **U-Net** | (named for its shape) | **U** = the diagram's shape: down one side, up the other | An **encoder-decoder CNN with skip connections** carrying fine spatial detail across the bottleneck; segmentation, medical imaging | A plain encoder-decoder; the VAE |
| **GAN** | Generative Adversarial Network | **A = Adversarial** | A **generator** makes fakes and a **discriminator** judges real from fake, in a minimax game (Goodfellow, 2014) | VAE, which learns by reconstruction instead |
| **Transformer** | "Self-Attention Model" | — | Token embeddings + positional encoding + self-attention + feed-forward blocks; no recurrence, no convolution (2017) | RNN/LSTM |
| **GPT** | Generative Pre-trained Transformer | **G** generates new text · **P** pre-trained on internet-scale text before any task · **T** the Transformer | Decoder-only | BERT; the P is *not* Perceptron |
| **LLM** | Large Language Model | **L = Large** | A very large transformer pre-trained to predict the next token | GPT, which is one example |
| **BERT** | (letters not unpacked in the course) | — | Devlin et al., 2018: a Transformer that **reads in both directions**, built for understanding tasks | GPT, which generates |
| **ViT** | Vision Transformer | **V = Vision** | A pure Transformer on sequences of **image patches**, no convolution (2020/21) | CNN |
| **Seq2Seq** | Sequence-to-sequence | — | An LSTM squeezes a whole sequence into **one fixed-length vector** and a decoder expands it (2014) | Attention, which removed that fixed bottleneck |
| **ResNet** | Residual Network | **Res = Residual** | Blocks add their input back to their output (H(x) = F(x) + x), trainable to 152+ layers (2015) | RNN; the U-Net's skip connections |
| **Diffusion** | "Denoising Diffusion Model" | — | Learn to **reverse a fixed noising process**: start from noise and iteratively denoise | VAE, which learns both directions |
| **DDPM** | (letters not unpacked in the course) | — | Ho, Jain & Abbeel, 2020: a diffusion model trained with a denoising objective | — |

`[2025 midterm, slide 2; lec 5, slide 7; lec 1, slide 70; GAI Lit 03, slide 5; GAI Lit 06, slide 8; infographics, slides 3-4, 7-8, 10-11, 14]`

?check id=abb-020
?check id=abb-023
?check id=abb-027

## Training and loss terms

| Term | Expansion | The letter that matters | One-line gloss | Confused with |
|---|---|---|---|---|
| **Loss (L)** | — | — | The single number saying how wrong the network is; training minimises it | RMSE, which is one instance of it |
| **RMSE** | Root Mean Squared Error | — | **One specific loss**, natural when the target is a **quantity** | The loss function in general; cross-entropy |
| **Cross-entropy** | — | — | The loss when the target is a **class**: it compares predicted probabilities with the one-hot label | RMSE |
| **Gradient (∇)** | — | — | The direction of **steepest ascent** of the loss, so you step **against** it | The loss itself |
| **GDA** | Gradient Descent Algorithm | **D = Descent** | Iterative and stepwise: find the downhill direction, take a small step, repeat | Backprop |
| **Backprop** | Backpropagation | **back** = error sent backwards | *"A specialisation of the GDA, applied to minimise prediction error for an MLP by tweaking weights and biases."* It **computes** the gradient; GDA **takes the step** | GDA; the optimiser variants |
| **SGD** | Stochastic Gradient Descent | **S = Stochastic** | One of the named variants: batch, stochastic, mini-batch, momentum, RMSProp, Adam | Backprop, which is **not** a variant |
| **ADAM** | Adaptive Moment Estimation | **A = Adaptive** | Gradient descent that adapts each step size (Kingma & Ba, 2014/15) | Backprop |
| **ReLU** | Rectified Linear Unit ("the ramp") | — | Returns the input if positive, else 0; the cheap modern default | Sigmoid, which it displaced |
| **Sigmoid** | — | — | Squashes one value into (0, 1); smooth, so it is differentiable, but gradients vanish in deep nets | Softmax; the hard step |
| **Softmax** | — | — | Turns a **vector** of raw scores into a probability distribution that **sums to 1** | Sigmoid, which acts on a single value |
| **η** | learning rate | — | The step size in w ← w − η∇L; a **hyperparameter** | A parameter (weights and biases are learned) |
| **UAT** | Universal Approximation Theorem | — | **One hidden layer** with finitely many neurons can approximate any continuous function on a closed bounded region (Cybenko, 1989) | "Deeper always fits better" |

**Parameter / weight / coefficient** are one thing, the numbers **learned**. **Inputs** are
supplied; **hyperparameters** are set before training.

> **Exam focus.** Two of the most-swapped pairs sit here: **GDA vs backprop** (general
> optimiser vs its specialisation to a layered network) and **loss vs RMSE** (the general
> idea vs one choice of it).

`[lec 5, slides 2-5; infographics, slides 12-13; Wolfram CNN, slide 8]`

?check id=abb-036
?check id=abb-038

## Generative-model shorthand

| Term | Expansion | One-line gloss | Confused with |
|---|---|---|---|
| **PRNG** | Pseudo-Random Number Generator | What you "crank" to get a new sample: randomness gives the variation, the learned map gives the structure | The prompt |
| **Prompt vs PRNG** | — | **The prompt is WHERE to navigate; the PRNG is WHICH possibility** | — |
| **TO / FRO** | the two mappings | **TO**: data → simple distribution. **FRO**: simple → data. **Generation always uses FRO**; TO only sets up the homework | Encoder / decoder, loosely |
| **ELBO** | Evidence Lower Bound | The VAE's objective: a **reconstruction** term plus a **regularisation** term | Reconstruction loss alone |
| **p(y\|x)** | discriminative | The probability of a label given an input: **predictive** | p(x) |
| **p(x), p(x\|c)** | generative | The distribution of the data itself, optionally given a condition c | p(y\|x) |
| **z** | the latent | A point in **latent space**, the compressed code at the bottleneck | x, the data |

The course's one-liners: predictive AI is **learn x → y, then crank the input**; generative AI
is **learn z → x, then crank the PRNG**.

`[infographics, slides 1-2, 4]`

?check id=abb-044
?check id=abb-046

## Language and application shorthand

| Term | Expansion | One-line gloss | Confused with |
|---|---|---|---|
| **NLP** | Natural Language Processing | Getting a computer to handle human language; METEO (1981) is the early example that shipped | — |
| **Token** | — | The unit text is chopped into before embedding: a word or a fragment of one | A letter |
| **Embedding** | — | A learned vector giving a word its **global-context** meaning | Attention |
| **Attention** | — | Weighs how much each other token bears on this one: **local-context** meaning, *in this sentence* | Embedding |
| **Q, K, V** | Query, Key, Value | **Q**: what a token is looking for. **K**: what each other token offers. **V**: the content passed along | — |
| **word2vec** | — | Mikolov, 2013: words as vectors, so *king − male + female ≈ queen* | The Transformer |
| **RAG** | Retrieval-Augmented Generation | Retrieve relevant documents first, then generate grounded on them. Belongs to the term project, not this exam | Plain generation |
| **HIWUTH** | How It Works Under The Hood | Stage 3 of the literacy roadmap: the six architectures | Stage 2, Timelines |

The four application modes, one verb each: **Chatbot talks with you · RAG searches for you ·
Agent does things for you · Wiki builds your knowledge.**

`[lec 1, slides 68, 70; GAI Lit 02, slide 14; GAI Lit 07, slides 2, 7; infographics, slide 15]`

?check id=abb-048
?check id=abb-050

## Hard terms that are not acronyms

| Term | One-line meaning | Confused with |
|---|---|---|
| **Retronym** | A new name for an older thing, needed because a newer thing took the original name (GOFAI; "predictive AI") | Acronym |
| **Verifiable ground truth** | A known right answer to score against; the standard for **predictive** AI | Plausible context relevance |
| **Plausible context relevance** | Convincing and relevant to its context; the standard for **generative** AI | Accuracy |
| **Interpolative vs extrapolative** | Recognition is an **interpolative** map within known territory; generation is an **extrapolative** map into unexplored territory | Each other |
| **Latent space** | The compressed representation at the bottleneck, where nearby points are similar inputs | Embedding space, loosely |
| **Bottleneck** | The narrow middle layer that forces the network to keep structure and discard accidents | Skip connection, which bypasses it |
| **Skip / residual connection** | A U-Net's skip carries encoder detail across to the decoder; a ResNet's residual adds a block's input back to its output | Each other |
| **Degradation** | A deeper net is worse **even on the training set**: the optimiser fails, not the capacity | Overfitting, where only validation error rises |

`[lec 1, slides 32, 59; 2024 midterm, slide 1; GAI Lit 07, slide 6; infographics, slides 3, 10]`

?check id=abb-052
?check id=abb-053
