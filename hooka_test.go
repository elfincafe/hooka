package webhook

import (
	"testing"
	"time"
)

const (
	TEAMS_WEBHOOK_URL = ""
	TEAMS_IMAGE_URL   = ""
	SLACK_WEBHOOK_URL = ""
)

func TestTeams(t *testing.T) {
	if len(TEAMS_WEBHOOK_URL) == 0 {
		t.Errorf("Webhook URL NOT Found.")
	}
	h := NewTeams(TEAMS_WEBHOOK_URL)
	p := h.Payload
	p.SetTitle("System Report")
	p.SetText("The following reports have come. Please ensure.")
	p.SetSummary("Report")
	p.SetColor("ff9955")
	s1 := p.NewSection("Error")
	s1.SetSubtitle("sub sub sub sub")
	s1.SetImage(TEAMS_IMAGE_URL)
	s1.SetFact("Occured Date", time.Now().Format(`2006/01/02 15:04:05 MST`))
	s1.SetFact("Message", "File Is Not Found.")
	s1.SetFact("Status", "Pendding")
	s2 := p.NewSection("Success")
	s2.SetSubtitle("sub sub sub sub")
	s2.SetImage(TEAMS_IMAGE_URL)
	s2.SetFact("Occured Date", time.Now().Format(`2006/01/02 15:04:05 MST`))
	s2.SetFact("Message", "System is running.")
	s2.SetFact("Status", "Running")
	err := h.Send()
	if err != nil {
		t.Errorf(err.Error())
	}
}

// func TestSlack(t *testing.T) {
// 	if len(SLACK_WEBHOOK_URL) == 0 {
// 		t.Errorf("Webhook URL NOT Found.")
// 	}
// 	h := NewSlack(SLACK_WEBHOOK_URL)
// 	h.SetText()
// }
