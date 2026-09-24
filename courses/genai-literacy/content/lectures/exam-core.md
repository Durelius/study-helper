---
id: exam-core
title: "Exam Core: What Will Be Asked"
deck: "GAI Lit 07 blueprint + all decks"
slides: 22
summary: "The fifteen questions most likely to be on Friday's paper, ranked, each with a two-line model answer, the clause that earns the mark and the trap that loses it."
---

## How to use this page

Marko Niinimaki sets the paper: **10 questions, two ruled lines each (25–35 words), closed
book, "no code, very little math."** His blueprint slide names *Concepts and Principles
(lecture 1)*, *neural networks and how they learn*, *MLP, CNN, VAE, ..*, the Wolfram U CNN
course, the playlist and the infographics. His own class questions are causal and
compositional (*"Why did X happen?"*, *"Can we have X without Y?"*, *"What are the main
parts of X?"*), and classmates report he stresses **telling concepts apart**.

So every answer below is written the same way: **the term, then the distinguishing
clause.** If you know more than fits, it will not be read. Spend the words on the clause
that separates the right answer from its neighbour.

`[GAI Lit 07, final slide]`

## Skip these

- **Hand-computed gradients, partial derivatives, gradient-descent arithmetic.** Keep the verbal definition only.
- **Any code**: Python, Keras, Colab, Wolfram function names (`NetEncoder`, `FlattenLayer`). The *concepts* behind the Wolfram course are examinable; the names are not.
- **Prompt engineering, RAG, agents.** Pushed past the midterm to the term project. One line of insurance: *RAG retrieves relevant documents first, then generates grounded on them.*
- **ResNet internals, LSTM, ViT, CLIP, flow matching** beyond recognising the name.
- **BAScii "Integrated Innovation" meta-questions, the drone sheet, Buddhism and the five aggregates.** Nacaskul's vocabulary, absent from Marko's blueprint.

## 1. Predictive AI vs generative AI, with an example of each

*The best-corroborated topic in the course: 2024 Q2 verbatim, two infographics, the lecture-1 blueprint line.*

> Predictive AI learns x → y and is judged on **accuracy against verifiable ground truths**
> (an MLP classifier). Generative AI learns the data distribution, z → x, and is judged on
> **plausible context relevance** (GPT).

**Mark earner:** the two evaluation standards, in the course's words. **Trap:** "accuracy vs
creativity". Backup framing: predictive asks *"What WILL be?"*, generative *"What COULD be?"*

`[2024 midterm p.1; infographics 2, 5]`

?check id=core-001
?check id=core-002

## 2. Name the six architectures and tell them apart

*Blueprint names three of them; the infographic, two sets of class notes and a page headed "6 architectures for midterms" name all six.*

> **MLP, CNN** predict; **VAE, U-Net** form the representation bridge; **Transformer,
> Diffusion** generate. Stage questions: *what is most likely? · what is the underlying
> structure? · what could be?*

| Architecture | One distinguishing clause |
|---|---|
| MLP | Stacked feedforward layers of perceptrons |
| CNN | Same learned filters slid over every position |
| VAE | Encodes to a *distribution*, so you can sample |
| U-Net | Encoder–decoder CNN with **skip connections** |
| Transformer | **Attention** weighs every token against every other |
| Diffusion | Learns to reverse a fixed noising process |

**Trap:** putting the GAN in the list, or forgetting U-Net.

`[infographics 3]`

?check id=core-004
?check id=core-006

## 3. How do neural networks learn? Gradient descent and backpropagation

*The blueprint's own phrase, plus an "exam sentence" dictated in class.*

> **Backpropagation is a specialisation of the gradient descent algorithm, applied to
> minimise prediction error for a multilayer perceptron by tweaking weights and biases.**
> The gradient points uphill (steepest ascent); training steps against it, repeatedly.

**Mark earner:** memorise the bold sentence verbatim. The loss (e.g. RMSE, cross-entropy)
is the single number being minimised. **Trap:** "follows the gradient"; or calling backprop
a variant like Adam. Adam, SGD and momentum are optimisers; backprop is *how the gradient
is computed*.

`[lec 1 slide 61; lec 5 slide 2; infographics 13]`

?check id=core-007
?check id=core-009

## 4. What is a latent space, and what makes a VAE generative?

*Marko's own lab slide asked "What is a latent space?"; the blueprint names VAE.*

> The latent space is the **compressed representation at an autoencoder's bottleneck**,
> where nearby points mean similar inputs. A plain AE maps each input to **one point**; a VAE
> encodes to a **distribution**, so you can **sample** a new point and decode a new example.

**Mark earner:** point vs distribution, then "sample and decode". **Trap:** saying the AE
"can't produce outputs". It can, but the gaps between its codes decode to nonsense.

`[GAI Lit 07 slide 6; infographics 4]`

?check id=core-010
?check id=core-011

## 5. How does a transformer decide which word it generates next?

*Marko's lab prompt, word for word, together with "What are the main parts of this transformer?"*

> It **embeds** and **positionally encodes** the tokens, uses **attention (query, key,
> value)** to weigh how much each token should inform the others, and outputs a
> **probability distribution over the vocabulary**; the next token is **sampled** from it.

**Mark earner:** ending on the distribution and the sample. **Paired distinction:** the
*embedding* gives a word its global meaning (the vocabulary); *attention* gives its
local-context meaning in *this* sentence (the grammar). **Trap:** "it picks the most likely
word". It samples, which is why the same prompt varies. The T is Vaswani et al., **2017**,
*Attention Is All You Need*.

`[GAI Lit 07 slide 7; lec 1 slide 70]`

?check id=core-013
?check id=core-014

## 6. What did Hubel and Wiesel find in the cat, and why does it matter?

*Flagged "exam for sure" in class and repeated in every set of classmates' notes.*

> Neurons in the cat's striate cortex fired for **oriented lines** (the slide's edge),
> not for the mouse on it. Line detectors combine into curves, shapes and objects: the
> **hierarchy of features** that the Neocognitron (1980) and the CNN copy.

