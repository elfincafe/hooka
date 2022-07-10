package teams

import (
	"encoding/json"
	"fmt"

	"github.com/elfincafe/webhook"
)

type Webhook struct {
	webhook.Webhook
}

type payload struct {
	title string `json:"title"`
	text  string `json:"text"`
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

func (w *Webhook) Send() error {
	println("Send")

	pl := payload{
		title: w.Title,
		text:  w.Text,
	}
	pl.title = w.Title
	pl.text = w.Text
	fmt.Println(pl)
	p, err := json.Marshal(pl)
	if err != nil {
		return err
	}
	fmt.Println(string(p))
	// resp, err := http.Post(w.Url, "application/json; charset=UTF-8", bytes.NewReader(p))
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()
	return nil
}
