package nav

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/caleb-sideras/goxstack/gox/utils"
)

type Nav struct {
	ListItems []ListItem
	Active    string
}

type ListItem struct {
	PageUrl string
	Id      string
	Name    string
}

var HomeRoutes map[string]string = map[string]string{
	"/":        "home",
	"examples": "example",
	"docs":     "docs",
}

type HomeData struct {
	ActiveTabId string
}

func getFirstDirectory(filePath string) (string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	if absPath == string(filepath.Separator) {
		return absPath, nil
	}

	dir := filepath.Dir(absPath)
	firstDir := filepath.Base(dir)

	return firstDir, nil
}

func Home_(w http.ResponseWriter, r *http.Request) {

	htmxUrl, err := utils.LastElementOfURL(utils.GetHtmxRequestURL(r))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	tmpl := template.Must(template.ParseFiles("templates/components/nav.html"))

	firstDir, err := getFirstDirectory(htmxUrl)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	active := ""
	if _, ok := HomeRoutes[firstDir]; ok {
		active = HomeRoutes[firstDir]
	}
	tmpl.ExecuteTemplate(w, "nav-main", HomeData{ActiveTabId: active})
}

var ExampleRoutes map[string]bool = map[string]bool{
	"/examples/todo":   true,
	"/examples/render": true,
	"/examples/data":   true,
	"/examples/pages":  true,
}

var ExampleItems []ListItem = []ListItem{
	{
		"/examples/todo",
		"todo",
		"Todo",
	},
	{
		"/examples/render",
		"render",
		"Render",
	},
	{
		"/examples/data",
		"data",
		"Data",
	},
	{
		"/examples/pages",
		"pages",
		"Pages",
	},
}

func Examples_(w http.ResponseWriter, r *http.Request) {

	htmxUrl, err := utils.LastElementOfURL(utils.GetHtmxRequestURL(r))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	tmpl := template.Must(template.ParseFiles("templates/components/list_nav.html"))

	active := ""
	if _, ok := ExampleRoutes[htmxUrl]; ok {
		active = filepath.Base(htmxUrl)
	}
	tmpl.Execute(w, Nav{ExampleItems, active})
}

var DocRoutes map[string]bool = map[string]bool{
	"/docs":         true,
	"/docs/routing": true,
	"/docs/pages":   true,
	"/docs/index":   true,
	"/docs/data":    true,
	"/docs/render":  true,
	"/docs/handle":  true,
	"/docs/router":  true,
}

var DocItems []ListItem = []ListItem{
	{
		"/docs",
		"docs",
		"Introduction",
	},
	{
		"/docs/routing",
		"routing",
		"Routing",
	},
	{
		"/docs/pages",
		"pages",
		"Pages",
	},
	{
		"/docs/index",
		"index",
		"Index",
	},
	{
		"/docs/data",
		"data",
		"Data",
	},
	{
		"/docs/render",
		"render",
		"Render",
	},
	{
		"/docs/handle",
		"handle",
		"Handling",
	},
	{
		"/docs/router",
		"router",
		"Gox Router",
	},
}

func Docs_(w http.ResponseWriter, r *http.Request) {

	htmxUrl, err := utils.LastElementOfURL(utils.GetHtmxRequestURL(r))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	tmpl := template.Must(template.ParseFiles("templates/components/list_nav.html"))

	active := ""
	if _, ok := DocRoutes[htmxUrl]; ok {
		active = filepath.Base(htmxUrl)
	}
	tmpl.Execute(w, Nav{DocItems, active})
}
