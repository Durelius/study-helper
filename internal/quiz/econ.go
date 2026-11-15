package quiz

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/wilhelmdurelius/chulastudy/internal/content"
	"github.com/wilhelmdurelius/chulastudy/internal/generate"
)

// The economics labs. Every number comes from internal/generate, whose tests pin it to
// the quiz packet; this file only decides what to ask and how to explain it.

const econDeck = "Quiz Prep"

// econGenerators are the computed drills this file provides, by generator name.
var econGenerators = map[string]func(count int, rnd *rand.Rand) ([]content.Question, error){
	"oligopoly":       generatedOligopoly,
	"oligopoly-steps": generatedOligopolySteps,
	"oligopoly-table": generatedOligopolyTable,
	"solow-computers": generatedComputers,
	"mixed":           generatedMixed,
	"solow":           generatedSolow,
}

var fmtN = content.FormatNumber

func numQ(topic, stem string, value float64, unit, expl string, slide int, tags ...string) content.Question {
	v := value
	return content.Question{
		Topic: topic, Type: "num", Stem: stem, Value: &v, Unit: unit,
		Explanation: expl, Source: content.Source{Deck: econDeck, Slide: slide},
		Difficulty: 2, Tags: append([]string{"generated", "formula"}, tags...), ExamFocus: true,
	}
}

func mcqQ(topic, stem, correct string, wrong []string, expl string, slide int, rnd *rand.Rand, tags ...string) (content.Question, bool) {
	choices, answer, ok := choose(correct, wrong, rnd)
	if !ok {
		return content.Question{}, false
	}
	return content.Question{
		Topic: topic, Type: "mcq", Stem: stem, Choices: choices, Answer: []int{answer},
		Explanation: expl, Source: content.Source{Deck: econDeck, Slide: slide},
		Difficulty: 2, Tags: append([]string{"generated"}, tags...), ExamFocus: true,
	}, true
}

// --- Oligopoly -------------------------------------------------------------------

var firmPairs = [][2]string{
	{"Grab", "Uber"}, {"Grab", "LINE MAN"}, {"Bolt", "Grab"}, {"foodpanda", "LINE MAN"},
	{"Robinhood", "Grab"}, {"Bolt", "inDrive"},
}

// randomMarket picks A and C so that A − C is a multiple of 12, which makes every
// quantity, price and profit in all five structures a whole number (and the CS
// triangles whole or half).
func randomMarket(rnd *rand.Rand) (generate.Market, [2]string) {
	gap := 12 * float64(8+rnd.Intn(13)) // 96 … 240
	c := float64(40 + rnd.Intn(161))    // 40 … 200
	return generate.Market{A: c + gap, C: c}, firmPairs[rnd.Intn(len(firmPairs))]
}

func marketSetup(m generate.Market, f [2]string) string {
	return fmt.Sprintf("**Setup.** %s and %s sell an identical service. Inverse market demand is **P = %s − Q** (Q in rides per minute, P in THB) and each firm's marginal cost is a constant **MC = %s THB**.",
		f[0], f[1], fmtN(m.A), fmtN(m.C))
}

// generatedOligopoly is a run of single questions, each on its own fresh market, so
// nothing carries over from the question before.
func generatedOligopoly(count int, rnd *rand.Rand) ([]content.Question, error) {
	var out []content.Question
	for tries := 0; len(out) < count && tries < count*20; tries++ {
		m, f := randomMarket(rnd)
		q, ok := oligopolyQuestion(m, f, rnd)
		if !ok {
			continue
		}
		q.ID = fmt.Sprintf("gen:olig:%d:%d", len(out), rnd.Int63())
		out = append(out, q)
	}
	if len(out) < count {
		return nil, fmt.Errorf("could not generate oligopoly questions")
	}
	return out, nil
}

