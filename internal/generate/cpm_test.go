package generate

import (
	"math/rand"
	"testing"
)

// fixture is a network worked out by hand, so a change in the solver shows up as a
// disagreement with arithmetic rather than with its own previous output.
//
//	A(3) ─┐                      forward: A 0→3   B 0→4   C 3→7
//	      ├─ C(4) ─┐                      D 4→6   E 7→13  F 6→7
//	B(4) ─┴─ D(2) ─┼─ E(6)       project duration 13
//	               └─ F(1)       critical A → C → E, float B=1 D=1 F=6
func fixture() []Activity {
	return []Activity{
		{Name: "A", Duration: 3},
		{Name: "B", Duration: 4},
		{Name: "C", Duration: 4, Predecessors: []string{"A"}},
		{Name: "D", Duration: 2, Predecessors: []string{"B"}},
		{Name: "E", Duration: 6, Predecessors: []string{"C", "D"}},
		{Name: "F", Duration: 1, Predecessors: []string{"D"}},
	}
}

func TestSolveMatchesHandCalculation(t *testing.T) {
	net, err := Solve(fixture())
	if err != nil {
		t.Fatalf("Solve: %v", err)
	}
	if net.Duration != 13 {
		t.Errorf("project duration = %d, want 13", net.Duration)
	}
	want := map[string]struct{ es, ef, ls, lf, float int }{
		"A": {0, 3, 0, 3, 0},
		"B": {0, 4, 1, 5, 1},
		"C": {3, 7, 3, 7, 0},
		"D": {4, 6, 5, 7, 1},
		"E": {7, 13, 7, 13, 0},
		"F": {6, 7, 12, 13, 6},
	}
	for _, a := range net.Activities {
		w, ok := want[a.Name]
		if !ok {
			t.Errorf("unexpected activity %q", a.Name)
			continue
		}
		if a.ES != w.es || a.EF != w.ef || a.LS != w.ls || a.LF != w.lf || a.Float != w.float {
			t.Errorf("%s = ES %d EF %d LS %d LF %d float %d, want ES %d EF %d LS %d LF %d float %d",
				a.Name, a.ES, a.EF, a.LS, a.LF, a.Float, w.es, w.ef, w.ls, w.lf, w.float)
		}
		if a.Critical != (w.float == 0) {
			t.Errorf("%s critical = %v, want %v", a.Name, a.Critical, w.float == 0)
		}
	}
	if got := net.CriticalPath; len(got) != 3 || got[0] != "A" || got[1] != "C" || got[2] != "E" {
		t.Errorf("critical path = %v, want [A C E]", got)
	}
}

func TestSolveIsOrderIndependent(t *testing.T) {
	shuffled := fixture()
	shuffled[0], shuffled[5] = shuffled[5], shuffled[0]
	shuffled[1], shuffled[3] = shuffled[3], shuffled[1]
	net, err := Solve(shuffled)
	if err != nil {
		t.Fatalf("Solve: %v", err)
	}
	if net.Duration != 13 {
		t.Errorf("duration = %d, want 13 regardless of input order", net.Duration)
	}
}

func TestSolveRejectsBrokenNetworks(t *testing.T) {
	if _, err := Solve([]Activity{
		{Name: "A", Duration: 1, Predecessors: []string{"B"}},
		{Name: "B", Duration: 1, Predecessors: []string{"A"}},
	}); err == nil {
		t.Error("a cycle should not solve")
	}
	if _, err := Solve([]Activity{
		{Name: "A", Duration: 1, Predecessors: []string{"Z"}},
	}); err == nil {
		t.Error("an unknown predecessor should not solve")
	}
	if _, err := Solve([]Activity{
		{Name: "A", Duration: 1},
		{Name: "A", Duration: 2},
	}); err == nil {
		t.Error("a duplicate activity name should not solve")
	}
}

func TestPathsAreLongestFirst(t *testing.T) {
	net, err := Solve(fixture())
	if err != nil {
		t.Fatalf("Solve: %v", err)
	}
	paths := Paths(net)
	if len(paths) != 3 {
		t.Fatalf("got %d paths, want 3 (A-C-E, B-D-E, B-D-F)", len(paths))
	}
	first := paths[0]
	if len(first) != 3 || first[0] != "A" || first[2] != "E" {
		t.Errorf("longest path = %v, want A C E", first)
	}
}

