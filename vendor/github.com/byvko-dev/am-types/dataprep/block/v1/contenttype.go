package block

type ContentType string

func (t *ContentType) Valid() bool {
	return *t != ContentTypeInvalid
}

// Available content type
var ContentTypeText ContentType = "text"
var ContentTypeIcon ContentType = "icon"
var ContentTypeImage ContentType = "image"
var ContentTypeBlocks ContentType = "blocks"
var ContentTypeInvalid ContentType = "invalid"

func (b *Block) GetContentType() ContentType {
	if b.ContentType == "" {
		return ContentTypeInvalid
	}
	return b.ContentType
}
