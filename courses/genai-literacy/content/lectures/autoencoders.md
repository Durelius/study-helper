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

## U-Net

A **U-Net** is an **encoder–decoder CNN with skip connections**. The encoder downsamples, the
decoder upsamples, and at each level a **skip connection** carries the encoder's output straight
across to the matching decoder stage — which is where the **U** in the diagram comes from.

Why it is needed: compressing to a bottleneck destroys exactly the **pixel-level precision** a
segmentation mask requires. The skip connections hand that detail back, so the output combines
**high-level context from the bottleneck** with **fine spatial information from the early
layers**. Typical tasks: **image segmentation, reconstruction, denoising, medical imaging**.

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
