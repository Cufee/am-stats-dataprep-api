package shared

import (
	"github.com/byvko-dev/am-core/helpers/slices"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

func LoadStyles(styleSheet style.Style, tags ...string) style.Style {
	if slices.Contains(tags, "gap25") > -1 {
		styleSheet = styleSheet.Merge(Gap25)
	}
	if slices.Contains(tags, "gap50") > -1 {
		styleSheet = styleSheet.Merge(Gap50)
	}
	if slices.Contains(tags, "growX") > -1 {
		styleSheet = styleSheet.Merge(GrowX)
	}
	if slices.Contains(tags, "growY") > -1 {
		styleSheet = styleSheet.Merge(GrowY)
	}
	return styleSheet
}
