package adaptive_card

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestTextBlockSetSpacing(t *testing.T) {
	// tb := &TextBlock{Version: "1.0", Text: "test", Type: "TextBlock"}
	tb := NewTextBlock("test")
	tb.SetSpacing(SpacingExtraLarge)
	j, _ := json.Marshal(tb)
	fmt.Println(string(j))
}

func TestTextBlockSetSubtle(t *testing.T) {
	tb := &TextBlock{Version: "1.0", Text: "test", Type: "TextBlock"}
	tb.SetSubtle(false)
	j, _ := json.Marshal(tb)
	fmt.Println(string(j))
}
