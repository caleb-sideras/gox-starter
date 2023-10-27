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

	err := os.Chdir("src")
	if err != nil {
		panic(err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "build":
			// Create a new instance of gox
			g := gox.NewGox(global.HTML_OUT_DIR)

			// Build your GoX app -> finds your handlers, creates your routes & renders static html
			g.Build(global.APP_DIR, global.PROJECT_PACKAGE_DIR)

		case "run":
			// gorilla/mux router instance
			r := mux.NewRouter()

			// Example - middleware
			r.Use(server.Middleware)

			// Serving my static folder
			http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../static/"))))

			// Create a new instance of gox
			g := gox.NewGox(global.HTML_OUT_DIR)

			// Run your GoX app -> binds handlers to routes
			g.Run(r, ":8000", global.HTML_SERVE_PATH)
		default:
			log.Println("Invalid argument. Use 'build' or 'run'.")
		}
	} else {
		log.Println("Please provide an argument: 'build' or 'run'.")
	}
}