func oligopolyQuestion(m generate.Market, f [2]string, rnd *rand.Rand) (content.Question, bool) {
	setup := marketSetup(m, f) + "\n\n"
	a, c, gap := fmtN(m.A), fmtN(m.C), m.Gap()
	mo, ca, be, co, st := m.Monopoly(), m.Cartel(), m.Bertrand(), m.Cournot(), m.Stackelberg()
	const T = "oligopoly"

	switch rnd.Intn(15) {
	case 0:
		return numQ(T, setup+fmt.Sprintf("If %s were a **monopoly**, what output would it choose?", f[0]),
			mo.Q, "rides/min",
			fmt.Sprintf("Revenue = (%s − Q)Q, so MR = %s − 2Q. Set MR = MC: %s − 2Q = %s ⇒ Q = %s/2 = %s.", a, a, a, c, fmtN(gap), fmtN(mo.Q)),
			28, "monopoly"), true
	case 1:
		return numQ(T, setup+"What price does a **monopolist** charge?", mo.P, "THB",
			fmt.Sprintf("Q = (%s − %s)/2 = %s from MR = MC, then read the price off the demand curve: P = %s − %s = %s. The trap is reading it off the MR line.", a, c, fmtN(mo.Q), a, fmtN(mo.Q), fmtN(mo.P)),
			28, "monopoly"), true
	case 2:
		return numQ(T, setup+"What is the **monopoly profit** per minute?", mo.Profit, "THB",
			fmt.Sprintf("Profit = (P − MC) × Q = (%s − %s) × %s = %s.", fmtN(mo.P), c, fmtN(mo.Q), fmtN(mo.Profit)),
			28, "monopoly"), true
	case 3:
		return numQ(T, setup+fmt.Sprintf("The two firms form a **cartel** and split the output equally. How much does %s produce?", f[0]),
			ca.Q1, "rides/min",
			fmt.Sprintf("A cartel acts as one monopolist: total Q = (%s − %s)/2 = %s, at P = %s. Split equally, each produces %s and earns (%s − %s) × %s = %s.", a, c, fmtN(ca.Q), fmtN(ca.P), fmtN(ca.Q1), fmtN(ca.P), c, fmtN(ca.Q1), fmtN(ca.Profit1)),
			29, "cartel"), true
	case 4:
		return numQ(T, setup+"The firms compete on **price** (Bertrand). What is the total market output?",
			be.Q, "rides/min",
			fmt.Sprintf("Each undercuts the other until P = MC = %s. Then demand gives Q = %s − %s = %s (%s each), and profit is zero, the same as the competitive outcome.", c, a, c, fmtN(be.Q), fmtN(be.Q1)),
			33, "bertrand"), true
	case 5:
		return numQ(T, setup+fmt.Sprintf("The firms choose quantities simultaneously (**Cournot**). %s's reaction function is q₁ = X − ½q₂. What is X?", f[0]),
			m.ReactionIntercept(), "",
			fmt.Sprintf("Revenue₁ = (%s − q₁ − q₂)q₁ ⇒ MR₁ = %s − 2q₁ − q₂. Set = %s: 2q₁ = %s − q₂ ⇒ q₁ = %s − ½q₂.", a, a, c, fmtN(gap), fmtN(m.ReactionIntercept())),
			35, "cournot", "reaction"), true
	case 6:
		return numQ(T, setup+fmt.Sprintf("In the **Cournot** equilibrium, how much does %s produce?", f[0]),
			co.Q1, "rides/min",
			fmt.Sprintf("Reaction functions q₁ = %s − ½q₂ and q₂ = %s − ½q₁. By symmetry q₁ = q₂ = q: q = %s − ½q ⇒ 1.5q = %s ⇒ q = %s, i.e. (A − C)/3.", fmtN(gap/2), fmtN(gap/2), fmtN(gap/2), fmtN(gap/2), fmtN(co.Q1)),
			36, "cournot"), true
	case 7:
		return numQ(T, setup+"What is the **Cournot** market price?", co.P, "THB",
			fmt.Sprintf("Each firm makes %s, so Q = %s and P = %s − %s = %s. It lies between the monopoly price %s and MC %s.", fmtN(co.Q1), fmtN(co.Q), a, fmtN(co.Q), fmtN(co.P), fmtN(mo.P), c),
			36, "cournot"), true
	case 8:
		return numQ(T, setup+"What profit does **each** firm earn in the **Cournot** equilibrium?", co.Profit1, "THB",
			fmt.Sprintf("q = %s each, P = %s. Profit each = (%s − %s) × %s = %s.", fmtN(co.Q1), fmtN(co.P), fmtN(co.P), c, fmtN(co.Q1), fmtN(co.Profit1)),
			36, "cournot"), true
	case 9:
		return numQ(T, setup+fmt.Sprintf("%s moves first as a **Stackelberg leader**. How much does the leader produce?", f[0]),
			st.Q1, "rides/min",
			fmt.Sprintf("The follower plays q₂ = %s − ½q₁. Substitute into demand: P = %s − q₁ − (%s − ½q₁) = %s − ½q₁. MR = %s − q₁ = %s ⇒ q₁ = %s.", fmtN(gap/2), a, fmtN(gap/2), fmtN(m.LeaderDemandIntercept()), fmtN(m.LeaderDemandIntercept()), c, fmtN(st.Q1)),
			40, "stackelberg"), true
	case 10:
		return numQ(T, setup+fmt.Sprintf("%s leads (**Stackelberg**). What is the market price?", f[0]),
			st.P, "THB",
			fmt.Sprintf("Leader %s, follower %s − ½·%s = %s, so Q = %s and P = %s − %s = %s.", fmtN(st.Q1), fmtN(gap/2), fmtN(st.Q1), fmtN(st.Q2), fmtN(st.Q), a, fmtN(st.Q), fmtN(st.P)),
			40, "stackelberg"), true
	case 11:
		return numQ(T, setup+"What is **consumer surplus** under **monopoly**?", mo.CS, "THB",
			fmt.Sprintf("The triangle under demand and above price: ½ × Q × (A − P) = ½ × %s × (%s − %s) = %s. (Area A on slide 26.)", fmtN(mo.Q), a, fmtN(mo.P), fmtN(mo.CS)),
			26, "surplus"), true
	case 12:
		return numQ(T, setup+"What is the **deadweight loss** of monopoly, compared with the competitive (Bertrand) outcome?", mo.DWL, "THB",
			fmt.Sprintf("The triangle between demand and MC from Q = %s to Q = %s: ½ × (%s − %s) × (%s − %s) = %s. (Area B on slide 26.)", fmtN(mo.Q), fmtN(be.Q), fmtN(be.Q), fmtN(mo.Q), fmtN(mo.P), c, fmtN(mo.DWL)),
			26, "surplus"), true
	case 13:
		order := []generate.Outcome{be, st, co, ca}
		pick := order[rnd.Intn(len(order))]
		return numQ(T, setup+fmt.Sprintf("The packet's results table reports **Total Value = P × Q**. What is it under **%s**?", structureName(pick.Name)),
			pick.Value, "THB/min",
			fmt.Sprintf("%s: Q = %s, P = %s, so P × Q = %s. Total Value falls as output falls: Bertrand > Stackelberg > Cournot > Collusion, while profit rises in the opposite order.", structureName(pick.Name), fmtN(pick.Q), fmtN(pick.P), fmtN(pick.Value)),
			41, "table"), true
	default:
		// Cheating output is 3(A − C)/8: keep it whole.
		if math.Mod(gap, 24) != 0 {
			return content.Question{}, false
		}
		ch := m.Cheat()
		return numQ(T, setup+fmt.Sprintf("The firms agree a cartel. %s keeps to its quota of %s, but %s **cheats** and plays its best response (its Cournot reaction function) instead. How much does %s produce?", f[1], fmtN(ca.Q1), f[0], f[0]),
			ch.Q1, "rides/min",
			fmt.Sprintf("Best response: q₁ = %s − ½ × %s = %s. Price falls to %s − %s = %s, and the cheater earns %s against %s from keeping the deal, while the loyal firm drops to %s. That is why each member has an incentive to cheat (slide 31).", fmtN(gap/2), fmtN(ca.Q1), fmtN(ch.Q1), a, fmtN(ch.Q), fmtN(ch.P), fmtN(ch.Profit1), fmtN(ca.Profit1), fmtN(ch.Profit2)),
			31, "cartel"), true
	}
}