// TestCPMProducesUsableProblems checks the properties every generated question relies
// on: one unambiguous answer, and internally consistent schedule figures.
func TestCPMProducesUsableProblems(t *testing.T) {
	for seed := int64(0); seed < 200; seed++ {
		p := CPM(rand.New(rand.NewSource(seed)))
		if p == nil {
			t.Fatalf("seed %d: generator gave up", seed)
		}
		byName := map[string]Activity{}
		for _, a := range p.Network.Activities {
			byName[a.Name] = a
			if a.LS-a.ES != a.Float || a.LF-a.EF != a.Float {
				t.Fatalf("seed %d: %s float %d disagrees with LS-ES=%d, LF-EF=%d",
					seed, a.Name, a.Float, a.LS-a.ES, a.LF-a.EF)
			}
			if a.EF-a.ES != a.Duration {
				t.Fatalf("seed %d: %s EF-ES is not its duration", seed, a.Name)
			}
		}
		longest := pathLength(p.Paths[0], byName)
		if longest != p.Network.Duration {
			t.Fatalf("seed %d: longest path %d but project duration %d", seed, longest, p.Network.Duration)
		}
		if runnerUp := pathLength(p.Paths[1], byName); runnerUp >= longest {
			t.Fatalf("seed %d: critical path is not unique (%d vs %d)", seed, longest, runnerUp)
		}
		// Every activity on the longest path must be one of the zero-float ones, and
		// there must be no others: that is what makes "float = 0" a safe rule to teach.
		onPath := map[string]bool{}
		for _, n := range p.Paths[0] {
			onPath[n] = true
			if !byName[n].Critical {
				t.Fatalf("seed %d: %s is on the longest path but has float %d", seed, n, byName[n].Float)
			}
		}
		for _, a := range p.Network.Activities {
			if a.Critical && !onPath[a.Name] {
				t.Fatalf("seed %d: %s has float 0 but is not on the critical path", seed, a.Name)
			}
		}
	}
}

// TestSolveMatchesTheDecksOwnExample runs the network printed on slide 30 of the
// schedule deck, whose boxes the lecture leaves blank for students to fill in.
//
//	START → A(6), B(4), F(10);  A→C(3);  B→D(4), B→E(3);
//	C→H(2), D→H;  E→G(3), F→G;  G, H → FINISH
//
// If the solver and the slide ever disagree, the app is teaching the wrong method.
func TestSolveMatchesTheDecksOwnExample(t *testing.T) {
	net, err := Solve([]Activity{
		{Name: "A", Duration: 6},
		{Name: "B", Duration: 4},
		{Name: "F", Duration: 10},
		{Name: "C", Duration: 3, Predecessors: []string{"A"}},
		{Name: "D", Duration: 4, Predecessors: []string{"B"}},
		{Name: "E", Duration: 3, Predecessors: []string{"B"}},
		{Name: "G", Duration: 3, Predecessors: []string{"E", "F"}},
		{Name: "H", Duration: 2, Predecessors: []string{"C", "D"}},
	})
	if err != nil {
		t.Fatalf("Solve: %v", err)
	}
	if net.Duration != 13 {
		t.Errorf("project duration = %d, want 13", net.Duration)
	}
	want := map[string]struct{ es, ef, ls, lf, float int }{
		"A": {0, 6, 2, 8, 2},
		"B": {0, 4, 3, 7, 3},
		"C": {6, 9, 8, 11, 2},
		"D": {4, 8, 7, 11, 3},
		"E": {4, 7, 7, 10, 3},
		"F": {0, 10, 0, 10, 0},
		"G": {10, 13, 10, 13, 0},
		"H": {9, 11, 11, 13, 2},
	}
	for _, a := range net.Activities {
		w := want[a.Name]
		if a.ES != w.es || a.EF != w.ef || a.LS != w.ls || a.LF != w.lf || a.Float != w.float {
			t.Errorf("%s = ES %d EF %d LS %d LF %d float %d, want ES %d EF %d LS %d LF %d float %d",
				a.Name, a.ES, a.EF, a.LS, a.LF, a.Float, w.es, w.ef, w.ls, w.lf, w.float)
		}
	}
	if got := net.CriticalPath; len(got) != 2 || got[0] != "F" || got[1] != "G" {
		t.Errorf("critical path = %v, want [F G]", got)
	}
}
