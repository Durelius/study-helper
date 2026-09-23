---
id: ethics-impact
title: "Ethics and Impact"
deck: "playlist + lec 1"
slides: 12
summary: "AI slop and the economics that produce it, cognitive offloading and the erosion of critical thinking, and what AI literacy is a defence against."
---

## What the exam will ask from this lecture

This is the smallest topic on the paper and the one that moved most recently. In the last week
before the exam a **whole new row appeared on the assigned playlist — "GAI Ethics &
Post-Capitalism Economics"** — with a MustWatch video on the **erosion of critical thinking**,
and the 2025 paper's first bonus question was exactly this territory. The blueprint points at
the playlist explicitly.

So: **medium confidence as a regular question, higher as a bonus.** One or two clean two-line
answers is the right investment.

## AI slop, and why it is economics rather than taste

**AI slop** is low-value, mass-produced AI-generated content flooding the web. The problem is
not that any one piece is terrible; it is that producing it is nearly free, so the volume
swamps the material that was worth reading.

The mechanism is a **broken payment loop**:

1. The old web paid creators through **clicks and advertising** — you visited the page, the page
   earned, the creator kept writing.
2. **AI search now answers the question without a visit.**
3. So the revenue that funded high-quality, trustworthy content **evaporates**, and with it the
   incentive to produce any.
4. The models were trained on that content. As the supply degrades, so does the training data —
   each generation learning from the previous generation's output, with distortions amplified
   rather than corrected.

**Countermeasures are limited.** Detection is unreliable and the economics push the other way,
so responsibility for judging what is worth trusting has moved to **the reader**. That is the
practical reason AI literacy matters to non-specialists.

> **Exam focus.** *"What is meant by 'AI slop', and why is it an economic problem and not just a
> quality one?"* Lead with the definition, then the broken loop: search answers without a visit,
> so nobody funds the good content.

`[slides 1-4]`

?check id=eth-002

## Cognitive offloading and collective stupidity

**Cognitive offloading** is handing thinking work to an external tool. A calculator offloads
arithmetic you already understand, and that is fine. Offloading the **reasoning itself** is not,
because the judgement needed to **check** the answer is exactly the judgement that never
develops.

That makes it a feedback loop rather than a one-off trade:

> Critical thinking is exercised by struggling with a question. An instant answer removes the
> struggle, so the capacity to evaluate the answer weakens at the same rate the answers become
> easier to get.

The course's second worry is quieter and broader: **passivity**. If the machine translates
better than you can, there is no everyday reason to learn the language — and the capability
quietly disappears from the population. Nobody takes it away; we simply stop showing up.

`[slides 5-7]`

?check id=eth-011

## AI Literacy vs "Artificial Literacy"

| | **AI Literacy** | **"Artificial Literacy"** |
|---|---|---|
| What it is | You can **code** it, **safeguard** it, and **explore or exploit** it | Outsourcing your **essay**, your **pitch** and your **reading** to a chatbot |
| Result | A competence you gain | The appearance of a competence you never formed |

And the **bare minimum**, in the 2025 paper's own words — memorise the three clauses:

> As a **global digital citizen**, appreciating **how AI got where it is**, **how it is trained
> from data**, and the **capability as well as the pitfalls** of AI solutions — **so as not to
> apply AI blindly**.

Note that all three clauses point at the last one. Literacy here is the ability to judge when an
AI answer should not be trusted, not the ability to produce one.

`[slides 8-9]`

?check id=eth-005

## Why generative models state falsehoods confidently

This belongs here as much as in the architecture notes, because it is the thing literacy is a
defence against. A generative model is optimised for **plausible context relevance**, not for
**accuracy against a verifiable ground truth**, so **a fluent invention scores as well as a
fact**. A predictive model can be wrong and be caught by its ground truth; a generative model
has no ground truth to fail against.

Two consequences:

- **AI has no common sense.** The model has the statistics of the world, not a model of it.
  Nothing in its training says a hand has five fingers or that a date must precede its
  consequences, so the sanity check has to come from outside.
- **Not verifiable is not the same as not judgeable.** Six fingers on a hand, or a citation to a
  paper that does not exist, are failures you can point at without any ground truth at all.

`[slide 10]`

## Bias, and where it comes from

Train a model on historical hiring decisions and use it to screen applicants, and the problem is
not the architecture. **Parameters are calibrated from data**, and the data is a record of what
people **did**, not of what was **right**. Whatever bias the old decisions carried is exactly
what the model learns to reproduce — now with a machine's authority behind it.

This is why *"how it is trained from data"* is one of the three clauses of minimum literacy. It
is the clause that catches this case.

`[slide 11]`

## Deepfakes

The consequence of the **GAN**'s design, not an abuse of it. The generator improves precisely
because a discriminator keeps catching it, so the fakes get better until a human cannot tell
either. Deepfakes are what made adversarial networks famous, and they are the concrete reason
the ethics discussion attached itself to generative AI before ChatGPT existed.

`[slide 12]`

## Two framings worth quoting

- On what AI actually changed: **"AI did not invent abstraction; it industrialized it — turning
  learned generalities into algorithmic capability at scale."** The ladder is **animal**
  (adaptive abstraction, learned from experience), **human** (deliberate, taught as
  disciplines), **AI** (scalable, extracted algorithmically from massive data).
- On where the consciousness argument attaches: the fifth element of the Intelligent Being, the
  **consciousness marker and program register** — a momentary, sequential awareness of which
  computation step is being performed. The course offers it as the place the debate **starts**,
  not as a claim that machines are conscious.

`[slides 10-12]`
