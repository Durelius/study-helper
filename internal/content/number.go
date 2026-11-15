package content

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

var numberPattern = regexp.MustCompile(`(-?\d*\.?\d+)(?:/(\d*\.?\d+))?`)

// Numeric answers are typed the way they are written in an exam booklet: "4,096",
// "4096 THB", "1/2", "0.5", "50%". Grading is exact up to a tolerance, never fuzzy —
// an edit-distance match would pass 65 for 64.

// ParseNumber reads a typed number. It ignores thousands separators, surrounding
// units and spaces, takes a simple fraction a/b, and reads a trailing % as hundredths.
func ParseNumber(s string) (float64, bool) {
	s = strings.TrimSpace(strings.ToLower(s))
	s = strings.NewReplacer("−", "-", "–", "-", "½", "1/2", "¼", "1/4", "¾", "3/4", ",", "", " ", "").Replace(s)
	percent := strings.Contains(s, "%")
	// The first number in what was typed, optionally a fraction a/b. Anything after it
	// (a unit like "THB" or "rides/min") is ignored.
	m := numberPattern.FindStringSubmatch(s)
	if m == nil {
		return 0, false
	}
	v, err := strconv.ParseFloat(m[1], 64)
	if err != nil {
		return 0, false
	}
	if m[2] != "" {
		d, err := strconv.ParseFloat(m[2], 64)
		if err != nil || d == 0 {
			return 0, false
		}
		v /= d
	}
	if percent {
		v /= 100
	}
	return v, true
}

// tolerance is how far an answer may be from the key. A whole-number key must be hit
// exactly (an arithmetic slip of one is still a slip); any other key takes rounding to
// two decimal places, or 0.1% of a large one.
func (q Question) tolerance() float64 {
	if q.Tolerance > 0 {
		return q.Tolerance
	}
	v := *q.Value
	if v == math.Trunc(v) {
		return 1e-9
	}
	return math.Max(0.006, math.Abs(v)*0.001)
}

// CorrectNumber reports whether a typed number is within tolerance of the key. A
// percentage typed for a proportion ("50%" for 0.5) passes, and so does the
// proportion typed for a percentage key.
func (q Question) CorrectNumber(typed string) bool {
	if q.Value == nil {
		return false
	}
	v, ok := ParseNumber(typed)
	if !ok {
		return false
	}
	want, tol := *q.Value, q.tolerance()
	if math.Abs(v-want) <= tol {
		return true
	}
	// "50" for a key of 0.5 when the reader meant per cent, only if they wrote no %.
	if !strings.Contains(typed, "%") && strings.Contains(q.Unit, "%") {
		return math.Abs(v/100-want) <= tol
	}
	return false
}

// FormatNumber writes a key the way the explanation would: whole numbers with
// thousands separators, everything else to at most four places.
func FormatNumber(v float64) string {
	if v == math.Trunc(v) && math.Abs(v) < 1e15 {
		n := strconv.FormatInt(int64(math.Abs(v)), 10)
		var b strings.Builder
		for i, r := range n {
			if i > 0 && (len(n)-i)%3 == 0 {
				b.WriteByte(',')
			}
			b.WriteRune(r)
		}
		if v < 0 {
			return "−" + b.String()
		}
		return b.String()
	}
	s := strconv.FormatFloat(math.Abs(v), 'f', 4, 64)
	s = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	whole, frac, _ := strings.Cut(s, ".")
	n, _ := strconv.ParseFloat(whole, 64)
	out := FormatNumber(n)
	if frac != "" {
		out += "." + frac
	}
	if v < 0 {
		return "−" + out
	}
	return out
}

// Reveal is what the reader is shown as the answer once they have submitted a typed
// question: the accepted words of a text question, or the key of a numeric one.
func (q Question) Reveal() []string {
	if q.Type == "num" && q.Value != nil {
		s := FormatNumber(*q.Value)
		if q.Unit != "" {
			s += " " + q.Unit
		}
		return []string{s}
	}
	return q.Accept
}
