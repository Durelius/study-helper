package content

import (
	"regexp"
	"testing"
	"unicode/utf8"
)

func matchQuestion() Question {
	return Question{
		ID: "m-1", Type: "match",
		Items:   []string{"Barolo", "Barbaresco", "Taurasi"},
		Choices: []string{"Campania", "Piemonte", "Puglia", "Sicilia", "Veneto"},
		Answer:  []int{1, 1, 0},
	}
}

// TestMatchQuestionsValidate: a shared choice is allowed, a pool with fewer than two spares is not,
// because then the last pair falls out by elimination.
func TestMatchQuestionsValidate(t *testing.T) {
	if err := validate(matchQuestion()); err != nil {
		t.Errorf("a match with two items sharing a choice should validate: %v", err)
	}
	short := matchQuestion()
	short.Choices = short.Choices[:3]
	if err := validate(short); err == nil {
		t.Error("a pool with only one spare choice should not validate")
	}
	wrongLen := matchQuestion()
	wrongLen.Answer = []int{1, 0}
	if err := validate(wrongLen); err == nil {
		t.Error("a match needs one answer per item")
	}
	outOfRange := matchQuestion()
	outOfRange.Answer = []int{1, 1, 9}
	if err := validate(outOfRange); err == nil {
		t.Error("an answer outside the pool should not validate")
	}
}

// TestMatchIsGradedPairByPair: the order is the pairing, so the right choices against
// the wrong items is wrong.
func TestMatchIsGradedPairByPair(t *testing.T) {
	q := matchQuestion()
	if !q.Correct([]int{1, 1, 0}) {
		t.Error("the exact pairing should be correct")
	}
	if q.Correct([]int{1, 0, 1}) {
		t.Error("the same choices on the wrong items should be wrong")
	}
	if q.Correct([]int{1, 1, -1}) {
		t.Error("an unpaired row should be wrong")
	}
}

var (
	absolutes = regexp.MustCompile(`(?i)\b(always|never|only|all|must|none|every)\b`)
	hedges    = regexp.MustCompile(`(?i)\b(usually|often|may|generally|tends?)\b`)
)

// TestWineedHasNoFormatTells keeps the Wine Education bank honest about what its user
// asked for: no answer findable from its shape. A right answer that is usually the
// longest, a key parked at one position, a true/false paper that is mostly True, or
// "never" appearing mostly in false statements all let you score without knowing wine.
func TestWineedHasNoFormatTells(t *testing.T) {
	set, ok := loadEveryCourse(t)["wineed"]
	if !ok {
		t.Skip("no wineed course")
	}
	var mcq, longest, shortest int
	positions := make([]int, 4)
	var tfTrue, tfFalse, absTrue, absFalse, hedgeTrue, hedgeFalse int
	for _, q := range set.Questions {
		switch q.Type {
		case "mcq":
			mcq++
			right := q.Answer[0]
			positions[right]++
			n := utf8.RuneCountInString(q.Choices[right])
			isLongest, isShortest := true, true
			for i, c := range q.Choices {
				if i == right {
					continue
				}
				if utf8.RuneCountInString(c) >= n {
					isLongest = false
				}
				if utf8.RuneCountInString(c) <= n {
					isShortest = false
				}
			}
			if isLongest {
				longest++
			}
			if isShortest {
				shortest++
			}
		case "tf":
			abs, hedge := absolutes.MatchString(q.Stem), hedges.MatchString(q.Stem)
			if q.Answer[0] == 0 {
				tfTrue++
				absTrue += b2i(abs)
				hedgeTrue += b2i(hedge)
			} else {
				tfFalse++
				absFalse += b2i(abs)
				hedgeFalse += b2i(hedge)
			}
		}
	}
	pct := func(a, b int) float64 { return 100 * float64(a) / float64(b) }
	if mcq > 0 {
		// Both ways: a key that is almost never the longest choice is as much a tell as
		// one that usually is, since "never pick the longest" then wins marks.
		if p := pct(longest, mcq); p > 30 || p < 8 {
			t.Errorf("the right answer is the longest choice in %.0f%% of MCQs (keep it within 8-30%%)", p)
		}
		if p := pct(shortest, mcq); p > 30 || p < 8 {
			t.Errorf("the right answer is the shortest choice in %.0f%% of MCQs (keep it within 8-30%%)", p)
		}
		for i, n := range positions {
			if p := pct(n, mcq); p < 15 {
				t.Errorf("the answer is at position %d in only %.0f%% of MCQs (floor 15%%)", i+1, p)
			}
		}
	}
	if tf := tfTrue + tfFalse; tf > 0 {
		if p := pct(tfTrue, tf); p < 40 || p > 60 {
			t.Errorf("%.0f%% of true/false items are True; keep it between 40%% and 60%%", p)
		}
		if d := pct(absTrue, tfTrue) - pct(absFalse, tfFalse); d > 15 || d < -15 {
			t.Errorf("absolutes appear in %.0f%% of true and %.0f%% of false items; an absolute should not signal the answer",
				pct(absTrue, tfTrue), pct(absFalse, tfFalse))
		}
		if d := pct(hedgeTrue, tfTrue) - pct(hedgeFalse, tfFalse); d > 15 || d < -15 {
			t.Errorf("hedges appear in %.0f%% of true and %.0f%% of false items; a hedge should not signal the answer",
				pct(hedgeTrue, tfTrue), pct(hedgeFalse, tfFalse))
		}
	}
}

func b2i(b bool) int {
	if b {
		return 1
	}
	return 0
}
