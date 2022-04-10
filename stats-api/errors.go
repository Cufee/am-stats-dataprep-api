package statsapi

import (
	"strings"

	errors "github.com/byvko-dev/am-types/errors/v1/stats"
)

func parseStatsError(err error) error {
	if strings.Contains(err.Error(), "mongo: no documents in result") {
		return errors.ErrNoSessionAvailable
	}
	return err
}
