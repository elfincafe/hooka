package adaptive_card

type (
	Spacing             uint8
	HorizontalAlignment uint8
)

const (
	SpacingNone               Spacing             = 0
	SpacingSmall              Spacing             = 1
	SpacingDefault            Spacing             = 2
	SpacingMedium             Spacing             = 3
	SpacingLarge              Spacing             = 4
	SpacingExtraLarge         Spacing             = 5
	SpacingPadding            Spacing             = 6
	HorizontalAlignmentNone   HorizontalAlignment = 0
	HorizontalAlignmentLeft   HorizontalAlignment = 1
	HorizontalAlignmentCenter HorizontalAlignment = 2
	HorizontalAlignmentRight  HorizontalAlignment = 3
)
