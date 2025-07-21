package adaptive_card

type(
	Container struct {
		Version float32 `json:"-"`
		Type string `json:"type"`
		Id string `json:"id,omitempty"`
	}
)

func NewContainer()Container{
	c := Container{
		Version: 1.0,
		Type: "Container",
		Id :"",
	}
	return c
}