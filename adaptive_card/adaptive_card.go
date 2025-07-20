package adaptive_card

type (
	Element interface {
		GetVersion() string
		GetType() string
	}
	Container interface {
	}
)
