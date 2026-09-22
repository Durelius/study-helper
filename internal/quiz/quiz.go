// Package quiz decides which questions a run is made of and grades the answers.
//
// The selection rules are the study strategy made concrete: interleave topics rather
// than block them, put the questions someone got wrong ahead of the ones they have
// already nailed, and keep the exam simulation shaped like the real paper.
package quiz

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sort"
	"strings"

	"github.com/wilhelmdurelius/chula-valuechain/internal/content"
)

// Mode is one thing a person can sit down and do.
type Mode struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Blurb string `json:"blurb"`
	// Count is how many questions a run holds; 0 means "ask the caller".
	Count int `json:"count"`
	// TimeLimitSec drives the countdown in the runner. 0 is untimed.
	TimeLimitSec int `json:"timeLimitSec"`
	// NeedsTopic marks the modes where the picker has to offer a topic list.
	NeedsTopic bool `json:"needsTopic"`
	// Why explains what this mode is for, shown under its name. Studying badly is
	// mostly a matter of not knowing which drill answers which weakness.
	Why string `json:"why"`

	topics    []string
	tags      []string
	types     []string
	examFocus bool
	generator string
	// weak draws from the player's own wrong answers.
	weak bool
}

// Modes is the catalogue, in the order the picker shows them.
var Modes = []Mode{
	{
		ID: "quick", Title: "Quick 10", Count: 10,
		Blurb: "Ten questions, every topic mixed.",
		Why:   "Interleaving beats studying one lecture at a time — it forces you to work out which idea applies before you apply it.",
	},
	{
		ID: "exam", Title: "Exam simulation", Count: 30, TimeLimitSec: 45 * 60,
		Blurb: "Timed, and shaped like the real paper.",
		Why:   "Weighted toward the five named questions: critical path, risk and knowledge areas, BPM and quality, a case, and project planning.",
	},
	{
		ID: "weak", Title: "Weak spots", Count: 12, weak: true,
		Blurb: "The ones you got wrong, confident mistakes first.",
		Why:   "Being sure and wrong is the most expensive state to walk into an exam with, so those come back first.",
	},
	{
		ID: "topic", Title: "Topic drill", Count: 12, NeedsTopic: true,
		Blurb: "One lecture at a time.",
		Why:   "Use this straight after reading a topic, then switch back to a mixed mode.",
	},
	{
		ID: "cpm", Title: "Critical path lab", Count: 8, generator: "cpm",
		Blurb: "Fresh networks, every time.",
		Why:   "Question 1 of the exam. Generated, so you have to run the forward and backward pass rather than recognise an answer.",
	},
	{
		ID: "risk", Title: "Risk triage", Count: 15,
		topics: []string{"risk-comm-quality"}, tags: []string{"risk", "classification", "defect", "issue", "constraint", "assumption", "dependency"},
		Blurb: "Risk, issue, constraint — and which source.",
		Why:   "Question 2 of the exam hands you a situation and asks what kind of risk it is.",
	},
	{
		ID: "knowledge-areas", Title: "Knowledge areas", Count: 12,
		topics: []string{"pm-foundations"}, tags: []string{"knowledge-area", "pmbok", "process-group", "life-cycle-vs"},
		Blurb: "Which PMBOK area does this belong to?",
		Why:   "The other half of question 2. Ten areas, and the boundaries between them are where marks are lost.",
	},
	{
		ID: "visual", Title: "Spot the error", Count: 12, types: []string{"hotspot"},
		Blurb: "Click what is wrong with the diagram.",
		Why:   "The paper shows you a model and asks you to circle the mistake. This is the same thing: read the diagram, point at the error.",
	},
	{
		ID: "bpmn", Title: "BPMN fix", Count: 12,
		topics: []string{"process-mapping"}, examFocus: true,
		Blurb: "Spot the broken modelling rule.",
		Why:   "Question 3 shows you a process model with one error in it.",
	},
	{
		ID: "quality", Title: "Quality & KPIs", Count: 10,
		topics: []string{"risk-comm-quality"}, tags: []string{"quality", "kpi", "metric", "cost-of-quality"},
		Blurb: "Is that a real metric or a wish?",
		Why:   "The second half of question 3: a quality statement is only worth marks if it has a metric, a target and a check.",
	},
	{
		ID: "glossary", Title: "Glossary sprint", Count: 15, generator: "glossary",
		Blurb: "Term to definition, fast.",
		Why:   "Cheap recall practice for the true/false and ABCD bulk of the paper.",
	},
	{
		ID: "case", Title: "Case study", Count: 0, generator: "case",
		Blurb: "A scenario with questions hanging off it.",
		Why:   "Question 4. Read the case once, then answer without scrolling back — that is the exam condition.",
	},
}

