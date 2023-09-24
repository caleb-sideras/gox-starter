package docs

import (
	"html/template"
	"log"

	"github.com/caleb-sideras/goxstack/gox/utils"
	"github.com/caleb-sideras/goxstack/src/global"
)

func Render() utils.RenderTemplateType {
	parentDocsTemplates := []string{
		"pages/index.html",
		"templates/components/nav.html",
		"pages/docs/docs.html",
	}

	markdownDocsTemplates := []string{
		"pages/docs/_markdown/introduction.md",
		"pages/docs/_markdown/routing-fundamentals.md",
		"pages/docs/_markdown/page-based-routing.md",
		"pages/docs/_markdown/index.md",
		"pages/docs/_markdown/data.md",
		"pages/docs/_markdown/render.md",
		"pages/docs/_markdown/custom-handling.md",
		"pages/docs/_markdown/gox-router.md",
		// "pages/docs/_markdown/client-side-routing.md",
		// "pages/docs/_markdown/spa-vs-mpa.md",
	}

	parentTmpl := template.Must(template.ParseFiles(parentDocsTemplates...))
	var childTmpl string

	for _, mkdnPath := range markdownDocsTemplates {
		markdownTmpl := template.Must(global.MarkdownToHTML(mkdnPath))
		childTmpl += markdownTmpl.Tree.Root.String() + "<hr>"
	}
	_, err := parentTmpl.New("markdown").Parse(childTmpl)

	if err != nil {
		log.Fatal(err)
	}

	markdownDocsContent := DocsData{
		ActiveTabId: "docs",
		LargeCards:  []global.LargeCard{},
	}
	return utils.RenderTemplateType{markdownDocsContent, parentTmpl, ""}
}
