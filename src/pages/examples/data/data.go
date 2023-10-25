package example_data

import (
	"github.com/caleb-sideras/goxstack/gox/data"
	"github.com/caleb-sideras/goxstack/src/pages/examples"
)

type DataContent struct {
	examples.ExampleContent
	Number int
}

var Data data.Page = data.Page{
	examples.ExampleContent{
		ActiveTabId:     "example",
		ActiveExampleId: "data",
		Title:           "Data Fetching",
		Description:     "Per-request data fetching for your pages",
		ExampleUrl:      "/examples/data/rn?index=false",
		CodeUrl:         "/examples/data/code",
	},
	[]string{
		"templates/components/nav.html",
		"pages/examples/_components/example.html",
	},
}
