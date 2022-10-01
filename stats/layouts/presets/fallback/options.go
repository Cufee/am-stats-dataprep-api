package fallback

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"

	"byvko.dev/repo/am-stats-dataprep-api/stats/helpers"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/logic"
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

var wrapperStyle style.Style = style.Style{
	PaddingLeft:   1,
	PaddingRight:  1,
	PaddingTop:    1,
	PaddingBottom: 1,
	Gap:           0.5,
}

func Init() {
	fmt.Println("fallback.Init()")

	img, err := helpers.LoadImage("./assets/bg_ukraine.jpg")
	if err != nil {
		fmt.Println("fallback.Init() - error:", err)
		img, err = helpers.LoadImage("./assets/bg_default.png")
	}
	if err != nil {
		panic(err)
	}

	img = helpers.BlurImage(img, 30)

	buf := new(bytes.Buffer)
	err = png.Encode(buf, img)
	if err != nil {
		panic(err)
	}
	wrapperStyle.BackgroundImage = base64.StdEncoding.EncodeToString(buf.Bytes())

	Preset = logic.LayoutOptions{
		WrapperStyle:         shared.AlignVertical.Merge(wrapperStyle),
		LayoutName:           "fallback",
		AccountStatus:        nil,
		Notifications:        nil,
		Challenges:           nil,
		PlayerInfo:           &PlayerName,
		RatingOverview:       &OverviewRating,
		RandomOverview:       &OverviewRandom,
		VehiclesFullOverview: &VehiclesDetailed,
		VehiclesSlimOverview: &VehiclesSlim,
	}

}

var Preset logic.LayoutOptions
