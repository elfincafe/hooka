package webhook

type Webhook struct {
	Url   string
	Title string
	Text  string
}

type Payload struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
