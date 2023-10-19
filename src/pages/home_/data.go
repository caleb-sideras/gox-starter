package home

import (
	"github.com/caleb-sideras/goxstack/gox/data"
	"net/http"
)

var Content HomeContent = HomeContent{
	ActiveTabId:  "home",
	HomeCards:    VarHomeCards,
	HomeSections: VarHomeSections,
}
var Templates []string = []string{
	"templates/components/nav.html",
	"templates/components/card.html",
}

var pageData data.Page = data.Page{
	Content:   Content,
	Templates: Templates,
}

func Data(w http.ResponseWriter, r *http.Request) data.PageReturn {
	return data.PageReturn{pageData, nil}
}
