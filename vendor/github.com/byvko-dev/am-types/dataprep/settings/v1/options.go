package settings

import (
	"image/color"

	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type BlockOptions struct {
	GenerationTag   string `json:"generationTag,omitempty"`
	LocalizationTag string `json:"localizationTag,omitempty"`

	IconDictOverwrite  map[string]string `json:"iconDictOverwrite,omitempty"`  // By default icons are arrows
	IconColorOverWrite color.RGBA        `json:"iconColorOverwrite,omitempty"` // By default, icons are red,green,yellow based on session vs all time comparison
	HasIcon            bool              `json:"hasIcon,omitempty"`
	HasInvisibleIcon   bool              `json:"hasInvisibleIcon,omitempty"`
	IconPosition       int               `json:"iconPosition,omitempty"`
}

func (b BlockOptions) WithIcon() BlockOptions {
	b.HasIcon = true
	b.IconPosition = style.IconPositionLeft
	return b
}

func (b BlockOptions) WithInvisibleIcon() BlockOptions {
	b.HasInvisibleIcon = true
	return b
}

func (b BlockOptions) IconLeft() BlockOptions {
	b.IconPosition = style.IconPositionLeft
	return b
}
func (b BlockOptions) IconRight() BlockOptions {
	b.IconPosition = style.IconPositionRight
	return b
}

type Options struct {
	Locale string `json:"locale"`

	// Status icons
	AccountStatus StatusIconsOptions `json:"accountStatus,omitempty" firestore:"accountStatus"`

	// Notifications
	Notifications NotificationsOptions `json:"notifications,omitempty" firestore:"notifications"`

	// Challenges
	Challenges ChallengesOptions `json:"challenges,omitempty" firestore:"challenges"`

	// Player info and general stats
	Player         PlayerOptions   `json:"player,omitempty" firestore:"player"`
	RatingBattles  OverviewOptions `json:"ratingBattles,omitempty" firestore:"ratingBattles"`
	RegularBattles OverviewOptions `json:"regularBattles,omitempty" firestore:"regularBattles"`

	// Per vehicle stats
	VehiclesFull VehicleOptions `json:"vehiclesFull,omitempty" firestore:"vehiclesFull"`
	VehiclesSlim VehicleOptions `json:"vehiclesSlim,omitempty" firestore:"vehiclesSlim"`

	WrapperStyle style.Style `json:"wrapperStyle,omitempty" firestore:"wrapperStyle"`
}

type StatusIconsOptions struct {
	Include bool `json:"include,omitempty" firestore:"include"`
	Limit   int  `json:"limit,omitempty" firestore:"limit"`
}

type NotificationsOptions struct {
	Include bool           `json:"include,omitempty" firestore:"include"`
	Blocks  []BlockOptions `json:"blocks,omitempty" firestore:"blocks"`
}

type ChallengesOptions struct {
	Include     bool           `json:"include,omitempty" firestore:"include"`
	Limit       int            `json:"limit,omitempty" firestore:"limit"`
	Types       []string       `json:"types,omitempty" firestore:"types"`
	Blocks      []BlockOptions `json:"blocks,omitempty" firestore:"blocks"`
	style.Style `json:"style,omitempty" firestore:"style"`
}

type PlayerOptions struct {
	Include     bool `json:"include,omitempty" firestore:"include"`
	WithName    bool `json:"withName,omitempty" firestore:"withName"`
	WithClanTag bool `json:"withClanTag,omitempty" firestore:"withClanTag"`
	WithPins    bool `json:"withPins,omitempty" firestore:"withPins"`
}

type OverviewOptions struct {
	Include          bool           `json:"include,omitempty" firestore:"include"`
	WithIcons        bool           `json:"withIcons,omitempty" firestore:"withIcons"`
	WithTitle        bool           `json:"withTitle,omitempty" firestore:"withTitle"`
	WithLabels       bool           `json:"withLabels,omitempty" firestore:"withLabels"`
	WithAllTimeStats bool           `json:"withAllTimeStats,omitempty" firestore:"withAllTimeStats"`
	Type             string         `json:"type,omitempty" firestore:"type"`
	Blocks           []BlockOptions `json:"blocks,omitempty" firestore:"blocks"`
}

type VehicleOptions struct {
	Include bool `json:"include,omitempty" firestore:"include"`
	Limit   int  `json:"limit,omitempty" firestore:"limit"`

	WithVehicleTier  bool           `json:"withVehicleTier,omitempty" firestore:"withVehicleTier"`
	WithVehicleName  bool           `json:"withVehicleName,omitempty" firestore:"withVehicleName"`
	WithAllTimeStats bool           `json:"withAllTimeStats,omitempty" firestore:"withAllTimeStats"`
	WithLabels       bool           `json:"withLabels,omitempty" firestore:"withLabels"`
	WithIcons        bool           `json:"withIcons,omitempty" firestore:"withIcons"`
	Offset           int            `json:"offset,omitempty" firestore:"offset"`
	Blocks           []BlockOptions `json:"blocks,omitempty" firestore:"blocks"`
}

const (
	OverviewTypeRegular = "regular"
	OverviewTypeRating  = "rating"
)
