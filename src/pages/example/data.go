package example

import (
	"github.com/caleb-sideras/goxstack/gox/data"
)

var Content ExampleContent = ExampleContent{
	ActiveTabId: "example",
	Tasks:       []Task{},
}
var Templates []string = []string{
	"templates/components/nav.html",
	"pages/example/_components/todo_ssr.html",
}

var Data data.Page = data.Page{
	Content:   Content,
	Templates: Templates,
}