**Mark earner:** "lines, not objects" plus "hierarchy". **Trap:** spending words on the
date. Say "around 1959".

`[lec 1 slide 50; Wolfram CNN lesson 12]`

?check id=core-016
?check id=core-017

## 7. Why can a CNN find an object anywhere in the image when an MLP cannot?

*The CNN is named in the blueprint, and the examiner assigned the Wolfram course.*

> A CNN slides the **same learned filters over every region** (weight sharing), so a feature
> is detected wherever it sits: **translational covariance**. An MLP gives each pixel
> position its own weight, so a shifted object looks like a different input.

**Mark earner:** "same filter everywhere". **Trap:** "a CNN has more layers". Depth is not
the difference.

`[lec 5 slide 1; Wolfram CNN lesson 5]`

?check id=core-018

## 8. Why did mobile phones start recognising images in the 2010s?

*Marko's own question from deck 02.*

> Three things converged: **large labelled datasets** (ImageNet, Fei-Fei Li), **GPU
> compute** for the matrix multiplications, and **deep CNNs**. AlexNet (2012) cut ImageNet
> error from about 25% to single digits.

**Mark earner:** naming all three. **Trap:** one cause only, or swapping ImageNet (the
data) with AlexNet (the network).

`[GAI Lit 02 slide 15; GAI Lit 03 slide 5]`

?check id=core-020
?check id=core-021

## 9. What does a neuron compute, and why stack them?

*The blueprint names MLP; Marko's labs drill XOR and bias; "count parameters" was flagged twice in class.*

> A perceptron takes a **weighted sum of its inputs plus a bias** and passes it through an
> **activation**. It draws one straight boundary, so it fails on **XOR**, which is not
> linearly separable. Stacking layers (the MLP) bends the boundary.

**Mark earner:** "weights scale inputs, **bias shifts the threshold**". **Count:** one
weight per connection plus one bias per neuron, so **2 inputs, 1 output = 3 parameters**.
Inputs and hyperparameters never count.

`[lec 1 slide 60; GAI Lit 03 slides 6, 9]`

?check id=core-023
?check id=core-024

## 10. Give three senses of the "I" (and of the "A")

*The lecture-1 blueprint line; "three forms (at least) of intelligence" was flagged as an exam notice.*

> **I:** human-like (Turing's imitation game) · advantageous information (military,
> business intelligence) · solution path to a problem (computational intelligence).
> **A:** artifact/artifice · artificial semblance ("seems about right") · artful creation
> ("truth manufactured").

**Trap:** crossing the lists. "Advantageous information" is an I, not an A. Giving only
"human-like" earns a third.

`[lec 1 slides 19-21]`

?check id=core-025
?check id=core-026

## 11. In what way is ML larger than AI, and AI larger than ML?

*Written down in class as "Exam question", verbatim.*

> ML is larger: its data-driven parameter calibration reaches far outside AI (statistics,
> optimisation, forecasting). AI is larger: it includes rule-based GOFAI, search and planning
> that involve **no learning at all**.

**Trap:** the textbook "ML ⊂ AI" diagram, which answers half. Give one example per direction.

