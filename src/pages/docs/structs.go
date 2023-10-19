package docs

import "github.com/caleb-sideras/goxstack/src/global"

type DocsData struct {
	ActiveTabId  string
	ActiveDocsId string
	LargeCards   []global.LargeCard
	HoverCards   []global.HoverCard
}
