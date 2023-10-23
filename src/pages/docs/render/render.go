package render

import (
	"html/template"

	"github.com/caleb-sideras/goxstack/gox/render"
	"github.com/caleb-sideras/goxstack/src/global"
	"github.com/caleb-sideras/goxstack/src/pages/docs"
)

func Render() render.DynamicT {
	tmpls := []string{"templates/components/nav.html"}

	markdownDocsContent := docs.DocsData{
		ActiveTabId:  "docs",
		ActiveDocsId: "render",
	}

	tmpl := template.Must(global.MarkdownToHTML("pages/docs/_markdown/render.md"))

	return render.DynamicT{tmpls, markdownDocsContent, tmpl}
}
