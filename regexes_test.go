package vstruct

import "testing"

func TestRegexMatcherAlpha(t *testing.T) {
	arr := []string{
		"hello world",
	}
	for _, phrase := range arr {
		if !GetRegex().Alpha(phrase) {
			t.Errorf("Failed to match %v", phrase)
		}
	}
}
