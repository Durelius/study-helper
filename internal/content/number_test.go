package content

import "testing"

func f(v float64) *float64 { return &v }

func TestParseNumber(t *testing.T) {
	cases := map[string]float64{
		"4,096": 4096, "4096 THB": 4096, " 1/2 ": 0.5, "0.5": 0.5, "50%": 0.5,
		"−3": -3, "-12.5": -12.5, "½": 0.5, "211 baht": 211, "2.5": 2.5,
		"144 rides/min": 144, "28,224 THB/min": 28224, "1 / 2": 0.5, "p = 0.5": 0.5,
	}
	for in, want := range cases {
		got, ok := ParseNumber(in)
		if !ok || got != want {
			t.Errorf("ParseNumber(%q) = %v, %v; want %v", in, got, ok, want)
		}
	}
	for _, bad := range []string{"", "abc", "1/0", "--", "/"} {
		if _, ok := ParseNumber(bad); ok {
			t.Errorf("ParseNumber(%q) should fail", bad)
		}
	}
}

func TestCorrectNumber(t *testing.T) {
	q := Question{ID: "x", Type: "num", Value: f(4096)}
	for _, ok := range []string{"4096", "4,096", "4096 THB"} {
		if !q.CorrectNumber(ok) {
			t.Errorf("%q should pass for 4096", ok)
		}
	}
	for _, bad := range []string{"4095", "4069", "40.96", "", "four thousand"} {
		if q.CorrectNumber(bad) {
			t.Errorf("%q should fail for 4096", bad)
		}
	}
	half := Question{ID: "y", Type: "num", Value: f(0.5)}
	for _, ok := range []string{"1/2", "0.5", ".5", "50%"} {
		if !half.CorrectNumber(ok) {
			t.Errorf("%q should pass for 0.5", ok)
		}
	}
	third := Question{ID: "z", Type: "num", Value: f(1.0 / 3)}
	if !third.CorrectNumber("0.33") || !third.CorrectNumber("1/3") || third.CorrectNumber("0.3") {
		t.Error("a third should take 0.33 and 1/3 but not 0.3")
	}
	if got := FormatNumber(28224); got != "28,224" {
		t.Errorf("FormatNumber = %q", got)
	}
	if got := FormatNumber(-5852.25); got != "−5,852.25" {
		t.Errorf("FormatNumber = %q", got)
	}
	if got := FormatNumber(2.5); got != "2.5" {
		t.Errorf("FormatNumber = %q", got)
	}
}
