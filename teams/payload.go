package teams

type Payload struct {
	Type       string     `json:"@type"`
	Context    string     `json:"@context"`
	ThemeColor string     `json:"themeColor"`
	Summary    string     `json:"summary"`
	Sections   []*Section `json:"sections"`
}

type Section struct {
	ActivityTitle    string  `json:"activityTitle"`
	ActivitySubtitle string  `json:"activitySubtitle,omitempty"`
	ActivityImage    string  `json:"activityImage,omitempty"`
	Facts            []*Fact `json:"facts"`
	Markdown         bool    `json:"markdown,omitempty"`
}

type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (s *Section) SetTitle(title string) {
	s.ActivityTitle = title
}

func (s *Section) SetSubtitle(subtitle string) {
	s.ActivitySubtitle = subtitle
}

func (s *Section) SetImage(url string) {
	s.ActivityImage = url
}

func (s *Section) Set(name, value string) {
	f := new(Fact)
	f.Name = name
	f.Value = value
	s.Facts = append(s.Facts, f)
}
