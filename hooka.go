package hooka

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

const (
	Unknown service = 0
	Teams   service = 1
	Slack   service = 2
)

type (
	service uint8
	Payload interface {
		Marshal() ([]byte, error)
	}
	Element interface {
		Marshal() ([]byte, error)
	}
	Hooka struct {
		service     service
		url         *url.URL
		attachments []Element
	}
)

func New(uri string) *Hooka {
	h := new(Hooka)
	u, err := url.Parse(uri)
	if err != nil {
		h.service = Unknown
	} else {
		// Service Type
		if strings.Contains(u.Host, "azure.com") {
			h.service = Teams
		} else if strings.Contains(u.Host, "slack.com") {
			h.service = Slack
		}
	}
	h.url = u
	return h
}

func (h *Hooka) Set(elem Element) {
	h.attachments = append(h.attachments, elem)
}

func (h *Hooka) Send(p Payload) error {
	var err error
	var payload string
	switch h.service {
	case Slack:
		payload, err = h.createSlackPayload()
	case Teams:
		payload, err = h.createTeamsPayload()
	default:
		err = fmt.Errorf("invalid service")
	}
	if err != nil {
		return err
	}

	fmt.Println(string(payload))
	return nil
}

func (h *Hooka) createTeamsPayload() (string, error) {
	payload, err := json.Marshal(
		struct {
			Type        string    `json:"type"`
			Attachments []Element `json:"attachments"`
		}{
			"message", h.attachments,
		},
	)
	if err != nil {
		return "", err
	}
	return string(payload), nil
}

func (h *Hooka) createSlackPayload() (string, error) {
	var payload []byte
	return string(payload), nil
}
