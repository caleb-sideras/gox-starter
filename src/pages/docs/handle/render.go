package handle

import (
	"html/template"
	"log"

	"github.com/caleb-sideras/goxstack/gox/render"
	"github.com/caleb-sideras/goxstack/src/global"
	"github.com/caleb-sideras/goxstack/src/pages/docs"
)

func Render() render.DynamicT {
	parentTmpl := template.Must(template.ParseFiles("templates/components/nav.html"))
	childTmpl := template.Must(global.MarkdownToHTML("pages/docs/_markdown/handle.md"))

	_, err := parentTmpl.New("page").Parse(childTmpl.Tree.Root.String())
	if err != nil {
		log.Fatal(err)
	}

	markdownDocsContent := docs.DocsData{
		ActiveTabId:  "docs",
		ActiveDocsId: "handling",
	}
	return render.DynamicT{parentTmpl, markdownDocsContent}
}
