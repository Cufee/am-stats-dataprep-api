package generators

import (
	"errors"

	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/settings/v1"
	api "github.com/byvko-dev/am-types/stats/v1"
)

func GenerateStatusIcons(stats *api.PlayerRawStats, options settings.StatusIconsOptions) ([]block.Block, error) {
	return []block.Block{}, errors.New("not implemented")
}
