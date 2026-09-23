---
id: autoencoders
title: "Autoencoders, VAEs and Latent Space"
deck: "GAI Lit 06/07"
slides: 17
summary: "The bottleneck, the latent space, and the single change that turns a reconstructor into a generator — the hinge of the whole course."
---

## What the exam will ask from this lecture

The blueprint names **VAE** explicitly, and the examiner's own lab was an autoencoder with the
prompt **"What is a latent space?"** written on the slide. This is also where the course turns:
everything before it predicts, everything after it generates, and the autoencoder is the hinge.

Two questions to be ready for above all others: **what is a latent space?** and **what makes a
VAE generative when a plain autoencoder is not?**

## The autoencoder

An **autoencoder** is trained to **reproduce its own input**: x → **encoder** → a compressed
code → **decoder** → x̂, with the loss measuring how far x̂ is from x.

That sounds pointless until you notice the constraint. The middle layer — the **bottleneck** —
holds **fewer numbers than the input**, so the network **cannot copy the input through**. It has
to throw something away, and to reconstruct well it must throw away the **accidents** and keep
the **structure**. The bottleneck is not a limitation of the design; it *is* the design.

Two consequences worth stating:

- **No labels are needed.** The target is the input itself, so an autoencoder can be trained on
  any pile of unlabelled data. This is **dimensionality reduction** done by learning.
- **The compression is lossy and specific.** A zip file is lossless and general — it restores
  any file exactly. An autoencoder learns what matters **in one kind of data** and discards the
  rest, which is why the code is a **representation** rather than a copy.

`[slides 5-8]`

?check id=ae-002

## Latent space

The **latent space** is the compressed internal representation at the bottleneck: each point
stands for an input, and **nearby points stand for similar inputs**. Its dimensions are
**learned**, not chosen, and it is **continuous** — which is the property that matters.

Train an autoencoder on handwritten digits with a two-dimensional latent space and the points
arrange themselves into clusters by digit. Walk a straight line from one cluster to another and
decode as you go, and you get a sequence of images **morphing smoothly from one digit into the
other**. Every point in between decodes to *something*, including points no training image ever
occupied.

> **Exam focus.** *"What is a latent space?"* is written on the examiner's own lab slide. Two
> lines: the compressed representation the encoder produces and the decoder reads, in which
> position encodes meaning and nearby points decode to similar outputs.

`[slides 6, 8]`

?check id=ae-004

**The picture used in class.** Houses scattered along a winding road can be located by **how
far along the road** each one sits, rather than by two map coordinates. One number instead of
two, and nothing you needed is lost — what is thrown away is the **empty space** the houses
never occupied. Ten inputs squeezed to two and expanded back is the same move: *it does not
take ten degrees of freedom to carry the information*.

**And the risk that follows.** The latent space stands for *our sense of what is realistic*, so
generating a new example means picking a random point in it. But a plain autoencoder is only
ever asked to reproduce the points it was shown, so a small step away from a real example may
land inside the region of realistic data or **outside** it — and the space between examples can
decode to nonsense. That is exactly the gap the VAE's regularisation term closes.

?check id=ae-031
?check id=ae-032

## AE vs VAE — the distinction the marks sit in

| | **Autoencoder** | **Variational Autoencoder** |
|---|---|---|
| The encoder outputs | **One fixed point** per input | A **probability distribution** over the latent space |
| The latent space | Has **gaps**; a point in a gap decodes to nothing meaningful | Trained to resemble a **simple distribution**, so it is filled in |
| What you can do with it | **Reconstruct** an input you supply | **Sample** a new point and decode it into a **new, realistic, previously unseen example** |
| Verdict | Representation | **Generative** |

Generation is sampling. After a VAE is trained, nothing is encoded at generation time at all:
you draw **z from the simple latent distribution** and push it through the **decoder**. That is
precisely the course's one-line summary of generative AI — **learn z → x, then crank the
PRNG**.

The VAE's training objective therefore has **two competing parts**: a **reconstruction** term
keeping the output close to the input, and a **regularisation** term keeping the latent
distribution close to a simple prior. Reconstruction alone would scatter the codes and leave
the space full of holes; regularisation alone would collapse everything to one blur. You need
both — a space you can sample from **and** a decoder that makes something of what you sample.

`[slides 8-9]`

?check id=ae-006

## "Our first generative AI?"

The examiner's slide asks this about the autoencoder lab, and the honest answer is a qualified
yes. The autoencoder is the first architecture in the course that **produces an output
resembling its training data** and contains a **decoder you could in principle sample from**. But
a plain AE only reconstructs. It takes the VAE's **probabilistic latent space** to make it
genuinely generative.

Useful non-generative use of the same machinery: **anomaly detection**. Train on correct parts
only, watch the **reconstruction error** in production; a defective part reconstructs badly and
the error spikes, because the latent space has room only for the regularities of normal parts.

`[slides 7-8]`

## Encoder–decoder in general

The autoencoder is a special case where the target happens to be the input. In general,
**encoder–decoder** means **compress to a latent, then expand to the target** — and the target
need not be the input:

