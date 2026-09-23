package content

import "testing"

// TestFreeTextGradingIsLiberal pins the tolerance down, because "liberal" has to mean
// something specific: a slip of the pen passes, a different answer does not.
func TestFreeTextGradingIsLiberal(t *testing.T) {
	q := Question{ID: "t-1", Type: "text", Accept: []string{"transformer"}}

	accept := []string{
		"transformer",   // exact
		"Transformer",   // case
		"TRANSFORMER",   // shouting
		"  transformer", // stray whitespace
		"transformer.",  // full stop
		"transformer,",  // comma
		"transfomer",    // dropped letter
		"transformr",    // another dropped letter
		"transsformer",  // doubled letter
		"transformers",  // plural
		"the transformer",
	}
	for _, in := range accept {
		if !q.CorrectText(in) {
			t.Errorf("%q should have been accepted", in)
		}
	}

	reject := []string{
		"",
		"tranny",
		"transducer",    // a real, different word
		"transfer",      // close in spelling, wrong answer
		"recurrent",
		"attention",
	}
	for _, in := range reject {
		if q.CorrectText(in) {
			t.Errorf("%q should have been rejected", in)
		}
	}
}

// TestFreeTextShortAnswersStillTolerateASlip guards the floor: a 20% tolerance on a
// four-letter word rounds to zero edits, which would fail an obvious typo.
func TestFreeTextShortAnswersStillTolerateASlip(t *testing.T) {
	q := Question{ID: "t-2", Type: "text", Accept: []string{"bias"}}
	for _, in := range []string{"bias", "Bias", "bais", "biass"} {
		if !q.CorrectText(in) {
			t.Errorf("%q should have been accepted", in)
		}
	}
	for _, in := range []string{"basis", "weight", "variance"} {
		if q.CorrectText(in) {
			t.Errorf("%q should have been rejected", in)
		}
	}
}

// TestFreeTextTakesAnyAcceptedAnswer covers questions with more than one right answer.
func TestFreeTextTakesAnyAcceptedAnswer(t *testing.T) {
	q := Question{ID: "t-3", Type: "text", Accept: []string{"generative", "generative pre-trained"}}
	for _, in := range []string{"generative", "Generative", "generativ", "generative pre trained"} {
		if !q.CorrectText(in) {
			t.Errorf("%q should have been accepted", in)
		}
	}
	if q.CorrectText("discriminative") {
		t.Error("a different word should have been rejected")
	}
}

func TestTextQuestionsValidate(t *testing.T) {
	if err := validate(Question{ID: "t-4", Type: "text", Accept: []string{"x"}}); err != nil {
		t.Errorf("a text question with an accepted answer should validate: %v", err)
	}
	if err := validate(Question{ID: "t-5", Type: "text"}); err == nil {
		t.Error("a text question with no accepted answer should not validate")
	}
	if err := validate(Question{ID: "t-6", Type: "text", Accept: []string{"  "}}); err == nil {
		t.Error("a blank accepted answer should not validate")
	}
}
