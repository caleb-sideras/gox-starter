package todo

import (
	"github.com/caleb-sideras/goxstack/gox/data"
	"github.com/caleb-sideras/goxstack/src/pages/examples"
)

type Task struct {
	Text string
}

type TodoContent struct {
	examples.ExampleContent
	Tasks []Task
}

var Data data.Page = data.Page{
	Content: TodoContent{
		examples.ExampleContent{
			ActiveTabId:     "example",
			ActiveExampleId: "todo",
			Title:           "TODO",
			Description:     "The following TODO component shows you how to return static and dynamic HTML",
			// func Render</span> and <span class="font-bold">Func Handle</span> for sub-routes</div>
			ExampleUrl: "/examples/todo/ssr",
			CodeUrl:    "/examples/todo/code",
		},
		[]Task{},
	},
	Templates: []string{
		"templates/components/nav.html",
		"pages/examples/_components/todo_ssr.html",
		"pages/examples/_components/example.html",
	},
}
