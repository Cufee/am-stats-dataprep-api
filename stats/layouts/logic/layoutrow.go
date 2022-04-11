package logic

import (
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type LayoutRow struct {
	Style style.Style  `json:"style"`
	Items []LayoutItem `json:"items"`
}

func (row *LayoutRow) ToBlock(lt Layout) block.Block {
	blocks := make([]block.Block, 0, len(row.Items))
	for _, item := range row.Items {
		blocks = append(blocks, item.ToBlock(lt.Values, lt.Style.Merge(row.Style)))
	}
	return block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     blocks,
		Style:       lt.Style,
	}
}