// ModeByID finds a mode, reporting whether it exists.
func ModeByID(id string) (Mode, bool) {
	for _, m := range Modes {
		if m.ID == id {
			return m, true
		}
	}
	return Mode{}, false
}

// examWeights shapes the simulation like the real paper: the five named questions carry
// most of the marks, and the rest is spread over everything else.
var examWeights = map[string]float64{
	"schedule-cost":     0.24,
	"risk-comm-quality": 0.22,
	"pm-foundations":    0.16,
	"process-mapping":   0.12,
	"initiation":        0.10,
	"value-chain":       0.06,
	"operations":        0.05,
	"scm":               0.03,
	"logistics":         0.02,
}

// History is what the app knows about one player, used to order their practice.
type History struct {
	// Misses are question ids they have got wrong, most worth revisiting first.
	Misses []string
	// Mastered are ids they answered correctly while marking themselves sure.
	Mastered map[string]bool
}

// Build assembles a run. count overrides the mode's default when it is greater than zero.
func Build(set *content.Set, mode Mode, topic string, count int, hist History, rnd *rand.Rand) ([]content.Question, error) {
	if count <= 0 {
		count = mode.Count
	}

	switch mode.generator {
	case "cpm":
		return generatedCPM(count, rnd)
	case "glossary":
		return generatedGlossary(set, count, rnd)
	case "case":
		return caseQuestions(set, topic, rnd)
	}

	pool := candidates(set, mode, topic)
	if len(pool) == 0 {
		return nil, fmt.Errorf("no questions available for %s yet", mode.Title)
	}

	if mode.weak {
		return weakSpots(set, pool, hist, count, rnd), nil
	}
	if mode.ID == "exam" {
		return examPaper(set, count, hist, rnd), nil
	}

	picked := pickFresh(pool, hist, count, rnd)
	return interleave(picked, rnd), nil
}

