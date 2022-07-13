package webhook

import (
	"testing"
	"time"
)

const (
	TEAMS_WEBHOOK_URL = ""
	TEAMS_IMAGE_URL   = ""
)

func TestTeams(t *testing.T) {
	h := NewTeams(TEAMS_WEBHOOK_URL)
	h.SetSummary("Test Summary")
	h.SetColor("ff5555")
	s := h.NewSection("Contact has come!")
	s.SetSubtitle("sub sub sub sub")
	s.SetImage(TEAMS_IMAGE_URL)
	s.Set("Occured Date", time.Now().Format(`2006/01/02 15:04:05 MST`))
	s.Set("Message", "File Is Not Found.")
	s.Set("Status", "Pendding")
	h.SetSection(s)
	err := h.Send()
	if err != nil {
		t.Errorf(err.Error())
	}
}
