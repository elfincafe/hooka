package adaptive_card

type (
	FactSet struct {
		Version float32 `json:"-"`
		Type    string  `json:"type"`
		Id      string  `json:"id,omitempty"`
	}
)

func NewFactSet() FactSet {
	fs := FactSet{
		Version: 1.0,
		Type:    "FactSet",
		Id:      "",
	}
	return fs
}

func (fs FactSet) GetVersion() float32 {
	return fs.Version
}

func (fs FactSet) GetType() string {
	return fs.Type
}

func (fs FactSet) SetId(id string) {
	fs.Id = id
}
