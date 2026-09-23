---
id: neural-networks
title: "How Neural Networks Learn"
deck: "lec 5 + GAI Lit 03"
slides: 26
summary: "Perceptron to MLP, weights and bias, loss, gradient descent and backpropagation — the examiner's own blueprint line, in words rather than arithmetic."
---

## What the exam will ask from this lecture

The blueprint line is **"Neural networks and how they learn"** — the examiner's own phrase.
There is one sentence in this lecture that was dictated in class as an exam sentence, and it
should be written out verbatim if the question comes up.

**No arithmetic.** The 2025 paper spent five questions computing gradients by hand; the
current examiner says **"no code, very little math"**, and two ruled lines cannot hold
working. The only counting worth doing is *how many parameters does a 2-input, 1-output
perceptron have?* — two weights plus a bias, **three**.

## The perceptron

A **perceptron** is a **mathematical-computational model of a generic neuron**. It **sums
weighted incoming signals** — excitatory or inhibitory — **and fires if that sum exceeds an
internal threshold**.

In modern terms: weighted sum of inputs, **plus a bias**, through an **activation function**,
out to the next layer. The **bias is the decision threshold** — it shifts where the neuron
fires, while the weights scale how much each input counts.

A single perceptron draws one straight boundary, so it can only separate **linearly
separable** data. **XOR** is the standard counterexample: no single line separates the true
cases from the false ones. That is the whole argument for **layers**.

> **Exam focus.** 2025 asked *"What does MLP stand for? Explain the P."* The content survives
> even if the wording does not: **Multilayer Perceptron**, and the Perceptron is the model of
> a neuron that sums weighted inputs and fires past a threshold.

`[slides 5-9]`

?check id=nn-001

## MLP and the network vocabulary

An **MLP** is **feedforward** layers of perceptrons stacked: each layer's output is the next
layer's input, **no recurrence and no skipping layers**. In a **feedforward network (FFN)**
the signal travels one way only; in a **recurrent network (RNN)** it loops back.

The vocabulary the class was told to know cold:

| Term | What it is |
|---|---|
| **Weight / coefficient / parameter** | The same thing — the numbers **learned** from data |
| **Input / variable** | The x values you supply |
| **Hyperparameter** | Set **before** training — layer sizes, learning rate |
| **Variable selection** | Choosing *which* x's to use |
| **Feature engineering** | **Building a new input from existing ones** by a logical function, e.g. x₁/x₃ |

The canonical worked example is the handwritten-digit network: a **28×28 pixel grid → 784
input neurons → two hidden layers of 16 → 10 output neurons**, about **13,000 weights and
biases** — "knobs and dials". The progression across layers is **pixels → edges → patterns →
the identified digit**, which is the same hierarchical abstraction Hubel and Wiesel found in
the cat. The output uses **one-hot encoding**: exactly one of the ten slots on, the rest off.

`[slides 10-14]`

?check id=nn-008

## The loss function

The network needs a single number saying how wrong it currently is. That number is the
**loss**, **L**.

- **RMSE** — Root Mean Squared Error — is **one specific loss**, natural when the target is a
  quantity. The loss function is the general idea; RMSE is an instance of it.
- **Cross-entropy** is the natural loss when the target is a **class**, because it compares
  two probability distributions rather than two numbers.
- The **cost function** in the canonical treatment is the **sum of squared differences between
  the actual output activations and the desired ones, averaged over tens of thousands of
  examples**.

`[slides 15-17]`

## Gradient descent

The **gradient** is the vector pointing in the direction of **steepest ascent** of the loss.
So the algorithm **steps against it** — downhill — and repeats. It is **iterative and
stepwise**: from wherever you are, find the direction that gets you down fastest, take a small
step, look again.

The intended mental model is a **ball rolling down a surface**. The Mathematica demo makes it
literal: an **RMSE gauge** to match and an **RMSE surface to descend**. That surface *is* the
loss landscape, and dragging the sliders down it by hand *is* gradient descent.

Two honest caveats the course makes:
- It settles in **a** local minimum, not the global one, and not in anything resembling human
  abstraction. Trained networks reach **96–98%** accuracy and still do not "understand".
- **ADAM = Adaptive Moment Estimation**, a refinement of plain gradient descent that adapts
  the step size.

`[slides 18-21]`

?check id=nn-014

## Backpropagation

Memorise this literally — it was dictated in class as an exam sentence:

> **BACKPROPAGATION IS A SPECIALIZATION OF THE GRADIENT DESCENT ALGORITHM, APPLIED TO MINIMIZE
> PREDICTION ERROR FOR A MULTILAYER PERCEPTRON — by tweaking weights and biases.**

The distinction the marks sit in: **gradient descent is the general "step downhill"
optimiser**; **backpropagation is gradient descent specialised to a layered network**, sending
the error backwards through the layers so each weight learns how much it contributed.

> **Exam focus.** If asked to "state the relationship between gradient descent and
> backpropagation", the sentence above *is* the answer, and it fits on two lines. Add the
> gradient's direction if there is room.

`[slides 22-24]`

## Activation functions

| Function | What it does | Why it matters |
|---|---|---|
| **Sigmoid** | Squashes any real input into **(0, 1)** | Suffers from **vanishing gradients** in deep networks |
| **ReLU** (the **ramp**) | Returns the input if positive, else **0** | Cheap nonlinearity; the modern default, and it displaced sigmoid |
| **Softmax** | Turns a vector of raw activations into a **probability distribution over classes** (sums to 1) | The output stage of a classifier |

Note the distinction between the last two: **softmax works across a whole vector**; **sigmoid
squashes one value at a time**.

Why any activation at all? Without a nonlinearity, stacking layers gains nothing — a chain of
linear maps is still a linear map. The activation is what lets a deep network learn a complex
decision boundary.

`[slides 12-13]`

?check id=nn-019

## Two things to be able to say about what a network is

- **(AI =) ANN = f(parameters) = architecture.** The network computes x → f() → y; what varies
  between architectures is the shape of f, and what training does is fit the parameters.
- **Universal approximation** says a sufficiently wide network *can* represent essentially any
  function. It does **not** say that gradient descent will find it, or that the result will
  generalise, or that the representation will resemble a human concept.

## Overfitting

**Overfitting**: training error keeps falling while **validation error plateaus or rises** —
the model is learning the training set rather than the pattern. **Underfitting** is the other
shape: **both** errors stay high, and the model lacks the capacity or the training to capture
anything.

This is the cheapest kind of two-line answer in the paper: name the symptom in terms of the
two error curves, then name one remedy.

`[slides 24-26]`
