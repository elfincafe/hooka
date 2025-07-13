package adaptive_card

type (
	Size   uint8
	Weight uint8
	Color  uint8
)

const (
	SizeSmall      Size   = 1
	SizeDefault    Size   = 2
	SizeMedium     Size   = 3
	SizeLarge      Size   = 4
	SizeExtraLarge Size   = 5
	WeightLighter  Weight = 1
	WeightDefault  Weight = 2
	WeightBolder   Weight = 3
	ColorDefault   Color  = 1
	ColorDark      Color  = 2
	ColorLight     Color  = 3
	ColorAccent    Color  = 4
	ColorGood      Color  = 5
	ColorWarning   Color  = 6
	ColorAttention Color  = 7
)
