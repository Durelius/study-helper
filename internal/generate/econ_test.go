package generate

import (
	"math"
	"testing"
)

func near(a, b float64) bool { return math.Abs(a-b) < 1e-9 }

// TestMarketMatchesThePacketTable checks every structure against the Grab/Uber
// numbers on slides 28–41 and the results table on slide 41. If the formulas and the
// slides ever disagree, this fails rather than the drill quietly teaching other numbers.
func TestMarketMatchesThePacketTable(t *testing.T) {
	m := Market{A: 339, C: 147}
	table := []struct {
		o                 Outcome
		q, p, profit, val float64
	}{
		{m.Bertrand(), 192, 147, 0, 28224},
		{m.Stackelberg(), 144, 195, 6912, 28080},
		{m.Cournot(), 128, 211, 8192, 27008},
		{m.Cartel(), 96, 243, 9216, 23328},
		{m.Monopoly(), 96, 243, 9216, 23328},
	}
	for _, row := range table {
		o := row.o
		if !near(o.Q, row.q) || !near(o.P, row.p) || !near(o.Profit, row.profit) || !near(o.Value, row.val) {
			t.Errorf("%s: got Q=%v P=%v profit=%v value=%v, want %v %v %v %v",
				o.Name, o.Q, o.P, o.Profit, o.Value, row.q, row.p, row.profit, row.val)
		}
	}

	c := m.Cournot()
	if !near(c.Q1, 64) || !near(c.Q2, 64) || !near(c.Profit1, 4096) {
		t.Errorf("Cournot per firm: %+v, want 64/64 and 4096 each (slide 36)", c)
	}
	s := m.Stackelberg()
	if !near(s.Q1, 96) || !near(s.Q2, 48) || !near(s.Profit1, 4608) || !near(s.Profit2, 2304) {
		t.Errorf("Stackelberg: %+v, want 96/48 and 4608/2304 (slide 40)", s)
	}
	k := m.Cartel()
	if !near(k.Q1, 48) || !near(k.Profit1, 4608) {
		t.Errorf("Cartel: %+v, want 48 each and 4608 each (slide 29)", k)
	}
	if b := m.Bertrand(); !near(b.Q1, 96) {
		t.Errorf("Bertrand per firm %v, want 96 (slide 33)", b.Q1)
	}
	if !near(m.ReactionIntercept(), 96) {
		t.Errorf("reaction intercept %v, want 96 (slide 35)", m.ReactionIntercept())
	}
	if !near(m.LeaderDemandIntercept(), 243) {
		t.Errorf("leader residual demand intercept %v, want 243 (slide 39)", m.LeaderDemandIntercept())
	}
	// Slide 23: MR of the 101st ride is about 138; the derivative at q=100 is 139 and
	// the discrete step 100→101 is 138. The drill uses the derivative.
	if !near(m.MonopolyMR(96), 147) {
		t.Errorf("MR at the monopoly output should equal MC")
	}
	// Slide 26: the MR line hits zero at 169.5.
	if !near(m.A/2, 169.5) {
		t.Errorf("MR intercept")
	}
}

func TestCheatingPaysTheCheater(t *testing.T) {
	m := Market{A: 339, C: 147}
	ch := m.Cheat()
	if !near(ch.Q1, 72) || !near(ch.P, 219) || !near(ch.Profit1, 5184) || !near(ch.Profit2, 3456) {
		t.Errorf("cheat: %+v", ch)
	}
	if ch.Profit1 <= m.Cartel().Profit1 || ch.Profit2 >= m.Cartel().Profit2 {
		t.Error("cheating should pay the cheater and cost the loyal firm")
	}
}

// TestMixedMatchesSlide44 is the packet's game: p = q = ½ and payoff 2.5 for both.
func TestMixedMatchesSlide44(t *testing.T) {
	g := Game{
		Row: [2][2]float64{{5, 0}, {4, 1}},
		Col: [2][2]float64{{1, 0}, {4, 5}},
	}
	mx, ok := g.SolveMixed()
	if !ok || !near(mx.P, 0.5) || !near(mx.Q, 0.5) || !near(mx.RowPay, 2.5) || !near(mx.ColPay, 2.5) {
		t.Errorf("mixed: %+v %v", mx, ok)
	}
	ne := g.PureNash()
	if len(ne) != 2 || ne[0] != [2]int{0, 0} || ne[1] != [2]int{1, 1} {
		t.Errorf("pure NE %v, want (U,L) and (D,R)", ne)
	}
}

func TestSolow(t *testing.T) {
	// Steady state of y = √k with s = 0.24 and δ + n = 0.06: k* = 16.
	s := Solow{Alpha: 0.5, S: 0.24, Delta: 0.05, N: 0.01}
	if !near(s.SteadyK(), 16) || math.Abs(s.DeltaK(16)) > 1e-12 {
		t.Errorf("steady k %v", s.SteadyK())
	}
	if s.DeltaK(9) <= 0 || s.DeltaK(25) >= 0 {
		t.Error("capital should grow below k* and shrink above it")
	}
	// Golden Rule of y = √k with δ = 0.05: MPK = 1/(2√k) = 0.05 ⇒ k = 100, s = ½.
	g := Solow{Alpha: 0.5, Delta: 0.05}
	if !near(g.GoldenK(), 100) || !near(g.GoldenS(), 0.5) || !near(g.MPK(100), 0.05) {
		t.Errorf("golden %v %v", g.GoldenK(), g.GoldenS())
	}
	if c := (Solow{Alpha: 0.3, Delta: 0.1}); !near(c.GoldenS(), 0.3) {
		t.Errorf("s_gold should equal α, got %v", c.GoldenS())
	}
}

// TestComputersMatchSlides88and89: 13 and 9 computers to add.
func TestComputersMatchSlides88and89(t *testing.T) {
	add, eff := ComputersToAdd(10, 0.1, 0.2, 1.0)
	if !near(add, 13) || !near(eff, 22) {
		t.Errorf("slide 88: add %v for %v effective workers, want 13 for 22", add, eff)
	}
	add, eff = ComputersToAdd(10, 0.2, 0.2, 0.5)
	if !near(add, 9) || !near(eff, 17) {
		t.Errorf("slide 89: add %v for %v effective workers, want 9 for 17", add, eff)
	}
}
