package gox

import (
	"github.com/caleb-sideras/goxstack/gox/data"
	"net/http"
)

type PageData struct {
	Data                data.Page
	AdditionalTemplates []string
}

type DataRender struct {
	Data                data.PageFunc
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
