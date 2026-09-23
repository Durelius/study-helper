package api

import (
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"
	"time"

	"github.com/wilhelmdurelius/chulastudy/internal/content"
	"github.com/wilhelmdurelius/chulastudy/internal/store"
)

// heaviestTopic is the topic with the largest share of the exam, used when a course
// has no computed drill to fall back on.
func (s *Server) heaviestTopic() string {
	best, bestWeight := "", 0.0
	for id, w := range s.Set.Course().Weights() {
		if w > bestWeight {
			best, bestWeight = id, w
		}
	}
	return best
}

// step is one recommended action on the dashboard.
type step struct {
	Action string `json:"action"` // read | quiz
	Mode   string `json:"mode,omitempty"`
	Topic  string `json:"topic,omitempty"`
	Label  string `json:"label"`
	// Why is the reason this is next, which is the part that makes the suggestion
	// worth following rather than guessing.
	Why string `json:"why"`
}

// handlePlan answers the only question that matters the night before an exam: what
// should I do right now?
//
// The rules, in order of precedence:
//   - a topic never touched is worth more than one already at 80%
//   - a confidently wrong answer outranks a merely unseen question
//   - inside the last twelve hours, switch from learning to rehearsing
func (s *Server) handlePlan(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimSpace(r.URL.Query().Get("player"))
	now := time.Now()
	remaining := s.Exam.Sub(now)

	var stats []store.TopicStat
	var misses []store.Miss
	if player != "" {
		stats, _ = s.DB.TopicStats(player)
		misses, _ = s.DB.Misses(player, 40)
	}
	seen := map[string]store.TopicStat{}
	for _, t := range stats {
		seen[t.Topic] = t
	}

	confidentlyWrong := 0
	for _, m := range misses {
		confidentlyWrong += m.Confident
	}

	var steps []step

	// Anything never opened comes first — you cannot be tested on what you have not
	// met, and reading is faster than discovering the gap in a quiz.
	var untouched []content.Topic
	for _, t := range s.Set.Topics {
		if t.Questions == 0 && !t.HasNotes {
			continue
		}
		if seen[t.ID].Seen == 0 {
			untouched = append(untouched, t)
		}
	}
	if len(untouched) > 0 {
		t := untouched[0]
		if t.HasNotes {
			steps = append(steps, step{
				Action: "read", Topic: t.ID,
				Label: "Read " + t.Title,
				Why:   "You have not answered a single question on this topic yet. Read it once, then drill it.",
			})
		}
		if t.Questions > 0 {
			steps = append(steps, step{
				Action: "quiz", Mode: "topic", Topic: t.ID,
				Label: "Drill " + t.Title,
				Why:   "Testing yourself straight after reading is what makes it stick — rereading it a second time is not.",
			})
		}
	}

	// Then the weakest topic that has actually been sampled.
	type scored struct {
		topic    content.Topic
		accuracy float64
		seen     int
	}
	var ranked []scored
	for _, t := range s.Set.Topics {
		st := seen[t.ID]
		if st.Seen < 3 {
			continue
		}
		ranked = append(ranked, scored{t, float64(st.Correct) / float64(st.Seen), st.Seen})
	}
	sort.Slice(ranked, func(i, j int) bool { return ranked[i].accuracy < ranked[j].accuracy })
	if len(ranked) > 0 && ranked[0].accuracy < 0.8 {
		steps = append(steps, step{
			Action: "quiz", Mode: "topic", Topic: ranked[0].topic.ID,
			Label: fmt.Sprintf("Drill %s — you are at %.0f%%", ranked[0].topic.Title, ranked[0].accuracy*100),
			Why:   "Your weakest topic of the ones you have sampled. Points are cheapest here.",
		})
	}

	if confidentlyWrong >= 3 {
		steps = append(steps, step{
			Action: "quiz", Mode: "weak",
			Label: fmt.Sprintf("Weak spots — %d confident mistakes", confidentlyWrong),
			Why:   "You were sure and wrong on these. Those are the answers you will put down again in the exam unless you overwrite them now.",
		})
	}

	// A computed drill only exists where the course teaches it — suggesting the
	// critical-path lab to a course with no critical path was a real bug.
	if s.Set.Course().Has("cpm") {
		steps = append(steps, step{
			Action: "quiz", Mode: "cpm",
			Label: "Critical path lab",
			Why:   "Question 1 is a critical path problem and the networks here are generated fresh, so it is real practice rather than recall.",
		})
	} else if heaviest := s.heaviestTopic(); heaviest != "" {
		steps = append(steps, step{
			Action: "quiz", Mode: "topic", Topic: heaviest,
			Label: "Drill " + s.Set.Title(heaviest),
			Why:   "The topic carrying the most marks on this paper, so it repays practice right up to the door.",
		})
	}

	if remaining < 12*time.Hour && remaining > 0 {
		steps = append([]step{{
			Action: "quiz", Mode: "exam",
			Label: "Sit a full exam simulation",
			Why:   "Under twelve hours to go. Stop learning new material and rehearse the paper instead — timed, mixed, no notes.",
		}}, steps...)
	}

	// Two rules can land on the same suggestion — the weakest topic is often also the
	// heaviest one — and a list that says the same thing twice reads like a bug.
	suggested := map[string]bool{}
	unique := steps[:0]
	for _, st := range steps {
		key := st.Action + "|" + st.Mode + "|" + st.Topic
		if suggested[key] {
			continue
		}
		suggested[key] = true
		unique = append(unique, st)
	}
	steps = unique

	// Overall accuracy, counted across topics rather than sessions, so retrying the
	// same question does not flatter it.
	totalSeen, totalCorrect := 0, 0
	for _, t := range stats {
		totalSeen += t.Seen
		totalCorrect += t.Correct
	}
	accuracy := 0.0
	if totalSeen > 0 {
		accuracy = float64(totalCorrect) / float64(totalSeen)
	}

	// Coverage is how much of the bank has been seen once — the honest answer to
	// "how far through am I?".
	bank := 0
	for _, t := range s.Set.Topics {
		bank += t.Questions
	}
	coverage := 0.0
	if bank > 0 {
		coverage = math.Min(1, float64(totalSeen)/float64(bank))
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"exam":             s.Exam.Format(time.RFC3339),
		"secondsRemaining": int64(remaining.Seconds()),
		"steps":            steps,
		"seen":             totalSeen,
		"bank":             bank,
		"accuracy":         accuracy,
		"coverage":         coverage,
		"confidentlyWrong": confidentlyWrong,
	})
}
