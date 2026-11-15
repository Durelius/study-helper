package quiz

import (
	"math/rand"
	"testing"

	"github.com/wilhelmdurelius/chulastudy/internal/content"
	"github.com/wilhelmdurelius/chulastudy/internal/course"
)

// TestEveryGlobaleconModeBuilds starts each drill the economics course offers, the mock
// quiz with its generated parts included, so a mode that cannot fill itself fails here
// rather than greyed out or erroring on the day.
func TestEveryGlobaleconModeBuilds(t *testing.T) {
	c, err := course.Load("../../courses/globalecon")
	if err != nil {
		t.Fatal(err)
	}
	set, err := content.Load(c)
	if err != nil {
		t.Fatal(err)
	}
	for _, m := range Modes(c) {
		// Topic drills need a topic; spot-the-error needs diagram questions this course
		// does not have, and is greyed out rather than offered.
		if m.NeedsTopic || m.ID == "visual" {
			continue
		}
		if Available(set, m) == 0 {
			t.Errorf("mode %s has nothing to draw on", m.ID)
			continue
		}
		for seed := int64(0); seed < 20; seed++ {
			qs, err := Build(set, m, "", 0, History{Mastered: map[string]bool{}}, rand.New(rand.NewSource(seed)))
			if err != nil {
				t.Fatalf("mode %s seed %d: %v", m.ID, seed, err)
			}
			if len(qs) == 0 {
				t.Fatalf("mode %s built an empty run", m.ID)
			}
			if m.ID == "exam" && len(qs) < 30 {
				t.Errorf("mock quiz has only %d questions", len(qs))
			}
		}
	}
}
