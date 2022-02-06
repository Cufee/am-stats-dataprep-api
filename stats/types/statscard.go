package types

// export interface StatsCard {
//   rows: StatsCardRow[];
//   tags: string[];
// }
type StatsCard struct {
	Rows []StatsCardRow `json:"rows"`
	Tags []string       `json:"tags"`
}

// export interface StatsCardRow {
//   blocks: StatsBlock[];
// }
type StatsCardRow struct {
	Blocks []StatsBlock `json:"blocks"`
}

// export interface StatsBlock {
//   rows: StatsBlockRow[];
//   tags: string[];
// }
type StatsBlock struct {
	Rows []StatsBlockRow `json:"rows"`
	Tags []string        `json:"tags"`
}

// export interface StatsBlockRow {
//   content: StatsBlockRowContent | StatsBlockRowContent[];
// }
type StatsBlockRow struct {
	Content []StatsBlockRowContent `json:"content"`
}

// export interface StatsBlockRowContent {
//   tags: string[];
//   content: any;
//   isLocalized: boolean;
// }
type StatsBlockRowContent struct {
	Tags        []string    `json:"tags"`
	Content     interface{} `json:"content"`
	Type        ContentType `json:"type"`
	IsLocalized bool        `json:"isLocalized"`
}

type ContentType string

const (
	ContentTypeText  ContentType = "text"
	ContentTypeImage ContentType = "image"
)
