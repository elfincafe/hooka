package adaptive_card

type (
	ImageSet struct{
		Version float32 `json:"-"`
		Type string `json:"type"`
		Id string `json:"id,omitempty"`
	}
)

func NewImageSet() ImageSet{
	is := ImageSet{
		Version:1.0,
		Type: "ImageSet",
		Id: "",
	}
	return is
}

func (is ImageSet)GetVersion()float32{
	return is.Version
}

func (is ImageSet)GetType()string{
	return is.Type
}

func (is ImageSet)SetId(id string){
	is.Id = id
}
