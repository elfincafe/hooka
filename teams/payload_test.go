package teams

/**
 * About Cards
 * https://learn.microsoft.com/ja-jp/microsoftteams/platform/task-modules-and-cards/cards-and-task-modules
 */
import (
	"reflect"
	"testing"
)

func TestTeamsNew(t *testing.T) {
	cases := []struct {
		expected string
	}{
		{
			"*teams.Payload",
		},
	}
	for k, v := range cases {
		p := New()
		typ := reflect.TypeOf(p).String()
		if typ != v.expected {
			t.Errorf("[Case%d] %s (%s)", k, typ, v.expected)
		}
	}
}

func TestTeamsMessage(t *testing.T) {
	cases := []struct {
		title    string
		text     string
		expected string
	}{
		{
			"",
			"",
			`{"@type":"MessageCard"}`,
		},
		{
			"Test Title",
			"Text Text Body",
			`{"title":"Test Title","text":"Text Text Body","@type":"MessageCard"}`,
		},
	}
	for k, v := range cases {
		p := New()
		p.Title = v.title
		p.Text = v.text
		s := p.Message()
		if s != v.expected {
			t.Errorf("[Case%d] %s (%s)", k, s, v.expected)
		}
	}
}
