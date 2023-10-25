package example_pages

import (
	"github.com/caleb-sideras/goxstack/gox/render"
	"github.com/caleb-sideras/goxstack/src/global"
	"html/template"
)

func Code_() render.StaticT {
	tmpl := template.Must(global.MarkdownToHTML("pages/examples/_markdown/data.md"))

	wrapperTemplateString := `<script src="/static/js/prism.js"></script> <div id="markdown">{{template "code" .}}</div>`
	wrapperTmpl := template.Must(template.New("markdown").Parse(wrapperTemplateString))
	wrapperTmpl.New("code").Parse(tmpl.Tree.Root.String())

	return render.StaticT{wrapperTmpl, nil, ""}
}

func Example_() render.StaticF {

	type PageBody struct {
		FullPage string
		Body     string
	}
	var content PageBody = PageBody{
		FullPage: "pages/examples?index=true",
		Body:     "pages/examples?index=false",
	}

	return render.StaticF{[]string{"pages/examples/_components/body.html"}, content, ""}
}
