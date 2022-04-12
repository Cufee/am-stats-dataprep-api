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

func (lt *Layout) ToBlock(printer func(string) string) *block.Block {
	blocks := make([]block.Block, 0, len(lt.Rows))
	for _, row := range lt.Rows {
		b := row.ToBlock(*lt, printer)
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
		Style:       lt.Style,
	}
}
