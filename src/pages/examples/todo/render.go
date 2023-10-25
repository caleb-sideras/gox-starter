package todo

import (
	"github.com/caleb-sideras/goxstack/gox/render"
	"github.com/caleb-sideras/goxstack/src/global"
	"html/template"
)

func Code_() render.StaticT {
	tmpl := template.Must(global.MarkdownToHTML("pages/examples/_markdown/todo.md"))

	wrapperTemplateString := `<script src="/static/js/prism.js"></script> <div id="markdown">{{template "code" .}}</div>`
	wrapperTmpl := template.Must(template.New("markdown").Parse(wrapperTemplateString))
	wrapperTmpl.New("code").Parse(tmpl.Tree.Root.String())

	return render.StaticT{wrapperTmpl, nil, ""}
}

// func Csr_() render.StaticF {
// 	return render.StaticF{[]string{"pages/examples/_components/todo_csr.html"}, nil, "todo-csr"}
// }

func Ssr_() render.StaticF {
	return render.StaticF{[]string{"pages/examples/_components/todo_ssr.html"}, nil, "todo-ssr"}
}
