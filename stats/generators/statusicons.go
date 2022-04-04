package generators

import (
	"errors"

	"github.com/byvko-dev/am-types/dataprep/v1/block"
	"github.com/byvko-dev/am-types/dataprep/v1/settings"
	api "github.com/byvko-dev/am-types/stats/v1"
)

func GenerateStatusIcons(stats *api.PlayerRawStats, options settings.StatusIconsOptions) ([]block.Block, error) {
	return []block.Block{}, errors.New("not implemented")
}
