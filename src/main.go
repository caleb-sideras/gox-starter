package main

import (
	"github.com/caleb-sideras/goxstack/gox"
	"github.com/caleb-sideras/goxstack/src/global"
	"github.com/caleb-sideras/goxstack/src/server"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "build":
			// gorilla/mux router instance
			r := mux.NewRouter()

			// Example - middleware
			r.Use(server.Middleware)

			// ---- Define your own custom handlers here ----
			//
			// Example - rolling your own Auth
			r.Handle("/protected", server.AuthenticationMiddleware(
				http.HandlerFunc(
					func(w http.ResponseWriter, r *http.Request) {
						log.Println("You are authenticated!")
					}),
			))
			//
			// Custom docs handler
			// r.PathPrefix("/docs/").HandlerFunc(docs.Handler_)
			//
			// -----------------------------------------------

			// Serving my static folder
			http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static/"))))
			// Create a new instance of gox
			g := gox.NewGox(global.HTML_OUT_DIR)
			// Build your GoX app -> finds your handlers, creates your routes & renders static html
			g.Build(global.APP_DIR, global.PROJECT_PACKAGE_DIR)
			// Run your GoX app -> binds handlers to routes
			g.Run(r, ":8000", global.HTML_SERVE_PATH)

		case "dev":
			// Build your own development env
			log.Println("Dev not implemented.")
		default:
			log.Println("Invalid argument. Use 'build' or 'dev'.")
		}
	} else {
		log.Println("Please provide an argument: 'build' or 'dev'.")
	}
}
