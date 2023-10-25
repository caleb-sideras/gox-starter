package example_render

import (
	"github.com/caleb-sideras/goxstack/gox/data"
	"github.com/caleb-sideras/goxstack/src/pages/examples"
)

var Data data.Page = data.Page{
	Content: examples.ExampleContent{
		ActiveTabId:     "example",
		ActiveExampleId: "render",
		Title:           "Custom Rendering",
		Description:     "Have a build step for your HTML",
		ExampleUrl:      "/examples/render/markdown",
		CodeUrl:         "/examples/render/code",
	},

	Templates: []string{
		"templates/components/nav.html",
		"pages/examples/_components/example.html",
	},
}
