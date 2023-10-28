package home

import (
	"github.com/caleb-sideras/gox-starter/.gox/data"
)

type HomeContent struct {
	Title       string
	Description string
}

var Content HomeContent = HomeContent{
	Title:       "Welcome",
	Description: "Go, HTMX and beyond!",
}

var Data data.Page = data.Page{
	Content:   Content,
	Templates: []string{},
}
