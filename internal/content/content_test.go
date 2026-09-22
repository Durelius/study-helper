package content

import (
	"encoding/json"
	"testing"
)

// TestRealContentLoads is the guard that makes an authoring mistake a failed build
// rather than a wrong mark during someone's revision.
func TestRealContentLoads(t *testing.T) {
	set, err := Load("../../content")
	if err != nil {
		t.Fatalf("loading the real content: %v", err)
	}
	if len(set.Questions) < 300 {
		t.Errorf("only %d questions loaded; the bank should be larger", len(set.Questions))
	}
	for _, topic := range set.Topics {
		if topic.Questions == 0 && topic.HasNotes {
			t.Errorf("%s has notes but no questions", topic.ID)
		}
	}
}

// hotspotDiagram is the part of a visual question's diagram this test needs.
type hotspotDiagram struct {
	Kind     string `json:"kind"`
	Hotspots []struct {
		Target string `json:"target"`
		Label  string `json:"label"`
	} `json:"hotspots"`
	Nodes []struct {
		ID string `json:"id"`
	} `json:"nodes"`
	Flows []struct {
		From string `json:"from"`
		To   string `json:"to"`
	} `json:"flows"`
	Rows []struct {
		ID    string   `json:"id"`
		Cells []string `json:"cells"`
	} `json:"rows"`
}

// TestHotspotQuestionsAreAnswerable checks the thing that makes a visual question mean
// anything: clicking region N and choosing option N must be the same answer.
//
// An earlier question listed a choice naming a task that did not appear in its own
// diagram, which is unanswerable by clicking and misleading by reading.
func TestHotspotQuestionsAreAnswerable(t *testing.T) {
	set, err := Load("../../content")
	if err != nil {
		t.Fatalf("loading: %v", err)
	}
	seen := 0
	for _, q := range set.Questions {
		if q.Type != "hotspot" {
			continue
		}
		seen++
		var d hotspotDiagram
		if err := json.Unmarshal(q.Data, &d); err != nil {
			t.Errorf("%s: diagram does not parse: %v", q.ID, err)
			continue
		}
		if len(d.Hotspots) != len(q.Choices) {
			t.Errorf("%s: %d hotspots but %d choices", q.ID, len(d.Hotspots), len(q.Choices))
			continue
		}

		ids := map[string]bool{}
		for _, n := range d.Nodes {
			ids[n.ID] = true
		}
		for _, f := range d.Flows {
			ids[f.From+">"+f.To] = true
		}
		for _, r := range d.Rows {
			ids[r.ID] = true
			for c := range r.Cells {
				ids[r.ID+":"+itoa(c)] = true
			}
		}

		targets := map[string]bool{}
		for i, h := range d.Hotspots {
			if h.Label != q.Choices[i] {
				t.Errorf("%s: hotspot %d is %q but choice %d is %q — clicking and choosing would mean different things",
					q.ID, i+1, h.Label, i+1, q.Choices[i])
			}
			if targets[h.Target] {
				t.Errorf("%s: two hotspots point at %q, so one of them cannot be clicked", q.ID, h.Target)
			}
			targets[h.Target] = true
			if !ids[h.Target] {
				t.Errorf("%s: hotspot %d points at %q, which is not in the diagram", q.ID, i+1, h.Target)
			}
		}
	}
	if seen == 0 {
		t.Error("no hotspot questions found; the visual quiz would be empty")
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
