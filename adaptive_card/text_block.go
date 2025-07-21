package adaptive_card

import (
	"strings"
)

type (
	TextBlock struct {
		version             float64
		Type                string `json:"type"`
		Text                string `json:"text"`
		Color               string `json:"color,omitempty"`
		HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
		Subtle              bool   `json:"isSubtle,omitempty"`
		MaxLines            int    `json:"maxLines,omitempty"`
		Size                string `json:"size,omitempty"`
		Weight              string `json:"weight,omitempty"`
		Wrap                bool   `json:"wrap,omitempty"`
		Separator           bool   `json:"separator,omitempty"`
		Spacing             string `json:"spacing,omitempty"`
		Id                  string `json:"id,omitempty"`
	}
)

func NewTextBlock(text string) *TextBlock {
	tb := &TextBlock{
		version:             1.0,
		Type:                "TextBlock",
		Text:                text,
		Color:               "",
		HorizontalAlignment: "",
		Subtle:              false,
		MaxLines:            0,
		Size:                "",
		Weight:              "",
		Wrap:                false,
		Separator:           false,
		Spacing:             "",
		Id:                  "",
	}
	return tb
}

func (tb *TextBlock) GetVersion() float64 {
	return tb.version
}

func (tb *TextBlock) GetType() string {
	return tb.Type
}

func (tb *TextBlock) SetId(id string) {
	tb.Id = id
}

func (tb *TextBlock) SetSpacing(spacing string) {
	spacing = strings.ToLower(spacing)
	switch spacing {
	case "default":
		tb.Spacing = spacing
	case "none":
		tb.Spacing = spacing
	case "small":
		tb.Spacing = spacing
	case "medium":
		tb.Spacing = spacing
	case "large":
		tb.Spacing = spacing
	case "extralarge":
		tb.Spacing = "extraLarge"
	case "padding":
		tb.Spacing = spacing
	default:
		tb.Spacing = ""
	}
}

func (tb *TextBlock) SetSeparator(separator bool) {
	tb.Separator = separator
}

func (tb *TextBlock) SetHorizontalAlignment(horizontalAlignment string) {
	horizontalAlignment = strings.ToLower(horizontalAlignment)
	switch horizontalAlignment {
	case "left":
		tb.HorizontalAlignment = horizontalAlignment
	case "center":
		tb.HorizontalAlignment = horizontalAlignment
	case "right":
		tb.HorizontalAlignment = horizontalAlignment
	default:
		tb.HorizontalAlignment = ""
	}
}

func (tb *TextBlock) SetWrap(wrap bool) {
	tb.Wrap = wrap
}

func (tb *TextBlock) SetMaxLines(maxLines int) {
	if maxLines < 0 {
		maxLines = 0
	}
	tb.MaxLines = maxLines
}

func (tb *TextBlock) SetSize(size string) {
	size = strings.ToLower(size)
	switch size {
	case "default":
		tb.Size = size
	case "small":
		tb.Size = size
	case "medium":
		tb.Size = size
	case "large":
		tb.Size = size
	case "extralarge":
		tb.Size = "extraLarge"
	default:
		tb.Size = ""
	}
}

func (tb *TextBlock) SetWeight(weight string) {
	weight = strings.ToLower(weight)
	switch weight {
	case "default":
		tb.Weight = weight
	case "lighter":
		tb.Weight = weight
	case "bolder":
		tb.Weight = weight
	default:
		tb.Weight = ""
	}
}

func (tb *TextBlock) SetColor(color string) {
	color = strings.ToLower(color)
	switch color {
	case "default":
		tb.Color = color
	case "dark":
		tb.Color = color
	case "light":
		tb.Color = color
	case "accent":
		tb.Color = color
	case "good":
		tb.Color = color
	case "warning":
		tb.Color = color
	case "attention":
		tb.Color = color
	default:
		tb.Color = ""
	}
}

func (tb *TextBlock) SetSubtle(subtle bool) {
	tb.Subtle = subtle
}
