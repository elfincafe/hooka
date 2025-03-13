package hooka

import "encoding/json"

type (
	AdaptiveCard struct {
		Schema  string     `json:"$schema"`
		Type    string     `json:"type"`
		Version string     `json:"version"`
		Bodies  []BodyElem `json:"body"`
	}
	BodyElem interface {
		GetType() string
	}
)

func NewAdaptiveCard() *AdaptiveCard {
	ac := new(AdaptiveCard)
	ac.Schema = "http://adaptivecards.io/schemas/adaptive-card.json"
	ac.Version = "1.0"
	ac.Bodies = []BodyElem{}
	return ac
}

func (ac *AdaptiveCard) GetType() string {
	return ac.Type
}

func (ac *AdaptiveCard) AppendBody(elem BodyElem) {
	ac.Bodies = append(ac.Bodies, elem)
}

func (ac *AdaptiveCard) Marshal() ([]byte, error) {
	payload := struct {
		ContentType string        `json:"contentType"`
		Content     *AdaptiveCard `json:"content"`
	}{
		"application/vnd.microsoft.card.adaptive",
		ac,
	}
	return json.Marshal(payload)
}
