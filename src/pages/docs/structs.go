package docs

import "calebsideras.com/gox/src/global"

type DocsData struct {
	HomeActive    bool
	ExampleActive bool
	DocsActive    bool
	ActiveTabId   string
	LargeCards    []global.LargeCard
}
