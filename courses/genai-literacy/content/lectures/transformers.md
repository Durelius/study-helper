---
id: transformers
title: "Transformers, Attention and LLMs"
deck: "GAI Lit 07 + lec 1"
slides: 16
summary: "Tokens, embeddings, attention and positional encoding — and the one-sentence answer to how a transformer decides which word it generates next."
---

## What the exam will ask from this lecture

The examiner wrote two questions on his own lab slide:

> ***"How does the transformer decide which word it generates next? What are the main parts of
> this transformer?"***

Those are the questions to prepare. The formally examined learning outcome for this paper is
*"explain key GAI architectures using intuitive mental models — how Transformer and Diffusion
Models work under the hood"*, so a mechanism in one sentence and a consequence in the second is
exactly the right shape.

## Words are not numbers

Neural networks take numbers, so a sentence has to be converted twice over:

1. **Tokenise** — chop the text into units the model has a vocabulary for. A token is usually a
   word or a fragment of one. Everything downstream happens over tokens, not letters.
2. **Embed** — map each token to a **vector of numbers**, its **embedding**. GPT's embedding
   dimension is about **12,288**: that is how many numbers one token occupies.

Because meaning lives in the geometry, you can do **word algebra**: **king − male + female ≈
queen**. Directions in the space correspond to consistent relationships, so gender, tense and
plurality each show up as roughly the same offset wherever you look. Nothing about analogies
was programmed in; it fell out of training on the task of **guessing a missing word** from its
neighbours — which is what **word2vec** (Mikolov, 2013) did.

A finite vocabulary meets infinite language, so a word-level model reserves a special
**unknown token** for anything outside its vocabulary. The input stays well-formed; the meaning
of that word is simply lost.

`[slides 1-4]`

?check id=tra-002

## Embedding vs attention — the distinction to memorise

| | **Embedding** | **Attention** |
|---|---|---|
| Supplies | **Global-context** meaning | **Local-context** interpretation |
| Scope | What the word means in general, the same everywhere | What it means **in this sentence**, given the other words present |

The embedding of *bank* is one vector wherever it occurs. Attention is what decides that in
*"the bank of the river"* the neighbouring words should pull that meaning towards the shore
rather than the vault.

> **Exam focus.** This row is on the distinctions table and it is the cleanest way to show you
> understand the architecture rather than the acronym. Global meaning versus local
> interpretation, in one sentence each.

`[slides 5-7]`

?check id=tra-003

## Attention, in plain terms

Attention is built from three roles, written **Q, K, V**:

- The **query** is what a token is looking for.
- The **key** is what each other token offers.
- The **value** is the content passed along when a query and a key match well.

Matching queries against keys produces the **weights**; the weighted mixture of **values** is
what each token takes away. So attention is a **learned, content-based lookup** rather than a
fixed wiring — which is why the same architecture works for any relationship the data happens
to contain.

**Self-attention** is the case where the queries, keys and values all come from the same
sequence: every token weighing every other token in its own sentence.

`[slides 7-8]`

## Positional encoding

Attention weighs relationships, not sequence — it is **order-blind by construction**. Without
help, *"the dog bit the man"* and *"the man bit the dog"* would be the same bag of tokens.

**Positional encoding** adds the missing information to each token's vector, so word order
survives into the representation. It is a separate component precisely because attention itself
cannot supply it.

`[slide 7]`

?check id=tra-008

## The main parts of a transformer

1. **Token embedding** — tokens become vectors.
2. **Positional encoding** — order is added.
3. **Self-attention layers** — every token weighs every other.
4. **Feedforward blocks** — an MLP applied at each position, stacked with the attention layers.
5. **Output layer and softmax** — a **probability distribution over the whole vocabulary**.

Note what is **not** there: **no recurrence and no convolution**. That is the claim in the title
*Attention Is All You Need* (Vaswani et al., **2017**) — not that attention does everything, but
that the sequential and local mechanisms previous sequence models were built on could be thrown
away.

`[slides 7-9]`

## How a transformer decides the next word

> It **embeds** and **positionally encodes** the tokens, uses **attention** to weigh how much
> every other token should inform each one, and outputs a **probability distribution over the
> vocabulary** via **softmax**, from which the next token is **sampled**.

The model never picks a word outright. It produces probabilities for all of them and one is
drawn — which is why the same prompt can give different answers, and why always taking the most
probable token makes output repetitive rather than better.

Generation is therefore **repeated prediction**: predict the next token, append it, predict
again. That is the loop behind the course's claim that **generative AI is built on predictive
AI** rather than opposed to it.

> **Exam focus.** This is the examiner's own question, near-verbatim. Write the four steps in
> one sentence and the sampling consequence in the second.

`[slides 7-9]`

?check id=tra-006

## GPT and LLMs

**GPT = Generative Pre-trained Transformer.**

- **Generative** — it produces new text rather than labelling existing text.
- **Pre-trained** — trained on **internet-scale text before any particular task**, so it arrives
  already knowing the language. Everything afterwards adapts a model that already exists.
- **Transformer** — the architecture above.

A **large language model** is that combination plus scale: a very large transformer, pre-trained
to predict the next token on a very large corpus. Everything it appears to know is a by-product
of getting good at that one prediction task.

GPT is **decoder-only**, whereas a translation transformer has both an encoder and a decoder.
The shape of the task sets the shape of the model: two sequences in two languages call for
encode-then-decode; one continuing sequence needs a single stack attending to what came before.

`[slides 10-14]`

## Why language ended up central

Language turned out to be the general medium. Once text could be **embedded** and **attended
over**, the same machinery could be pointed at images, audio and code — which is how the
transformer became the backbone of multimodal models, and how a **text** encoder ended up
steering an **image** diffusion model. Even a picture can be described in words.

`[slides 15-16]`
