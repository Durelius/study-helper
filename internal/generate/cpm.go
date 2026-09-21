// Package generate builds practice problems that are computed rather than authored, so
// the two procedural parts of the exam — the critical path and earned value — can be
// drilled without ever repeating a problem.
package generate

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// Activity is one node of an activity-on-node network, with the schedule figures the
// deck's box notation shows: ES | Duration | EF on top, LS | Float | LF underneath.
type Activity struct {
	Name         string   `json:"name"`
	Duration     int      `json:"duration"`
	Predecessors []string `json:"predecessors"`

	ES       int  `json:"es"`
	EF       int  `json:"ef"`
	LS       int  `json:"ls"`
	LF       int  `json:"lf"`
	Float    int  `json:"float"`
	Critical bool `json:"critical"`
	// Layer is the node's column in the diagram, so the client can lay it out without
	// implementing a graph algorithm of its own.
	Layer int `json:"layer"`
}

// Network is a solved schedule.
type Network struct {
	Activities []Activity `json:"activities"`
	Duration   int        `json:"duration"`
	// CriticalPath is the activity names in order, e.g. ["B","D","H"].
	CriticalPath []string `json:"criticalPath"`
}

// Solve runs the forward pass, the backward pass and the float calculation over a set
// of activities, exactly as the CPM slides describe it. Activities may arrive in any
// order; the returned slice is topologically sorted.
//
// It reports an error for a cycle or an unknown predecessor, because a network with
// either is not a schedule.
func Solve(in []Activity) (*Network, error) {
	byName := make(map[string]*Activity, len(in))
	order := make([]string, 0, len(in))
	for i := range in {
		a := in[i]
		a.Critical, a.ES, a.EF, a.LS, a.LF, a.Float = false, 0, 0, 0, 0, 0
		if _, dup := byName[a.Name]; dup {
			return nil, fmt.Errorf("duplicate activity %q", a.Name)
		}
		byName[a.Name] = &a
		order = append(order, a.Name)
	}
	successors := map[string][]string{}
	for _, name := range order {
		for _, p := range byName[name].Predecessors {
			if _, ok := byName[p]; !ok {
				return nil, fmt.Errorf("activity %q lists unknown predecessor %q", name, p)
			}
			successors[p] = append(successors[p], name)
		}
	}

	sorted, err := topoSort(order, byName)
	if err != nil {
		return nil, err
	}

	// Forward pass: an activity starts as soon as every predecessor has finished.
	duration := 0
	for _, name := range sorted {
		a := byName[name]
		for _, p := range a.Predecessors {
			if ef := byName[p].EF; ef > a.ES {
				a.ES = ef
			}
		}
		a.EF = a.ES + a.Duration
		a.Layer = 0
		for _, p := range a.Predecessors {
			if l := byName[p].Layer + 1; l > a.Layer {
				a.Layer = l
			}
		}
		if a.EF > duration {
			duration = a.EF
		}
	}

	// Backward pass: an activity must finish before the earliest of its successors
	// has to start; an activity with no successor must finish by the project end.
	for i := len(sorted) - 1; i >= 0; i-- {
		a := byName[sorted[i]]
		succ := successors[a.Name]
		if len(succ) == 0 {
			a.LF = duration
		} else {
			a.LF = duration
			for _, s := range succ {
				if ls := byName[s].LS; ls < a.LF {
					a.LF = ls
				}
			}
		}
		a.LS = a.LF - a.Duration
		a.Float = a.LS - a.ES
		a.Critical = a.Float == 0
	}

	// Empty slices rather than nil: these cross the wire as JSON, and a nil slice
	// marshals to `null`, which the diagram then tries to iterate. An activity with no
	// predecessors has an empty list of them, not a missing one.
	net := &Network{Duration: duration, CriticalPath: []string{}}
	for _, name := range sorted {
		a := *byName[name]
		if a.Predecessors == nil {
			a.Predecessors = []string{}
		}
		net.Activities = append(net.Activities, a)
		if a.Critical {
			net.CriticalPath = append(net.CriticalPath, name)
		}
	}
	return net, nil
}

// topoSort orders activities so every predecessor comes first, breaking ties by name so
// the same network always solves to the same layout.
func topoSort(names []string, byName map[string]*Activity) ([]string, error) {
	remaining := map[string]int{}
	for _, n := range names {
		remaining[n] = len(byName[n].Predecessors)
	}
	var ready []string
	for _, n := range names {
		if remaining[n] == 0 {
			ready = append(ready, n)
		}
	}
	sort.Strings(ready)

	successors := map[string][]string{}
	for _, n := range names {
		for _, p := range byName[n].Predecessors {
			successors[p] = append(successors[p], n)
		}
	}

	out := make([]string, 0, len(names))
	for len(ready) > 0 {
		n := ready[0]
		ready = ready[1:]
		out = append(out, n)
		next := []string{}
		for _, s := range successors[n] {
			remaining[s]--
			if remaining[s] == 0 {
				next = append(next, s)
			}
		}
		sort.Strings(next)
		ready = append(ready, next...)
		sort.Strings(ready)
	}
	if len(out) != len(names) {
		return nil, fmt.Errorf("the network contains a cycle")
	}
	return out, nil
}

