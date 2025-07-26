package adaptive_card

import "strings"

type (
	ImageSet struct {
		Version             float32  `json:"-"`
		Type                string   `json:"type"`
		Id                  string   `json:"id,omitempty"`
		Images              []*Image `json:"images"`
		ImageSize           string   `json:"image_size,omitempty"`
		Spacing             string   `json:"spacing,omitempty"`
		Separator           bool     `json:"separator,omitempty"`
		HorizontalAlignment string   `json:"horizontalAlignment,omitempty"`
	}
)

func NewImageSet() *ImageSet {
	is := &ImageSet{
		Version:   1.0,
		Type:      "ImageSet",
		Id:        "",
		Images:    []*Image{},
		ImageSize: "",
		Separator: false,
		Spacing:   "",
	}
	return is
}

func (is *ImageSet) GetVersion() float32 {
	return is.Version
}

func (is *ImageSet) GetType() string {
	return is.Type
}

func (is *ImageSet) SetId(id string) {
	is.Id = id
}

func (is *ImageSet) Append(image *Image) {
	is.Images = append(is.Images, image)
}

func (is *ImageSet) SetImageSize(imageSize string) {
	imageSize = strings.ToLower(imageSize)
	switch imageSize {
	case "auto":
		is.ImageSize = imageSize
	case "stretch":
		is.ImageSize = imageSize
	case "small":
		is.ImageSize = imageSize
	case "medium":
		is.ImageSize = imageSize
	case "large":
		is.ImageSize = imageSize
	default:
		is.ImageSize = ""
	}
}

func (is *ImageSet) SetSpacing(spacing string) {
	spacing = strings.ToLower(spacing)
	switch spacing {
	case "default":
		is.Spacing = spacing
	case "none":
		is.Spacing = spacing
	case "small":
		is.Spacing = spacing
	case "medium":
		is.Spacing = spacing
	case "large":
		is.Spacing = spacing
	case "extralarge":
		is.Spacing = "extraLarge"
	case "padding":
		is.Spacing = spacing
	default:
		is.Spacing = ""
	}
}

func (is *ImageSet) SetSeparator(separator bool) {
	is.Separator = separator
}
