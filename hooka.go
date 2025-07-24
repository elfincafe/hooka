package hooka

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	serviceName string  = "Hooka"
	Unknown     service = 0
	Teams       service = 1
	Slack       service = 2
)

type (
	service uint8
	Message struct {
		typ  int
		data []byte
	}
	Hooka struct {
		service service
		uri     *url.URL
		message Message
	}
)

func New(uri string) *Hooka {
	h := new(Hooka)
	u, err := url.Parse(uri)
	if err != nil {
		h.service = Unknown
		return h
	}
	// Service Type
	if strings.Contains(u.Host, "azure.com") {
		h.service = Teams
	} else if strings.Contains(u.Host, "slack.com") {
		h.service = Slack
	} else {
		h.service = Unknown
	}
	h.uri = u
	return h
}

func (h *Hooka) Send(data []byte) error {
	// Request
	req, err := http.NewRequest("POST", h.uri.String(), bytes.NewReader(data))
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
	fmt.Println(res)
	if res.StatusCode >= 400 {
		return fmt.Errorf("[%d] %s", res.StatusCode, res.Status)
	}
	content, _ := io.ReadAll(res.Body)
	fmt.Println(string(content))
	return nil
}