func structureName(n string) string {
	switch n {
	case "Bertrand":
		return "Bertrand competition"
	case "Stackelberg":
		return "Stackelberg (sequential)"
	case "Cournot":
		return "Cournot (simultaneous)"
	case "Cartel":
		return "collusion (cartel)"
	}
	return n
}

// generatedOligopolySteps is a worked problem: one market and one structure, solved in
// the order the booklet expects, one number per step. The scenario stays on screen.
func generatedOligopolySteps(count int, rnd *rand.Rand) ([]content.Question, error) {
	var out []content.Question
	for p := 0; len(out) < count; p++ {
		m, f := randomMarket(rnd)
		steps, title := oligopolySteps(m, f, rnd.Intn(5))
		data, _ := json.Marshal(map[string]string{"title": title, "scenario": marketSetup(m, f)})
		for i, s := range steps {
			s.ID = fmt.Sprintf("gen:oligsteps:%d:%d:%d", p, i, rnd.Int63())
			s.Stem = fmt.Sprintf("**Step %d of %d.** %s", i+1, len(steps), s.Stem)
			s.Data = data
			s.Difficulty = 3
			out = append(out, s)
		}
	}
	return out, nil
}

func oligopolySteps(m generate.Market, f [2]string, which int) ([]content.Question, string) {
	a, c, gap := fmtN(m.A), fmtN(m.C), m.Gap()
	const T = "oligopoly"
	s := func(stem string, v float64, unit, expl string, slide int) content.Question {
		return numQ(T, stem, v, unit, expl, slide, "steps")
	}
	switch which {
	case 0:
		o := m.Monopoly()
		return []content.Question{
			s(fmt.Sprintf("%s is the only firm. Its revenue is R = (%s − Q)Q. At what Q is **MR = 0**?", f[0], a), m.A/2, "rides/min",
				fmt.Sprintf("MR = dR/dQ = %s − 2Q. MR = 0 at Q = %s/2 = %s: MR falls twice as fast as demand.", a, a, fmtN(m.A/2)), 26),
			s("Set MR = MC. What output maximises profit?", o.Q, "rides/min",
				fmt.Sprintf("%s − 2Q = %s ⇒ Q = %s.", a, c, fmtN(o.Q)), 28),
			s("What price does it charge?", o.P, "THB", fmt.Sprintf("Off the demand curve: P = %s − %s = %s.", a, fmtN(o.Q), fmtN(o.P)), 28),
			s("What is its profit?", o.Profit, "THB", fmt.Sprintf("(%s − %s) × %s = %s.", fmtN(o.P), c, fmtN(o.Q), fmtN(o.Profit)), 28),
			s("What is consumer surplus?", o.CS, "THB", fmt.Sprintf("½ × %s × (%s − %s) = %s.", fmtN(o.Q), a, fmtN(o.P), fmtN(o.CS)), 26),
			s("What is the deadweight loss against the competitive outcome?", o.DWL, "THB",
				fmt.Sprintf("Competitive Q = %s at P = %s. DWL = ½ × (%s − %s) × (%s − %s) = %s.", fmtN(gap), c, fmtN(gap), fmtN(o.Q), fmtN(o.P), c, fmtN(o.DWL)), 26),
		}, "Worked problem: monopoly"
	case 1:
		o := m.Cartel()
		return []content.Question{
			s(fmt.Sprintf("%s and %s collude. With q₁ = q₂ = q, the joint demand is P = %s − 2q. What is the coefficient on q in each firm's MR (MR = %s − k·q)?", f[0], f[1], a, a), 4, "",
				fmt.Sprintf("R = (%s − 2q)q = %sq − 2q², so MR = %s − 4q (slide 29).", a, a, a), 29),
			s("Set MR = MC. How much does each firm produce?", o.Q1, "rides/min", fmt.Sprintf("%s − 4q = %s ⇒ q = %s/4 = %s.", a, c, fmtN(gap), fmtN(o.Q1)), 29),
			s("What is the cartel price?", o.P, "THB", fmt.Sprintf("Q = 2 × %s = %s, so P = %s − %s = %s, the monopoly price.", fmtN(o.Q1), fmtN(o.Q), a, fmtN(o.Q), fmtN(o.P)), 29),
			s("What is each firm's profit?", o.Profit1, "THB", fmt.Sprintf("(%s − %s) × %s = %s, half the monopoly profit %s.", fmtN(o.P), c, fmtN(o.Q1), fmtN(o.Profit1), fmtN(o.Profit)), 29),
		}, "Worked problem: cartel"
	case 2:
		o := m.Bertrand()
		return []content.Question{
			s("The firms set **prices** (Bertrand) and the products are identical. What is the equilibrium price?", o.P, "THB",
				fmt.Sprintf("Any price above MC is undercut for the whole market, so the only equilibrium is P = MC = %s.", c), 32),
			s("What is total market output?", o.Q, "rides/min", fmt.Sprintf("%s = %s − Q ⇒ Q = %s.", c, a, fmtN(o.Q)), 33),
			s("How much does each firm produce?", o.Q1, "rides/min", fmt.Sprintf("Split evenly: %s/2 = %s.", fmtN(o.Q), fmtN(o.Q1)), 33),
			s("What is each firm's profit?", 0, "THB", "P = MC, so (P − MC) × q = 0. Nash–Bertrand is the competitive outcome.", 33),
			s("What is consumer surplus?", o.CS, "THB", fmt.Sprintf("½ × %s × (%s − %s) = %s: the whole triangle A + B + C on slide 26.", fmtN(o.Q), a, c, fmtN(o.CS)), 26),
		}, "Worked problem: Bertrand"
	case 3:
		o := m.Cournot()
		return []content.Question{
			s(fmt.Sprintf("Cournot. %s's revenue is R₁ = (%s − q₁ − q₂)q₁, so MR₁ = %s − 2q₁ − q₂. Setting MR₁ = MC gives q₁ = X − ½q₂. What is X?", f[0], a, a), m.ReactionIntercept(), "",
				fmt.Sprintf("%s − 2q₁ − q₂ = %s ⇒ 2q₁ = %s − q₂ ⇒ q₁ = %s − ½q₂.", a, c, fmtN(gap), fmtN(m.ReactionIntercept())), 34),
			s(fmt.Sprintf("Substitute %s's identical reaction function. What is q₁?", f[1]), o.Q1, "rides/min",
				fmt.Sprintf("q₁ = %s − ½(%s − ½q₁) = %s + ¼q₁ ⇒ ¾q₁ = %s ⇒ q₁ = %s.", fmtN(gap/2), fmtN(gap/2), fmtN(gap/4), fmtN(gap/4), fmtN(o.Q1)), 36),
			s("What is total market output?", o.Q, "rides/min", fmt.Sprintf("2 × %s = %s.", fmtN(o.Q1), fmtN(o.Q)), 36),
			s("What is the market price?", o.P, "THB", fmt.Sprintf("%s − %s = %s.", a, fmtN(o.Q), fmtN(o.P)), 36),
			s("What is each firm's profit?", o.Profit1, "THB", fmt.Sprintf("(%s − %s) × %s = %s.", fmtN(o.P), c, fmtN(o.Q1), fmtN(o.Profit1)), 36),
		}, "Worked problem: Cournot"
	default:
		o := m.Stackelberg()
		return []content.Question{
			s(fmt.Sprintf("Stackelberg: %s leads, %s follows. The follower's reaction function is q₂ = X − ½q₁. What is X?", f[0], f[1]), m.ReactionIntercept(), "",
				fmt.Sprintf("Same as Cournot: MR₂ = %s − 2q₂ − q₁ = %s ⇒ q₂ = %s − ½q₁.", a, c, fmtN(m.ReactionIntercept())), 38),
			s("Substitute it into demand. The leader faces P = Y − ½q₁. What is Y?", m.LeaderDemandIntercept(), "",
				fmt.Sprintf("P = %s − q₁ − (%s − ½q₁) = %s − ½q₁.", a, fmtN(gap/2), fmtN(m.LeaderDemandIntercept())), 39),
			s("The leader sets MR = Y − q₁ equal to MC. What is q₁?", o.Q1, "rides/min",
				fmt.Sprintf("%s − q₁ = %s ⇒ q₁ = %s.", fmtN(m.LeaderDemandIntercept()), c, fmtN(o.Q1)), 40),
			s("What does the follower produce?", o.Q2, "rides/min", fmt.Sprintf("q₂ = %s − ½ × %s = %s.", fmtN(gap/2), fmtN(o.Q1), fmtN(o.Q2)), 40),
			s("What is the market price?", o.P, "THB", fmt.Sprintf("Q = %s, P = %s − %s = %s.", fmtN(o.Q), a, fmtN(o.Q), fmtN(o.P)), 40),
			s("What is the leader's profit?", o.Profit1, "THB",
				fmt.Sprintf("(%s − %s) × %s = %s, double the follower's %s: moving first pays.", fmtN(o.P), c, fmtN(o.Q1), fmtN(o.Profit1), fmtN(o.Profit2)), 40),
		}, "Worked problem: Stackelberg"
	}
}

