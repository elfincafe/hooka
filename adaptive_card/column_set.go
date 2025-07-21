package adaptive_card

import (
	"strings"
)

type (
	ColumnSet struct {
		Version             float32
		Type                string
		Id string
		Columns             []Column
		Style               string
		Spacing             string
		Separator           bool
		HorizontalAlignment string
	}
)

func NewColumnSet() ColumnSet {
	cs := ColumnSet{
		Version:             1.0,
		Type:                "ColumnSet",
		Id: "",
		Columns: []Column{},
		Style:               "",
		Spacing:             "",
		Separator:           false,
		HorizontalAlignment: "",
	}
	return cs
}

func (cs ColumnSet)GetVersion()float32{
	return cs.Version
}

func (cs ColumnSet)GetType()string{
	return cs.Type
}

func (cs ColumnSet)SetId(id string){
	cs.Id = id
}

func (cs ColumnSet) AppendColumn(column Column) {
	cs.Columns = append(cs.Columns, column)
}

func (cs ColumnSet) SetStyle(style string) {
	style = strings.ToLower(style)
	switch style {
	case "default":
		cs.Style="default"
	case "emphasis":
		cs.Style = "emphasis"
	default:
		cs.Style = ""
	}
}

func (cs ColumnSet)SetSpacing{

}