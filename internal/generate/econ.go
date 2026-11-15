package generate

import "math"

// The economics drills: closed-form solutions for the models in the Economic Concepts
// quiz packet, so a generated question can never disagree with the method on the
// slides. econ_test.go pins every one of them to the packet's own worked numbers.

// Market is a linear market with two identical firms: inverse demand P = A − Q and a
// constant marginal cost C for each firm. The packet's Grab/Uber case is A=339, C=147.
type Market struct {
	A, C float64
}

// Outcome is one market structure's equilibrium.
type Outcome struct {
	Name string
	// Q1 and Q2 are the two firms' outputs; Q is the market total.
	Q1, Q2, Q float64
	P         float64
	// Profit1, Profit2 are each firm's profit, Profit the total: (P − C) × Q.
	Profit1, Profit2, Profit float64
	// CS is consumer surplus, ½ × Q × (A − P), the triangle under demand above price.
	CS float64
	// DWL is the deadweight loss against the competitive outcome.
	DWL float64
	// Value is what the packet's results table calls "Total Value": P × Q.
	Value float64
}

// Gap is A − C, the distance between the demand intercept and marginal cost. Every
// equilibrium quantity in the packet is a simple fraction of it.
func (m Market) Gap() float64 { return m.A - m.C }

func (m Market) outcome(name string, q1, q2 float64) Outcome {
	q := q1 + q2
	p := m.A - q
	comp := m.Gap() // competitive output
	return Outcome{
		Name: name, Q1: q1, Q2: q2, Q: q, P: p,
		Profit1: (p - m.C) * q1, Profit2: (p - m.C) * q2, Profit: (p - m.C) * q,
		CS:    0.5 * q * (m.A - p),
		DWL:   0.5 * (comp - q) * (p - m.C),
		Value: p * q,
	}
}

// Monopoly: one firm sets MR = A − 2Q equal to C, so Q = (A − C)/2. The whole output
// is the single firm's, so Q2 is zero.
func (m Market) Monopoly() Outcome { return m.outcome("Monopoly", m.Gap()/2, 0) }

// Cartel: two firms collude on the monopoly output and split it evenly.
func (m Market) Cartel() Outcome { return m.outcome("Cartel", m.Gap()/4, m.Gap()/4) }

// Bertrand: price competition with identical products undercuts down to P = MC, so the
// market produces the competitive output A − C, split evenly.
func (m Market) Bertrand() Outcome { return m.outcome("Bertrand", m.Gap()/2, m.Gap()/2) }

// Cournot: each firm's reaction function is q_i = (A − C)/2 − q_j/2; solving the pair
// gives (A − C)/3 each.
func (m Market) Cournot() Outcome { return m.outcome("Cournot", m.Gap()/3, m.Gap()/3) }

// Stackelberg: firm 1 leads, maximising against the follower's reaction function. Its
// residual demand is P = (A + C)/2 − q1/2, so q1 = (A − C)/2 and the follower plays
// (A − C)/4.
func (m Market) Stackelberg() Outcome { return m.outcome("Stackelberg", m.Gap()/2, m.Gap()/4) }

// ReactionIntercept is the constant in q_i = ReactionIntercept − ½ q_j.
func (m Market) ReactionIntercept() float64 { return m.Gap() / 2 }

// LeaderDemandIntercept is the constant in the Stackelberg leader's residual demand,
// P = A − (q1 + (Gap/2 − q1/2)) = (A + C)/2 − q1/2.
func (m Market) LeaderDemandIntercept() float64 { return m.A - m.Gap()/2 }

// Cheat is what happens when one cartel member keeps to its quota of Gap/4 and the
// other plays its best response to that instead. The cheater produces 3·Gap/8.
func (m Market) Cheat() Outcome {
	loyal := m.Gap() / 4
	cheat := m.ReactionIntercept() - loyal/2
	return m.outcome("Cheating", cheat, loyal)
}

// MonopolyMR is the monopolist's marginal revenue at output q: d(P·q)/dq = A − 2q.
func (m Market) MonopolyMR(q float64) float64 { return m.A - 2*q }

// Game is a 2×2 game in normal form. Row picks U or D, column picks L or R. Payoffs
// are indexed [row][col], with row 0 = U, col 0 = L.
type Game struct {
	Row [2][2]float64
	Col [2][2]float64
}

