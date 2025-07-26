package adaptive_card

import "strings"

type (
	Container struct {
		Version   float32   `json:"-"`
		Type      string    `json:"type"`
		Id        string    `json:"id,omitempty"`
		Items     []Element `json:"items"`
		Style     string    `json:"style,omitempty"`
		Separator bool      `json:"separator,omitempty"`
		Spacing   string    `json:"spacing,omitempty"`
	}
)

func NewContainer() *Container {
	c := &Container{
		Version:   1.0,
		Type:      "Container",
		Id:        "",
		Items:     []Element{},
		Style:     "",
		Separator: false,
		Spacing:   "",
	}
	return c
}

func (c *Container) GetVersion() float32 {
	return c.Version
}

func (c *Container) GetType() string {
	return c.Type
}

func (c *Container) SetId(id string) {
	c.Id = id
}

func (c *Container) Append(item Element) {
	c.Items = append(c.Items, item)
}

func (c *Container) SetStyle(style string) {
	style = strings.ToLower(style)
	switch style {
	case "default":
		c.Style = style
	case "emphasis":
		c.Style = style
	default:
		c.Style = ""
	}
}

func (c *Container) SetSeparator(separator bool) {
	c.Separator = separator
}

func (c *Container) SetSpacing(spacing string) {
	spacing = strings.ToLower(spacing)
	switch spacing {
	case "default":
		c.Spacing = spacing
	case "none":
		c.Spacing = spacing
	case "small":
		c.Spacing = spacing
	case "medium":
		c.Spacing = spacing
	case "large":
		c.Spacing = spacing
	case "extralarge":
		c.Spacing = "extraLarge"
	case "padding":
		c.Spacing = spacing
	default:
		c.Spacing = ""
	}
}
