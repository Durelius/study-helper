---
id: generative-models
title: "The Generative Architectures Compared"
deck: "infographics + lec 5"
slides: 18
summary: "The six foundational architectures on one arc, and how GAN, VAE, diffusion and flow matching differ as learning games with the same goal."
---

## What the exam will ask from this lecture

The blueprint line is **"Architectures: Multi-Layer Perceptron, Convolutional NN, VAE, .."**,
and the "…" is doing work — Transformer and Diffusion sit under it. Class notes record the six
foundational architectures twice, both times with the instruction **"be able to tell them
apart"**. Expect two to four of the ten questions from this territory.

## The six, in order

**MLP → CNN → VAE → U-Net → Transformer → Diffusion.**

| # | Architecture | Key idea | Typical tasks |
|---|---|---|---|
| 1 | **MLP** (Multi-Layer Perceptron) | The foundational network: fully connected layers learn nonlinear mappings from input to output | Regression, classification, tabular prediction |
| 2 | **CNN** (Convolutional Neural Network) | Exploit spatial structure with **local filters and weight sharing** | Image classification, object detection |
| 3 | **VAE** (Variational Autoencoder) | Encode into a **probabilistic latent space**, then decode to generate new, similar examples | Representation learning, anomaly detection, data generation |
| 4 | **U-Net** (Encoder–Decoder CNN) | **Skip connections** preserve fine spatial detail across the bottleneck | Segmentation, reconstruction, denoising, medical imaging |
| 5 | **Transformer** (Self-Attention Model) | **Self-attention** models long-range dependencies and scales | Language modelling, text generation, multimodal, LLMs |
| 6 | **Diffusion** (Denoising Diffusion Model) | **Learn to reverse noise** — gradually denoise randomness into data | High-quality image generation, text-to-image, editing |

And the three stages they fall into:

| Stage | Architectures | Its question |
|---|---|---|
| **Predictive AI** | MLP, CNN | **"What is most likely?"** — discriminative, p(y\|x) |
| **Representation bridge** | VAE, U-Net | **"What is the underlying structure?"** |
| **Generative AI** | Transformer, Diffusion | **"What could be?"** — generative, p(x) or p(x\|c) |

> **Exam focus.** *"Name the six foundational architectures, from predictive to generative."*
> List them, then add the three-stage grouping in the second line. Note that the **GAN is not
> one of the six** even though it matters historically — a tempting wrong answer.

`[slides 1-6]`

?check id=gen-001

## The key takeaway

> **The same core idea — neural networks that learn useful representations — powers both
> predictive AI and generative AI. Each architecture adds a new capability, moving us from
> understanding the world to creating new possibilities.**

Nothing was abandoned along the way. A diffusion model still contains a CNN; a transformer
still learns by gradient descent. That is why the arc is a progression rather than a sequence
of replacements.

## What "generative" means, precisely

The definition the examiner put on his own slide, twice:

> **Generative AI refers to systems that capture the underlying probability distribution of a
> dataset in order to synthesize new data samples that appear to originate from that same
> distribution.**

That is why the evaluation standard has to be **plausible context relevance** rather than
**accuracy against verifiable ground truths**: a newly generated face has nothing to be
accurate against. It is also why the characteristic failure is a hand with six fingers or a
citation to a paper that does not exist — **locally convincing, globally wrong**.

`[slides 7-8]`

?check id=gen-021

## GAN — the adversarial origin of the "G"

Two networks in a **minimax game**:

- The **generator** makes fakes from random input.
- The **discriminator** judges real against fake.

Each improvement raises the bar for the other. The training signal is not a label but the
opponent's judgement, which is what makes it **adversarial**. Goodfellow and colleagues, **2014**.

**The irony, and it was flagged in class as an exam note:** the adversarial setup was devised as
a device to **sharpen a predictive, discriminative model** — the generator was there to give the
discriminator harder cases. The generator became the star instead, **deepfakes** made it famous,
and "generative" took on its modern popular meaning from this architecture before GPT inherited
the word.

`[slides 9-10]`

?check id=gen-005

**How the idea was posed in class**, as a second question following a first:

1. *Can a model look at a photograph of a person and produce the judgement that it is that
   person?* That is **recognition** — a predictive model.
2. *Can another engine produce an image good enough that the first model judges it a real
   photograph of them?* That is the **GAN**.

The judge was built as a predictive model, and turning its verdict into a training signal for a
generator is what makes the adversarial network both ironic and effective.

`[slides 8-10]`

?check id=gen-033

## What carried the field from predictive to generative

A question recorded in class in those words, and worth two rehearsed lines:

> Networks first learned **compact representations of what the data are**. Once that
> representation existed, the same machinery could be run the other way — sample the simple
> distribution and decode — so **prediction and generation turned out to share one engine**.

Nothing was discarded in the move. **Universal approximation** and **gradient-based learning**
sit under both halves of the arc; what the VAE and the U-Net added was the **representation
bridge**, and what the Transformer and the diffusion model added was the means to exploit it.

The warning that goes with it: generation **extrapolates beyond the basis of validation**. That
is precisely what makes it creative, and precisely what means **there is no ground truth left
to check the output against** — which is why a generative model is judged on plausible context
relevance rather than accuracy.

`[slides 1-3]`

?check id=gen-034
?check id=gen-035

## Diffusion

A **fixed forward process** gradually corrupts data with **Gaussian noise**, step by step. The
model **learns to reverse it** — to denoise. Generation starts from **pure noise** and
iteratively denoises, optionally **conditioned on a text embedding** (Stable Diffusion, DALL·E).

