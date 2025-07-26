package adaptive_card

type (
	AdaptiveCard struct {
		ContentType string              `json:"contentType"`
		Content     AdaptiveCardContent `json:"content"`
	}
	AdaptiveCardContent struct {
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
		ContentType: "application/vnd.microsoft.card.adaptive",
		Content: AdaptiveCardContent{
			Schema:          "http://adaptivecards.io/schemas/adaptive-card.json",
			Type:            "AdaptiveCard",
			Version:         1.0,
			Bodies:          []Element{},
			FallbackText:    "",
			BackgroundImage: "",
			Speak:           "",
			Lang:            "",
		},
	}
	return ac
}

func (ac *AdaptiveCard) GetVersion() float32 {
	return ac.Content.Version
}

func (ac *AdaptiveCard) GetType() string {
	return ac.Content.Type
}

func (ac *AdaptiveCard) Append(elem Element) {
	ac.Content.Bodies = append(ac.Content.Bodies, elem)
	if elem.GetVersion() > ac.GetVersion() {
		ac.Content.Version = elem.GetVersion()
	}
}

func (ac *AdaptiveCard) SetLang(lang string) {
	ac.Content.Lang = lang
}