// generatedOligopolyTable is the quiz's 16-cell results table: given only demand and
// marginal cost, fill output, price, Total Value (P × Q) and total profit for each
// structure, in the packet's column order. One market per table; the scenario stays
// on screen for all sixteen cells.
func generatedOligopolyTable(count int, rnd *rand.Rand) ([]content.Question, error) {
	var out []content.Question
	for p := 0; len(out) < count; p++ {
		m, f := randomMarket(rnd)
		cells := oligopolyTable(m)
		scenario := marketSetup(m, f) + "\n\nFill the results table: **total output, price, Total Value (P × Q) and total profit** under Bertrand, Stackelberg, Cournot and collusion."
		data, _ := json.Marshal(map[string]string{"title": "Fill the table", "scenario": scenario})
		for i, c := range cells {
			c.ID = fmt.Sprintf("gen:oligtable:%d:%d:%d", p, i, rnd.Int63())
			c.Stem = fmt.Sprintf("**Cell %d of %d.** %s", i+1, len(cells), c.Stem)
			c.Data = data
			out = append(out, c)
		}
	}
	return out, nil
}

// oligopolyTable returns the sixteen cells column by column: Bertrand, Stackelberg,
// Cournot, collusion, each as total output, price, Total Value and total profit.
func oligopolyTable(m generate.Market) []content.Question {
	a, c, gap := fmtN(m.A), fmtN(m.C), fmtN(m.Gap())
	cols := []struct {
		name  string
		o     generate.Outcome
		qExpl string
		slide int
	}{
		{"Bertrand", m.Bertrand(), fmt.Sprintf("Price competition drives P down to MC = %s, so %s = %s − Q ⇒ Q = A − C = %s.", c, c, a, fmtN(m.Bertrand().Q)), 33},
		{"Stackelberg", m.Stackelberg(), fmt.Sprintf("Leader (A − C)/2 = %s/2 = %s, follower (A − C)/4 = %s; total ¾ × %s = %s.", gap, fmtN(m.Stackelberg().Q1), fmtN(m.Stackelberg().Q2), gap, fmtN(m.Stackelberg().Q)), 40},
		{"Cournot", m.Cournot(), fmt.Sprintf("Each firm (A − C)/3 = %s/3 = %s; total ⅔ × %s = %s.", gap, fmtN(m.Cournot().Q1), gap, fmtN(m.Cournot().Q)), 36},
		{"Collusion", m.Monopoly(), fmt.Sprintf("The cartel acts as one monopolist: MR = %s − 2Q = %s ⇒ Q = (A − C)/2 = %s.", a, c, fmtN(m.Monopoly().Q)), 41},
	}
	var out []content.Question
	for _, col := range cols {
		o := col.o
		q := func(stem string, v float64, unit, expl string) content.Question {
			x := numQ("oligopoly", stem, v, unit, expl, col.slide, "table")
			x.Difficulty = 2
			return x
		}
		out = append(out,
			q(col.name+": what is **total output** Q?", o.Q, "rides/min", col.qExpl),
			q(col.name+": what is the **price**?", o.P, "THB", fmt.Sprintf("Off the demand curve: P = %s − %s = %s.", a, fmtN(o.Q), fmtN(o.P))),
			q(col.name+": what is **Total Value** (P × Q)?", o.Value, "THB", fmt.Sprintf("Total Value = P × Q = %s × %s = %s. It is total spending, not consumer surplus.", fmtN(o.P), fmtN(o.Q), fmtN(o.Value))),
			q(col.name+": what is **total profit** (both firms)?", o.Profit, "THB", fmt.Sprintf("(P − MC) × Q = (%s − %s) × %s = %s.", fmtN(o.P), c, fmtN(o.Q), fmtN(o.Profit))),
		)
	}
	return out
}

