package example

import (
	"github.com/caleb-sideras/goxstack/gox/render"
	"github.com/caleb-sideras/goxstack/src/global"
	"html/template"
)

func TodoCsr_() render.StaticF {
	return render.StaticF{[]string{"pages/example/_components/todo_csr.html"}, nil, "todo-csr"}
}
func TodoSsr_() render.StaticF {
	return render.StaticF{[]string{"pages/example/_components/todo_ssr.html"}, nil, "todo-ssr"}
}

func TodoCode_() render.StaticT {
	tmpl := template.Must(global.MarkdownToHTML("pages/example/todo.md"))

	wrapperTemplateString := `<script src="/static/js/prism.js"></script> <div id="markdown">{{template "code" .}}</div>`
	wrapperTmpl := template.Must(template.New("markdown").Parse(wrapperTemplateString))
	wrapperTmpl.New("code").Parse(tmpl.Tree.Root.String())

	return render.StaticT{wrapperTmpl, nil, ""}
}
