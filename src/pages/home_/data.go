package home

import (
	"github.com/caleb-sideras/goxstack/gox/utils"
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

var data utils.PageData = utils.PageData{
	Content:   Content,
	Templates: Templates,
}

func Data(w http.ResponseWriter, r *http.Request) utils.DataReturnType {
	return utils.DataReturnType{data, nil}
}
