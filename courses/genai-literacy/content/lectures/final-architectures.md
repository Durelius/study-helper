---
id: final-architectures
title: "Final Architectures: MLP, CNN, Diffusion"
deck: "Nov 13 quiz"
slides: 0
summary: "The architecture questions the Nov 13 quiz asked, with the one clause that earns each mark."
---

## Pick the architecture from the data

The first question in the quiz is a choice of tool. **Predicting a number from a table of features** (floor area, district, bedrooms) is a job for a **Multi-Layer Perceptron (MLP)**: a stack of fully connected layers, each neuron reading every neuron before it. A table has no grid and no word order, so neither a CNN nor a Transformer adds anything.

> **Exam focus.** Match the data shape to the architecture. Table: MLP. Image grid: CNN. Token sequence: Transformer. Generated samples from a latent space: VAE. Noise removed step by step: diffusion.

`[Nov 13 quiz, Q1]`

## Why a CNN convolves

A **convolution layer** slides a small **kernel** (a grid of learned weights, often 3×3) across the image. Two reasons are both marked correct in the quiz:

1. **Fixed-size kernels reused everywhere.** The same weights are applied at every position, so a CNN has far fewer parameters than an MLP on the same picture. This is **weight sharing**. It is why the answer is not "the image gets bigger", which is wrong.
2. **Hierarchical feature learning.** Early layers find edges, later layers combine them into textures, parts and objects. Each layer builds on the one before.

The trap is the phrase "guarantee that dimensions increase exponentially". Convolution changes the size of the feature maps in a controlled way and never guarantees growth.

`[Nov 13 quiz, Q2]`

## Diffusion: noise in, image out

A **diffusion model** has two processes. The **forward process** is fixed: it adds a little Gaussian noise to a clean image at each of a set number of time steps until nothing but noise remains. The **learned part** is the reverse: a network that predicts and removes noise at each step, turning random noise into an image.

The second correct answer in the quiz is conditioning. **Embeddings** of a text prompt tell the denoising network what to make, so the same noise can become many different pictures. Diffusion does not compress images; that is the VAE's job.

> **Exam focus.** Two ruled lines are enough for: forward adds noise over fixed steps; the network learns to reverse it; embeddings steer the result.

`[Nov 13 quiz, Q3]`

## VAE and generative versus predictive

A **variational autoencoder** encodes an input into a **latent space** that is a probability distribution, then decodes samples from it. Because you can sample, a VAE generates. The V in VAE stands for variational.

The distinction the quiz tests is **predictive versus generative**. Predictive models map an input to a label or value: the MLP and CNN classifiers. Generative models learn the data itself and produce new samples: VAE, diffusion, and a Transformer that writes text. Keep the question shape in mind: a list of architectures with a mix of both.

`[Nov 13 quiz, Q4]`

## What the MLP is not

An MLP is the right answer for the house-price question and the wrong answer for a photograph. A fully connected layer links every input value to every neuron, so the parameter count grows with image size and nothing favours nearby pixels. That is why the course moves to the CNN for images.

`[Nov 13 quiz, Q1 and Q2]`

?check id=fa-001
?check id=fa-002
?check id=fa-005
?check id=fa-010
