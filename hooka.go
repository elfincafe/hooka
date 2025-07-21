package hooka

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	serviceName = "Hooka"
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
		} else {
			h.service = Unknown
		}
	}
	h.url = u
	return h
}

func (h *Hooka) Set(elem Element) {
	h.attachments = append(h.attachments, elem)
}

func (h *Hooka) Send(data []byte) error {
	return h.post(data)
}

func (h *Hooka)SendAdaptiveCards(ac... adaptive_card.AdaptiveCard)error{
	if h.service != Teams {
		return fmt.Errorf("%s is not Teams mode now", serviceName)
	}
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

func (h *Hooka) post(data string) error {
	// Request
	req, err := http.NewRequest("POST", h.url.String(), strings.NewReader(data))
	if err != nil {
		return err
	}
	// Header
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	req.Header = headers
	// Send
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 && res.StatusCode == 204 {
		return fmt.Errorf("[%d] %s", res.StatusCode, res.Status)
	}
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
	return nil
}
