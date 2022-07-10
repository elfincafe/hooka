package webhook

type Webhook struct {
	Url   string
	Title string `json:"title"`
	Text  string `json:"text"`
}
