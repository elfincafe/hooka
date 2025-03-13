package hooka

import "encoding/json"

const (
	FontTypeDefault           FontType            = 1
	FontTypeMonospace         FontType            = 2
	SizeSmall                 Size                = 1
	SizeDefault               Size                = 2
	SizeMedium                Size                = 3
	SizeLarge                 Size                = 4
	SizeExtraLarge            Size                = 5
	WeightLighter             Weight              = 1
	WeightDefault             Weight              = 2
	WeightBolder              Weight              = 3
	ColorDefault              Color               = 1
	ColorDark                 Color               = 2
	ColorLight                Color               = 3
	ColorAccent               Color               = 4
	ColorGood                 Color               = 5
	ColorWarning              Color               = 6
	ColorAttention            Color               = 7
	SpacingNone               Spacing             = 0
	SpacingSmall              Spacing             = 1
	SpacingDefault            Spacing             = 2
	SpacingMedium             Spacing             = 3
	SpacingLarge              Spacing             = 4
	SpacingExtraLarge         Spacing             = 5
	SpacingPadding            Spacing             = 6
	HorizontalAlignmentNone   HorizontalAlignment = 0
	HorizontalAlignmentLeft   HorizontalAlignment = 1
	HorizontalAlignmentCenter HorizontalAlignment = 2
	HorizontalAlignmentRight  HorizontalAlignment = 3
	HeightNone                Height              = 0
	HeightAutomatic           Height              = 1
	HeightStretch             Height              = 2
)

type (
	FontType            uint8
	Size                uint8
	Weight              uint8
	Color               uint8
	Spacing             uint8
	HorizontalAlignment uint8
	Height              uint8
	TextBlock           struct {
		Type                string `json:"type"`
		Text                string `json:"text"`
		FontType            string `json:"fontType,omitempty"`
		Size                string `json:"size,omitempty"`
		Weight              string `json:"weight,omitempty"`
		Color               string `json:"color,omitempty"`
		Subtle              bool   `json:"isSubtle,omitempty"`
		Spacing             string `json:"spacing,omitempty"`
		Separator           bool   `json:"separator,omitempty"`
		HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
		Height              string `json:"height,omitempty"`
		Wrap                bool   `json:"wrap,omitempty"`
		MaxLines            int    `json:"maxLines,omitempty"`
	}
)

func NewTextBlock(text string) *TextBlock {
	tb := new(TextBlock)
	tb.Type = "TextBlock"
	tb.Text = text
	return tb
}

func (tb *TextBlock) GetType() string {
	return tb.Type
}

func (tb *TextBlock) SetText(text string) {
	tb.Text = text
}

func (tb *TextBlock) SetFontType(fontType FontType) {
	switch fontType {
	case FontTypeMonospace:
		tb.FontType = "Monospace"
	case FontTypeDefault:
		tb.FontType = "Default"
	}
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

func (tb *TextBlock) SetHeight(height Height) {
	switch height {
	case HeightAutomatic:
		tb.Height = "Automatic"
	case HeightStretch:
		tb.Height = "Stretch"
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

func (tb *TextBlock) Marshal() ([]byte, error) {
	return json.Marshal(tb)
}
