package adaptive_card

import (
	"encoding/json"
)

type (
	AdaptiveCard struct {
		Schema          string    `json:"$schema"`
		Type            string    `json:"type"`
		Version         float32   `json:"version"`
		Bodies          []Element `json:"body"`
		FallbackText    string    `json:"fallbackText,omitempty"`
		BackgroundImage string    `json:"backgroundImage,omitempty"`
		Speak           string    `json:"speak,omitempty"`
		Lang            string    `json:"lang,omitempty"`
	}
	Element interface {
		GetVersion() float32
		GetType() string
	}
)

func NewAdaptiveCard() *AdaptiveCard {
	ac := &AdaptiveCard{
		Schema:          "http://adaptivecards.io/schemas/adaptive-card.json",
		Type:            "AdaptiveCard",
		Version:         1.0,
		Bodies:          []Element{},
		FallbackText:    "",
		BackgroundImage: "",
		Speak:           "",
		Lang:            "",
	}
	return ac
}

func (ac *AdaptiveCard) GetVersion() float32 {
	return ac.Version
}

func (ac *AdaptiveCard) GetType() string {
	return ac.Type
}

func (ac *AdaptiveCard) Append(elem Element) {
	ac.Bodies = append(ac.Bodies, elem)
	if elem.GetVersion() > ac.GetVersion() {
		ac.Version = elem.GetVersion()
	}
}

func (ac *AdaptiveCard) SetLang(lang string) {
	ac.Lang = lang
}

func (ac *AdaptiveCard) Marshal() ([]byte, error) {
	return json.Marshal(ac)
}
