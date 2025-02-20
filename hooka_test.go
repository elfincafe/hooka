package hooka

import "testing"

func TestNew(t *testing.T) {
	// Test Case
	cases := []struct {
		url string
	}{
		{
			url: "https://example.com/webhook",
		},
	}
	// Test
	for i, c := range cases {
		h := New(c.url)
		if h.url != c.url {
			t.Errorf(`[Case%d] URL Result: %v, Expected: %v`, i, h.url, c.url)
		}
	}
}
