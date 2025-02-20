package teams

import "encoding/json"

type (
	AdaptiveCard struct {
		Schema  string `json:"$schema"`
		Type    string `json:"type"`
		Version string `json:"version"`
		Body    []BodyElem
	}
	BodyElem interface {
		GetType() string
	}
	TextBlock struct {
		Type    string `json:"type"`
		Text    string `json:"text"`
		Size    string `json:"size,omitempty"`
		Spacing string `json:"spacing,omitempty"`
	}
)

func NewAdaptiveCard() *AdaptiveCard {
	ac := new(AdaptiveCard)
	ac.Schema = "http://adaptivecards.io/schemas/adaptive-card.json"
	ac.Version = "1.0"
	ac.Body = []BodyElem{}
	return ac
}

func (ac *AdaptiveCard) GetType() string {
	return ac.Type
}

func (ac *AdaptiveCard) AppendBody(elem BodyElem) {
	ac.Body = append(ac.Body, elem)
}

func (ac *AdaptiveCard) Marshal() ([]byte, error) {
	return json.Marshal(ac)
}

func NewAdaptiveCardTextBlock(text string) *TextBlock {
	tb := new(TextBlock)
	tb.Type = "TextBlock"
	tb.Text = text
	return tb
}

func (tb *TextBlock) GetType() string {
	return tb.Type
}
