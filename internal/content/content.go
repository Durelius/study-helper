// Package content loads the study material — lecture notes, question banks, cases and
// the glossary — and answers questions about it.
//
// The material is authored as files under content/ and read once at boot. In
// development it can be read from disk instead, so editing a question file only needs
// a server restart rather than a rebuild.
package content

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"

	valuechain "github.com/wilhelmdurelius/chula-valuechain"
)

// Source is the slide a fact came from, so a wrong answer can send the reader back to it.
type Source struct {
	Deck  string `json:"deck"`
	Slide int    `json:"slide"`
}

// Question is one item in the bank. Answer and Explanation never reach the browser
// before the reader has submitted — see the api package's wire types.
type Question struct {
	ID          string   `json:"id"`
	Topic       string   `json:"topic"`
	Type        string   `json:"type"` // mcq | tf | multi
	Stem        string   `json:"stem"`
	Choices     []string `json:"choices"`
	Answer      []int    `json:"answer"`
	Explanation string   `json:"explanation"`
	Source      Source   `json:"source"`
	Difficulty  int      `json:"difficulty"`
	Tags        []string `json:"tags"`
	ExamFocus   bool     `json:"examFocus"`
	// CaseID is set on questions that belong to a case study, so the runner knows to
	// show the scenario alongside them.
	CaseID string `json:"caseId,omitempty"`
	// Data carries a generated question's diagram — an activity network, a set of
	// earned-value figures — for the client to draw. Authored questions leave it nil.
	Data json.RawMessage `json:"data,omitempty"`
}

// Correct reports whether a set of chosen indices matches the key exactly. Order does
// not matter; a multi-answer question needs every correct choice and no extras.
func (q Question) Correct(given []int) bool {
	if len(given) != len(q.Answer) {
		return false
	}
	want := map[int]bool{}
	for _, i := range q.Answer {
		want[i] = true
	}
	for _, i := range given {
		if !want[i] {
			return false
		}
		delete(want, i)
	}
	return len(want) == 0
}

// Lecture is one topic's condensed notes.
type Lecture struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Deck    string `json:"deck"`
	Slides  int    `json:"slides"`
	Summary string `json:"summary"`
	// Body is markdown. Lines of the form "?check id=xx-001" mark inline retrieval
	// checks; the client turns them into a question card.
	Body string `json:"body"`
}

// Case is a long scenario with its own questions, mirroring the exam's case study.
type Case struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Topics    []string   `json:"topics"`
	Scenario  string     `json:"scenario"`
	Questions []Question `json:"questions"`
}

// Term is a glossary entry.
type Term struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Topic      string `json:"topic"`
	Source     Source `json:"source"`
}

// Topic is a lecture-sized chunk of the course, in the order it was taught.
type Topic struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Questions int    `json:"questions"`
	HasNotes  bool   `json:"hasNotes"`
}

// order is the canonical topic order and the authority on titles. A question file for
// an unknown topic is a mistake worth noticing, so loading rejects it.
var order = []Topic{
	{ID: "value-chain", Title: "Value Chain Fundamentals"},
	{ID: "operations", Title: "Operations Management"},
	{ID: "process-mapping", Title: "Process Mapping & BPMN"},
	{ID: "scm", Title: "Supply Chain Management"},
	{ID: "logistics", Title: "Supply Chain Strategy & Global Logistics"},
	{ID: "pm-foundations", Title: "PM Foundations & PMBOK"},
	{ID: "initiation", Title: "Project Initiation & Scope"},
	{ID: "schedule-cost", Title: "Schedule & Cost Planning"},
	{ID: "risk-comm-quality", Title: "Risk, Communication & Quality"},
}

// Set is the whole loaded library.
type Set struct {
	Topics    []Topic
	Lectures  map[string]Lecture
	Questions map[string]Question
	ByTopic   map[string][]string
	Cases     []Case
	Glossary  []Term
	// BPMNRules feeds the spot-the-error generator.
	BPMNRules []BPMNRule
}

