---
id: oligopoly
title: "Oligopoly: Grab vs Uber"
deck: "Quiz Prep"
slides: 30
summary: "Five ways two firms can share one market (monopoly, cartel, Bertrand, Cournot, Stackelberg), each solved from P = 339 − Q and MC = 147, plus the surplus areas and why cartels break down."
---

## The setup and the one idea behind every answer

Two firms, Grab and Uber, sell an identical ride in Bangkok. The packet fixes:

| | |
|---|---|
| Inverse market demand | **P = 339 − Q** (Q = rides per minute, P in THB) |
| Marginal cost, each firm | **MC = 147 THB**, constant |

**Marginal cost** is the extra cost of one more unit. **Marginal revenue (MR)** is the
extra revenue from selling one more unit. Every quantity-setting firm in this chapter does
the same thing: it **produces until MR = MC**. The structures differ only in *whose*
revenue the firm is maximising and *what it assumes the rival does*.

A useful number to hold: **A − C = 339 − 147 = 192**. Every answer in the packet is a
simple fraction of 192.

`[slides 19–22]`

## Math primer: marginal revenue by derivative

You need one calculus rule, the **power rule**: the derivative of a·qⁿ is n·a·qⁿ⁻¹.

- d/dq of 339q = 339
- d/dq of q² = 2q
- d/dq of 2q² = 4q
- d/dq of q·q_U, holding q_U fixed, = q_U

So to get MR, write revenue as **R = P × q**, substitute demand for P, multiply out, and
differentiate term by term:

```
R  = (339 − q)·q = 339q − q²
MR = dR/dq     = 339 − 2q
```

**MR falls twice as fast as demand.** Demand hits zero at Q = 339; MR hits zero at 169.5.
That is the 169.5 marked on the surplus graphs.

The packet also checks this by hand: the 1st ride sells for 338, so MR ≈ 338. Going from
100 to 101 rides, revenue goes 23,900 → 24,038, so MR ≈ 138, already *below* MC = 147.
Selling the 101st ride loses money.

`[slides 23, 28]`

?check id=olig-004

## 1 · Monopoly (one firm)

The monopolist's demand *is* the market demand.

```
MR = 339 − 2Q = 147 = MC
2Q = 192           ⇒  Q = 96
P  = 339 − 96      =  243      ← read P off DEMAND, not off MR
Profit = (243 − 147) × 96 = 9,216 THB/min
```

> **Exam focus.** The commonest slip is reading the price off the MR line (339 − 2·96 =
> 147). Q comes from MR = MC; **P always comes from the demand curve.**

**Competitive benchmark.** If Grab were one of many firms, price would equal MC: 147 =
339 − Q ⇒ **Q = 192, P = 147, profit 0**. Monopoly halves output.

`[slides 24–25, 28]`

## 2 · Cartel (collusion)

Grab and Uber agree to act as one monopolist and split the spoils. With q_G = q_U = q, total
Q = 2q, so P = 339 − 2q.

```
R  = (339 − 2q)·q = 339q − 2q²
MR = 339 − 4q = 147  ⇒  4q = 192  ⇒  q = 48 each
Q  = 96, P = 339 − 96 = 243        (the monopoly price)
Profit = 96 × 96 = 9,216 total, 4,608 each
```

**Why cartels break down.** At the cartel price (243) each firm's price is far above its MC
(147), so one more ride is profitable *for the firm that sells it*. Each member has an
incentive to cheat by producing more, and gets away with it if others cannot tell who
produced the extra output. As firms leave, the price falls. Demand is unchanged, so
cheating is **zero-sum**: the cheater gains at the loyal firm's expense.

With the Cournot reaction function you can put numbers on it: if Uber keeps its quota of
48 and Grab best-responds, Grab makes 96 − ½·48 = **72**, the price falls to 339 − 120 =
**219**, Grab earns 72 × 72 = **5,184** (vs 4,608) and Uber drops to 72 × 48 = **3,456**.

`[slides 29–31]`

?check id=olig-017

## 3 · Bertrand (compete on price)

Firms set **prices**, not quantities, and let consumers decide how much to buy. With
identical products, whoever is cheaper takes the whole market, so each undercuts the
other **until P = MC**.

```
P = MC = 147
147 = 339 − Q  ⇒  Q = 192, so 96 each
Profit = 0
```

The **Nash–Bertrand equilibrium equals the competitive outcome**. Output is double the
monopoly's. Two firms are enough for perfect competition here.

`[slides 32–33]`

## 4 · Cournot (choose quantities at the same time)

Each firm picks its output **simultaneously**, taking the rival's as given. Now Grab's
price depends on both outputs: P = 339 − q_G − q_U.

