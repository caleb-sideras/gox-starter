package introduction

import (
	"github.com/caleb-sideras/goxstack/gox/data"
	"github.com/caleb-sideras/goxstack/src/global"
	"github.com/caleb-sideras/goxstack/src/pages/docs"
)

var content docs.DocsData = docs.DocsData{
	ActiveTabId:  "docs",
	ActiveDocsId: "intro",

	LargeCards: []global.LargeCard{
		{
			Title:       "Introduction",
			Description: "GoX is a framework designed to make working with HTMX and Go easier",
			Image:       "https://lh3.googleusercontent.com/wle2URFsybXPDp10jnrUay6k4sw14Hg8Ajb91djUCIHa1Z7hBDFAbjPhnm-mCmjeoI-6_FPQYW6G9H2Z8kL5LfXfx_FpDlAS3DZX4wZjrbEEe2arzQ=w2400-rj",
		},
	},
	HoverCards: []global.HoverCard{
		{
			ImageText:   "index.html",
			Title:       "Index",
			Description: "Shared UI and State",
			Image:       "https://lh3.googleusercontent.com/wle2URFsybXPDp10jnrUay6k4sw14Hg8Ajb91djUCIHa1Z7hBDFAbjPhnm-mCmjeoI-6_FPQYW6G9H2Z8kL5LfXfx_FpDlAS3DZX4wZjrbEEe2arzQ=w2400-rj",
			Link:        "/docs/index",
		},
		{
			ImageText:   "page.html",
			Title:       "Pages",
			Description: "GoX handles your routing",
			Image:       "https://lh3.googleusercontent.com/wle2URFsybXPDp10jnrUay6k4sw14Hg8Ajb91djUCIHa1Z7hBDFAbjPhnm-mCmjeoI-6_FPQYW6G9H2Z8kL5LfXfx_FpDlAS3DZX4wZjrbEEe2arzQ=w2400-rj",
			Link:        "/docs/pages",
		},
		{
			ImageText:   "func() Data",
			Title:       "Data",
			Description: "Dynamic data with your pages",
			Image:       "https://lh3.googleusercontent.com/wle2URFsybXPDp10jnrUay6k4sw14Hg8Ajb91djUCIHa1Z7hBDFAbjPhnm-mCmjeoI-6_FPQYW6G9H2Z8kL5LfXfx_FpDlAS3DZX4wZjrbEEe2arzQ=w2400-rj",
			Link:        "/docs/data",
		},
		{
			ImageText:   "render.go",
			Title:       "Render",
			Description: "Custom rendering at build time",
			Image:       "https://lh3.googleusercontent.com/wle2URFsybXPDp10jnrUay6k4sw14Hg8Ajb91djUCIHa1Z7hBDFAbjPhnm-mCmjeoI-6_FPQYW6G9H2Z8kL5LfXfx_FpDlAS3DZX4wZjrbEEe2arzQ=w2400-rj",
			Link:        "/docs/render",
		},
		{
			ImageText:   "handle.go",
			Title:       "Handle",
			Description: "Return dynamic HTML partials",
			Image:       "https://lh3.googleusercontent.com/wle2URFsybXPDp10jnrUay6k4sw14Hg8Ajb91djUCIHa1Z7hBDFAbjPhnm-mCmjeoI-6_FPQYW6G9H2Z8kL5LfXfx_FpDlAS3DZX4wZjrbEEe2arzQ=w2400-rj",
			Link:        "/docs/handle",
		},
		{
			ImageText:   "Gox Router",
			Title:       "Router",
			Description: "Learn more about the Gox Router",
			Image:       "https://lh3.googleusercontent.com/wle2URFsybXPDp10jnrUay6k4sw14Hg8Ajb91djUCIHa1Z7hBDFAbjPhnm-mCmjeoI-6_FPQYW6G9H2Z8kL5LfXfx_FpDlAS3DZX4wZjrbEEe2arzQ=w2400-rj",
			Link:        "/docs/router",
		},
	},
}
var templates []string = []string{
	"templates/components/nav.html",
	"templates/components/large_card2.html",
	"templates/components/hover_card.html",
	"templates/components/hover_card.html",
}

var Data data.Page = data.Page{
	Content:   content,
	Templates: templates,
}
