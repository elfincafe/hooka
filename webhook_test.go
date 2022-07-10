package webhook

import (
	"github.com/elfincafe/webhook/teams"
)

func teamsTest() {
	t := teams.New("test")
	t.SetTitle("タイトル")
	t.SetText("テキスト")
	t.Send()
}
