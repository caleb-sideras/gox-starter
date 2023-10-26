package index

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
		ActiveDocsId: "index",
	}

	tmpl := template.Must(global.MarkdownToHTML("pages/docs/_markdown/index.md"))

	return render.DynamicT{tmpls, markdownDocsContent, tmpl}
}
