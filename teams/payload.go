package teams

type Payload struct {
	Title      string     `json:"title,omitempty"`
	Text       string     `json:"text,omitempty"`
	Type       string     `json:"@type,omitempty"`
	Context    string     `json:"@context,omitempty"`
	ThemeColor string     `json:"themeColor,omitempty"`
	Summary    string     `json:"summary,omitempty"`
	Sections   []*Section `json:"sections,omitempty"`
}

type Section struct {
	ActivityTitle    string  `json:"activityTitle,omitempty"`
	ActivitySubtitle string  `json:"activitySubtitle,omitempty"`
	ActivityImage    string  `json:"activityImage,omitempty"`
	Facts            []*Fact `json:"facts,omitempty"`
	Markdown         bool    `json:"markdown,omitempty"`
}

type Fact struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (p *Payload) SetTitle(title string) {
	p.Title = title
}

func (p *Payload) SetText(text string) {
	p.Text = text
}

func (p *Payload) SetSummary(s string) {
	p.Summary = s
}

func (p *Payload) SetSection(s *Section) {
	p.Sections = append(p.Sections, s)
}

func (p *Payload) SetColor(color string) {
	p.ThemeColor = color
}

func (p *Payload) NewSection(title string) *Section {
	s := new(Section)
	s.SetTitle(title)
	s.Markdown = true
	p.Sections = append(p.Sections, s)
	return s
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

func (s *Section) SetFact(name, value string) {
	f := new(Fact)
	f.Name = name
	f.Value = value
	s.Facts = append(s.Facts, f)
}
