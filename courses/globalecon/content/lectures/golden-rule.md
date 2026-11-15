---
id: golden-rule
title: "Golden Rule, Population & Technology"
deck: "Quiz Prep"
slides: 18
summary: "Which steady state is best for consumers (MPK = δ), what a policymaker does above or below it, and how population growth n and labour-augmenting technology g join depreciation in break-even investment (δ + n + g)k."
---

## Saving or consumption: what matters?

People do not care about capital, or even output, for its own sake. They care about what
they can **consume**. (We do care about saving too, because it smooths consumption into
the future, including donations.) So a policymaker who wants to maximise well-being
picks the steady state with the **highest consumption per worker**: the **Golden Rule
level of capital, k\*gold**.

In the steady state, investment equals depreciation (i = δk\*), so:

```
c* = f(k*) − δk*
```

`[slides 72–73]`

## The condition: MPK = δ

Steady-state consumption is the **gap** between the production curve f(k\*) and the
depreciation line δk\*.

- **Below k\*gold**: f(k) is steeper than δk, so raising k raises output by more than
  depreciation. The gap (consumption) **grows**.
- **Above k\*gold**: raising k adds less output than depreciation, so the gap **shrinks**.

The gap is widest where the two slopes are equal. The slope of f is MPK and the slope of
the δk line is δ:

```
Golden Rule:   MPK = δ      (equivalently MPK − δ = 0)
```

**Intuition.** Move from k\* to k\* + 1: output rises by f(k\* + 1) − f(k\*) = **MPK**, and
extra depreciation is **δ**. Keep adding capital while MPK − δ > 0. At the Golden Rule the
**marginal product of capital net of depreciation is zero**.

`[slides 74–75]`

?check id=gr-006

## Math primer: MPK for y = k^α

MPK is the derivative of f(k) (power rule):

```
f(k) = k^α        ⇒   MPK = α · k^(α − 1)
f(k) = √k = k^½   ⇒   MPK = ½ · k^(−½) = 1 / (2√k)
```

**Worked example.** y = √k, δ = 0.05.

```
1/(2√k) = 0.05   ⇒   2√k = 20   ⇒   √k = 10   ⇒   k*gold = 100
y = √100 = 10,   δk = 0.05 × 100 = 5,   c*gold = 10 − 5 = 5
saving rate needed: s·f(k) = δk  ⇒  s = 5/10 = 0.5
```

**Shortcut** (going beyond the slides, but a handy check): for y = k^α the Golden Rule saving rate
is always **s_gold = α**. With y = √k that is ½; with y = k^0.3 it is 0.3.

## Policy: too much or too little capital

**Too much capital (k\* above k\*gold): lower s.**

- At the moment s falls, **consumption jumps up** and investment falls by the same amount.
- Over time, as capital falls, output, consumption and investment fall together.
- But the new steady state has **higher consumption than the initial one**. Every
  generation is better off, starting with the current one.

**Too little capital (k\* below k\*gold): raise s.**

- At the moment s rises, **consumption dips** and investment jumps by the same amount.
- Over time, as capital grows, output, consumption and investment rise together.
- The new steady state has **higher consumption** than the initial one, but the
  **current generation pays** with the dip.

> **Exam focus.** Above the Golden Rule, every generation gains from saving less,
> including the current one. Below it, saving more means the current generation pays
> first with a consumption dip and later generations gain. That trade-off is the
> "sacrifice" the Q5 quote ignores.

`[slides 76–78]`

?check id=gr-019

## Population growth: break-even investment (δ + n)k

At the steady state, capital accumulation alone **cannot explain sustained growth**:
capital per worker, output per worker and total output stay the same. So add population
growth: L grows at rate **n** (US ≈ 1%: 150 million workers → 151.5 → 153.015 million).

Now three forces change k: investment raises it, depreciation lowers it, and **new
workers dilute it**. To keep k constant you must replace depreciated capital (**δk**) and
equip the new workers (**nk**):

```
Δk = s·f(k) − (δ + n)k
break-even investment = (δ + n)k
```

The steady state is where s·f(k) crosses (δ + n)k.

**Prediction:** countries with **higher population growth have lower capital per worker
and therefore lower income** per person. The (δ + n)k line is steeper, so it cuts sf(k) at
a lower k\*.

`[slides 79–81]`

?check id=gr-025

## Technological progress: (δ + n + g)k

Technology enters as **labour-augmenting** efficiency:

```
Y = F(K, L × E)
```

**E** is the **efficiency of labour**: society's knowledge about production methods
(computers in the 21st century, the assembly line in the 20th). **L × E** is the number
of **effective workers**. Increasing E is like increasing L: technology does not raise the
actual number of workers, but it raises the effective number.

| Grows at | |
|---|---|
| L | n |
| E | g |
| L × E | **n + g** |

Redefine **k = K/(L × E)** (capital per effective worker) and **y = Y/(L × E)**:

```
Δk = s·f(k) − (δ + n + g)k
```

Break-even investment has three parts: **δk** replaces worn-out capital, **nk** equips
new workers, **gk** equips the new *effective* workers created by technology. At k\*,
capital and output **per effective worker** are constant. Labour-augmenting technology does
not substantially change the shape of the steady state.

`[slides 84–87]`

?check id=gr-033

## Why the rates add: the computers example

Start with 10 computers and 10 workers, so k = 1. Let δ = 0.1, n = 0.2, g = 1.0.

1. Depreciation: 10 → **9 computers**.
2. Population: 10 → **12 workers**.
3. Technology: only the **original 10** workers benefit (the 2 new ones were not in the
   system yet), so they work like **20 effective workers**. The new 2 count as 2.
4. Effective workers = 20 + 2 = **22**. To keep k = 1 you need 22 computers: **add 22 − 9 = 13**.
5. Check: (δ + n + g) × 10 = (0.1 + 0.2 + 1.0) × 10 = 1.3 × 10 = **13**.

The rates **add** because the new workers do not get the technology boost. If they did,
you would multiply (1.2 × 2.0) instead.

**Second example:** δ = 0.2, n = 0.2, g = 0.5. Depreciation leaves 8 computers; 10
workers become 15 effective plus 2 new = **17 effective workers**; add 17 − 8 = **9**
computers = (0.2 + 0.2 + 0.5) × 10. (The slide's closing line says "22 computers for 22
effective workers" here, which was copied from the first example. The right figure is
17 for 17.)

`[slides 88–89]`

?check id=gr-040