// --- Mixed strategies ------------------------------------------------------------

// randomGame builds a 2×2 game with two pure equilibria on one diagonal and an
// interior mixed one, like the packet's: each player prefers a different equilibrium.
func randomGame(rnd *rand.Rand) (generate.Game, generate.Mixed) {
	for {
		var g generate.Game
		for r := 0; r < 2; r++ {
			for c := 0; c < 2; c++ {
				g.Row[r][c] = float64(rnd.Intn(7))
				g.Col[r][c] = float64(rnd.Intn(7))
			}
		}
		// Each strategy must matter to the opponent, or an expected payoff is just a
		// constant read off the table.
		if g.Row[0][0] == g.Row[0][1] || g.Row[1][0] == g.Row[1][1] ||
			g.Col[0][0] == g.Col[1][0] || g.Col[0][1] == g.Col[1][1] {
			continue
		}
		mx, ok := g.SolveMixed()
		if !ok || len(g.PureNash()) != 2 {
			continue
		}
		// Keep the fractions friendly: denominators a calculator-free student can handle.
		if !friendly(mx.P) || !friendly(mx.Q) || !friendly(mx.RowPay) || !friendly(mx.ColPay) {
			continue
		}
		// Strict equilibria only, so "which cells are equilibria" has one answer.
		strict := true
		for _, ne := range g.PureNash() {
			r, c := ne[0], ne[1]
			if g.Row[r][c] == g.Row[1-r][c] || g.Col[r][c] == g.Col[r][1-c] {
				strict = false
			}
		}
		if strict {
			return g, mx
		}
	}
}

func friendly(v float64) bool {
	for _, d := range []float64{1, 2, 3, 4, 5, 6, 8, 10} {
		if math.Abs(v*d-math.Round(v*d)) < 1e-9 {
			return true
		}
	}
	return false
}

func gameTable(g generate.Game) string {
	cell := func(r, c int) string { return fmtN(g.Row[r][c]) + ", " + fmtN(g.Col[r][c]) }
	return "| | **L** (q) | **R** (1 − q) |\n|---|---|---|\n" +
		"| **U** (p) | " + cell(0, 0) + " | " + cell(0, 1) + " |\n" +
		"| **D** (1 − p) | " + cell(1, 0) + " | " + cell(1, 1) + " |\n\n" +
		"Payoffs are (Player 1, Player 2). Player 1 picks the row, Player 2 the column."
}

