package adaptive_card

type(
	Column struct{
		Version float32
		Type string
	}
)

func NewColumn(elem Element)Column{
	c := Column{
		Version: 1.0,
		Type:"Column",
	}
	return c
}

func (c Column)GetVersion()float32{
	return c.Version
}

func (c Column) GetType()string{
	return c.Type
}
