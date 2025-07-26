package adaptive_card

type (
	Fact struct {
		Title string `json:"title"`
		Value string `json:"value"`
	}
)

func NewFact(title, value string) *Fact {
	f := &Fact{
		Title: title,
		Value: value,
	}
	return f
}