var cellName = [2][2]string{{"(U, L)", "(U, R)"}, {"(D, L)", "(D, R)"}}

func generatedMixed(count int, rnd *rand.Rand) ([]content.Question, error) {
	var out []content.Question
	const T = "mixed-strategy"
	for len(out) < count {
		g, mx := randomGame(rnd)
		tbl := gameTable(g) + "\n\n"
		r, c := g.Row, g.Col
		pExpl := fmt.Sprintf("p is set to make **Player 2** indifferent, so use Player 2's payoffs. Payoff(L) = %s·p + %s·(1 − p); Payoff(R) = %s·p + %s·(1 − p). Setting them equal gives p = %s.",
			fmtN(c[0][0]), fmtN(c[1][0]), fmtN(c[0][1]), fmtN(c[1][1]), fmtN(mx.P))
		qExpl := fmt.Sprintf("q is set to make **Player 1** indifferent, so use Player 1's payoffs. Payoff(U) = %s·q + %s·(1 − q); Payoff(D) = %s·q + %s·(1 − q). Setting them equal gives q = %s.",
			fmtN(r[0][0]), fmtN(r[0][1]), fmtN(r[1][0]), fmtN(r[1][1]), fmtN(mx.Q))
		var q content.Question
		switch rnd.Intn(5) {
		case 0:
			q = numQ(T, tbl+"In the mixed-strategy Nash equilibrium, with what probability **p** does Player 1 play **U**?", mx.P, "", pExpl+" The trap is using Player 1's own payoffs.", 45, "mixed")
		case 1:
			q = numQ(T, tbl+"In the mixed-strategy Nash equilibrium, with what probability **q** does Player 2 play **L**?", mx.Q, "", qExpl, 46, "mixed")
		case 2:
			q = numQ(T, tbl+"What is **Player 1's expected payoff** in the mixed-strategy equilibrium?", mx.RowPay, "",
				qExpl+fmt.Sprintf(" Then Payoff(U) = %s·%s + %s·%s = %s (and Payoff(D) is the same: that is what indifference means).", fmtN(r[0][0]), fmtN(mx.Q), fmtN(r[0][1]), fmtN(1-mx.Q), fmtN(mx.RowPay)), 47, "mixed")
		case 3:
			q = numQ(T, tbl+"What is **Player 2's expected payoff** in the mixed-strategy equilibrium?", mx.ColPay, "",
				pExpl+fmt.Sprintf(" Then Payoff(L) = %s·%s + %s·%s = %s.", fmtN(c[0][0]), fmtN(mx.P), fmtN(c[1][0]), fmtN(1-mx.P), fmtN(mx.ColPay)), 47, "mixed")
		default:
			ne := g.PureNash()
			right := cellName[ne[0][0]][ne[0][1]] + " and " + cellName[ne[1][0]][ne[1][1]]
			var other [][2]int
			for rr := 0; rr < 2; rr++ {
				for cc := 0; cc < 2; cc++ {
					if [2]int{rr, cc} != ne[0] && [2]int{rr, cc} != ne[1] {
						other = append(other, [2]int{rr, cc})
					}
				}
			}
			wrong := []string{
				cellName[other[0][0]][other[0][1]] + " and " + cellName[other[1][0]][other[1][1]],
				cellName[ne[0][0]][ne[0][1]] + " only",
				"None: the game has only a mixed equilibrium",
			}
			var ok bool
			q, ok = mcqQ(T, tbl+"Which cells are **pure-strategy** Nash equilibria?", right, wrong,
				fmt.Sprintf("Underline each player's best reply. In %s and %s neither player gains by switching alone. Because there are two pure equilibria, players can also randomise, which is the mixed equilibrium p = %s, q = %s.", cellName[ne[0][0]][ne[0][1]], cellName[ne[1][0]][ne[1][1]], fmtN(mx.P), fmtN(mx.Q)),
				44, rnd, "mixed")
			if !ok {
				continue
			}
		}
		q.ID = fmt.Sprintf("gen:mixed:%d:%d", len(out), rnd.Int63())
		out = append(out, q)
	}
	return out, nil
}

// --- Solow -----------------------------------------------------------------------

var (
	breakEvenRates = []float64{0.04, 0.05, 0.06, 0.08, 0.1}
	goldenDeltas   = []float64{0.025, 0.05, 0.1, 0.125}
)

func pct(v float64) string { return fmtN(math.Round(v*1000)/10) + "%" }

func generatedSolow(count int, rnd *rand.Rand) ([]content.Question, error) {
	var out []content.Question
	for len(out) < count {
		q, ok := solowQuestion(rnd)
		if !ok {
			continue
		}
		q.ID = fmt.Sprintf("gen:solow:%d:%d", len(out), rnd.Int63())
		out = append(out, q)
	}
	return out, nil
}

