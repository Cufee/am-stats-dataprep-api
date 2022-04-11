package logic

import (
	"byvko.dev/repo/am-stats-dataprep-api/stats/layouts/shared"
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type ItemType string

const (
	ItemTypeText     ItemType = "text"
	ItemTypeIcon     ItemType = "icon"
	ItemTypeTemplate ItemType = "template"
)

type LayoutItem struct {
	Type  ItemType    `json:"type"`
	Data  interface{} `json:"data"`
	Style style.Style `json:"style"`
}

func (item *LayoutItem) ToBlock(values Values, stl style.Style) block.Block {
	var b block.Block
	switch item.Type {
	case ItemTypeIcon:
		item, ok := item.Data.(Icon)
		if !ok {
			return shared.InvalidBlock
		}
		b.Style = stl.Merge(item.GetStyle(values))
		b.Content = item.GetName(values)
		b.ContentType = block.ContentTypeIcon
	case ItemTypeTemplate:
		item, ok := item.Data.(Template)
		if !ok {
			return shared.InvalidBlock
		}
		b.Style = stl
		b.Content = item.Evaluate(values)
		b.ContentType = block.ContentTypeText
	case ItemTypeText:
		text, ok := item.Data.(Text)
		if !ok {
			return shared.InvalidBlock
		}
		text.Style = stl.Merge(text.Style)
		return text.ToBlock()
	default:
		return shared.InvalidBlock
	}
	return b
}
