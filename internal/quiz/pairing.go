package quiz

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"

	"github.com/wilhelmdurelius/chulastudy/internal/content"
)

// The AIS food-and-wine pairing chart (Wine Education, Analysis deck slides 55-64).
//
// The dish is scored 0-10 on six axes (red) and the wine on four (blue). The chart is
// laid out so each wine axis faces the food axes it has to answer:
//
//	food sweet tendency, fattiness   → wine acidity / effervescence   (opposite)
//	food juiciness, oiliness         → wine alcoholicity / tannicity  (opposite)
//	food saltiness, bitter, sour     → wine sweetness / roundness     (opposite: softness)
//	food persistence, spice, aroma   → wine T-O intensity / persistence (parallel)
//
// When a group holds several components, the leading (highest) one is plotted, which
// is what slides 59-60 say. A wine that completes the dish scores the same on its axis
// as the leading food component it faces, so the two polygons come out the same shape
// and size — the "compatible" reading of slides 62-64.

func init() {
	econGenerators["pairing"] = generatedPairing
}

const pairingTopic = "pairing"

// PairingChart is what the client draws. A nil side is not drawn: that is the polygon
// the reader has to complete.
type PairingChart struct {
	Kind string       `json:"kind"` // always "pairing"
	Food *FoodProfile `json:"food,omitempty"`
	Wine *WineProfile `json:"wine,omitempty"`
}

type FoodProfile struct {
	SweetTendency int `json:"sweetTendency"`
	Fattiness     int `json:"fattiness"`
	Juiciness     int `json:"juiciness"`
	Oiliness      int `json:"oiliness"`
	Hardness      int `json:"hardness"`    // saltiness / bitter / sour tendency / sweetness
	Persistence   int `json:"persistence"` // T-O persistence / spiciness / aromaticity
}

type WineProfile struct {
	Softness  int `json:"softness"`  // sweetness / roundness
	Intensity int `json:"intensity"` // T-O intensity / persistence
	AlcTannin int `json:"alcTannin"` // alcoholicity / tannicity
	Acidity   int `json:"acidity"`   // acidity / effervescence / saltiness
}

// Complete is the wine that balances a dish on the chart.
func (f FoodProfile) Complete() WineProfile {
	return WineProfile{
		Softness:  f.Hardness,
		Intensity: f.Persistence,
		AlcTannin: max(f.Juiciness, f.Oiliness),
		Acidity:   max(f.SweetTendency, f.Fattiness),
	}
}

func (w WineProfile) String() string {
	return fmt.Sprintf("Softness %d · Intensity %d · Alcohol/tannin %d · Acidity/effervescence %d",
		w.Softness, w.Intensity, w.AlcTannin, w.Acidity)
}

func (f FoodProfile) String() string {
	return fmt.Sprintf("Sweet tendency %d · Fattiness %d · Juiciness %d · Oiliness %d · Salt/bitter/sour %d · Persistence/spice/aroma %d",
		f.SweetTendency, f.Fattiness, f.Juiciness, f.Oiliness, f.Hardness, f.Persistence)
}

func clamp(n int) int {
	if n < 0 {
		return 0
	}
	if n > 10 {
		return 10
	}
	return n
}

func (w WineProfile) shift(d int) WineProfile {
	return WineProfile{clamp(w.Softness + d), clamp(w.Intensity + d), clamp(w.AlcTannin + d), clamp(w.Acidity + d)}
}

func (f FoodProfile) shift(d int) FoodProfile {
	return FoodProfile{clamp(f.SweetTendency + d), clamp(f.Fattiness + d), clamp(f.Juiciness + d),
		clamp(f.Oiliness + d), clamp(f.Hardness + d), clamp(f.Persistence + d)}
}

