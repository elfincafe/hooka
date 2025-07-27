package hooka

import (
	"encoding/json"
	"errors"
	"hooka/adaptive_card"
	"net/url"
)

type (
	teamsMessage struct {
		Type        string                        `json:"type"`
		Attachments []*adaptive_card.AdaptiveCard `json:"attachments"`
	}
	Teams struct {
		uri     *url.URL
		message *teamsMessage
	}
)

func NewTeams(uri string) (*Teams, error) {
	u, err := parseUri(uri, "azure.com")
	if err != nil {
		return nil, err
	}
	t := &Teams{
		uri: u,
		message: &teamsMessage{
			Type:        "message",
			Attachments: []*adaptive_card.AdaptiveCard{},
		},
	}
	return t, nil
}

func (t *Teams) Send(data []byte) error {
	res, err := send(data, t.uri)
	if err != nil {
		return err
	}
	if res.StatusCode >= 400 {
		return errors.New("Error")
	}
	return nil
}

func (t *Teams) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.message)
}

func (t *Teams) Attach(adaptiveCard *adaptive_card.AdaptiveCard) {
	t.message.Attachments = append(t.message.Attachments, adaptiveCard)
}
