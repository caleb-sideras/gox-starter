package gox

import (
	"github.com/caleb-sideras/goxstack/gox/utils"
	"net/http"
)

type PageData struct {
	Data                utils.PageData
	AdditionalTemplates []string
}

type RenderData struct {
	Data                utils.DataReturnFunc
	AdditionalTemplates []string
}

type HandlerDefaultFunc func(http.ResponseWriter, *http.Request)
type HandlerDefault struct {
	Path    string
	Handler HandlerDefaultFunc
}

type RenderCustomFunc func() error
type RenderCustom struct {
	Handler RenderCustomFunc
}

type RenderDefault struct {
	Path    string
	Handler interface{}
}
