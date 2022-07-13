package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Slack struct {
	Webhook
}

func NewSlack(url string) *Slack {
	println("New")
	h := new(Slack)
	h.Url = url
	return h
}

func (h *Slack) SetTitle(title string) {
	println("SetTitle")
	h.Title = title
}

func (h *Slack) SetText(text string) {
	println("SetText")
	h.Text = text
}

func (h *Slack) Send() error {
	pl := new(Payload)
	pl.Title = h.Title
	pl.Text = h.Text
	p, err := json.Marshal(pl)
	if err != nil {
		return err
	}
	resp, err := http.Post(h.Url, "application/json; charset=UTF-8", bytes.NewReader(p))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
