package logic

import (
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type LayoutRow struct {
	Style style.Style  `json:"style"`
	Items []LayoutItem `json:"items"`
}

func (row *LayoutRow) ToBlock(lt Layout, printer func(string) string) *block.Block {
	blocks := make([]block.Block, 0, len(row.Items))
	for _, item := range row.Items {
		b := item.ToBlock(lt.Values, lt.Style.Merge(row.Style), printer)
		if b != nil {
			blocks = append(blocks, *b)
		}
	}
	if len(blocks) == 0 {
		return nil
	}
	return &block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     blocks,
		Style:       row.Style,
	}
}
