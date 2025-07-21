package adaptive_card

import "strings"

type (
	Image struct {
		Version             float32 `json:"-"`
		Type                string  `json:"type"`
		Url                 string  `json:"url"`
		Id                  string  `json:"id,omitempty"`
		AltText             string  `json:"altText,omitempty"`
		Spacing             string  `json:"spacing,omitempty"`
		Separator           bool    `json:"separator,omitempty"`
		HorizontalAlignment string  `json:"horizontalAlignment,omitempty"`
		Size                string  `json:"size,omitempty"`
		Style               string  `json:"style,omitempty"`
	}
)

func NewImage(url string) Image {
	i := Image{
		Version:             1.0,
		Type:                "Image",
		Id:                  "",
		Url:                 url,
		AltText:             "",
		Spacing:             "",
		Separator:           false,
		HorizontalAlignment: "",
		Size:                "",
		Style:               "",
	}
	return i
}

func (i Image) GetVersion() float32 {
	return i.Version
}

func (i Image) GetType() string {
	return i.Type
}

func (i Image) SetId(id string) {
	i.Id = id
}

func (i Image) SetAltText(altText string) {
	i.AltText = altText
}

func (i Image) SetSpacing(spacing string) {
	spacing = strings.ToLower(spacing)
	switch spacing {
	case "none":
		i.Spacing = "None"
	case "small":
		i.Spacing = "Small"
	case "medium":
		i.Spacing = "Medium"
	case "large":
		i.Spacing = "Large"
	case "extralarge":
		i.Spacing = "ExtraLarge"
	case "padding":
		i.Spacing = "Padding"
	default:
		i.Spacing = ""
	}
}

func (i Image) SetSeparator(separator bool) {
	i.Separator = separator
}

func (i Image) SetHorizontalAlignment(horizontalAlignment string) {
	horizontalAlignment = strings.ToLower(horizontalAlignment)
	switch horizontalAlignment {
	case "left":
		i.HorizontalAlignment = "Left"
	case "center":
		i.HorizontalAlignment = "Center"
	case "right":
		i.HorizontalAlignment = "Right"
	default:
		i.HorizontalAlignment = ""
	}
}

func (i Image) SetSize(size string) {
	size = strings.ToLower(size)
	switch size {
	case "stretch":
		i.Size = "Stretch"
	case "small":
		i.Size = "Small"
	case "medium":
		i.Size = "Medium"
	case "large":
		i.Size = "Large"
	default:
		i.Size = ""
	}
}

func (i Image) SetStyle(style string) {
	style = strings.ToLower(style)
	switch style {
	case "person":
		i.Style = "Person"
	default:
		i.Style = ""
	}
}
