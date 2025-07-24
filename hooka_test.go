package hooka

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	// Test Case
	cases := []struct {
		uri     string
		service service
		none    bool
	}{
		{
			uri:     "https://azure.com:443/workflows",
			service: Teams,
			none:    false,
		},
		{
			uri:     "https://slack.com/workflow",
			service: Slack,
			none:    false,
		},
		{
			uri:     "https://example.com/webhook",
			service: Unknown,
			none:    false,
		},
		{
			uri:     "webhook_url",
			service: Unknown,
			none:    true,
		},
	}
	// Test
	for i, c := range cases {
		h := New(c.uri)
		if h.uri.String() != c.uri {
			t.Errorf(`[Case%d] URL Result: %v, Expected: %v`, i, h.uri, c.uri)
		}
		if h.service != c.service {
			t.Errorf(`[Case%d] Service Result: %v, Expected: %v`, i, h.service, c.service)
		}
		if c.none && h.uri.Hostname() != "" {
			t.Errorf(`[Case%d] URL Nil Expected: %v`, i, h.uri.Hostname())
		}
		if !c.none && fmt.Sprintf("%T", h) != "*hooka.Hooka" {
			t.Errorf(`[Case%d] *Hooka Expected: %v`, i, fmt.Sprintf("%T", h))
		}
	}
}
