# Custom Handling

GoX provides an easy way to define and handle routes that will be receving and sending html partials.  


## handle.go

Within your desired route, add a `handle.go` file and define a function that takes both a `http.ResponseWriter` and a `*http.Request`. The name of the function will resolve to your desired path.  
\
Given the following folder structure:
```
app
 ├─ index.html
 └─ example
    ├─ page.html
    └─ handle.go
``` 

And contents of the handle.go file:

```go
package example

import (
	"html/template"
	"net/http"
)

func AddTask(w http.ResponseWriter, r *http.Request) {
	task := r.PostFormValue("task")
	tmpl := template.Must(template.ParseFiles("templates/todo_ssr.html"))
	tmpl.ExecuteTemplate(w, "task-list-element", Task{Text: task})
}
```
\
The route will resolve to `/example/addtask`  
\
Subsequently, any folder post-fixed with an underscore **_** will use the parent folder in the route
```
app
 ├─ index.html
 └─ home_
    ├─ page.html
    └─ handle.go
``` 
\
This route will resolve to `/addtask`  

## Ignoring Handlers

If you want GoX to ignore an exported function in your `handle.go` file, simply post-fix your function name with an underscore **_**. 

```go
package example

func AddTask_(w http.ResponseWriter, r *http.Request) {
  // Perform an operation
}
```  
\

Then you have the freedom to handle the routing on your own

```go
// main.go

import "path-to-project/app/example"

r.Handle("/example/").HandlerFunc(example.AddTask_)
```
