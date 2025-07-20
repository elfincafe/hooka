package adaptive_card

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTextBlockSetSpacing(t *testing.T) {
	tb := NewTextBlock("test")
	tb.SetSpacing("ExtraLarge")
	j, _ := json.Marshal(tb)
	fmt.Println(string(j))
}

func TestTextBlockSetSubtle(t *testing.T) {
	tb := NewTextBlock("test")
	tb.SetSubtle(false)
	j, _ := json.Marshal(tb)
	fmt.Println(string(j))
}
