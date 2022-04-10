package minimal

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/dataprep/icons"
	"byvko.dev/repo/am-stats-dataprep-api/stats/presets/shared"

	"github.com/byvko-dev/am-types/dataprep/settings/v1"
)

var (
	BlockBattles            = settings.BlockOptions{GenerationTag: shared.GenerationTagBattles, LocalizationTag: "localized_battles"}
	BlockWinrateWithBattles = settings.BlockOptions{GenerationTag: shared.GenerationTagWinrateWithBattles, LocalizationTag: "localized_winrate_with_battles"}
	BlockWinrate            = settings.BlockOptions{GenerationTag: shared.GenerationTagWinrate, LocalizationTag: "localized_winrate"}
	BlockDamageDone         = settings.BlockOptions{GenerationTag: shared.GenerationTagDamageDone, LocalizationTag: "localized_damage_done"}
	BlockAverageDamage      = settings.BlockOptions{GenerationTag: shared.GenerationTagAverageDamage, LocalizationTag: "localized_average_damage"}
	BlockShotAccuracy       = settings.BlockOptions{GenerationTag: shared.GenerationTagShotAccuracy, LocalizationTag: "localized_shot_accuracy"}
	BlockWN8Rating          = settings.BlockOptions{GenerationTag: shared.GenerationTagWN8Rating, LocalizationTag: "localized_wn8_rating", IconDictOverwrite: icons.IconsCircle}
)