// candidates narrows the bank to what a mode is allowed to ask.
func candidates(set *content.Set, mode Mode, topic string) []content.Question {
	var out []content.Question
	for _, q := range set.Questions {
		if q.CaseID != "" {
			continue // case questions only come up inside their case
		}
		if topic != "" && q.Topic != topic {
			continue
		}
		if len(mode.topics) > 0 && !contains(mode.topics, q.Topic) {
			continue
		}
		if mode.examFocus && !q.ExamFocus {
			continue
		}
		if len(mode.tags) > 0 && !hasAnyTag(q, mode.tags) {
			continue
		}
		if len(mode.types) > 0 && !contains(mode.types, q.Type) {
			continue
		}
		out = append(out, q)
	}
	// A map has no order, so sort before any sampling to keep runs reproducible
	// from a seed.
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// pickFresh prefers questions the player has not already mastered, and only falls back
// to mastered ones when there is nothing else left to ask.
func pickFresh(pool []content.Question, hist History, count int, rnd *rand.Rand) []content.Question {
	var fresh, seen []content.Question
	for _, q := range pool {
		if hist.Mastered[q.ID] {
			seen = append(seen, q)
		} else {
			fresh = append(fresh, q)
		}
	}
	shuffle(fresh, rnd)
	shuffle(seen, rnd)
	out := append(fresh, seen...)
	if len(out) > count {
		out = out[:count]
	}
	return out
}

// weakSpots rebuilds the misses in order, topping up with fresh questions from the same
// topics when someone has not made enough mistakes to fill a run.
func weakSpots(set *content.Set, pool []content.Question, hist History, count int, rnd *rand.Rand) []content.Question {
	var out []content.Question
	taken := map[string]bool{}
	for _, id := range hist.Misses {
		if q, ok := set.Questions[id]; ok && !taken[id] {
			out = append(out, q)
			taken[id] = true
		}
		if len(out) >= count {
			break
		}
	}
	if len(out) < count {
		// Bias the top-up toward the topics they are already struggling with.
		weak := map[string]bool{}
		for _, q := range out {
			weak[q.Topic] = true
		}
		var near, rest []content.Question
		for _, q := range pool {
			if taken[q.ID] || hist.Mastered[q.ID] {
				continue
			}
			if weak[q.Topic] {
				near = append(near, q)
			} else {
				rest = append(rest, q)
			}
		}
		shuffle(near, rnd)
		shuffle(rest, rnd)
		for _, q := range append(near, rest...) {
			if len(out) >= count {
				break
			}
			out = append(out, q)
		}
	}
	return out
}

// examPaper draws a weighted spread across topics and leans on the exam-focus flag, so
// a simulation feels like the real thing rather than a random sample.
func examPaper(set *content.Set, count int, hist History, rnd *rand.Rand) []content.Question {
	var out []content.Question
	taken := map[string]bool{}

	for topic, weight := range examWeights {
		want := int(float64(count)*weight + 0.5)
		if want == 0 {
			continue
		}
		var focus, other []content.Question
		for _, id := range set.ByTopic[topic] {
			q := set.Questions[id]
			if q.ExamFocus {
				focus = append(focus, q)
			} else {
				other = append(other, q)
			}
		}
		shuffle(focus, rnd)
		shuffle(other, rnd)
		for _, q := range append(focus, other...) {
			if want == 0 {
				break
			}
			if taken[q.ID] {
				continue
			}
			out = append(out, q)
			taken[q.ID] = true
			want--
		}
	}

	// Rounding and thin topics leave gaps; fill them from whatever is left.
	if len(out) < count {
		var rest []content.Question
		for _, q := range set.Questions {
			if !taken[q.ID] && q.CaseID == "" {
				rest = append(rest, q)
			}
		}
		sort.Slice(rest, func(i, j int) bool { return rest[i].ID < rest[j].ID })
		shuffle(rest, rnd)
		for _, q := range rest {
			if len(out) >= count {
				break
			}
			out = append(out, q)
			taken[q.ID] = true
		}
	}
	if len(out) > count {
		out = out[:count]
	}
	return interleave(out, rnd)
}

// interleave spreads topics apart so consecutive questions rarely come from the same
// lecture. Blocked practice feels easier and is measurably worse.
func interleave(qs []content.Question, rnd *rand.Rand) []content.Question {
	if len(qs) < 3 {
		return qs
	}
	byTopic := map[string][]content.Question{}
	var order []string
	for _, q := range qs {
		if _, ok := byTopic[q.Topic]; !ok {
			order = append(order, q.Topic)
		}
		byTopic[q.Topic] = append(byTopic[q.Topic], q)
	}
	sort.Strings(order)
	shuffleStrings(order, rnd)

	out := make([]content.Question, 0, len(qs))
	for len(out) < len(qs) {
		progressed := false
		for _, topic := range order {
			bucket := byTopic[topic]
			if len(bucket) == 0 {
				continue
			}
			out = append(out, bucket[0])
			byTopic[topic] = bucket[1:]
			progressed = true
		}
		if !progressed {
			break
		}
	}
	return out
}

func contains(list []string, want string) bool {
	for _, s := range list {
		if s == want {
			return true
		}
	}
	return false
}

// hasAnyTag matches a mode's tags against a question's by prefix, so a mode asking for
// "knowledge-area" picks up "knowledge-areas" and "knowledge-area-integration" alike.
// The banks are written by hand and their tag spellings vary; an exact match silently
// empties a drill, which is the one failure nobody would notice until the night before.
func hasAnyTag(q content.Question, tags []string) bool {
	for _, t := range q.Tags {
		for _, want := range tags {
			if strings.HasPrefix(strings.ToLower(t), strings.ToLower(want)) {
				return true
			}
		}
	}
	return false
}

func shuffle(qs []content.Question, rnd *rand.Rand) {
	rnd.Shuffle(len(qs), func(i, j int) { qs[i], qs[j] = qs[j], qs[i] })
}

func shuffleStrings(s []string, rnd *rand.Rand) {
	rnd.Shuffle(len(s), func(i, j int) { s[i], s[j] = s[j], s[i] })
}

// caseQuestions returns one case's questions in their authored order, with the scenario
// attached to each so the runner can keep it on screen.
func caseQuestions(set *content.Set, id string, rnd *rand.Rand) ([]content.Question, error) {
	if len(set.Cases) == 0 {
		return nil, fmt.Errorf("no case studies are written yet")
	}
	chosen := set.Cases[rnd.Intn(len(set.Cases))]
	for _, c := range set.Cases {
		if c.ID == id {
			chosen = c
		}
	}
	if len(chosen.Questions) == 0 {
		return nil, fmt.Errorf("case %q has no questions", chosen.ID)
	}
	out := make([]content.Question, 0, len(chosen.Questions))
	for _, q := range chosen.Questions {
		data, _ := json.Marshal(map[string]string{"scenario": chosen.Scenario, "title": chosen.Title})
		q.Data = data
		out = append(out, q)
	}
	return out, nil
}

// Available reports how many questions a mode can currently draw on, so the picker can
// grey out a drill whose material has not been written yet rather than offering it and
// then failing.
//
// Generated modes are effectively unlimited; they report a large number.
func Available(set *content.Set, mode Mode) int {
	switch mode.generator {
	case "cpm":
		return 999
	case "glossary":
		if len(set.Glossary) < 4 {
			return 0
		}
		return len(set.Glossary)
	case "case":
		n := 0
		for _, c := range set.Cases {
			n += len(c.Questions)
		}
		return n
	}
	if mode.weak {
		// Weak spots tops itself up from the whole bank, so it is available as soon
		// as anything is.
		return len(set.Questions)
	}
	return len(candidates(set, mode, ""))
}
