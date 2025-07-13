package adaptive_card

import (
	"encoding/json"
	"fmt"
)

type (
	TextBlock struct {
		Version string `json:"-"`
		Type    string `json:"type"`
		Text    string `json:"text"`
		// Layout
		Spacing             string `json:"spacing,omitempty"`
		Separator           bool   `json:"separator,omitempty"`
		HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
		Wrap                bool   `json:"wrap,omitempty"`
		MaxLines            int    `json:"maxLines,omitempty"`
		// Style
		Size   string `json:"size,omitempty"`
		Weight string `json:"weight,omitempty"`
		Color  string `json:"color,omitempty"`
		Subtle bool   `json:"isSubtle,omitempty"`
	}
)

func NewTextBlock(text string) *TextBlock {
	tb := new(TextBlock)
	tb.Version = "1.0"
	tb.Type = "TextBlock"
	tb.Text = text
	return tb
}

func (tb *TextBlock) GetType() string {
	return tb.Type
}

func (tb *TextBlock) SetSize(size Size) {
	switch size {
	case SizeSmall:
		tb.Size = "Small"
	case SizeDefault:
		tb.Size = "Default"
	case SizeMedium:
		tb.Size = "Medium"
	case SizeLarge:
		tb.Size = "Large"
	case SizeExtraLarge:
		tb.Size = "ExtraLarge"
	}
}

func (tb *TextBlock) SetWeight(weight Weight) {
	switch weight {
	case WeightLighter:
		tb.Weight = "Lighter"
	case WeightDefault:
		tb.Weight = "Default"
	case WeightBolder:
		tb.Weight = "Bolder"
	}
}

func (tb *TextBlock) SetColor(color Color) {
	switch color {
	case ColorDark:
		tb.Color = "Dark"
	case ColorLight:
		tb.Color = "Light"
	case ColorAccent:
		tb.Color = "Accent"
	case ColorGood:
		tb.Color = "Good"
	case ColorWarning:
		tb.Color = "Warning"
	case ColorAttention:
		tb.Color = "Attention"
	}
}

func (tb *TextBlock) SetSubtle(subtle bool) {
	tb.Subtle = subtle
	fmt.Println(json.Marshal((tb)))
}

func (tb *TextBlock) SetSpacing(spacing Spacing) {
	switch spacing {
	case SpacingNone:
		tb.Spacing = "None"
	case SpacingSmall:
		tb.Spacing = "Small"
	case SpacingDefault:
		tb.Spacing = "Default"
	case SpacingMedium:
		tb.Spacing = "Medium"
	case SpacingLarge:
		tb.Spacing = "Large"
	case SpacingExtraLarge:
		tb.Spacing = "ExtraLarge"
	case SpacingPadding:
		tb.Spacing = "Padding"
	}
}

func (tb *TextBlock) SetHorizontalAlignment(ha HorizontalAlignment) {
	switch ha {
	case HorizontalAlignmentLeft:
		tb.HorizontalAlignment = "Left"
	case HorizontalAlignmentCenter:
		tb.HorizontalAlignment = "Center"
	case HorizontalAlignmentRight:
		tb.HorizontalAlignment = "Right"
	}
}

func (tb *TextBlock) SetSeparator(separator bool) {
	tb.Separator = separator
}

func (tb *TextBlock) SetWrap(wrap bool) {
	tb.Wrap = wrap
}

func (tb *TextBlock) SetMaxLines(maxLines int) {
	tb.MaxLines = maxLines
}
