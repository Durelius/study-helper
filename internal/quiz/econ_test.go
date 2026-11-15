package quiz

import (
	"math"
	"math/rand"
	"strings"
	"testing"

	"github.com/wilhelmdurelius/chulastudy/internal/content"
	"github.com/wilhelmdurelius/chulastudy/internal/generate"
)

// TestEconGeneratorsAreSound runs every economics generator over many seeds: each
// question must pass the authoring checks, carry an explanation, have a finite key
// that its own explanation states, and grade its own key as correct.
func TestEconGeneratorsAreSound(t *testing.T) {
	for name, gen := range econGenerators {
		for seed := int64(0); seed < 200; seed++ {
			qs, err := gen(12, rand.New(rand.NewSource(seed)))
			if err != nil {
				t.Fatalf("%s seed %d: %v", name, seed, err)
			}
			if len(qs) < 12 {
				t.Fatalf("%s seed %d: only %d questions", name, seed, len(qs))
			}
			ids := map[string]bool{}
			for _, q := range qs {
				if err := content.Validate(q); err != nil {
					t.Fatalf("%s seed %d: %v", name, seed, err)
				}
				if ids[q.ID] {
					t.Fatalf("%s: duplicate id %s", name, q.ID)
				}
				ids[q.ID] = true
				if q.Explanation == "" || q.Stem == "" || q.Source.Slide == 0 {
					t.Fatalf("%s: incomplete question %+v", name, q)
				}
				if q.Type != "num" {
					continue
				}
				v := *q.Value
				if math.IsNaN(v) || math.IsInf(v, 0) {
					t.Fatalf("%s: key is %v in %q", name, v, q.Stem)
				}
				key := content.FormatNumber(v)
				if !q.CorrectNumber(key) {
					t.Fatalf("%s: own key %q does not grade correct", name, key)
				}
				if !strings.Contains(q.Explanation, key) {
					t.Fatalf("%s: explanation does not state the key %s:\n%s\n%s", name, key, q.Stem, q.Explanation)
				}
				// Two decimals is the most a booklet answer should need.
				if math.Abs(v*100-math.Round(v*100)) > 1e-6 && !q.CorrectNumber(content.FormatNumber(math.Round(v*100)/100)) {
					t.Fatalf("%s: key %v cannot be typed to two decimals", name, v)
				}
			}
		}
	}
}

// TestOligopolyTableMatchesPacket pins the 16-cell table to the packet's Grab/Uber
// results table (P = 339 − Q, MC = 147), column by column.
func TestOligopolyTableMatchesPacket(t *testing.T) {
	want := []float64{
		192, 147, 28224, 0, // Bertrand
		144, 195, 28080, 6912, // Stackelberg
		128, 211, 27008, 8192, // Cournot
		96, 243, 23328, 9216, // Collusion
	}
	cells := oligopolyTable(generate.Market{A: 339, C: 147})
	if len(cells) != len(want) {
		t.Fatalf("got %d cells, want %d", len(cells), len(want))
	}
	for i, c := range cells {
		if *c.Value != want[i] {
			t.Errorf("cell %d (%s): got %v, want %v", i+1, c.Stem, *c.Value, want[i])
		}
	}
}
