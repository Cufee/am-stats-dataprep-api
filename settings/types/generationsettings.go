package types

import (
	"errors"
	"fmt"
	"time"

	"byvko.dev/repo/am-stats-dataprep-api/stats/types"
)

type GenerationSettings struct {
	Options     types.Options `json:"options" firestore:"options"`
	StylePreset string        `json:"stylePreset" firestore:"stylePreset"`
	LastUsed    time.Time     `json:"lastUsed,omitempty" firestore:"lastUsed,omitempty"`

	Player struct {
		ID    int    `json:"id" firestore:"id"`
		Realm string `json:"realm" firestore:"realm"`
	} `json:"player" firestore:"player"`

	Locale string `json:"locale" firestore:"locale"`
}

func (s *GenerationSettings) Validate() error {
	if s.Player.ID == 0 {
		return errors.New("player ID is not set")
	}
	if !CheckRealm(s.Player.Realm) {
		return errors.New("realm is not valid")
	}
	var blank types.Options
	if fmt.Sprintf("%v", s.Options) == fmt.Sprintf("%v", blank) {
		return errors.New("options are not set")
	}
	return nil
}

func CheckRealm(realm string) bool {
	switch realm {
	case "NA":
		return true
	case "EU":
		return true
	case "RU":
		return true
	case "ASIA":
		return true
	default:
		return false
	}

}
