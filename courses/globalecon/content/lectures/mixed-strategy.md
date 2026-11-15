---
id: mixed-strategy
title: "Mixed-Strategy Nash Equilibrium"
deck: "Quiz Prep"
slides: 6
summary: "When a 2×2 game has two pure equilibria, players can randomise: how to find the pure equilibria, then p, q and the expected payoffs by making the opponent indifferent."
---

## The words

- A **Nash equilibrium** is a pair of strategies where neither player can gain by
  changing their own strategy alone.
- A **pure strategy** is one action played for sure. A **mixed strategy** is a
  *probability distribution* over actions, used to avoid being predictable.
- In a **strictly mixed** equilibrium of a 2×2 game, the row player plays U with
  probability **p > 0** and the column player plays L with probability **q > 0**, each
  a best response to the other.

The key idea: **you choose your probability to make your *opponent* indifferent**
between their two actions. If they are indifferent, they cannot exploit any pattern in
what you do. Making the opponent indifferent in expected payoff = minimising their
ability to recognise and exploit your behaviour.

`[slides 43–44]`

## The packet's game

|  | **L** (q) | **R** (1 − q) |
|---|---|---|
| **U** (p) | 5, 1 | 0, 0 |
| **D** (1 − p) | 4, 4 | 1, 5 |

Payoffs are (Player 1, Player 2). Player 1 chooses the row, Player 2 the column.

## Step 1: find the pure equilibria

For each cell ask: *would Player 1 switch rows? Would Player 2 switch columns?*

| Cell | P1 compares | P2 compares | Equilibrium? |
|---|---|---|---|
| (U, L) | 5 vs 4: stays | 1 vs 0: stays | **yes** |
| (U, R) | 0 vs 1: switches | | no |
| (D, L) | 4 vs 5: switches | | no |
| (D, R) | 1 vs 0: stays | 5 vs 4: stays | **yes** |

Two pure equilibria, and each player prefers a different one (P1 wants (U, L), P2 wants
(D, R)). With **multiple pure Nash equilibria**, the players can keep each other
guessing with a mixed strategy.

`[slide 44]`

?check id=mix-006

## Step 2: find p from PLAYER 2's payoffs

Player 1 picks **p** to make **Player 2** indifferent between L and R. So you write
Player 2's expected payoffs, which are the *second* number in each cell:

```
Payoff₂(L) = 1·p + 4·(1 − p) = 4 − 3p
Payoff₂(R) = 0·p + 5·(1 − p) = 5 − 5p
4 − 3p = 5 − 5p   ⇒   2p = 1   ⇒   p = ½
```

> **Exam focus.** The one trap: **p comes from the opponent's payoffs, not your own.**
> If you set Player 1's own payoffs equal you get q instead, and if you then call it p
> the answer is marked wrong even though the arithmetic was right.

`[slide 45]`

## Step 3: find q from PLAYER 1's payoffs

Player 2 picks **q** to make **Player 1** indifferent between U and D:

```
Payoff₁(U) = 5·q + 0·(1 − q) = 5q
Payoff₁(D) = 4·q + 1·(1 − q) = 1 + 3q
5q = 1 + 3q   ⇒   2q = 1   ⇒   q = ½
```

`[slide 46]`

?check id=mix-016

## Step 4: the expected payoffs

Substitute back into either side of the indifference equation (they are equal, which is
the whole point):

- Player 1: Payoff(U) = 5q = 5 × ½ = **2.5** (and Payoff(D) = 1 + 3·½ = 2.5)
- Player 2: Payoff(L) = 4 − 3p = 4 − 1.5 = **2.5**

So the game has a **unique mixed equilibrium (p = ½, q = ½) with expected payoffs (2.5,
2.5)**. At these probabilities both players are indifferent between their strategies and
have no incentive to deviate.

The general recipe for any 2×2 game:

```
p = (b_DR − b_DL) / (b_UL − b_DL − b_UR + b_DR)     ← b = Player 2's payoffs
q = (a_DR − a_UR) / (a_UL − a_UR − a_DL + a_DR)     ← a = Player 1's payoffs
```

But in the booklet, write the two indifference equations. That is what earns the marks.

`[slide 47]`

## Where it shows up

Operating-system providers (Android, iOS) can use mixed strategies when choosing which
new features to put in each update: an element of randomness in the iterative development
decisions, giving periodic updates with varying features. The **stable point** is where
neither provider has much incentive to make abrupt changes to its strategy, because both
must keep their competitive position and satisfy users and developers.

`[slide 48]`
