package example

import "github.com/caleb-sideras/goxstack/gox/utils"

var Content ExampleContent = ExampleContent{
	ActiveTabId: "example",
	Tasks:       []Task{},
}
var Templates []string = []string{
	"templates/components/nav.html",
	"pages/example/_components/todo_ssr.html",
}

var Data utils.PageData = utils.PageData{
	Content:   Content,
	Templates: Templates,
}