func solowQuestion(rnd *rand.Rand) (content.Question, bool) {
	const T = "solow"
	switch rnd.Intn(11) {
	case 0, 1: // returns to scale
		exps := []float64{0.2, 0.25, 0.3, 0.4, 0.5, 0.6, 0.7, 0.75}
		a := exps[rnd.Intn(len(exps))]
		b := 1 - a
		switch rnd.Intn(3) {
		case 1:
			b += 0.1 * float64(1+rnd.Intn(2))
		case 2:
			b -= 0.1 * float64(1+rnd.Intn(2))
		}
		b = math.Round(b*100) / 100
		sum := math.Round((a+b)*100) / 100
		right := "Constant returns to scale"
		if sum > 1 {
			right = "Increasing returns to scale"
		} else if sum < 1 {
			right = "Decreasing returns to scale"
		}
		var wrong []string
		for _, w := range []string{"Constant returns to scale", "Increasing returns to scale", "Decreasing returns to scale", "It depends on the saving rate"} {
			if w != right {
				wrong = append(wrong, w)
			}
		}
		return mcqQ(T, fmt.Sprintf("Y = F(K, L) = K^%s · L^%s. What returns to scale does this production function have?", fmtN(a), fmtN(b)),
			right, wrong,
			fmt.Sprintf("Scale both inputs by z: (zK)^%s(zL)^%s = z^(%s+%s)·K^%sL^%s = z^%s·Y. The exponents sum to %s, so %s (CRS needs exactly z¹·Y).", fmtN(a), fmtN(b), fmtN(a), fmtN(b), fmtN(a), fmtN(b), fmtN(sum), fmtN(sum), strings.ToLower(right)),
			59, rnd, "crs")
	case 2: // per-worker function
		exps := []float64{0.2, 0.25, 0.3, 0.35, 0.4, 0.5, 0.6}
		a := exps[rnd.Intn(len(exps))]
		return numQ(T, fmt.Sprintf("Y = K^%s · L^%s. The per-worker production function is y = k^x. What is x?", fmtN(a), fmtN(1-a)), a, "",
			fmt.Sprintf("Divide by L: Y/L = K^%s · L^%s / L¹ = K^%s / L^%s = (K/L)^%s, so y = k^%s. The exponent on capital survives; the labour exponent disappears because they sum to 1.", fmtN(a), fmtN(1-a), fmtN(a), fmtN(a), fmtN(a), fmtN(a)),
			71, "per-worker"), true
	case 3, 4, 5: // steady state
		x := breakEvenRates[rnd.Intn(len(breakEvenRates))]
		r := float64(2 + rnd.Intn(4)) // k* = r²
		s := generate.Solow{Alpha: 0.5, S: r * x}
		n := []float64{0, 0.01, 0.02}[rnd.Intn(3)]
		s.N = n
		s.Delta = math.Round((x-n)*1000) / 1000
		k := s.SteadyK()
		y := s.Y(k)
		setup := fmt.Sprintf("An economy has y = √k (that is, k^0.5), saving rate s = %s, depreciation δ = %s and population growth n = %s, with no technological progress.\n\n", pct(s.S), pct(s.Delta), pct(s.N))
		base := fmt.Sprintf("Steady state: s·f(k) = (δ + n)k ⇒ %s·√k = %s·k ⇒ √k = %s/%s = %s ⇒ k* = %s.", fmtN(s.S), fmtN(x), fmtN(s.S), fmtN(x), fmtN(r), fmtN(k))
		switch rnd.Intn(4) {
		case 0:
			return numQ(T, setup+"What is the steady-state capital per worker **k\\***?", k, "", base, 63, "steady-state"), true
		case 1:
			return numQ(T, setup+"What is steady-state output per worker **y\\***?", y, "", base+fmt.Sprintf(" Then y* = √%s = %s.", fmtN(k), fmtN(y)), 63, "steady-state"), true
		case 2:
			return numQ(T, setup+"What is steady-state **consumption** per worker **c\\***?", (1-s.S)*y, "",
				base+fmt.Sprintf(" y* = %s, and c* = (1 − s)y* = %s × %s = %s.", fmtN(y), fmtN(1-s.S), fmtN(y), fmtN((1-s.S)*y)), 61, "steady-state"), true
		default:
			return numQ(T, setup+"What is steady-state **investment** per worker **i\\***?", s.S*y, "",
				base+fmt.Sprintf(" i* = s·y* = %s × %s = %s, which equals break-even investment (δ + n)k* = %s × %s, as it must.", fmtN(s.S), fmtN(y), fmtN(s.S*y), fmtN(x), fmtN(k)), 80, "steady-state"), true
		}
	case 6: // Δk off the steady state
		x := breakEvenRates[rnd.Intn(len(breakEvenRates))]
		r := float64(2 + rnd.Intn(4))
		s := generate.Solow{Alpha: 0.5, S: r * x, Delta: x}
		kstar := s.SteadyK()
		ks := []float64{kstar / 4, kstar * 4, kstar / 4 * 9 / 4}
		k := ks[rnd.Intn(len(ks))]
		dk := s.DeltaK(k)
		return numQ(T, fmt.Sprintf("y = √k, s = %s, δ = %s, no population growth. Capital per worker is currently k = %s. What is **Δk** this period?", pct(s.S), pct(s.Delta), fmtN(k)), dk, "",
			fmt.Sprintf("Δk = s·f(k) − δk = %s × √%s − %s × %s = %s − %s = %s. It is %s because k is %s the steady state k* = %s.", fmtN(s.S), fmtN(k), fmtN(s.Delta), fmtN(k), fmtN(s.S*s.Y(k)), fmtN(s.Delta*k), fmtN(dk), map[bool]string{true: "positive", false: "negative"}[dk > 0], map[bool]string{true: "below", false: "above"}[dk > 0], fmtN(kstar)),
			63, "steady-state"), true
	case 7, 8: // Golden Rule
		d := goldenDeltas[rnd.Intn(len(goldenDeltas))]
		s := generate.Solow{Alpha: 0.5, Delta: d}
		k := s.GoldenK()
		y := s.Y(k)
		setup := fmt.Sprintf("y = √k, depreciation δ = %s, no population growth or technological progress.\n\n", pct(d))
		base := fmt.Sprintf("Golden Rule: MPK = δ. For y = k^0.5, MPK = 0.5·k^(−0.5) = 1/(2√k). Set 1/(2√k) = %s ⇒ √k = %s ⇒ k_gold = %s.", fmtN(d), fmtN(math.Sqrt(k)), fmtN(k))
		switch rnd.Intn(3) {
		case 0:
			return numQ(T, setup+"What is the **Golden Rule** capital per worker?", k, "", base, 75, "golden-rule"), true
		case 1:
			return numQ(T, setup+"What saving rate reaches the Golden Rule steady state?", 0.5, "",
				base+fmt.Sprintf(" In steady state s·√k = δk ⇒ s = δ√k = %s × %s = 0.5. For any y = k^α the answer is s_gold = α.", fmtN(d), fmtN(math.Sqrt(k))), 75, "golden-rule"), true
		default:
			return numQ(T, setup+"What is consumption per worker at the Golden Rule?", y-d*k, "",
				base+fmt.Sprintf(" c = f(k) − δk = %s − %s × %s = %s, the largest steady-state consumption available.", fmtN(y), fmtN(d), fmtN(k), fmtN(y-d*k)), 73, "golden-rule"), true
		}
	case 9: // computers
		return computersQuestion(rnd, 2), true
	default: // break-even investment
		k := float64([]int{4, 9, 16, 25, 36}[rnd.Intn(5)])
		d := []float64{0.04, 0.05, 0.06}[rnd.Intn(3)]
		n := []float64{0.01, 0.02}[rnd.Intn(2)]
		g := []float64{0.01, 0.02, 0.03}[rnd.Intn(3)]
		be := (d + n + g) * k
		return numQ("golden-rule", fmt.Sprintf("Capital per effective worker is k = %s, with δ = %s, n = %s and g = %s. What is **break-even investment** per effective worker?", fmtN(k), pct(d), pct(n), pct(g)),
			be, "",
			fmt.Sprintf("(δ + n + g)k = (%s + %s + %s) × %s = %s × %s = %s. δk replaces worn-out capital, nk equips new workers and gk equips the new effective workers created by technology.", fmtN(d), fmtN(n), fmtN(g), fmtN(k), fmtN(d+n+g), fmtN(k), fmtN(be)),
			86, "break-even"), true
	}
}

