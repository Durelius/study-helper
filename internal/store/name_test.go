package store

import (
	"errors"
	"testing"
)

func TestNormaliseName(t *testing.T) {
	cases := []struct{ in, want string }{
		{"wilhelm", "Wilhelm"},
		{"WILHELM", "Wilhelm"},
		{"WiLhElM", "Wilhelm"},
		{"  wilhelm  ", "Wilhelm"},
		{"wilhelm durelius", "Wilhelm Durelius"},
		{"WILHELM DURELIUS", "Wilhelm Durelius"},
		{"wilhelm   durelius", "Wilhelm Durelius"},
		{"\tbecka\n", "Becka"},
		// Swedish vowels survive, and capitalise correctly.
		{"åsa", "Åsa"},
		{"ÅSA", "Åsa"},
		{"örjan ängström", "Örjan Ängström"},
		{"MALMÖ", "Malmö"},
	}
	for _, c := range cases {
		got, err := NormaliseName(c.in)
		if err != nil {
			t.Errorf("NormaliseName(%q) errored: %v", c.in, err)
			continue
		}
		if got != c.want {
			t.Errorf("NormaliseName(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

// TestNormaliseNameMatchesAcrossSpellings is the property that actually matters: two
// spellings of one person must collapse to one leaderboard row.
func TestNormaliseNameMatchesAcrossSpellings(t *testing.T) {
	for _, pair := range [][2]string{
		{"wilhelm", "Wilhelm"},
		{"SHAYAN", "shayan"},
		{" becka ", "Becka"},
		{"anna karin", "ANNA   KARIN"},
	} {
		a, err1 := NormaliseName(pair[0])
		b, err2 := NormaliseName(pair[1])
		if err1 != nil || err2 != nil {
			t.Fatalf("unexpected error: %v %v", err1, err2)
		}
		if a != b {
			t.Errorf("%q and %q normalised to %q and %q — they should be the same player", pair[0], pair[1], a, b)
		}
	}
}

func TestNormaliseNameRejects(t *testing.T) {
	cases := []struct {
		in   string
		want error
	}{
		{"", ErrEmptyName},
		{"   ", ErrEmptyName},
		{"\t\n", ErrEmptyName},
		{"wilhelm2", ErrNameCharacters},
		{"wilhelm!", ErrNameCharacters},
		{"o'brien", ErrNameCharacters},
		{"anna-lena", ErrNameCharacters},
		{"wilhelm@example.com", ErrNameCharacters},
		{"🎉", ErrNameCharacters},
		{"عبدالله", ErrNameCharacters},
		{"this name is far too long to fit on the board", ErrNameTooLong},
	}
	for _, c := range cases {
		got, err := NormaliseName(c.in)
		if !errors.Is(err, c.want) {
			t.Errorf("NormaliseName(%q) = %q, %v — wanted %v", c.in, got, err, c.want)
		}
	}
}
