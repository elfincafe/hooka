package adaptive_card

type (
	Column struct {
	}
	ColumnSet struct {
		Version             string
		Type                string
		Columns             Column
		Style               string
		Spacing             string
		Separator           bool
		HorizontalAlignment string
	}
)

func NewColumnSet() *ColumnSet {
	cs := &ColumnSet{
		Version:             "1.0",
		Type:                "ColumnSet",
		Style:               "",
		Spacing:             "",
		Separator:           false,
		HorizontalAlignment: "",
	}
	return cs
}

func (cs *ColumnSet) SetElement(elem Element) {

}
