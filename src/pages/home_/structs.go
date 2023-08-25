package home

import "html/template"

type HomeCard struct {
	Title       string
	Description template.HTML
	Link        string
}

type HomeSection struct {
	Title       string
	Description string
	Section     int
	Cards       []HomeCard
}

type HomeContent struct {
	HomeActive    bool
	ExampleActive bool
	DocsActive    bool
	ActiveTabId   string
	HomeCards     []HomeCard
	HomeSections  []HomeSection
}
