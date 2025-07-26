package hooka

import (
	"encoding/json"
	"errors"
	"hooka/adaptive_card"
	"net/url"
)

type (
	Teams struct {
		uri     *url.URL
		message struct {
			Type        string                        `json:"message"`
			Attachments []*adaptive_card.AdaptiveCard `json:"attachments"`
		}
	}
)

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

func (t *Teams) Marshal() ([]byte, error) {
	data, err := json.Marshal(t.message)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (t *Teams) Attach(adaptiveCard adaptive_card.AdaptiveCard) {
	t.message.Attachments = append(t.message.Attachments, &adaptiveCard)
}
