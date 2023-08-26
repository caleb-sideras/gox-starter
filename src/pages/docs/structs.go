package docs

import "github.com/caleb-sideras/goxstack/src/global"

type DocsData struct {
	HomeActive    bool
	ExampleActive bool
	DocsActive    bool
	ActiveTabId   string
	LargeCards    []global.LargeCard
}
