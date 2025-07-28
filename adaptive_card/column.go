package adaptive_card

import (
	"encoding/json"
	"strconv"
	"strings"
)

type (
	Column struct {
		version   float64
		Type      string    `json:"type"`
		Id        string    `json:"id,omitempty"`
		Items     []Element `json:"items"`
		Separator bool      `json:"separator,omitempty"`
		Spacing   string    `json:"spacing,omitempty"`
		Style     string    `json:"style,omitempty"`
		width     string
	}
	ColumnStr struct {
		Column
		Width string `json:"width,omitempty"`
	}
	ColumnInt struct {
		Column
		Width int `json:"width,omitempty"`
	}
)

func NewColumn() *Column {
	c := &Column{
		version:   1.0,
		Type:      "Column",
		Id:        "",
		Items:     []Element{},
		Separator: false,
		Spacing:   "",
		Style:     "",
		width:     "",
	}
	return c
}

func (c *Column) GetVersion() float64 {
	return c.version
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
			c.width = width
		case "stretch":
			c.width = width
		default:
			c.width = ""
		}
	} else {
		if widthInt > 0 {
			c.width = width
		} else {
			c.width = "0"
		}
	}
}

// func (c *Column) SetWidth(width string) {
// 	width = strings.ToLower(width)
// 	widthInt, err := strconv.Atoi(width)
// 	if err != nil {
// 		switch width {
// 		case "auto":
// 			c.Width = width
// 		case "stretch":
// 			c.Width = width
// 		default:
// 			c.Width = ""
// 		}
// 	} else {
// 		c.Width = widthInt
// 	}
// }

func (c *Column) MarshalJSON() ([]byte, error) {
	var data []byte
	var err error
	var widthInt int
	widthInt, err = strconv.Atoi(c.width)
	if err != nil {
		col := ColumnStr{Column: *c, Width: c.width}
		data, err = json.Marshal(col)
	} else {
		col := ColumnInt{Column: *c, Width: widthInt}
		data, err = json.Marshal(col)
	}
	return data, err
}
