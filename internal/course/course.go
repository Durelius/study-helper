// Package course describes one course the app can serve.
//
// Everything that differs between courses lives in a course.json beside that course's
// content: its topics and their weight in the exam simulation, the drills that only
// make sense for this syllabus, the countdown target, and what the app calls itself.
// The Go code stays the same for every course.
package course

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Topic is one lecture-sized chunk, in the order it is taught.
type Topic struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	// Weight is this topic's share of an exam simulation. Weights are normalised on
	// load, so they can be written as percentages, fractions or raw marks.
	Weight float64 `json:"weight"`
}

// Mode is a drill that only makes sense for one syllabus — "BPMN fix", "Risk triage".
// The generic modes (a quick ten, weak spots, the exam simulation) are built in.
type Mode struct {
	ID     string   `json:"id"`
	Title  string   `json:"title"`
	Count  int      `json:"count"`
	Blurb  string   `json:"blurb"`
	Why    string   `json:"why"`
	Topics []string `json:"topics"`
	Tags   []string `json:"tags"`
	Types  []string `json:"types"`
	// TimeLimitSec puts a clock on the mode. Used when a course redefines the built-in
	// exam simulation to match the length of its own paper.
	TimeLimitSec int `json:"timeLimitSec"`
	// ExamFocus restricts the mode to questions flagged as exam-critical.
	ExamFocus bool `json:"examFocus"`
	// Parts lays a run out as the real paper is laid out: each part draws its own
	// questions, in order, rather than one mix across the whole bank.
	Parts []Part `json:"parts"`
}

// Part is one section of a paper-shaped mode. Its filters combine like a mode's; Case
// draws one case study whole instead, for a part that is a set of questions about one
// thing in front of you (a label, a scenario).
type Part struct {
	Title  string   `json:"title"`
	Count  int      `json:"count"`
	Topics []string `json:"topics"`
	Tags   []string `json:"tags"`
	Types  []string `json:"types"`
	Case   bool     `json:"case"`
	// CaseID pins a case part to one named case, so a mock always draws the same
	// case study. Without it the part draws any case the player has not mastered.
	CaseID string `json:"caseId"`
	// Generator fills the part with computed questions ("oligopoly-steps", "mixed"…)
	// so a mock exam gets fresh numbers. It must be listed in the course's generators.
	Generator string `json:"generator"`
}

// Course is one course's configuration.
type Course struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Code      string `json:"code"`
	ShortName string `json:"shortName"`
	// Exam is the countdown target, in RFC 3339 with the course's own offset.
	Exam string `json:"exam"`
	// ExamName is what the countdown's assessment is called, such as "midterm" or
	// "final". It defaults to "midterm" when a course does not say.
	ExamName string  `json:"examName"`
	Topics   []Topic `json:"topics"`
	Modes    []Mode  `json:"modes"`
	// Generators names the computed drills this course offers, such as "cpm". A
	// course that does not teach critical path simply omits it and the lab disappears.
	Generators []string `json:"generators"`

	// Dir is where this course's content and audio live.
	Dir     string    `json:"-"`
	ExamAt  time.Time `json:"-"`
	weights map[string]float64
}

// Load reads a course from its directory.
func Load(dir string) (*Course, error) {
	raw, err := os.ReadFile(filepath.Join(dir, "course.json"))
	if err != nil {
		return nil, fmt.Errorf("reading the course config: %w", err)
	}
	var c Course
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, fmt.Errorf("course.json: %w", err)
	}
	c.Dir = dir
	if c.ID == "" {
		return nil, fmt.Errorf("course.json: needs an id")
	}
	if len(c.Topics) == 0 {
		return nil, fmt.Errorf("course.json: needs at least one topic")
	}
	seen := map[string]bool{}
	total := 0.0
	for _, t := range c.Topics {
		if t.ID == "" || t.Title == "" {
			return nil, fmt.Errorf("course.json: every topic needs an id and a title")
		}
		if seen[t.ID] {
			return nil, fmt.Errorf("course.json: duplicate topic %q", t.ID)
		}
		seen[t.ID] = true
		total += t.Weight
	}
	// Normalising means weights can be written however suits the course — percentages,
	// fractions, or the marks each topic carries on the real paper.
	c.weights = map[string]float64{}
	for _, t := range c.Topics {
		if total > 0 {
			c.weights[t.ID] = t.Weight / total
		} else {
			c.weights[t.ID] = 1 / float64(len(c.Topics))
		}
	}
	for _, m := range c.Modes {
		for _, topic := range m.Topics {
			if !seen[topic] {
				return nil, fmt.Errorf("course.json: mode %q refers to unknown topic %q", m.ID, topic)
			}
		}
		for _, p := range m.Parts {
			if p.Generator != "" && !c.Has(p.Generator) {
				return nil, fmt.Errorf("course.json: mode %q part %q uses generator %q, which is not in generators", m.ID, p.Title, p.Generator)
			}
			for _, topic := range p.Topics {
				if !seen[topic] {
					return nil, fmt.Errorf("course.json: mode %q part %q refers to unknown topic %q", m.ID, p.Title, topic)
				}
			}
		}
	}

	if c.Exam != "" {
		c.ExamAt, err = time.Parse(time.RFC3339, c.Exam)
		if err != nil {
			return nil, fmt.Errorf("course.json: cannot read exam time: %w", err)
		}
	}
	if c.ShortName == "" {
		c.ShortName = c.Title
	}
	if c.ExamName == "" {
		c.ExamName = "midterm"
	}
	return &c, nil
}

// Weights is each topic's normalised share of an exam simulation.
func (c *Course) Weights() map[string]float64 { return c.weights }

// Title returns a topic's display name, falling back to its id.
func (c *Course) TopicTitle(id string) string {
	for _, t := range c.Topics {
		if t.ID == id {
			return t.Title
		}
	}
	return id
}

// Has reports whether this course offers a computed drill.
func (c *Course) Has(generator string) bool {
	for _, g := range c.Generators {
		if g == generator {
			return true
		}
	}
	return false
}

// ContentDir, AudioDir and MaterialsDir are the conventional locations inside a course
// directory. MaterialsDir holds the source slides rendered by tools/materials.py.
func (c *Course) ContentDir() string   { return filepath.Join(c.Dir, "content") }
func (c *Course) AudioDir() string     { return filepath.Join(c.Dir, "audio") }
func (c *Course) MaterialsDir() string { return filepath.Join(c.Dir, "materials") }

// LabelsDir holds the label photos a case can show. Kept out of git like materials.
func (c *Course) LabelsDir() string { return filepath.Join(c.Dir, "labels") }
