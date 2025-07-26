package adaptive_card

import "strings"

type (
	Image struct {
		Version             float32 `json:"-"`
		Type                string  `json:"type"`
		Url                 string  `json:"url"`
		AltText             string  `json:"altText,omitempty"`
		HorizontalAlignment string  `json:"horizontalAlignment,omitempty"`
		Size                string  `json:"size,omitempty"`
		Style               string  `json:"style,omitempty"`
		Separator           bool    `json:"separator,omitempty"`
		Spacing             string  `json:"spacing,omitempty"`
		Id                  string  `json:"id,omitempty"`
	}
)

func NewImage(url string) *Image {
	i := &Image{
		Version:             1.0,
		Type:                "Image",
		Url:                 url,
		AltText:             "",
		HorizontalAlignment: "",
		Size:                "",
		Style:               "",
		Separator:           false,
		Spacing:             "",
		Id:                  "",
	}
	return i
}

func (i *Image) GetVersion() float32 {
	return i.Version
}

func (i *Image) GetType() string {
	return i.Type
}

func (i *Image) SetId(id string) {
	i.Id = id
}

func (i *Image) SetAltText(altText string) {
	i.AltText = altText
}

func (i *Image) SetSeparator(separator bool) {
	i.Separator = separator
}

func (i *Image) SetHorizontalAlignment(horizontalAlignment string) {
	horizontalAlignment = strings.ToLower(horizontalAlignment)
	switch horizontalAlignment {
	case "left":
		i.HorizontalAlignment = horizontalAlignment
	case "center":
		i.HorizontalAlignment = horizontalAlignment
	case "right":
		i.HorizontalAlignment = horizontalAlignment
	default:
		i.HorizontalAlignment = ""
	}
}

func (i *Image) SetSize(size string) {
	size = strings.ToLower(size)
	switch size {
	case "auto":
		i.Size = size
	case "stretch":
		i.Size = size
	case "small":
		i.Size = size
	case "medium":
		i.Size = size
	case "large":
		i.Size = size
	default:
		i.Size = ""
	}
}

func (i *Image) SetStyle(style string) {
	style = strings.ToLower(style)
	switch style {
	case "default":
		i.Style = "default"
	case "person":
		i.Style = "person"
	default:
		i.Style = ""
	}
}

func (i *Image) SetSpacing(spacing string) {
	spacing = strings.ToLower(spacing)
	switch spacing {
	case "default":
		i.Spacing = spacing
	case "none":
		i.Spacing = spacing
	case "small":
		i.Spacing = spacing
	case "medium":
		i.Spacing = spacing
	case "large":
		i.Spacing = spacing
	case "extralarge":
		i.Spacing = "extraLarge"
	case "padding":
		i.Spacing = spacing
	default:
		i.Spacing = ""
	}
}
