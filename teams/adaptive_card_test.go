package teams

import (
	"reflect"
	"testing"
)

func TestNewAdaptiveCard(t *testing.T) {
	ac := NewAdaptiveCard()
	expected := "http://adaptivecards.io/schemas/adaptive-card.json"
	if ac.Schema != expected {
		t.Errorf("[Schema] Result: %v, Expected: %v", ac.Schema, expected)
	}
	expected = "1.0"
	if ac.Schema != expected {
		t.Errorf("[Version] Result: %v, Expected: %v", ac.Version, expected)
	}
	r := reflect.TypeOf(ac)
	expected = "*hooka.AdaptiveCard"
	if r.String() != expected {
		t.Errorf("[Struct] Result: %v, Expected: %v", r.String(), expected)
	}
}

// func TestTextAdaptiveCardMarshal(t *testing.T) {
// 	// Test Case
// 	cases := []struct {
// 		bodies   []BodyElem
// 		expected string
// 	}{
// 		{
// 			[]BodyElem{
// 				NewAdaptiveCardTextBlock("TestText"),
// 			},
// 			`{"$schema":"http://adaptivecards.io/schemas/adaptive-card.json","type":"","version":"1.0","Body":[{"type":"TextBlock","text":"TestText"}]}`,
// 		},
// 	}
// 	// Test
// 	for i, c := range cases {
// 		ac := NewAdaptiveCard()
// 		for _, elem := range c.bodies {
// 			ac.AppendBody(elem)
// 		}
// 		b, _ := ac.Marshal()
// 		if !bytes.Equal(b, []byte(c.expected)) {
// 			t.Errorf(`[Case%d] Result: %v, Expected: %v`, i, string(b), c.expected)
// 		}
// 	}
// }
