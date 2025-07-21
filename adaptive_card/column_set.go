package adaptive_card

import "strings"

type (
	ColumnSet struct {
		version             float64
		Type                string    `json:"type"`
		Id                  string    `json:"id,omitempty"`
		Columns             []*Column `json:"columns,omitempty"`
		HorizontalAlignment string    `json:"horizontalAlignment,omitempty"`
		Spacing             string    `json:"spacing,omitempty"`
		Separator           bool      `json:"separator,omitempty"`
	}
)

func NewColumnSet() *ColumnSet {
	cs := &ColumnSet{
		version:             1.0,
		Type:                "ColumnSet",
		Id:                  "",
		Columns:             []*Column{},
		Spacing:             "",
		Separator:           false,
		HorizontalAlignment: "",
	}
	return cs
}

func (cs *ColumnSet) GetVersion() float64 {
	return cs.version
}

func (cs *ColumnSet) GetType() string {
	return cs.Type
}

func (cs *ColumnSet) SetId(id string) {
	cs.Id = id
}

func (cs *ColumnSet) Append(column *Column) {
	cs.Columns = append(cs.Columns, column)
	if column.GetVersion() > cs.GetVersion() {
		cs.version = column.GetVersion()
	}
}

func (cs *ColumnSet) SetSpacing(spacing string) {
	spacing = strings.ToLower(spacing)
	switch spacing {
	case "default":
		cs.Spacing = spacing
	case "none":
		cs.Spacing = spacing
	case "small":
		cs.Spacing = spacing
	case "medium":
		cs.Spacing = spacing
	case "large":
		cs.Spacing = spacing
	case "extralarge":
		cs.Spacing = "extraLarge"
	case "padding":
		cs.Spacing = spacing
	default:
		cs.Spacing = ""
	}
}

func (cs *ColumnSet) SetSeparator(separator bool) {
	cs.Separator = separator
}

func (cs *ColumnSet) SetHorizontalAlignment(horizontalAlignment string) {
	horizontalAlignment = strings.ToLower(horizontalAlignment)
	switch horizontalAlignment {
	case "left":
		cs.HorizontalAlignment = horizontalAlignment
	case "center":
		cs.HorizontalAlignment = horizontalAlignment
	case "right":
		cs.HorizontalAlignment = horizontalAlignment
	default:
		cs.HorizontalAlignment = ""
	}
}