```
R_G  = (339 − q_G − q_U)·q_G = 339q_G − q_G² − q_U·q_G
MR_G = 339 − 2q_G − q_U          (q_U treated as a constant)
MR_G = MC:  339 − 2q_G − q_U = 147
            2q_G = 192 − q_U
            q_G  = 96 − ½q_U     ← Grab's REACTION FUNCTION
```

By symmetry Uber's is **q_U = 96 − ½q_G**. A **reaction (best-response) function** gives
the profit-maximising output for one firm *given its belief about the rival's*. The Nash
equilibrium is where the two cross: each is best-responding to the other.

```
q_G = 96 − ½(96 − ½q_G) = 48 + ¼q_G
¾q_G = 48  ⇒  q_G = 64, q_U = 64
Q = 128,  P = 339 − 128 = 211
Profit = (211 − 147) × 64 = 4,096 each, 8,192 total
```

On the best-response graph, Grab's curve runs from (0, 192) to (96, 0) and Uber's from
(192, 0) to (0, 96); they cross at **(64, 64)**.

> **Exam focus.** Cournot lies **between** monopoly and competition on every measure:
> 96 < 128 < 192 for output, 243 > 211 > 147 for price.

`[slides 34–37]`

?check id=olig-028

## 5 · Stackelberg (leader moves first)

Grab **leads**: it sets output first, knowing Uber will then play its reaction function
q_U = 96 − ½q_G. So Grab substitutes Uber's response into the demand it faces:

```
P = 339 − q_G − (96 − ½q_G) = 243 − ½q_G     ← leader's residual demand
R = 243q_G − ½q_G²
MR = 243 − q_G = 147  ⇒  q_G = 96
q_U = 96 − ½·96 = 48
Q = 144,  P = 339 − 144 = 195
Profit: Grab (195 − 147) × 96 = 4,608;  Uber 48 × 48 = 2,304;  total 6,912
```

Outputs are **asymmetric**: the leader produces twice the follower's output and earns twice
its profit. Moving first pays, because committing to a large output forces the
follower to shrink.

`[slides 38–40]`

## The results table

| | Bertrand (competitive) | Stackelberg (sequential) | Cournot (simultaneous) | Collusion (monopoly) |
|---|---|---|---|---|
| Total output | **192** | **144** | **128** | **96** |
| Price | 147 | 195 | 211 | 243 |
| Total Value (= P × Q) | 28,224 | 28,080 | 27,008 | 23,328 |
| Total profit | 0 | 6,912 | 8,192 | 9,216 |

Read it left to right: **output and Total Value fall, price and profit rise.** The more
competition, the more value is created for the market; the more collusion, the more
profit for the business.

Note that the packet's **"Total Value" is P × Q**, total spending on rides (147 × 192 =
28,224). It is *not* consumer surplus. Use P × Q if the quiz asks for it from the table.

**General formulas** for P = A − Q, MC = C (check: A − C = 192):

| | Each firm's q | Total Q | Price |
|---|---|---|---|
| Monopoly | (A − C)/2 | (A − C)/2 | (A + C)/2 |
| Cartel | (A − C)/4 | (A − C)/2 | (A + C)/2 |
| Cournot | (A − C)/3 | 2(A − C)/3 | C + (A − C)/3 |
| Stackelberg | leader (A − C)/2, follower (A − C)/4 | 3(A − C)/4 | C + (A − C)/4 |
| Bertrand | (A − C)/2 | A − C | C |

`[slides 21, 41–42]`

?check id=olig-046

## Surplus on the graph

- **Consumer surplus (CS)**: the value consumers enjoy from paying less than they were
  willing to. It is the area **under demand, above the price**: ½ × Q × (A − P).
- **Producer surplus (PS)**: what producers get from selling above marginal cost. With
  constant MC it is the rectangle **(P − MC) × Q** (= profit here).
- **Deadweight loss (DWL)**: surplus that nobody gets because output is below the
  competitive level. It is the triangle between demand and MC from the monopoly Q to the competitive Q.

On the consumer-surplus graph (A = top triangle above 243, C = rectangle 147–243 over 0–96,
B = triangle 96–192 between 147 and demand, D = area under MC):

| Area | Competitive (Bertrand) | Monopoly |
|---|---|---|
| CS | **A + C + B** = ½ × 192 × 192 = 18,432 | **A** = ½ × 96 × 96 = 4,608 |
| PS | **0** | **C** = 96 × 96 = 9,216 |
| DWL | 0 | **B** = ½ × 96 × 96 = 4,608 |

Monopoly converts rectangle C from consumers to the producer and destroys triangle B. The
PS graph uses different letters: there the monopoly PS rectangle is **B** and the
deadweight triangle is **C**, so **read the letters on the graph you are given**.

`[slides 26–27]`

?check id=olig-055