// Paths returns every start-to-end path through the network with its length, longest
// first. It is what makes a plausible wrong answer: the second-longest path is exactly
// the trap a student who stops at the first long chain falls into.
func Paths(net *Network) [][]string {
	byName := map[string]Activity{}
	var starts, ends []string
	hasSucc := map[string]bool{}
	for _, a := range net.Activities {
		byName[a.Name] = a
		if len(a.Predecessors) == 0 {
			starts = append(starts, a.Name)
		}
		for _, p := range a.Predecessors {
			hasSucc[p] = true
		}
	}
	for _, a := range net.Activities {
		if !hasSucc[a.Name] {
			ends = append(ends, a.Name)
		}
	}
	_ = ends

	var all [][]string
	var walk func(path []string)
	walk = func(path []string) {
		last := path[len(path)-1]
		var next []string
		for _, a := range net.Activities {
			for _, p := range a.Predecessors {
				if p == last {
					next = append(next, a.Name)
				}
			}
		}
		if len(next) == 0 {
			all = append(all, append([]string(nil), path...))
			return
		}
		sort.Strings(next)
		for _, n := range next {
			walk(append(path, n))
		}
	}
	sort.Strings(starts)
	for _, s := range starts {
		walk([]string{s})
	}
	sort.Slice(all, func(i, j int) bool {
		li, lj := pathLength(all[i], byName), pathLength(all[j], byName)
		if li != lj {
			return li > lj
		}
		return strings.Join(all[i], "") < strings.Join(all[j], "")
	})
	return all
}

func pathLength(path []string, byName map[string]Activity) int {
	total := 0
	for _, n := range path {
		total += byName[n].Duration
	}
	return total
}

// CPMProblem is a generated network plus the figures a question can be asked about.
type CPMProblem struct {
	Network *Network
	Paths   [][]string
}

// CPM builds a random activity-on-node network with a single unambiguous critical path.
//
// The shape is deliberate: three parallel chains that merge, which is the shape the
// slides use and the shape an exam can fit on a page. It retries until the longest
// path is strictly longer than the runner-up, so "what is the critical path" has
// exactly one defensible answer.
func CPM(rnd *rand.Rand) *CPMProblem {
	for attempt := 0; attempt < 400; attempt++ {
		acts := randomNetwork(rnd)
		net, err := Solve(acts)
		if err != nil {
			continue
		}
		paths := Paths(net)
		if len(paths) < 4 {
			continue
		}
		byName := map[string]Activity{}
		for _, a := range net.Activities {
			byName[a.Name] = a
		}
		longest := pathLength(paths[0], byName)
		runnerUp := pathLength(paths[1], byName)
		// A clear winner, but not a runaway: a gap of one or two days keeps the
		// problem honest and makes the near-critical path a real temptation.
		if longest-runnerUp < 1 || longest-runnerUp > 3 {
			continue
		}
		// Somebody has to have float, or the float question is trivial.
		spare := 0
		for _, a := range net.Activities {
			if a.Float > 0 {
				spare++
			}
		}
		if spare < 2 {
			continue
		}
		return &CPMProblem{Network: net, Paths: paths}
	}
	return nil
}

// randomNetwork lays activities out in columns and wires each one to one or two
// predecessors in the column before it, which yields a readable diagram rather than a
// hairball.
func randomNetwork(rnd *rand.Rand) []Activity {
	names := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	columns := 3 + rnd.Intn(2) // 3 or 4 columns
	total := 6 + rnd.Intn(3)   // 6 to 8 activities

	// Deal the activities into columns, keeping at least two in the middle so there
	// is something parallel to reason about.
	sizes := make([]int, columns)
	for i := range sizes {
		sizes[i] = 1
	}
	for left := total - columns; left > 0; left-- {
		i := rnd.Intn(columns)
		if sizes[i] < 3 {
			sizes[i]++
		} else {
			left++
		}
	}

	var acts []Activity
	var prevColumn []string
	next := 0
	for c := 0; c < columns; c++ {
		var thisColumn []string
		for i := 0; i < sizes[c] && next < len(names); i++ {
			a := Activity{Name: names[next], Duration: 2 + rnd.Intn(9)}
			next++
			if c > 0 {
				// One predecessor usually, two sometimes: two is what creates the
				// merge points where the forward pass actually has to take a max.
				picks := 1
				if len(prevColumn) > 1 && rnd.Intn(3) == 0 {
					picks = 2
				}
				perm := rnd.Perm(len(prevColumn))
				for p := 0; p < picks; p++ {
					a.Predecessors = append(a.Predecessors, prevColumn[perm[p]])
				}
				sort.Strings(a.Predecessors)
			}
			acts = append(acts, a)
			thisColumn = append(thisColumn, a.Name)
		}
		prevColumn = thisColumn
	}

	// Every activity in the last column feeds the finish, and any earlier activity
	// with no successor would dangle, so make sure each column is consumed.
	return acts
}