The crucial asymmetry: **the forward process is prescribed, not learned.** Because the
corruption was applied deliberately, the true noise at every step is **known**, which turns
generation into an ordinary supervised problem repeated many times over.

> **Exam focus.** The distinctions table row: **a VAE learns both directions** — an encoder into
> the simple distribution and a decoder back out. **A diffusion model prescribes the forward
> noising and learns only the reverse.**

`[slides 11-13]`

## Three learning games, one goal

| | **VAE** | **Diffusion** | **Flow matching** |
|---|---|---|---|
| **Core idea** | Learn **both** directions: encode to a simple distribution, decode back | Add noise by a **known** process, learn to **reverse** it | Define a **probability path** from noise to data and learn the **vector field** that transports samples along it |
| **After training** | Sample z from the simple distribution, decode | Start from pure noise, **iteratively denoise** | Start from noise, **solve the learned ODE** |

The unifying view, which is the whole sheet in one sentence:

> **VAE, diffusion and flow matching are three different learning games with the same goal — a
> learnable generator mapping a simple distribution to the data distribution.**

Each one **constructs a mapping** between data and a simple distribution, which **creates a
supervised learning task**, and then **samples from a PRNG and applies the learned
transformation** to make something new.

Why it works at all: **data live on a low-dimensional, structured manifold.** Real images
occupy a tiny, highly structured corner of all possible pixel arrangements, which is what makes
it possible to map a simple distribution onto them — the same assumption behind the
autoencoder's bottleneck. **Do not study flow matching beyond the line above.**

`[slides 14-17]`

?check id=gen-011

## TO and FRO: what is chosen and what is learned

The cleanest single frame for the whole generative family. Every one of these models involves
**two mappings**:

- the **TO** mapping, **data → a simple distribution**;
- the **FRO** mapping, **simple distribution → data**.

**Generation always uses FRO.** TO exists only to set up the learning problem:

> **TO sets the homework; FRO learns to solve it backwards.**

Because *we* chose or constructed the trip out to the simple distribution, we know the right
answer at every step, so training a generator becomes an ordinary supervised task. That is the
trick the whole field turns on.

| | **TO** (data → simple) | **FRO** (simple → data) | Where the supervision comes from |
|---|---|---|---|
| **VAE** | **Learned** encoder q(z\|x) | **Learned** decoder p(x\|z) | Reconstruction + KL regularisation |
| **Diffusion** | **Fixed** forward noising | **Learned** reverse denoiser | The noise we added ourselves |
| **Flow matching** | **Chosen** probability path | **Learned** vector field | The target velocity we designed |
| **GAN** | **None** — the exception | **Learned** generator | A **discriminator's** verdict (adversarial) |

> **Exam focus.** The GAN is the odd one out and that is a ready-made distinction question:
> **it has no explicit data-to-simple mapping at all.** VAE, diffusion and flow matching each
> construct a route out and learn the route back; the GAN learns its generator purely from an
> adversary's judgement, which is why it trains so differently and so unstably.

The **VAE's objective is the ELBO — Evidence Lower Bound** — which splits into a
**reconstruction** term keeping the output close to the input and a **latent regularisation**
term keeping the code close to the simple prior. Those are the two competing halves.

`[slides 14-17]`

?check id=gen-026
?check id=gen-028

## Conditioning: where to navigate, which possibility

Text-to-image is the case worth having a sentence ready for. A **text encoder** turns the
prompt into an embedding **c**; a **PRNG** supplies a random **z**; the generator produces a
sample from **p(x | c)**.

> **The prompt is WHERE to navigate. The PRNG is WHICH possibility.**

Semantics directs; randomness varies. That is why the same prompt run twice gives two
different pictures, and why the randomness is not a defect to be engineered away — it is what
makes the model generative rather than a lookup table.

`[slides 14-17]`

?check id=gen-029

## Which to reach for, and what you give up

- **VAE** — a compact, structured latent code you can inspect, interpolate in and sample
  cheaply, plus an explicit likelihood. Sample quality is the weak point.
- **Diffusion** — the best empirical sample quality and scaling, conditioned naturally on
  text. You give up the compact latent and pay in a long iterative sampling loop.
- **Flow matching** — a simple continuous-time view that unifies the other two. One line only.

A caveat the course is careful about: data are said to live on a **manifold**, a tiny
high-density region of an enormous space. That is a **useful modelling intuition**; the safer
statement is a **concentrated high-density data distribution**, not a proven low-dimensional
linear subspace.

`[slides 14-17]`

?check id=gen-032

## Can we have generative AI without neural networks?

One of the examiner's own bonus questions. The answer is **yes in principle, no in practice**:

> A Markov chain generates text from next-token statistics, and any sampler from a fitted
> distribution is generative. But without a learned representation it has no context, so the
> output is **locally plausible and globally incoherent**.

That answer is worth rehearsing because it tests whether you know what "generative" actually
means — a property of sampling from a distribution, not a property of neural networks.

`[slide 18]`

## Telling the generative pair apart

If asked which architecture suits a task, the split is clean:

- **Sequences and language** — the **Transformer**, because attention models long-range
  dependencies and scales to enormous corpora.
- **High-quality images from noise** — the **Diffusion** model, because iterative denoising
  produces realistic detail and takes conditioning from a text embedding naturally.
- **Unpaired image-to-image translation** (summer photos into winter ones, with no matched
  pairs) — a **Cycle GAN**, because adversarial training needs no reconstruction target.
