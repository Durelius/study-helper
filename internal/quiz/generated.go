package quiz

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"

	"github.com/wilhelmdurelius/chulastudy/internal/content"
	"github.com/wilhelmdurelius/chulastudy/internal/generate"
)

// choose assembles a multiple choice from one right answer and some wrong ones, drops
// any distractor that collides with the answer, shuffles, and reports where the answer
// landed. Nothing else in this file is allowed to decide an answer index by hand.
func choose(correct string, distractors []string, rnd *rand.Rand) ([]string, int, bool) {
	seen := map[string]bool{correct: true}
	options := []string{correct}
	for _, d := range distractors {
		if seen[d] {
			continue
		}
		seen[d] = true
		options = append(options, d)
		if len(options) == 4 {
			break
		}
	}
	if len(options) < 3 {
		return nil, 0, false // not enough genuinely different options to be a fair question
	}
	rnd.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
	for i, o := range options {
		if o == correct {
			return options, i, true
		}
	}
	return nil, 0, false
}

func days(n int) string {
	if n == 1 {
		return "1 day"
	}
	return fmt.Sprintf("%d days", n)
}

// generatedCPM builds a run of critical-path questions on fresh networks.
//
// Each network yields one question. Asking two questions about the same diagram would
// let the second be answered from the first, which is exactly the shortcut the lab
// exists to remove.
func generatedCPM(count int, rnd *rand.Rand) ([]content.Question, error) {
	var out []content.Question
	for len(out) < count {
		p := generate.CPM(rnd)
		if p == nil {
			return nil, fmt.Errorf("could not generate a network")
		}
		q, ok := cpmQuestion(p, len(out), rnd)
		if !ok {
			continue
		}
		out = append(out, q)
	}
	return out, nil
}

func cpmQuestion(p *generate.CPMProblem, n int, rnd *rand.Rand) (content.Question, bool) {
	net := p.Network
	data, err := json.Marshal(net)
	if err != nil {
		return content.Question{}, false
	}
	base := content.Question{
		ID:         fmt.Sprintf("gen:cpm:%d:%d", n, rnd.Int63()),
		Topic:      "schedule-cost",
		Type:       "mcq",
		Difficulty: 3,
		Tags:       []string{"cpm", "generated"},
		ExamFocus:  true,
		Source:     content.Source{Deck: "new 06", Slide: 29},
		Data:       data,
	}

	byName := map[string]generate.Activity{}
	var slack []generate.Activity
	for _, a := range net.Activities {
		byName[a.Name] = a
		if a.Float > 0 {
			slack = append(slack, a)
		}
	}
	sort.Slice(slack, func(i, j int) bool { return slack[i].Name < slack[j].Name })

	switch rnd.Intn(4) {
	case 0: // which path is critical
		format := func(path []string) string { return strings.Join(path, " → ") }
		var wrong []string
		for _, path := range p.Paths[1:] {
			wrong = append(wrong, format(path))
		}
		choices, answer, ok := choose(format(p.Paths[0]), wrong, rnd)
		if !ok {
			return content.Question{}, false
		}
		base.Stem = "Run the forward and backward pass over this network. Which sequence is the critical path?"
		base.Choices = choices
		base.Answer = []int{answer}
		base.Explanation = fmt.Sprintf(
			"The critical path is the longest path through the network: %s, totalling %s. Every activity on it has float 0, so any delay to one of them delays the whole project.",
			format(p.Paths[0]), days(net.Duration))
		return base, true

	case 1: // project duration
		d := net.Duration
		choices, answer, ok := choose(days(d), []string{days(d - 1), days(d + 2), days(d + 1), days(d - 3)}, rnd)
		if !ok {
			return content.Question{}, false
		}
		base.Stem = "What is the shortest time in which this project can finish?"
		base.Choices = choices
		base.Answer = []int{answer}
		base.Explanation = fmt.Sprintf(
			"The project cannot finish before its longest path is done. That path is %s = %s. Adding up every activity instead gives a much larger number and is the usual mistake — parallel work does not add.",
			strings.Join(p.Paths[0], " → "), days(d))
		return base, true

	case 2: // float of one activity
		if len(slack) == 0 {
			return content.Question{}, false
		}
		a := slack[rnd.Intn(len(slack))]
		choices, answer, ok := choose(days(a.Float), []string{days(0), days(a.Float + 1), days(a.Float + 2), days(a.EF)}, rnd)
		if !ok {
			return content.Question{}, false
		}
		base.Stem = fmt.Sprintf("What is the total float of activity %s?", a.Name)
		base.Choices = choices
		base.Answer = []int{answer}
		base.Explanation = fmt.Sprintf(
			"%s has ES %d and LS %d, so total float = LS − ES = %d − %d = %s. It can slip that much before the project end moves.",
			a.Name, a.ES, a.LS, a.LS, a.ES, days(a.Float))
		return base, true

	default: // knock-on effect of a delay
		if len(slack) == 0 {
			return content.Question{}, false
		}
		a := slack[rnd.Intn(len(slack))]
		delay := a.Float + 1 + rnd.Intn(3)
		slip := delay - a.Float
		choices, answer, ok := choose(days(slip),
			[]string{days(delay), days(0), days(a.Float), days(delay + a.Float)}, rnd)
		if !ok {
			return content.Question{}, false
		}
		base.Stem = fmt.Sprintf("Activity %s is delayed by %s. By how long does the project finish late?", a.Name, days(delay))
		base.Choices = choices
		base.Answer = []int{answer}
		base.Explanation = fmt.Sprintf(
			"%s has %s of float, so the first %s are absorbed. A %s delay overruns that float by %s, and only the overrun pushes the project end out.",
			a.Name, days(a.Float), days(a.Float), days(delay), days(slip))
		return base, true
	}
}

