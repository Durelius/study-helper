package store

import (
	"errors"
	"strings"
	"unicode"
)

// Leaderboard names are matched case-insensitively and stored in one canonical
// spelling, so "wilhelm", "WILHELM" and "Wilhelm" are one person with one row on the
// board rather than three.

// MaxNameLength is the longest name the board will show.
const MaxNameLength = 24

// ErrEmptyName and ErrNameCharacters explain a rejected name in words the person typing
// it can act on.
var (
	ErrEmptyName      = errors.New("A name is needed — it goes on the leaderboard.")
	ErrNameCharacters = errors.New("Letters and spaces only (Å, Ä and Ö are fine).")
	ErrNameTooLong    = errors.New("That name is too long for the board.")
)

// allowed reports whether a rune may appear in a name: the Latin alphabet, the Swedish
// vowels, and the space between given and family name. Everything else — digits,
// punctuation, emoji — is refused rather than silently stripped, so nobody wonders why
// the board is showing something they did not type.
func allowed(r rune) bool {
	switch {
	case r >= 'a' && r <= 'z', r >= 'A' && r <= 'Z':
		return true
	case r == 'å', r == 'ä', r == 'ö', r == 'Å', r == 'Ä', r == 'Ö':
		return true
	case r == ' ':
		return true
	}
	return false
}

// NormaliseName canonicalises a name: trimmed, single-spaced, and each word capitalised
// with the rest lower case. It reports why a name is unusable rather than quietly
// changing it into something else.
func NormaliseName(name string) (string, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return "", ErrEmptyName
	}
	for _, r := range name {
		if !allowed(r) {
			return "", ErrNameCharacters
		}
	}
	// Collapse runs of spaces, so "Anna   Karin" and "Anna Karin" are one person.
	words := strings.Fields(name)
	if len(words) == 0 {
		return "", ErrEmptyName
	}
	for i, w := range words {
		runes := []rune(strings.ToLower(w))
		runes[0] = unicode.ToUpper(runes[0])
		words[i] = string(runes)
	}
	out := strings.Join(words, " ")
	if len([]rune(out)) > MaxNameLength {
		return "", ErrNameTooLong
	}
	return out, nil
}
