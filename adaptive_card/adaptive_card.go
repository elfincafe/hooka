package adaptive_card

import "json"

type (
	AdaptiveCard struct {
		Schema   string    `json:"$schema"`
		Type     string    `json:"type"`
		Version  string    `json:"version"`
		Contents []Content `json:"content"`
	}
	Element interface {
		GetVersion() string
		GetType() string
		SetId(string)
	}
)

func NewAdaptiveCard() AdaptiveCard {
	ac := AdaptiveCard{
		Schema:   "http://adaptivecards.io/schemas/adaptive-card.json",
		Type:     "AdaptiveCard",
		Version:  1.0,
		Contents: []Content{},
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
