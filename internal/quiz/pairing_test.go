package quiz

import (
	"encoding/json"
	"math/rand"
	"strings"
	"testing"
)

// TestPairingCompletesTheDish pins the rule the drill teaches against the worked
// reading on the slides: each wine axis matches the leading food component it faces.
func TestPairingCompletesTheDish(t *testing.T) {
	f := FoodProfile{SweetTendency: 6, Fattiness: 3, Juiciness: 2, Oiliness: 7, Hardness: 4, Persistence: 5}
	got := f.Complete()
	want := WineProfile{Softness: 4, Intensity: 5, AlcTannin: 7, Acidity: 6}
	if got != want {
		t.Fatalf("Complete() = %+v, want %+v", got, want)
	}
}

// TestPairingQuestionsAreSound generates many runs: each must validate, each MCQ key
// must be the wine (or dish) the rule gives, and no two choices may be the same.
func TestPairingQuestionsAreSound(t *testing.T) {
	for seed := int64(0); seed < 500; seed++ {
		rnd := rand.New(rand.NewSource(seed))
		qs, err := generatedPairing(6, rnd)
		if err != nil {
			t.Fatal(err)
		}
		for _, q := range qs {
			var c PairingChart
			if err := json.Unmarshal(q.Data, &c); err != nil || c.Kind != "pairing" {
				t.Fatalf("%s: bad chart data %s", q.ID, q.Data)
			}
			switch q.Type {
			case "num":
				if q.Value == nil || *q.Value < 0 || *q.Value > 10 || c.Food == nil || c.Wine != nil {
					t.Fatalf("%s: numeric pairing question malformed", q.ID)
				}
			case "mcq":
				seen := map[string]bool{}
				for _, ch := range q.Choices {
					if seen[ch] {
						t.Fatalf("%s: duplicate choice %q", q.ID, ch)
					}
					seen[ch] = true
				}
				key := q.Choices[q.Answer[0]]
				switch {
				case c.Food != nil && c.Wine == nil:
					if key != c.Food.Complete().String() {
						t.Fatalf("%s: key %q is not the completing wine", q.ID, key)
					}
				case c.Wine != nil && c.Food == nil:
					if !strings.HasPrefix(key, "Sweet tendency") {
						t.Fatalf("%s: key %q is not a dish", q.ID, key)
					}
				}
			default:
				t.Fatalf("%s: unexpected type %s", q.ID, q.Type)
			}
		}
	}
}
