package docs

import (
	// "github.com/caleb-sideras/gox/utils"
	"log"
	"net/http"
	"path/filepath"
)

func Handler_(w http.ResponseWriter, r *http.Request) {
	docsGeneralHandler([]string{"templates/index.html", "templates/components/nav.html", "pages/docs/docs.html"}, w, r)
}

func docsGeneralHandler(parentPaths []string, w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	paths := map[string]bool{
		"/docs/spa-vs-mpa":           true,
		"/docs/client-side-routing":  true,
		"/docs/routing-fundamentals": true,
		"/docs/introduction":         true,
		"/docs/pre-rendered-routes":  true,
		"/docs/custom-handling":      true,
	}

	if _, ok := paths[path]; ok {
		// if utils.IsHtmxRequest(r) {
		// 	log.Println("Serving static: " + filepath.Base(r.URL.Path) + "-body.html")
		// 	http.ServeFile(w, r, "../static/html/"+filepath.Base(r.URL.Path)+"-body.html")
		// } else {
		log.Println("Serving static: " + filepath.Base(r.URL.Path) + ".html")
		http.ServeFile(w, r, "../static/html/"+filepath.Base(r.URL.Path)+".html")
		// }
	}
}
