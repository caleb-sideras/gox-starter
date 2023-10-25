package example_pages

import (
	"github.com/caleb-sideras/goxstack/gox/data"
	"github.com/caleb-sideras/goxstack/src/pages/examples"
)

var Data data.Page = data.Page{
	Content: examples.ExampleContent{
		ActiveTabId:     "example",
		ActiveExampleId: "pages",
		Title:           "Pages",
		Description:     "Access full or partial HTML from your pages",
		ExampleUrl:      "/examples/pages/example",
		CodeUrl:         "/examples/pages/code",
	},

	Templates: []string{
		"templates/components/nav.html",
		"pages/examples/_components/example.html",
	},
}
