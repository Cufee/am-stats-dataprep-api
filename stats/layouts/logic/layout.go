package logic

import (
	"github.com/byvko-dev/am-types/dataprep/block/v1"
	"github.com/byvko-dev/am-types/dataprep/style/v1"
)

type Values map[string]interface{}

type Layout struct {
	Rows   []LayoutRow `json:"rows"`
	Style  style.Style `json:"style"`
	Values Values      `json:"values"`
}

func (lt *Layout) ToBlock() block.Block {
	blocks := make([]block.Block, 0, len(lt.Rows))
	for _, row := range lt.Rows {
		blocks = append(blocks, row.ToBlock(*lt))
	}
	return block.Block{
		ContentType: block.ContentTypeBlocks,
		Content:     blocks,
		Style:       lt.Style,
	}
}