// generatedComputers is the company-level Solow drill: computers per effective worker
// at one firm, with fresh rates each time.
func generatedComputers(count int, rnd *rand.Rand) ([]content.Question, error) {
	out := make([]content.Question, 0, count)
	for len(out) < count {
		q := computersQuestion(rnd, rnd.Intn(4))
		q.ID = fmt.Sprintf("gen:computers:%d:%d", len(out), rnd.Int63())
		out = append(out, q)
	}
	return out, nil
}

// computersQuestion asks one step of the packet's computers example (slide 88):
// 0 computers left after depreciation, 1 effective workers, 2 computers to add,
// 3 computers needed in total.
func computersQuestion(rnd *rand.Rand, which int) content.Question {
	ns := []float64{10, 20, 50, 100}
	n := ns[rnd.Intn(len(ns))]
	d := []float64{0.1, 0.2}[rnd.Intn(2)]
	p := []float64{0.1, 0.2}[rnd.Intn(2)]
	g := []float64{0.5, 1.0}[rnd.Intn(2)]
	add, eff := generate.ComputersToAdd(n, d, p, g)
	left := n * (1 - d)
	setup := fmt.Sprintf("A company has %s computers and %s workers (one computer per worker, k = 1). Over one period computers depreciate at δ = %s, the workforce grows at n = %s and technology makes each existing worker g = %s more effective; new workers are not yet trained in the technology.\n\n", fmtN(n), fmtN(n), fmtN(d), fmtN(p), fmtN(g))
	chain := fmt.Sprintf("Depreciation leaves %s × (1 − %s) = %s computers. The original %s workers become %s × (1 + %s) = %s effective workers, plus %s new workers who count once each: %s effective workers. Keeping k = 1 needs %s computers, so add %s − %s = %s = (δ + n + g) × %s = %s × %s. The rates add, they do not multiply.",
		fmtN(n), fmtN(d), fmtN(left), fmtN(n), fmtN(n), fmtN(g), fmtN(n*(1+g)), fmtN(n*p), fmtN(eff), fmtN(eff), fmtN(eff), fmtN(left), fmtN(add), fmtN(n), fmtN(d+p+g), fmtN(n))
	switch which {
	case 0:
		return numQ("golden-rule", setup+"How many computers are **left after depreciation**?", left, "computers", chain, 88, "break-even", "computers")
	case 1:
		return numQ("golden-rule", setup+"How many **effective workers** are there at the end of the period?", eff, "workers", chain, 88, "break-even", "computers")
	case 3:
		return numQ("golden-rule", setup+"How many computers does the company need **in total** at the end of the period to keep one per effective worker?", eff, "computers", chain, 88, "break-even", "computers")
	default:
		return numQ("golden-rule", setup+"How many computers must be **added** to keep one computer per effective worker?", add, "computers", chain, 88, "break-even", "computers")
	}
}