| Task | Encoder input | Decoder output |
|---|---|---|
| Autoencoding | Image | The same image |
| **Machine translation** | Sentence in English | The sentence in French |
| **Segmentation** | Scan | A per-pixel mask |
| Denoising | Noisy image | Clean image |

Translation is where the idea became famous: the course's milestone list records
**encoder–decoder in 1997** and **"deep learning" in 2014**, used by **Google Translate**.

`[slides 5, 10]`

## The encoder–decoder genealogy

One definition covers the whole family:

> **An encoder–decoder architecture converts an input into a representation, and then uses
> that representation to construct an output.** x → z = E(x) → y = D(z).

Same pattern, many learning objectives and data types. The variants, and what distinguishes
each:

| Variant | Target | Distinguishing feature |
|---|---|---|
| **Autoencoder (AE)** | y = x | Usually a **dimensional bottleneck**; learns compact representations |
| **Denoising autoencoder (DAE)** | y = the *clean* x | Input is **deliberately corrupted**; forces **robust** features |
| **Variational autoencoder (VAE)** | y = x | z is a **distribution**, not a point — which is what enables generation |
| **General encoder–decoder** | **y ≠ x** | **No bottleneck required**; works for any data type — translation, summarisation, segmentation, captioning |
| **Attention-based encoder–decoder** | y ≠ x | The decoder attends to a **sequence** of encoder states; **removes the fixed-length bottleneck** |
| **Transformer encoder–decoder** | y ≠ x | Attention only; **self-attention** within each side, **cross-attention** from decoder to encoder |
| **U-Net** | y ≠ x | **Skip connections** preserve fine spatial detail |
| **Diffusion** | x̂ from noise | An **iterative** encoder–decoder (usually a U-Net), learning a generative transformation |

> **Exam focus.** The cleanest distinction here: **an autoencoder requires y = x and usually a
> bottleneck; a general encoder–decoder requires neither.** The bottleneck is a feature of the
> special case, not of the pattern.

**The timeline.** **1980s** — auto-associative networks trained to reproduce their own inputs,
whose bottleneck hidden layers were found to hold learned representations (Rumelhart, Hinton &
Williams 1986; Bourlard & Kamp 1988). **2006** — **deep autoencoders**, stacked for nonlinear
representation learning (Hinton & Salakhutdinov). **2008–2013** — **denoising** and **sparse**
variants, for robustness. **2013–2014** — the **general encoder–decoder for sequence
transduction**, i.e. neural machine translation (Cho et al.; Sutskever, Vinyals & Le).
**2014–2015** — **attention and the end of the fixed bottleneck** (Bahdanau et al.; Luong et
al.). **2017+** — the **Transformer encoder–decoder** scales the idea to large models.
**2015+** — beyond text: U-Net, image-to-image, VAEs, diffusion, multimodal.

**Attention's contribution, stated as the distinction it is:** before it, the decoder read a
**single fixed-length vector** summarising the whole input, and the longer the input the more
was lost. After it, the decoder **attends to every encoder state, weighted per output element**
— so nothing has to be squeezed through one code. Self-attention in the Transformer is the
direct descendant.

A caution the course is explicit about: **sequence data gave the encoder–decoder its historical
prominence, but the architecture is not restricted to sequences.** The same pattern runs on
text, images, audio, graphs and video. Sequences are historical, not fundamental.

`[slides 5, 10]`

?check id=ae-025
?check id=ae-027

## U-Net

A **U-Net** is an **encoder–decoder CNN with skip connections**. The encoder downsamples, the
decoder upsamples, and at each level a **skip connection** carries the encoder's output straight
across to the matching decoder stage — which is where the **U** in the diagram comes from.

Why it is needed: compressing to a bottleneck destroys exactly the **pixel-level precision** a
segmentation mask requires. The skip connections hand that detail back, so the output combines
**high-level context from the bottleneck** with **fine spatial information from the early
layers**. Typical tasks: **image segmentation, reconstruction, denoising, medical imaging** — and, as named in class, **autonomous vehicles**, which need the scene labelled pixel by pixel in real time.

> **Exam focus.** *Encoder–decoder vs U-Net* is a distinctions-table row. Encoder–decoder
> compresses then expands; a U-Net is an encoder–decoder CNN whose **skip connections** carry
> fine spatial detail across the bottleneck.

`[slides 11-13]`

?check id=ae-011

## Where this sits on the arc

**VAE and U-Net are the representation bridge** between predictive AI (MLP, CNN) and generative
AI (Transformer, Diffusion). The bridge's job is to **learn compact, structured representations
that capture the underlying factors of the data**, which is what makes both prediction and
generation possible from the same foundation. Its question is *"what is the underlying
structure?"*, against prediction's *"what is most likely?"* and generation's *"what could be?"*

One forward-looking distinction, worth having ready because it is a table row in its own right:
**a VAE learns both directions** — an encoder into the simple distribution and a decoder back
out. **A diffusion model prescribes the forward noising and learns only the reverse.**

`[slides 14-17]`
