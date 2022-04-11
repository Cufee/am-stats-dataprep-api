package fallback

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
)

func Loader(kind shared.LayoutKind) *logic.Layout {
	switch kind {
	case shared.Accuracy:
		return accuracy()
	case shared.Battles:
		return battles()
	default:
		return nil
	}
}
