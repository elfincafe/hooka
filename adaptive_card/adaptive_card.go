package adaptive_card

import "encoding/json"

type (
	AdaptiveCard struct {
		Schema   string    `json:"$schema"`
		Type     string    `json:"type"`
		Version  float32   `json:"version"`
		Contents []Element `json:"content"`
	}
	Element interface {
		GetVersion() float32
		GetType() string
		SetId(string)
	}
)

func NewAdaptiveCard() AdaptiveCard {
	ac := AdaptiveCard{
		Schema:   "http://adaptivecards.io/schemas/adaptive-card.json",
		Type:     "AdaptiveCard",
		Version:  1.0,
		Contents: []Element{},
	}
	return ac
}

func (ac AdaptiveCard) GetVersion() float32 {
	return ac.Version
}
func (ac AdaptiveCard) GetType() string {
	return ac.Type
}

func (ac AdaptiveCard) AppendContent(elem Element) {
	ac.Contents = append(ac.Contents, elem)
	if elem.GetVersion() > ac.GetVersion() {
		ac.Version = elem.GetVersion()
	}
}

func (ac AdaptiveCard) Marshal() ([]byte, error) {
	return json.Marshal(ac)
}
