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

## Counting parameters

Flagged **twice** in the class notes as an exam notice, which is the strongest signal in them.
It is arithmetic, not mathematics, so it survives "very little math".

**The convention:** a parameter is a number **training changes**. So

> **one weight per incoming connection, plus one bias per neuron.**

The **inputs do not count** — they are the data you supply. **Hyperparameters do not count** —
they are fixed before training starts.

| Network | Weights | Biases | Total |
|---|---|---|---|
| One neuron, **2 inputs** | 2 | 1 | **3** |
| One neuron, 3 inputs | 3 | 1 | **4** |
| One layer of 3 neurons, 4 inputs each | 12 | 3 | **15** |
| 2 inputs → hidden layer of 2 → 1 output | 4 + 2 = 6 | 2 + 1 = 3 | **9** |
| 784 → 16 → 16 → 10 (the digit classifier) | — | — | **≈ 13,000** |

The worked example dictated in class is the first row: **two input variables, one output, three
parameters.** Forgetting the biases is the standard slip — count layer by layer and add them as
you go.

A tidy bookkeeping trick you may see: the bias is written as a weight **w₀** on an input **x₀**
permanently set to **1**, which folds it into the weight vector so the whole neuron is a single
dot product. It is still a parameter and still shifts the threshold.

`[slides 5-8]`

?check id=nn-046
?check id=nn-048

## Why the hard edge is smoothed

A perceptron in its raw form fires or does not: a **step**. That is a problem, because a step
has **no useful derivative** — zero everywhere it exists — so gradient descent has nothing to
follow.

> **The sigmoid turns a hard-edged model into a differentiable one.**

That single move is what makes the whole architecture learnable, and it is the reason the
optimization principle can get a grip at all. It is also why the network's decisions come out
as **probabilities rather than rigid verdicts**: the smoothed boundary reports how far a point
sits from the partition, not merely which side it is on.

The geometric picture behind this: a perceptron *is* a **partition** — a line in two
dimensions, a plane in three, and in general a **hyperplane of one dimension fewer than the
data**. Thirty input variables are separated by a twenty-nine-dimensional hyperplane, exactly
as a plane is cut by a line. Machine learning is then *adjusting the parameters of those
partitions until the data are separated accurately*, and where one partition cannot do it,
you add another — which is what a layer of neurons is.

**Where the boundary sits is a judgement, not just a fit.** False positives and false negatives
are both errors with different costs: a **civil court** would rather acquit the guilty than
convict the innocent, so it minimises **false positives**; a **military command** would rather
investigate a harmless contact than miss a real one, so it minimises **false negatives**. No
placement of the line eliminates both.

`[slides 9-13]`

?check id=nn-050
?check id=nn-053

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

## The chain from calculus to backpropagation

Four steps, and each one is a sentence:

1. **Differential calculus** gives the **rate of change** of a one-variable function. Positive
   derivative, the function is rising; negative, falling; zero, a stationary point.
2. For a function of **many** variables the **gradient** — the vector of partial derivatives —
   gives the **direction of greatest increase**. Its negative gives the greatest decrease.
3. **Gradient descent** walks in that decreasing direction: **w ← w − η∇L(w)**, repeatedly.
   **η is the learning rate**, the step size, and it is a **hyperparameter** — set before
   training, not learned. Too small and training crawls; too large and the step overshoots.
4. Applied to a **feedforward network**, which is a **composition of layers**, gradient descent
   becomes **backpropagation**: the **chain rule** assembles the derivative of the loss with
   respect to any weight from a product of local derivatives, so **one backward sweep yields
   the gradient for every weight at once**.

Without the chain rule you would have to perturb each weight separately — hopeless for the
thirteen thousand weights of a digit classifier, impossible for a billion.

**Forward pass computes the activations and hence the loss. Backward pass computes the
gradients.** Backprop finds the direction; gradient descent takes the step.

**Attribution worth knowing:** **Werbos published backpropagation in a 1974 doctoral
dissertation**; **Rumelhart, Hinton & Williams (1986)** is where it became known to the field.
1986 is the popularisation, not the invention.

**Variants of gradient descent** named in this course: **batch**, **stochastic (SGD)**,
**mini-batch**, **momentum**, **RMSProp**, **Adam**. Backpropagation is *not* one of them — it
is how the gradient they all consume gets computed.

The loss surface of a real network is **high-dimensional and non-convex**, nothing like the
tidy bowl in the picture. Gradient descent settles in **a** local minimum; the empirical
finding is that in high dimensions those minima are usually good enough.