// generatedGlossary turns the glossary into definition-matching questions. It is the
// cheapest retrieval practice in the app and the closest thing to the paper's
// true/false and ABCD bulk.
func generatedGlossary(set *content.Set, count int, rnd *rand.Rand) ([]content.Question, error) {
	if len(set.Glossary) < 4 {
		return nil, fmt.Errorf("the glossary is too short to quiz from yet")
	}
	terms := append([]content.Term(nil), set.Glossary...)
	sort.Slice(terms, func(i, j int) bool { return terms[i].Term < terms[j].Term })
	rnd.Shuffle(len(terms), func(i, j int) { terms[i], terms[j] = terms[j], terms[i] })

	var out []content.Question
	for i := 0; i < len(terms) && len(out) < count; i++ {
		t := terms[i]
		// Distractors come from the same topic first: telling two terms from the same
		// lecture apart is the discrimination the exam actually tests.
		var near, far []string
		for _, other := range terms {
			if other.Term == t.Term {
				continue
			}
			if other.Topic == t.Topic {
				near = append(near, other.Term)
			} else {
				far = append(far, other.Term)
			}
		}
		rnd.Shuffle(len(near), func(a, b int) { near[a], near[b] = near[b], near[a] })
		rnd.Shuffle(len(far), func(a, b int) { far[a], far[b] = far[b], far[a] })
		choices, answer, ok := choose(t.Term, append(near, far...), rnd)
		if !ok {
			continue
		}
		out = append(out, content.Question{
			ID:          fmt.Sprintf("gen:glossary:%d:%s", i, t.Term),
			Topic:       t.Topic,
			Type:        "mcq",
			Stem:        fmt.Sprintf("Which term does this describe?\n\n> %s", t.Definition),
			Choices:     choices,
			Answer:      []int{answer},
			Explanation: fmt.Sprintf("%s — %s", t.Term, t.Definition),
			Source:      t.Source,
			Difficulty:  1,
			Tags:        []string{"glossary", "generated"},
		})
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("could not build glossary questions")
	}
	return out, nil
}
