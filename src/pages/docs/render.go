package docs

import (
	"github.com/caleb-sideras/gox/utils"
	"github.com/caleb-sideras/goxstack/src/global"
	"github.com/caleb-sideras/goxstack/src/pages/home_"
	"html/template"
	"log"
)

func PreRenderDocs_() error {

	parentDocsTemplates := []string{
		"pages/index.html",
		"templates/components/nav.html",
		"pages/docs/docs.html",
	}

	markdownDocsTemplates := map[string]string{
		"spa-vs-mpa":          "pages/docs/_markdown/spa-vs-mpa.md",
		"folder-structure":    "pages/docs/_markdown/folder-structure.md",
		"client-side-routing": "pages/docs/_markdown/client-side-routing.md",
	}

	htmlDocsTemplates := map[string][]string{
		"introduction": []string{"templates/components/large_card.html"},
	}

	htmlDocsContent := DocsData{
		DocsActive:  true,
		ActiveTabId: "docs",
		LargeCards:  VarLargeCards,
	}

	markdownDocsContent := DocsData{
		DocsActive:  true,
		ActiveTabId: "docs",
		LargeCards:  []global.LargeCard{},
	}

	err := utils.RenderGeneric("tmp.html", global.HTML_OUT_DIR, append([]string{"pages/index.html", "pages/home_/page.html"}, home.Data.Templates...), home.Data.Content, "")
	if err != nil {
		log.Panicln(err)
	}
	// Markdown to HTML
	for tmplName, mkdnPath := range markdownDocsTemplates {
		parentTmpl := template.Must(template.ParseFiles(parentDocsTemplates...))
		markdownTmpl := template.Must(global.MarkdownToHTML(mkdnPath))
		_ = template.Must(parentTmpl.AddParseTree("markdown", markdownTmpl.Tree))
		file, err := utils.CreateFile(tmplName+".html", global.HTML_OUT_DIR)
		fileBody, errBody := utils.CreateFile(tmplName+"-body.html", global.HTML_OUT_DIR)
		if err != nil {
			return err
		}
		if errBody != nil {
			return err
		}
		err = utils.WriteToFile("", file, parentTmpl, markdownDocsContent)
		errBody = utils.WriteToFile("body", fileBody, parentTmpl, markdownDocsContent)
		if err != nil {
			return err
		}
		if errBody != nil {
			return err
		}
	}

	// Template to HTML
	for tmplName, tmplPaths := range htmlDocsTemplates {
		parentTmpl := template.Must(template.ParseFiles(parentDocsTemplates...))
		_ = template.Must(parentTmpl.ParseFiles(tmplPaths...))
		file, err := utils.CreateFile(tmplName+".html", global.HTML_OUT_DIR)
		fileBody, errBody := utils.CreateFile(tmplName+"-body.html", global.HTML_OUT_DIR)
		if err != nil {
			return err
		}
		if errBody != nil {
			return err
		}
		err = utils.WriteToFile("", file, parentTmpl, htmlDocsContent)
		errBody = utils.WriteToFile("body", fileBody, parentTmpl, htmlDocsContent)
		if err != nil {
			return err
		}
		if errBody != nil {
			return err
		}
	}
	return nil
}
