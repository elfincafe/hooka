package webhook

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/elfincafe/webhook/teams"
)

type Teams struct {
	Webhook
	Payload *teams.Payload
}

func NewTeams(url string) *Teams {
	p := new(teams.Payload)
	p.Type = "MessageCard"
	p.Context = "http://schema.org/extensions"
	p.ThemeColor = "7fffd4"
	p.Summary = ""
	p.Sections = []*teams.Section{}
	h := new(Teams)
	h.Url = url
	h.Payload = p
	return h
}

func (h *Teams) SetSummary(s string) {
	h.Payload.Summary = s
}

func (h *Teams) SetSection(s *teams.Section) {
	h.Payload.Sections = append(h.Payload.Sections, s)
}

func (h *Teams) SetColor(color string) {
	h.Payload.ThemeColor = color
}

func (h *Teams) Send() error {
	p, err := json.Marshal(h.Payload)
	if err != nil {
		return err
	}
	resp, err := http.Post(h.Url, "application/json; charset=UTF-8", bytes.NewReader(p))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if !bytes.Equal(res, []byte("1")) {
		return errors.New(string(res))
	}

	return nil
}

func (h *Teams) NewSection(title string) *teams.Section {
	s := new(teams.Section)
	s.SetTitle(title)
	s.Markdown = true
	return s
}
