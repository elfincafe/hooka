package adaptive_card

import "testing"

func TestAdaptiveCardNewFact(t *testing.T) {
	cases := []struct {
		title string
		value string
	}{
		{"TestTitle", "TestValue"},
	}
	for i, c := range cases {
		f := NewFact(c.title, c.value)
		if f.Title != c.title {
			t.Errorf("[Case%d] Title Expected: %v, Result: %v", i+1, c.title, f.Title)
		}
		if f.Value != c.value {
			t.Errorf("[Case%d] Value Expected: %v, Result: %v", i+1, c.value, f.Value)
		}
	}
}