// Mixed is a game's fully mixed Nash equilibrium.
type Mixed struct {
	// P is the probability the row player plays U, Q the probability the column player
	// plays L.
	P, Q float64
	// RowPay and ColPay are each player's expected payoff at the equilibrium.
	RowPay, ColPay float64
}

// SolveMixed finds the mixed equilibrium by indifference. P is chosen to make the
// *column* player indifferent between L and R, so it is computed from the column
// player's payoffs; Q likewise from the row player's. ok is false if there is no
// strictly mixed equilibrium.
func (g Game) SolveMixed() (Mixed, bool) {
	// Column indifference: p·Col[U][L] + (1−p)·Col[D][L] = p·Col[U][R] + (1−p)·Col[D][R]
	dp := g.Col[0][0] - g.Col[1][0] - g.Col[0][1] + g.Col[1][1]
	// Row indifference: q·Row[U][L] + (1−q)·Row[U][R] = q·Row[D][L] + (1−q)·Row[D][R]
	dq := g.Row[0][0] - g.Row[0][1] - g.Row[1][0] + g.Row[1][1]
	if dp == 0 || dq == 0 {
		return Mixed{}, false
	}
	p := (g.Col[1][1] - g.Col[1][0]) / dp
	q := (g.Row[1][1] - g.Row[0][1]) / dq
	if p <= 0 || p >= 1 || q <= 0 || q >= 1 {
		return Mixed{}, false
	}
	return Mixed{
		P: p, Q: q,
		RowPay: q*g.Row[0][0] + (1-q)*g.Row[0][1],
		ColPay: p*g.Col[0][0] + (1-p)*g.Col[1][0],
	}, true
}

// PureNash lists the pure-strategy equilibria as [row, col] pairs: cells where neither
// player gains by switching alone.
func (g Game) PureNash() [][2]int {
	var out [][2]int
	for r := 0; r < 2; r++ {
		for c := 0; c < 2; c++ {
			if g.Row[r][c] >= g.Row[1-r][c] && g.Col[r][c] >= g.Col[r][1-c] {
				out = append(out, [2]int{r, c})
			}
		}
	}
	return out
}

// Solow is a Solow economy with Cobb–Douglas output per effective worker y = k^Alpha.
type Solow struct {
	Alpha, S, Delta, N, G float64
}

// BreakEven is the rate (δ + n + g) at which capital per effective worker wears away.
func (s Solow) BreakEven() float64 { return s.Delta + s.N + s.G }

// SteadyK solves s·k^α = (δ + n + g)·k for k* = (s / (δ + n + g))^(1/(1−α)).
func (s Solow) SteadyK() float64 {
	return math.Pow(s.S/s.BreakEven(), 1/(1-s.Alpha))
}

// Y is output per worker at capital k.
func (s Solow) Y(k float64) float64 { return math.Pow(k, s.Alpha) }

// MPK is the slope of the production function at k: α·k^(α−1).
func (s Solow) MPK(k float64) float64 { return s.Alpha * math.Pow(k, s.Alpha-1) }

// DeltaK is the change in capital per worker at k: s·f(k) − (δ + n + g)·k.
func (s Solow) DeltaK(k float64) float64 { return s.S*s.Y(k) - s.BreakEven()*k }

// GoldenK is the Golden Rule capital stock, where MPK = δ + n + g:
// α·k^(α−1) = δ + n + g  ⇒  k = (α / (δ + n + g))^(1/(1−α)).
func (s Solow) GoldenK() float64 {
	return math.Pow(s.Alpha/s.BreakEven(), 1/(1-s.Alpha))
}

// GoldenS is the saving rate that makes the Golden Rule stock the steady state. For
// Cobb–Douglas it is always α.
func (s Solow) GoldenS() float64 {
	k := s.GoldenK()
	return s.BreakEven() * k / s.Y(k)
}

// ComputersToAdd is the packet's slide-88 thought experiment: n computers for n
// workers (k = 1), and after one period depreciation, population growth and
// technology have acted. It returns how many computers must be added to keep one per
// effective worker, and the number of effective workers. The new workers do not get
// the technology gain, which is why the answer adds the rates rather than
// multiplying them.
func ComputersToAdd(n, delta, pop, g float64) (add, effective float64) {
	effective = n*(1+g) + n*pop
	remaining := n * (1 - delta)
	return effective - remaining, effective
}
