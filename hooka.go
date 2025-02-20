package hooka

import (
	"bytes"
	"fmt"
	"net/http"
)

type (
	Payload interface {
		Marshal() ([]byte, error)
	}
	Hooka struct {
		url string
	}
)

func New(url string) *Hooka {
	h := new(Hooka)
	h.url = url
	return h
}

func (h *Hooka) Send(p Payload) error {
	b, err := p.Marshal()
	if err != nil {
		return err
	}
	content := bytes.NewReader(b)
	res, err := http.Post(h.url, "application/json", content)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}
