package adaptive_card

import (
	"testing"
)

func TestAdaptiveCardNew(t *testing.T) {
	ac := New()
	if len(ac.Content.Bodies) != 0 {
		t.Errorf("[Case1] Expected: AdaptiveCard.Content.[]Bodies{}, Result: %v", ac.Content.Bodies)
	}
}

func TestAdaptiveCardAdaptiveCardGetVersion(t *testing.T) {
	cases := []struct {
		version float64
	}{
		{
			1.0,
		},
	}
	for i, c := range cases {
		ac := New()
		if ac.GetVersion() != c.version {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.version, ac.GetVersion())
		}
	}
}

func TestAdaptiveCardAdaptiveCardGetType(t *testing.T) {
	cases := []struct {
		typ string
	}{
		{
			"AdaptiveCard",
		},
	}
	for i, c := range cases {
		ac := New()
		if ac.GetType() != c.typ {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.typ, ac.GetType())
		}
	}
}

func TestAdaptiveCardAdaptiveCardAppend(t *testing.T) {
	cases := []struct {
		bodies []Element
	}{
		{
			[]Element{&Column{}},
		},
		{
			[]Element{&Column{}, &Column{}},
		},
	}
	for i, c := range cases {
		ac := New()
		for _, content := range c.bodies {
			ac.Append(content)
		}
		if len(ac.Content.Bodies) != len(c.bodies) {
			t.Errorf("[Case%d] Contents Count Expected: %v, Result: %v", i+1, len(c.bodies), len(ac.Content.Bodies))
		}
	}
}

func TestAdaptiveCardAdaptiveCardSetLang(t *testing.T) {
	cases := []struct {
		lang string
	}{
		{"ja"},
		{"en"},
		{"de"},
		{"fr"},
	}
	for i, c := range cases {
		ac := New()
		ac.SetLang(c.lang)
		if ac.Content.Lang != c.lang {
			t.Errorf("[Case%d] Expected: %v, Result: %v", i+1, c.lang, ac.Content.Lang)
		}
	}
}
