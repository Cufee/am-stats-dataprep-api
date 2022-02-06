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

// export default interface BlockRowContent {
//   type: "text" | "image";
//   isLocalized: boolean;
//   tags: string[];
//   content: any;
// }
type StatsBlockRowContent struct {
	Tags        []string    `json:"tags"`
	Content     interface{} `json:"content"`
	Type        ContentType `json:"type"`
	IsLocalized bool        `json:"isLocalized"`
}

type ContentIcon struct {
	Color string `json:"color"`
	Name  string `json:"name"`
	Size  int    `json:"size"`
}

type ContentImage struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type ContentType string

const (
	ContentTypeText  ContentType = "text"
	ContentTypeImage ContentType = "image"
	ContentTypeIcon  ContentType = "icon"
)