`[lec 1 slides 59, 76]`

?check id=core-027
?check id=core-028

## 12. What is symbolic AI? Expert systems?

*Marko's own bonus question, and MYCIN appears on his own slide.*

> **Hand-written if-then rules over symbols, with no learning from data**: expert systems
> such as **MYCIN (1975)**, diagnosing bacterial infections at about 65%, in LISP or Prolog.
> It stalled into the AI winter and was renamed **GOFAI**, a retronym.

`[GAI Lit 02 slide 15; lec 1 slide 59]`

?check id=core-029

## 13. The Wolfram CNN glossary: one-line definitions

*The blueprint names the course "(assignment, Wolfram U)"; section 1 of it is exactly this list.*

| Term | Two-line definition |
|---|---|
| **ReLU** | Returns the input if positive, else zero: cheap **nonlinearity** |
| **Pooling** | **Reduces the size** of the representation |
| **Softmax** | Raw activations → **probability distribution over classes** (sums to 1) |
| **Cross-entropy** | Loss: distance between predicted probabilities and the one-hot label |
| **Overfitting** | Training error falls while **validation error plateaus or rises** |
| **Dropout** | Randomly switches off units in training so no single path is relied on |
| **Receptive field** | The input region a neuron sees; **widens with depth** (edges → objects) |

**Trap:** softmax vs cross-entropy (transform vs loss); overfitting vs underfitting (both
errors high).

`[Wolfram CNN lessons 7-12]`

?check id=core-031
?check id=core-032

## 14. Can we have generative AI without neural networks?

*Marko's own bonus question.*

> Yes in principle: a **Markov chain** generates text from next-token statistics, and any
> sampler from a fitted distribution is generative. But with no learned representation it
> has **no context**: locally plausible, globally incoherent.

**Mark earner:** the "but". A bare "yes" or "no" gets nothing.

`[GAI Lit 02 slide 15]`

?check id=core-033

## 15. Diffusion and GAN: two lines each

*Diffusion is on the syllabus's examined outcome (CLO 2); the GAN's irony was flagged as an exam note.*

> **Diffusion:** a *fixed* forward process adds noise step by step; the network **learns
> only the reverse** (denoising). Generation starts from **pure noise** and denoises,
> optionally guided by text.
>
> **GAN:** a generator makes fakes, a discriminator judges real vs fake, and both improve.
> **Irony:** it was built to sharpen a *predictive* discriminator; the generator became the star.

**Trap:** "diffusion learns both directions". That is the VAE.

`[infographics 4; lec 5 slide 7]`

?check id=core-034
?check id=core-035

## Bonus insurance: AI slop

A whole ethics row was added to the playlist days before the exam, and 2025's bonus was this video.

> The old web paid creators through clicks and ads; AI search answers **without a visit**,
> so the incentive to produce trustworthy content evaporates, degrading the supply the
> models learned from.

`[playlist item 21]`

?check id=core-036

## Last-hour cram list

1. **Predictive** = accuracy vs verifiable ground truths; **generative** = plausible context relevance. WILL be / COULD be.
2. **MLP, CNN | VAE, U-Net | Transformer, Diffusion.** No GAN.
3. **Backprop is a specialisation of GDA, applied to minimise prediction error for an MLP by tweaking weights and biases.** Step *against* the gradient.
4. **AE → one point; VAE → a distribution you sample.**
5. **Transformer:** embed → position → attention (Q, K, V) → distribution over vocabulary → sample. Embedding = global, attention = local. Vaswani **2017**.
6. **Cat: lines, not the mouse → hierarchy → Neocognitron 1980 → CNN.**
7. **CNN: same filter everywhere = translational covariance.**
8. **2010s phones: ImageNet + GPUs + deep CNNs; AlexNet 2012, 25% → single digits.**
9. **Neuron = weighted sum + bias → activation.** Bias = threshold. XOR not linearly separable. **2 in, 1 out = 3 parameters.**
10. **I:** human-like · advantageous information · solution path. **A:** artifact · semblance · artful creation.
11. **ML ⊄ AI and AI ⊄ ML**: statistics/forecasting one way, GOFAI/search the other.
12. **Symbolic AI = hand-written if-then, no learning; MYCIN 1975, ~65%; GOFAI.**
13. **Softmax** = probabilities; **cross-entropy** = loss; **dropout** vs overfitting; **receptive field** widens with depth.
14. **GAI without NN:** yes (Markov chain), *but* no context.
15. **Diffusion:** fixed noising, learned denoising, start from noise. **GAN:** built to sharpen a discriminator.

Write the term first, then the clause. Write legibly.
