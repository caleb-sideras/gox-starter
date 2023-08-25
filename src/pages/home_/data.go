package home

import "github.com/caleb-sideras/gox/utils"

var Content HomeContent = HomeContent{
	HomeActive:   true,
	ActiveTabId:  "home",
	HomeCards:    VarHomeCards,
	HomeSections: VarHomeSections,
}
var Templates []string = []string{
	"templates/components/nav.html",
	"templates/components/card.html",
}

var Data utils.PageData = utils.PageData{
	Content:   Content,
	Templates: Templates,
}