// BPMNRule is one modelling rule from the process-mapping deck.
type BPMNRule struct {
	ID        string `json:"id"`
	Rule      string `json:"rule"`
	Violation string `json:"violation"`
	Source    Source `json:"source"`
}

// Load reads the library. An empty dir uses the embedded copy; otherwise dir is read
// from disk, which is what `-content ./content` is for during authoring.
func Load(dir string) (*Set, error) {
	var fsys fs.FS
	if dir == "" {
		sub, err := fs.Sub(valuechain.Files, "content")
		if err != nil {
			return nil, err
		}
		fsys = sub
	} else {
		fsys = os.DirFS(dir)
	}

	s := &Set{
		Lectures:  map[string]Lecture{},
		Questions: map[string]Question{},
		ByTopic:   map[string][]string{},
	}
	known := map[string]bool{}
	for _, t := range order {
		known[t.ID] = true
	}

	// Questions. A missing file is fine — a topic whose writer has not finished yet
	// simply has nothing to ask, and the rest of the app still works.
	entries, err := fs.ReadDir(fsys, "questions")
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		topic := strings.TrimSuffix(e.Name(), ".json")
		if !known[topic] {
			return nil, fmt.Errorf("questions/%s: %q is not a known topic", e.Name(), topic)
		}
		var qs []Question
		if err := readJSON(fsys, "questions/"+e.Name(), &qs); err != nil {
			return nil, err
		}
		for _, q := range qs {
			q.Topic = topic
			if err := validate(q); err != nil {
				return nil, fmt.Errorf("questions/%s: %w", e.Name(), err)
			}
			if _, dup := s.Questions[q.ID]; dup {
				return nil, fmt.Errorf("questions/%s: duplicate question id %q", e.Name(), q.ID)
			}
			s.Questions[q.ID] = q
			s.ByTopic[topic] = append(s.ByTopic[topic], q.ID)
		}
	}

	// Lectures.
	entries, err = fs.ReadDir(fsys, "lectures")
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		id := strings.TrimSuffix(e.Name(), ".md")
		if !known[id] {
			return nil, fmt.Errorf("lectures/%s: %q is not a known topic", e.Name(), id)
		}
		raw, err := fs.ReadFile(fsys, "lectures/"+e.Name())
		if err != nil {
			return nil, err
		}
		lec := parseLecture(string(raw))
		lec.ID = id
		s.Lectures[id] = lec
	}

	// Cases. Their questions join the main index so they can be graded like any other.
	entries, _ = fs.ReadDir(fsys, "cases")
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		var c Case
		if err := readJSON(fsys, "cases/"+e.Name(), &c); err != nil {
			return nil, err
		}
		for i, q := range c.Questions {
			q.CaseID = c.ID
			if q.Topic == "" && len(c.Topics) > 0 {
				q.Topic = c.Topics[0]
			}
			if err := validate(q); err != nil {
				return nil, fmt.Errorf("cases/%s: %w", e.Name(), err)
			}
			c.Questions[i] = q
			s.Questions[q.ID] = q
		}
		s.Cases = append(s.Cases, c)
	}

	// Visual questions live in their own directory because each one carries a whole
	// diagram and its topic varies file by file, unlike questions/<topic>.json.
	entries, _ = fs.ReadDir(fsys, "visuals")
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".json") {
			continue
		}
		var qs []Question
		if err := readJSON(fsys, "visuals/"+e.Name(), &qs); err != nil {
			return nil, err
		}
		for _, q := range qs {
			if !known[q.Topic] {
				return nil, fmt.Errorf("visuals/%s: %q is not a known topic", e.Name(), q.Topic)
			}
			if err := validate(q); err != nil {
				return nil, fmt.Errorf("visuals/%s: %w", e.Name(), err)
			}
			if _, dup := s.Questions[q.ID]; dup {
				return nil, fmt.Errorf("visuals/%s: duplicate question id %q", e.Name(), q.ID)
			}
			s.Questions[q.ID] = q
			s.ByTopic[q.Topic] = append(s.ByTopic[q.Topic], q.ID)
		}
	}

	if err := readJSON(fsys, "glossary.json", &s.Glossary); err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	if err := readJSON(fsys, "bpmn-rules.json", &s.BPMNRules); err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	for _, t := range order {
		ids := s.ByTopic[t.ID]
		sort.Strings(ids)
		s.ByTopic[t.ID] = ids
		_, hasNotes := s.Lectures[t.ID]
		s.Topics = append(s.Topics, Topic{ID: t.ID, Title: t.Title, Questions: len(ids), HasNotes: hasNotes})
	}
	return s, nil
}

