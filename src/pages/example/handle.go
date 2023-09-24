package example

import (
	"html/template"
	"net/http"
	"time"
)

func HandleAddTask_(w http.ResponseWriter, r *http.Request) {
	// To showcase loading state
	time.Sleep(1 * time.Second)
	task := r.PostFormValue("task")
	tmpl := template.Must(template.ParseFiles("pages/example/_components/todo_ssr.html"))
	tmpl.ExecuteTemplate(w, "task-list-element", Task{Text: task})
}
