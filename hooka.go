package hooka

import (
	"bytes"
	"fmt"
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

func parseUri(uri, domain string) (*url.URL, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	if !strings.Contains(u.Host, domain) {
		return nil, fmt.Errorf(`the passed uri doesn't contain "%s"`, domain)
	}
	return u, nil
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
