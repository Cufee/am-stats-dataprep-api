package logic

import (
	"fmt"

	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	str "github.com/byvko-dev/am-core/helpers/strings"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type Text struct {
	Localize bool                `json:"localize" bson:"localize"`
	String   string              `json:"string" bson:"string"`
	Style    style.Style         `json:"style" bson:"style"`
	Printer  func(string) string `json:"-"`
}

func (txt *Text) ToBlock(values Values) block.Block {
	if txt.Localize && txt.Printer != nil {
		return block.Block{
			ContentType: block.ContentTypeText,
			Content:     txt.Printer(str.Or(txt.String, fmt.Sprint(values[String]))),
			Style:       shared.DefaultFont.Merge(txt.Style),
		}
	}
	return block.Block{
		ContentType: block.ContentTypeText,
		Content:     str.Or(txt.String, fmt.Sprint(values[String])),
		Style:       shared.DefaultFont.Merge(txt.Style),
	}
}