// Title returns a topic's display name, falling back to the id for generated drills
// that have no lecture of their own.
func Title(id string) string {
	for _, t := range order {
		if t.ID == id {
			return t.Title
		}
	}
	return id
}

func readJSON(fsys fs.FS, name string, into any) error {
	raw, err := fs.ReadFile(fsys, name)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(raw, into); err != nil {
		return fmt.Errorf("%s: %w", name, err)
	}
	return nil
}

// validate catches the authoring mistakes that would otherwise show up as a question
// nobody can answer correctly: an out-of-range key, an empty key, a true/false item
// with four choices.
func validate(q Question) error {
	if q.ID == "" {
		return fmt.Errorf("question with no id")
	}
	if len(q.Choices) < 2 {
		return fmt.Errorf("%s: needs at least two choices", q.ID)
	}
	if len(q.Answer) == 0 {
		return fmt.Errorf("%s: no answer key", q.ID)
	}
	seen := map[int]bool{}
	for _, i := range q.Answer {
		if i < 0 || i >= len(q.Choices) {
			return fmt.Errorf("%s: answer index %d is outside its %d choices", q.ID, i, len(q.Choices))
		}
		if seen[i] {
			return fmt.Errorf("%s: answer index %d repeated", q.ID, i)
		}
		seen[i] = true
	}
	switch q.Type {
	case "tf":
		if len(q.Choices) != 2 || len(q.Answer) != 1 {
			return fmt.Errorf("%s: a true/false question needs two choices and one answer", q.ID)
		}
	case "mcq":
		if len(q.Answer) != 1 {
			return fmt.Errorf("%s: an mcq needs exactly one answer", q.ID)
		}
	case "multi":
		if len(q.Answer) < 2 {
			return fmt.Errorf("%s: a multi-answer question needs at least two answers", q.ID)
		}
	case "hotspot":
		// The choices are the diagram's clickable regions, so a hotspot question
		// without a diagram has nothing to click.
		if len(q.Data) == 0 {
			return fmt.Errorf("%s: a hotspot question needs a diagram in its data", q.ID)
		}
		if len(q.Answer) != 1 {
			return fmt.Errorf("%s: a hotspot question needs exactly one answer", q.ID)
		}
	default:
		return fmt.Errorf("%s: unknown type %q", q.ID, q.Type)
	}
	return nil
}

// parseLecture splits the leading --- front matter from the markdown body. The front
// matter is a handful of flat key: value pairs, so a YAML dependency is not worth it.
func parseLecture(raw string) Lecture {
	lec := Lecture{Body: raw}
	if !strings.HasPrefix(raw, "---") {
		return lec
	}
	rest := strings.TrimPrefix(raw, "---")
	end := strings.Index(rest, "\n---")
	if end < 0 {
		return lec
	}
	head, body := rest[:end], rest[end+4:]
	for _, line := range strings.Split(head, "\n") {
		key, value, ok := strings.Cut(line, ":")
		if !ok {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.Trim(strings.TrimSpace(value), `"`)
		switch key {
		case "title":
			lec.Title = value
		case "deck":
			lec.Deck = value
		case "summary":
			lec.Summary = value
		case "slides":
			lec.Slides, _ = strconv.Atoi(value)
		}
	}
	lec.Body = strings.TrimLeft(body, "\n")
	return lec
}
