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
	Type         ItemType          `json:"type"`
	Data         interface{}       `json:"data"`
	Style        style.Style       `json:"style"`
	AddCondition func(Values) bool `json:"-"`
}

func (item *LayoutItem) ToBlock(values Values, stl style.Style, printer func(string) string) *block.Block {
	if item.AddCondition != nil && !item.AddCondition(values) {
		return nil
	}
	var b block.Block
	switch item.Type {
	case ItemTypeIcon:
		data, ok := item.Data.(Icon)
		if !ok {
			return nil
		}
		iconStyle := data.GetStyle(values)
		b.Content = block.Block{
			Content: data.GetName(values),
			Style:   iconStyle,
		}
		b.ContentType = block.ContentTypeIcon
		b.Style = shared.DefaultFont.Merge(stl).Merge(item.Style)
	case ItemTypeTemplate:
		data, ok := item.Data.(Template)
		if !ok {
			return nil
		}
		b.Content = data.Evaluate(values)
		b.ContentType = block.ContentTypeText
		b.Style = shared.DefaultFont.Merge(stl).Merge(item.Style)
	case ItemTypeText:
		text, ok := item.Data.(Text)
		text.Printer = printer
		if !ok {
			return nil
		}
		text.Style = shared.DefaultFont.Merge(stl).Merge(text.Style)
		b = text.ToBlock(values)
	default:
		return nil
	}
	return &b
}
