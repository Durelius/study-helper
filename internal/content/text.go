package content

import (
	"strings"
	"unicode"
)

// Free-text answers are graded liberally on purpose. The exam this drills is written by
// hand under time pressure, and the point of the drill is whether the term is in the
// reader's head — not whether they can spell it. So case, punctuation, surrounding
// words and a near miss all pass.

// normalise strips everything that should not decide a mark: case, punctuation, extra
// whitespace, and the articles people put in without thinking.
func normalise(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) || r == '-' {
			b.WriteRune(r)
		}
	}
	fields := strings.Fields(b.String())
	kept := fields[:0]
	for _, f := range fields {
		if f == "the" || f == "a" || f == "an" {
			continue
		}
		kept = append(kept, f)
	}
	return strings.Join(kept, " ")
}

// distance is the Damerau-Levenshtein edit distance, which counts swapping two
// adjacent characters as one edit rather than two. Transposition is the commonest
// typo there is — "bais" for "bias" — and plain Levenshtein is unreasonably harsh
// about it on short words.
func distance(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	if len(ra) == 0 {
		return len(rb)
	}
	if len(rb) == 0 {
		return len(ra)
	}
	prev := make([]int, len(rb)+1)
	curr := make([]int, len(rb)+1)
	// grid keeps every row, which the transposition rule needs to look two back.
	grid := make([][]int, len(ra)+1)
	for i := range grid {
		grid[i] = make([]int, len(rb)+1)
	}
	for j := range prev {
		prev[j] = j
		grid[0][j] = j
	}
	for i := range grid {
		grid[i][0] = i
	}
	for i := 1; i <= len(ra); i++ {
		curr[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			curr[j] = min3(curr[j-1]+1, prev[j]+1, prev[j-1]+cost)
		}
		if i > 1 {
			// A transposition needs the row from two steps back, so it is applied
			// here against the grid rather than inside the loop above.
			for j := 2; j <= len(rb); j++ {
				if ra[i-1] == rb[j-2] && ra[i-2] == rb[j-1] {
					if d := grid[i-2][j-2] + 1; d < curr[j] {
						curr[j] = d
					}
				}
			}
		}
		copy(grid[i], curr)
		prev, curr = curr, prev
	}
	return grid[len(ra)][len(rb)]
}

func min3(a, b, c int) int {
	if b < a {
		a = b
	}
	if c < a {
		a = c
	}
	return a
}

// closeEnough reports whether a typed answer matches an accepted one.
//
// The tolerance is 20% of the accepted answer's length, with a floor of one edit so
// that short words like "bias" still survive a slip. A long word like "transformer"
// gets two edits, which covers the usual doubled or dropped letter without letting a
// genuinely different word through.
func closeEnough(typed, accepted string) bool {
	if typed == "" {
		return false
	}
	if typed == accepted {
		return true
	}
	allowed := len([]rune(accepted)) / 5
	if allowed < 1 {
		allowed = 1
	}
	return distance(typed, accepted) <= allowed
}

// CorrectText reports whether a typed answer matches any accepted answer.
func (q Question) CorrectText(typed string) bool {
	t := normalise(typed)
	if t == "" {
		return false
	}
	for _, a := range q.Accept {
		if closeEnough(t, normalise(a)) {
			return true
		}
	}
	return false
}
