package adaptive_card

import (
	"strconv"
	"strings"
)

type (
	Column struct {
		Version   float32
		Type      string    `json:"type"`
		Id        string    `json:"id,omitempty"`
		Items     []Element `json:"items"`
		Separator bool      `json:"separator,omitempty"`
		Spacing   string    `json:"spacing,omitempty"`
		Style     string    `json:"style,omitempty"`
		Width     string    `json:"width,omitempty"`
		WidthInt  int       `json:"width,omitempty"`
	}
)

func NewColumn() *Column {
	c := &Column{
		Version:   1.0,
		Type:      "Column",
		Id:        "",
		Items:     []Element{},
		Separator: false,
		Spacing:   "",
		Style:     "",
		Width:     "",
		WidthInt:  0,
	}
	return c
}

func (c *Column) GetVersion() float32 {
	return c.Version
}

func (c *Column) GetType() string {
	return c.Type
}

func (c *Column) SetId(id string) {
	c.Id = id
}

func (c *Column) Append(item Element) {
	c.Items = append(c.Items, item)
}

func (c *Column) SetSeparator(separator bool) {
	c.Separator = separator
}

func (c *Column) SetSpacing(spacing string) {
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

func (c *Column) SetStyle(style string) {
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

func (c *Column) SetWidth(width string) {
	width = strings.ToLower(width)
	widthInt, err := strconv.Atoi(width)
	if err != nil {
		switch width {
		case "auto":
			c.Width = width
			c.WidthInt = 0
		case "stretch":
			c.Width = width
			c.WidthInt = 0
		default:
			c.Width = ""
			c.WidthInt = 0
		}
	} else {
		c.Width = ""
		if widthInt > 0 {
			c.WidthInt = widthInt
		} else {
			c.WidthInt = 0
		}
	}
}
