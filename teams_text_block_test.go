package hooka

import (
	"testing"
)

func TestTextBlockMarshal(t *testing.T) {
	tb := NewTextBlock("test")
	tb.SetColor(ColorDark)
	j, err := tb.Marshal()
	t.Errorf("%v\n%v", string(j), err)
}
