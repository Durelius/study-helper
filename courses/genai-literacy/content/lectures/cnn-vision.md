---
id: cnn-vision
title: "CNNs and Machine Vision"
deck: "Wolfram CNN + lec 5"
slides: 24
summary: "Why convolution exists, what natural vision contributed to the design, and the layer-by-layer glossary the examiner's own assigned CNN course is built on."
---

## What the exam will ask from this lecture

Two independent reasons this topic is dense with marks. The blueprint names
**"Convolutional NN"** among the architectures, *and* it adds **"some topics covered in the
CNN course (assignment, Wolfram U), too"** — a course the examiner set and wrote the quizzes
for. So both the *why* (biological motivation, weight sharing) and the *glossary* (convolution,
ReLU, pooling, softmax, cross-entropy, dropout, receptive field) are live.

**Learn the concepts, not the function names.** The assigned course is Mathematica-based;
"no code" still applies.

## Why an MLP fails on pixels

Feed a photograph straight into an MLP and every pixel becomes its own input: x₁ through
x₂₅₀₀₀₀. Each input gets **its own weight**. So if the object moves three pixels to the left,
or the lighting changes its colour, **every number the network sees has changed** — and as far
as the model is concerned it is looking at a different thing entirely.

The fix is not more layers. It is to stop giving every position its own parameters.

`[slides 1-3]`

?check id=cnn-003

## The convolutional layer

A **convolutional layer** is a set of **convolution matrices — filters — with trainable
entries**, slid across the grid of pixel values to **detect underlying visual patterns**.

The consequence of sliding *one* filter over *every* position is **weight sharing**, and weight
sharing buys two things at once:

- **Translational covariance** — the network processes every region with the **same learned
  filters**, so a feature is detected consistently wherever it appears.
- A parameter count that depends on the **filter size and the number of filters**, not on how
  large the image is.

> **Exam focus.** The examiner's own question is *"Why did mobile phones start recognising
> images in the 2010s?"*, and the architectural half of that answer is here. The 2025 paper
> asked *"What does CNN stand for? Explain the C."* — **Convolutional Neural Network**, and the
> convolutional layer is a layer of convolution matrices with trainable entries applied to a
> grid of pixel values to detect visual patterns.

`[slides 5-6]`

## What natural vision contributes

The assigned course opens with the **human visual system**, and lists the properties of
natural images that the architecture is designed around:

| Principle | What it means |
|---|---|
| **Translational covariance** | Same filters everywhere, so position does not matter |
| **Scale covariance** | Objects appear at **different sizes depending on distance from the camera** |
| **Rotation invariance** | Turning the input should not change the label — *usually* |
| **Symmetry transformations** | Sometimes it *does* matter: flip **"bp"** and you get **"qd"** |
| **Inherent scale of features** | Features have a natural size; filters are sized to match |
| **Hierarchy of features** | **Leaves → trees → forest**; edges → shapes → objects |
| **Colour correlation length** | Neighbouring pixels have similar colours, and the similarity **falls off with distance** |
| **Motion detection** | Movement directs attention |
| **Feature binding** | The brain binds **spatially aligned** features into a single object |

The hierarchy row is the one that ties back to the timeline: **Hubel and Wiesel's cat** had
neurons that fired for **oriented lines**, and lines combine upward into curves, shapes and
objects. The CNN copies that ladder directly.

`[slides 1-2]`

?check id=cnn-007

## The layer glossary

These are the cheapest marks available: each is a clean two-line definition.

- **Encoder** — the component that converts raw images into the numeric **arrays/tensors** a
  network can consume. (A **tensor** is just the multi-dimensional array those numbers live
  in.)
- **Convolution layer** — trainable filters slid over the input to detect patterns.
- **ReLU / ramp** — returns the argument if positive, else zero. Placed after convolutional
  and linear layers to introduce **nonlinearity**, which is what lets the network learn complex
  decision boundaries.
- **Pooling** — **reduces the size of the representation**, summarising each neighbourhood into
  one value.
- **Flattening** — converts the **3-D tensor** from the convolutional stack into a **1-D
  vector** the fully-connected layers can take. Nothing is computed; the numbers are
  rearranged.
- **Softmax** — converts raw activations into a **probability distribution over classes**.
- **Cross-entropy loss** — the **distance between the predicted class probabilities and the
  true (one-hot encoded) label**.
- **One-hot encoding** — exactly one slot on, the rest off.

> **Exam focus.** Softmax and cross-entropy sit next to each other and are easy to confuse.
> Softmax is the **output transform**; cross-entropy is the **loss**. Say both halves.

`[slides 3-8]`

?check id=cnn-013

## Training pathologies and what to do about them

**Overfitting** shows as **training error continuing to fall while validation error plateaus or
rises**. **Underfitting** is both curves staying high.

Two remedies, and the distinction between them is exam-shaped:

| Remedy | Acts on | How |
|---|---|---|
| **Dropout** | The **network** | Randomly switches off a fraction of units during training, so the network cannot rely on any single path |
| **Data augmentation** | The **data** | Flips and small translations manufacture extra plausible training variation |

`[slides 9-10]`

?check id=cnn-018

## Looking inside a trained network

- **Receptive field** — the region of the input image a given neuron can actually **see**. Each
  layer pools information from the one below, so it **widens with depth** — which is exactly why
  early layers detect edges and deep layers detect whole objects.
- **Sensitivity map** — measured on a particular image, it shows **which input pixels most
  affect a given output**. The receptive field is architecture; the sensitivity map is
  measurement.
- **Input gradients** — the direction in pixel space that most changes the loss for a given
  class. The same machinery that trains the network, pointed at the input instead of the
  weights.
- **Optical illusions** — used in the assigned course to tie the whole thing back to biological
  vision: both systems are making inferences, and both can be fooled.

`[slides 11-14]`

## Beyond plain classification

The assigned course's later sections name architectures worth recognising even though the
details are out of scope: **Inception**, **ResNets**, **U-Nets** (lesson 17), **EfficientNet**,
**visual Transformers** (lesson 19), and **Cycle GAN** (lesson 23) among the applications —
image retrieval, object detection, reconstruction and segmentation.

Of these, **U-Net** matters most for this paper, because it is one of the six foundational
architectures: an **encoder–decoder CNN with skip connections** that preserve fine spatial
detail. See the autoencoder notes.

`[slides 15-24]`

## The two-line answers to have ready

- *Why can a CNN recognise an object wherever it appears, when an MLP cannot?* — A CNN slides
  the **same learned filters** over every region, so a feature is detected consistently
  regardless of position. An MLP gives every pixel position its own weight, so a shifted object
  looks entirely different to it.
- *What does the CNN borrow from biological vision?* — The **hierarchy of features** from Hubel
  and Wiesel's cat: line detectors combining into curve detectors and upward into shapes and
  objects, with the same filters applied everywhere.
- *What is a receptive field?* — The region of the input a neuron can see; it **widens with
  depth**, which is why early layers catch edges and deep layers catch objects.
