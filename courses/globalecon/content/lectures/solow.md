---
id: solow
title: "Solow Model: Capital, Saving & Steady State"
deck: "Quiz Prep"
slides: 25
summary: "The Solow growth model from the production function to the steady state: returns to scale, the per-worker form, i = sf(k), Δk = sf(k) − δk, the level effect of saving, and the two war questions."
---

## What the model is for

**Robert Solow** won the Nobel Prize for his theory of economic growth. In the 1950s he
built a mathematical model of how factors contribute to sustained national growth.
From the 1960s, the model's implication persuaded governments to fund **technological
research and development**.

The model shows how growth in the **capital stock (K)**, growth in the **labour force
(L)** and advances in **technology (A or E)** interact, and how they affect total output.
It is built in steps: first K changes with L and technology fixed, then population
growth, then technology.

- **Endogenous** variables are the ones the model explains (real GDP, investment, employment).
- **Exogenous** variables are taken as given and not explained (the saving rate s).

`[slides 50–51]`

## Assumptions

- **Y = C + I + G + NX** (the national income identity by expenditure). In the model
  G = 0 and NX = 0, so **Y = C + I**.
- People save a constant fraction **s** of income and consume the rest: **C = (1 − s)Y**.
- Capital depreciates at a constant rate **δ**.
- Every firm uses the same technology: **Y = F(K, L)**.
- Population grows at a constant rate (taken as 0 at first).

`[slide 52]`

## Math primer: exponents

You need four rules:

| Rule | Example |
|---|---|
| xᵃ · xᵇ = xᵃ⁺ᵇ | z^½ · z^½ = z¹ |
| xᵃ / xᵇ = xᵃ⁻ᵇ | L^0.7 / L¹ = L^−0.3 |
| (xy)ᵃ = xᵃ yᵃ | (zK)^½ = z^½ K^½ |
| xᵃ / yᵃ = (x/y)ᵃ | K^0.3 / L^0.3 = (K/L)^0.3 |

And x^½ = √x, so **k^0.5 = √k**.

## Returns to scale

**Constant returns to scale (CRS)**: multiply both inputs by any z > 0 and output is
multiplied by exactly z:

```
zY = F(zK, zL)
```

**Test for K^a · L^b**: (zK)ᵃ(zL)ᵇ = z^(a+b) · KᵃLᵇ = z^(a+b) · Y.

| a + b | Returns to scale |
|---|---|
| = 1 | **constant** |
| > 1 | increasing |
| < 1 | decreasing |

**Pop quiz Q1a** (Y = K^½ L^½): F(zK, zL) = (zK)^½(zL)^½ = z^½ z^½ (K^½ L^½) = z·F(K, L) =
zY. **Yes, it has CRS.**

`[slides 56, 58–59]`

?check id=sol-012

## The per-worker production function

CRS lets us study everything **per worker**. Set z = 1/L:

```
Y/L = F(K/L, 1)    ⇒    y = f(k)
```

where **y = Y/L** is output per worker and **k = K/L** is capital per worker.

**Recipe:** divide both sides by L and collect the exponents.

- **Q1b**: Y = K^½ L^½ ⇒ Y/L = K^½ L^½ / L = K^½ / L^½ = (K/L)^½ ⇒ **y = k^½ = √k**
- **Q4a**: Y = K^0.3 L^0.7 ⇒ Y/L = K^0.3 · L^0.7 / L¹ = K^0.3 / L^0.3 = (K/L)^0.3 ⇒ **y = k^0.3**

The exponent on capital survives, because the exponents add to 1.

`[slides 57–59, 70–71]`

?check id=sol-018

## The shape of f(k): diminishing returns

The production function y = f(k) has:

- a **positive first derivative**: more capital per worker always raises output (upward
  slope). The slope is the **marginal product of capital, MPK**: the extra output a
  worker produces when given one more unit of capital;
- a **negative second derivative**: the slope gets **flatter** as capital accumulates
  (**diminishing MPK**).

This is the regularity Solow set out to explain (the packet shows it with German data
from 1820).

`[slides 53–54, 60]`

## Saving, investment, depreciation

Output per worker splits into consumption and investment, **y = c + i**. With c = (1 − s)y:

```
y = (1 − s)y + i   ⇒   i = sy = s·f(k)
c = f(k) − s·f(k)
```

The saving rate **s** splits output between consumption and investment, and it can
differ across countries. Depreciation per worker is **δk**, a straight line through the
origin.

**Change in capital = investment − depreciation:**

```
Δk = s·f(k) − δk
```

- Below k*: investment > depreciation, so **k grows**.
- Above k*: depreciation > investment, so **k shrinks**.
- At the **steady state k\***: **Δk = 0**, so s·f(k*) = δk*, and capital no longer changes.

**Solving for k\*** with y = √k: s√k = δk ⇒ √k = s/δ ⇒ **k\* = (s/δ)²**. For example,
s = 0.3 and δ = 0.1 give k\* = 9, y\* = 3, c\* = 0.7 × 3 = 2.1.

`[slides 55, 61–63]`

?check id=sol-025

## Level effect, not growth effect

Raise the saving rate from s₁ to s₂: the investment curve shifts up, k\* rises to k₂\*,
and y\* rises. The economy **grows faster while it moves to the new steady state**, then
growth stops again.

- **Level effect**: a policy changes the *level* of income per person in the steady
  state. A higher saving rate has a level effect.
- **Growth effect**: a policy changes the *steady-state growth rate* of income per
  person. **A higher saving rate does not have one.**

So higher saving means faster growth **only temporarily**. Maintaining a high s
keeps a high k and a high level of output, but not a high rate of growth.

`[slides 64, 72]`

?check id=sol-031

## Wars: destroying capital vs killing workers

**Japan and Germany after 1945.** War destroyed much of the capital stock: k fell from k\*
to k₁. Output fell immediately, but because the saving rate stayed high (Japan 24%,
Germany 26%, US 22%), investment exceeded depreciation at the low k, so the economy grew
fast: **8.2% a year in Japan and 5.7% in Germany vs 2.2% in the US (1948–72)**. The high
growth lasts until the economy approaches its former steady state. This is the
**"level-effect" catch-up**.

**Pop quiz Q3: a war kills workers but leaves capital.**

- **(a) Immediately:** L falls, so **total output Y falls** (fewer workers), but k = K/L
  **rises**, so **output per worker rises** (each worker has more capital).
- **(b) Afterwards:** k is now *above* k\*. Depreciation exceeds investment, so k falls
  back to k\*, and **output per worker falls** from its higher level back to y\*. The
  growth rate of output per worker is **negative** (smaller than normal, which is zero
  in steady state without technological change) until the steady state is reached.

> **Exam focus.** Capital destroyed → k below k\* → **above-normal** growth. Workers
> killed → k above k\* → **below-normal (negative)** growth. They are mirror images.

`[slides 65–69]`

?check id=sol-040

## Model answer: Q5

*"Devoting a larger share of national output to investment would help restore rapid
productivity growth and rising living standards." Assume k is below the Golden Rule.*

1. **Immediately**, a higher investment share means a smaller consumption share, so
   **living standards (consumption) fall**.
2. Capital then accumulates faster, so the **growth rates of output and output per
   worker (productivity) rise**, but only during the transition.
3. In the **new steady state** output grows at n and output per worker at zero
   again. Because k started *below* the Golden Rule, the new steady state has **higher
   consumption**, so living standards end up higher.
4. So the claim is **half right**: productivity growth rises in the short run but not in
   the long run, and living standards fall first and rise only over time. The quotation
   stresses the growth and **ignores the sacrifice** needed to get it.

`[slides 82–83]`
