package logic

import (
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type Text struct {
	Localize bool                `json:"localize" bson:"localize"`
	String   string              `json:"string" bson:"string"`
	Style    style.Style         `json:"style" bson:"style"`
	Printer  func(string) string `json:"-"`
}

func (txt *Text) ToBlock() block.Block {
	if txt.Localize && txt.Printer != nil {
		return block.Block{
			ContentType: block.ContentTypeText,
			Content:     txt.Printer(txt.String),
			Style:       txt.Style,
		}
	}
	return block.Block{
		ContentType: block.ContentTypeText,
		Content:     txt.String,
		Style:       txt.Style,
	}
}
