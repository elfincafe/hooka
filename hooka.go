package hooka

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"strings"
)

type (
	service uint8
	Hooka   interface {
		Send([]byte) error
	}
)

func New(uri string) (Hooka, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	// Service Type
	if strings.Contains(u.Host, "azure.com") {
		return &Teams{uri: u}, nil
	}
	return nil, errors.New("Specified URL isn't supported")
}

func send(data []byte, uri *url.URL) (*http.Response, error) {
	// Request
	req, err := http.NewRequest("POST", uri.String(), bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	// Header
	headers := http.Header{}
	headers.Add("Content-Type", "application/json")
	req.Header = headers
	// Send
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
