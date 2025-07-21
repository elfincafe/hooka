package adaptive_card

import (
	"strings"
)

type (
	TextBlock struct {
		Version float32 `json:"-"`
		Type    string  `json:"type"`
		Text    string  `json:"text"`
		Id      string  `json:"id,omitempty"`
		Spacing             string `json:"spacing,omitempty"`
		Separator           bool   `json:"separator,omitempty"`
		HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
		Wrap                bool   `json:"wrap,omitempty"`
		MaxLines            int    `json:"maxLines,omitempty"`
		Size                string `json:"size,omitempty"`
		Weight              string `json:"weight,omitempty"`
		Color               string `json:"color,omitempty"`
		Subtle              bool   `json:"isSubtle,omitempty"`
	}
)

func NewTextBlock(text string) TextBlock {
	tb := TextBlock{
		Version:             1.0,
		Type:                "TextBlock",
		Text:                text,
		Id:                  "",
		Spacing:             "",
		Separator:           false,
		HorizontalAlignment: "",
		Wrap:                false,
		MaxLines:            0,
		Size:                "",
		Weight:              "",
		Color:               "",
		Subtle:              false,
	}
	return tb
}

func (tb TextBlock) GetVersion() float32 {
	return tb.Version
}

func (tb TextBlock) GetType() string {
	return tb.Type
}

func (tb TextBlock) SetId(id string){
	tb.Id = id
}

func (tb TextBlock) SetSpacing(spacing string) {
	spacing = strings.ToLower(spacing)
	switch spacing {
	case "none":
		tb.Spacing = "None"
	case "small":
		tb.Spacing = "Small"
	case "medium":
		tb.Spacing = "Medium"
	case "large":
		tb.Spacing = "Large"
	case "extralarge":
		tb.Spacing = "ExtraLarge"
	case "padding":
		tb.Spacing = "Padding"
	default:
		tb.Spacing = ""
	}
}

func (tb TextBlock) SetSeparator(separator bool) {
	tb.Separator = separator
}

func (tb TextBlock) SetHorizontalAlignment(horizontalAlignment string) {
	horizontalAlignment = strings.ToLower(horizontalAlignment)
	switch horizontalAlignment {
	case "left":
		tb.HorizontalAlignment = "Left"
	case "center":
		tb.HorizontalAlignment = "Center"
	case "right":
		tb.HorizontalAlignment = "Right"
	default:
		tb.HorizontalAlignment = ""
	}
}

func (tb TextBlock) SetWrap(wrap bool) {
	tb.Wrap = wrap
}

func (tb TextBlock) SetMaxLines(maxLines int) {
	if maxLines < 0 {
		maxLines = 0
	}
	tb.MaxLines = maxLines
}

func (tb TextBlock) SetSize(size string) {
	size = strings.ToLower(size)
	switch size {
	case "small":
		tb.Size = "Small"
	case "default":
		tb.Size = "Default"
	case "medium":
		tb.Size = "Medium"
	case "large":
		tb.Size = "Large"
	case "extralarge":
		tb.Size = "ExtraLarge"
	default:
		tb.Size = ""
	}
}

func (tb TextBlock) SetWeight(weight string) {
	weight = strings.ToLower(weight)
	switch weight {
	case "lighter":
		tb.Weight = "Lighter"
	case "default":
		tb.Weight = "Default"
	case "bolder":
		tb.Weight = "Bolder"
	default:
		tb.Weight = ""
	}
}

func (tb TextBlock) SetColor(color string) {
	color = strings.ToLower(color)
	switch color {
	case "dark":
		tb.Color = "Dark"
	case "light":
		tb.Color = "Light"
	case "accent":
		tb.Color = "Accent"
	case "good":
		tb.Color = "Good"
	case "warning":
		tb.Color = "Warning"
	case "attention":
		tb.Color = "Attention"
	default:
		tb.Color = ""
	}
}

func (tb TextBlock) SetSubtle(subtle bool) {
	tb.Subtle = subtle
}