`[slides 22-24]`

?check id=nn-040
?check id=nn-042

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

## Universal approximation, stated precisely

> For any continuous function on a closed bounded region, and any tolerance you like, there
> exists a **one-hidden-layer** feedforward network with a **finite** number of neurons that
> approximates it uniformly.

**One hidden layer.** That is the part people get wrong. **Cybenko (1989)** proved it for the
sigmoid, **Funahashi (1989)** for general non-polynomial activations, **Hornik (1991)** under
broad conditions.

So why go deep at all? Because depth is not about what *can* be represented, it is about what
can be *found*:

> **Expressivity is abundant; trainability is scarce.**

Deep networks reach the same expressivity **far more parameter-efficiently** and are **easier
to optimise**. The obstacle in practice is never the function class — it is optimisation and
**inductive bias**, the design choices that make good solutions reachable.

`[slides 24-26]`

?check id=nn-028
?check id=nn-029

## Three ways to make a network more capable

If expressivity is the question "what set of functions can this network represent?", the course
gives **three complementary answers**, and they can be combined:

| | **Enrich what the links compute** | **Create many computational paths** | **Make the network deep** |
|---|---|---|---|
| **Named as** | **Functional-link networks** (Pao, 1995) | **Residual networks** (He et al., 2015) | Traditional **sequential stacking** |
| **Core idea** | Add **explicit nonlinear transforms of the input** — polynomial, radial-basis, trigonometric terms — so the basis is richer | Add **identity skip connections**; a net with L residual blocks implicitly contains **2^L** input-output paths of varying length | Stack many nonlinear layers in series so the composition becomes more expressive |
| **Strength** | High expressivity at **shallow** depth; good with little data when you know the useful features | Trains **very deep** networks; ensemble-like robustness | Systematic hierarchical representations; scales with data and compute |
| **Limitation** | Needs good prior knowledge; input dimension can explode | Most of the paths are short, and only those carry gradient | **Vanishing / exploding gradients** without skip connections |

**None of the three changes the function class** — universal approximation already held. They
change the **optimisation landscape and the effective depth**, which is the practical
distinction between what a model *could* express and what training will actually deliver.

## The degradation problem, and the residual fix

By **early 2015** deep learning had hit a wall. Beyond roughly **twenty to thirty layers**,
making a network deeper stopped helping and then started hurting: a **30-layer model scored
16.59% error against a 14-layer model's 13.34%** — worse on the **training** set, not just on
validation.

That last detail is the whole diagnosis. **This is not overfitting.** Overfitting shows as a
widening gap between training and validation error. Here the deeper model fitted the data it
had already seen *worse*, which makes it an **optimisation failure**.

And it was confounding, because the solution provably existed: **take a trained 14-layer model,
add 16 pass-through (identity) layers, and performance must be at least as good**. That setting
of the parameters is sitting in the space and gradient descent never finds it.

**The fix, in one line:** instead of asking the layers to learn **H(x)** directly, let them
learn the **residual F(x) = H(x) − x**, and add the input back — so the block outputs
**H(x) = F(x) + x**.

- The shortcut lets **information and gradients flow** without passing through every weight.
- The block only has to learn a **correction**: *refine* rather than *re-create*.
- **Identity mapping is available by construction** — to do nothing, drive F(x) to zero.

The consequence was immediate: residual networks train at **18, 34, 50, 101, 152 layers and
beyond**, with error still falling where a plain network's rises. He, Zhang, Ren and Sun's
twelve-page **Deep Residual Learning for Image Recognition** (Microsoft Research, December
2015) became the **most cited paper of the twenty-first century**, and the residual block is
now the standard **backbone** underneath AlphaGo, AlphaFold and the large language models.

One reading worth a line: **ResNets behave like ensembles of relatively shallow networks.**
Unravel the shortcuts and a 110-block network trains mostly through paths of **10 to 34
blocks**; the long paths carry almost no gradient. **Effective depth is much smaller than
nominal depth**, which is why very deep residual networks train at all.

`[slides 24-26]`

?check id=nn-032
?check id=nn-034

## Overfitting

**Overfitting**: training error keeps falling while **validation error plateaus or rises** —
the model is learning the training set rather than the pattern. **Underfitting** is the other
shape: **both** errors stay high, and the model lacks the capacity or the training to capture
anything.

This is the cheapest kind of two-line answer in the paper: name the symptom in terms of the
two error curves, then name one remedy.

`[slides 24-26]`