// randomDish draws a dish whose leading components differ enough that a swapped wine
// is visibly wrong: the acidity side and the tannin side are at least 3 apart, and so
// are softness and intensity.
func randomDish(rnd *rand.Rand, lo, hi int) FoodProfile {
	gap := min(3, hi-lo)
	for {
		r := func() int { return lo + rnd.Intn(hi-lo+1) }
		f := FoodProfile{r(), r(), r(), r(), r(), r()}
		w := f.Complete()
		if abs(w.Acidity-w.AlcTannin) >= gap && abs(w.Softness-w.Intensity) >= min(2, gap) {
			return f
		}
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

const pairingRule = "Each wine axis answers the food axes it faces: acidity/effervescence answers sweet tendency and fattiness; " +
	"alcohol/tannin answers juiciness and oiliness; softness (sweetness, roundness) answers saltiness, bitter and sour tendency; " +
	"intensity/persistence runs parallel to the dish's persistence, spiciness and aromaticity. Plot the leading component of each group, " +
	"and the wine that completes the dish has the same score as what it faces."

func chartData(c PairingChart) json.RawMessage {
	c.Kind = "pairing"
	b, _ := json.Marshal(c)
	return b
}

func generatedPairing(count int, rnd *rand.Rand) ([]content.Question, error) {
	var out []content.Question
	for len(out) < count {
		var q content.Question
		var ok bool
		switch rnd.Intn(6) {
		case 0, 1:
			q, ok = pairingWineForFood(rnd)
		case 2:
			q, ok = pairingFoodForWine(rnd)
		case 3:
			q, ok = pairingReading(rnd)
		default:
			q, ok = pairingAxisScore(rnd)
		}
		if !ok {
			continue
		}
		q.ID = fmt.Sprintf("gen:pairing:%d:%d", len(out), rnd.Int63())
		out = append(out, q)
	}
	return out, nil
}

func pairingQ(stem string, data json.RawMessage, expl string, slide int) content.Question {
	return content.Question{
		Topic: pairingTopic, Stem: stem, Data: data, Explanation: expl,
		Source: content.Source{Deck: "Analysis", Slide: slide}, Difficulty: 2,
		Tags: []string{"generated", "polygon"}, ExamFocus: true,
	}
}

// The dish is drawn; pick the wine that completes it.
func pairingWineForFood(rnd *rand.Rand) (content.Question, bool) {
	f := randomDish(rnd, 1, 9)
	w := f.Complete()
	swapSides := WineProfile{w.Softness, w.Intensity, w.Acidity, w.AlcTannin}
	swapTop := WineProfile{w.Intensity, w.Softness, w.AlcTannin, w.Acidity}
	small, big := w.shift(-3), w.shift(3)
	wrong := []string{swapSides.String(), swapTop.String()}
	if rnd.Intn(2) == 0 {
		wrong = append(wrong, small.String(), big.String())
	} else {
		wrong = append(wrong, big.String(), small.String())
	}
	choices, ans, ok := choose(w.String(), wrong, rnd)
	if !ok {
		return content.Question{}, false
	}
	q := pairingQ("The red polygon is a dish scored on the pairing chart. Which wine completes it?",
		chartData(PairingChart{Food: &f}),
		fmt.Sprintf("The completing wine is %s. %s", w, pairingRule), 59)
	q.Type, q.Choices, q.Answer = "mcq", choices, []int{ans}
	return q, true
}

// The wine is drawn; pick the dish it completes.
func pairingFoodForWine(rnd *rand.Rand) (content.Question, bool) {
	f := randomDish(rnd, 1, 9)
	w := f.Complete()
	// Swap the dish's acidity-side and tannin-side components: same numbers, wrong axes.
	swapped := FoodProfile{f.Juiciness, f.Oiliness, f.SweetTendency, f.Fattiness, f.Hardness, f.Persistence}
	flipped := FoodProfile{f.SweetTendency, f.Fattiness, f.Juiciness, f.Oiliness, f.Persistence, f.Hardness}
	wrong := []string{swapped.String(), flipped.String(), f.shift(3).String()}
	if swapped.Complete() == w || flipped.Complete() == w {
		return content.Question{}, false
	}
	choices, ans, ok := choose(f.String(), wrong, rnd)
	if !ok {
		return content.Question{}, false
	}
	q := pairingQ("The blue polygon is a wine scored on the pairing chart. Which dish does it complete?",
		chartData(PairingChart{Wine: &w}),
		fmt.Sprintf("The dish is %s. %s", f, pairingRule), 60)
	q.Type, q.Choices, q.Answer = "mcq", choices, []int{ans}
	return q, true
}

var readings = []string{
	"Simple but compatible",
	"Important and well balanced",
	"Simple but incompatible",
	"The wine overshadows the food",
	"The food overshadows the wine",
}

var readingWhy = []string{
	"Both polygons are the same shape with low scores on every axis: a simple dish and a simple wine that suit each other.",
	"Both polygons are the same shape with high scores on every axis: an important dish and wine that complement each other.",
	"The shapes differ: the dish peaks on an axis where the wine stays low, so the wine cannot balance that element.",
	"Same shape, but the wine's polygon is larger: the wine overshadows the dish.",
	"Same shape, but the dish's polygon is larger: the food overshadows the wine.",
}

// Both polygons are drawn; read the result as slides 62-64 do.
func pairingReading(rnd *rand.Rand) (content.Question, bool) {
	k := rnd.Intn(len(readings))
	var f FoodProfile
	var w WineProfile
	switch k {
	case 0:
		f = randomDish(rnd, 1, 4)
		w = f.Complete()
	case 1:
		f = randomDish(rnd, 6, 9)
		w = f.Complete()
	case 2:
		// A simple wine against a dish with one or two strong peaks.
		f = randomDish(rnd, 1, 3)
		f.Juiciness, f.Hardness = 8+rnd.Intn(3), 7+rnd.Intn(3)
		w = WineProfile{1 + rnd.Intn(2), 1 + rnd.Intn(3), 1 + rnd.Intn(2), 1 + rnd.Intn(3)}
	case 3:
		f = randomDish(rnd, 2, 5)
		w = f.Complete().shift(4)
	case 4:
		f = randomDish(rnd, 5, 9)
		w = f.Complete().shift(-4)
	}
	var wrong []string
	for i, r := range readings {
		if i != k {
			wrong = append(wrong, r)
		}
	}
	rnd.Shuffle(len(wrong), func(i, j int) { wrong[i], wrong[j] = wrong[j], wrong[i] })
	choices, ans, ok := choose(readings[k], wrong, rnd)
	if !ok {
		return content.Question{}, false
	}
	q := pairingQ("Red is the dish and blue is the wine. How do you read this pairing?",
		chartData(PairingChart{Food: &f, Wine: &w}), readingWhy[k], 62+k/2)
	q.Type, q.Choices, q.Answer = "mcq", choices, []int{ans}
	return q, true
}

// The dish is drawn; give the wine's score on one axis.
func pairingAxisScore(rnd *rand.Rand) (content.Question, bool) {
	f := randomDish(rnd, 1, 9)
	w := f.Complete()
	axes := []struct {
		name  string
		value int
		why   string
	}{
		{"acidity / effervescence", w.Acidity, fmt.Sprintf("It answers the leading of sweet tendency (%d) and fattiness (%d).", f.SweetTendency, f.Fattiness)},
		{"alcoholicity / tannicity", w.AlcTannin, fmt.Sprintf("It answers the leading of juiciness (%d) and oiliness (%d).", f.Juiciness, f.Oiliness)},
		{"sweetness / roundness (softness)", w.Softness, fmt.Sprintf("Softness answers the dish's saltiness, bitter and sour tendency (%d).", f.Hardness)},
		{"taste-olfactory intensity / persistence", w.Intensity, fmt.Sprintf("Parallel pairing: it matches the dish's persistence, spiciness and aromaticity (%d).", f.Persistence)},
	}
	a := axes[rnd.Intn(len(axes))]
	v := float64(a.value)
	q := pairingQ(fmt.Sprintf("Complete the wine for the red dish. What score does the wine get on the **%s** axis?", a.name),
		chartData(PairingChart{Food: &f}),
		strings.TrimSpace(a.why+" "+pairingRule), 59)
	q.Type, q.Value, q.Tolerance, q.Unit = "num", &v, 1, "/ 10"
	return q, true
}
