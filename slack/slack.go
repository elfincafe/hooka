package slack

import (
	"github.com/elfincafe/webhook"
)

type Webhook struct {
	webhook.Webhook
}

func New(url string) *Webhook {
	println("New")
	w := new(Webhook)
	w.Url = url
	return w
}

func (w *Webhook) SetTitle(title string) {
	println("SetTitle")
	w.Title = title
}

func (w *Webhook) SetText(text string) {
	println("SetText")
	w.Text = text
}

func (w *Webhook) Send() bool {
	println("Send")
	ret := false
	return ret
}
