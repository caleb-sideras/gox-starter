# Custom Handling with GoX

*__GoX__* makes it easy to set up custom handlers for routes and sub-routes. With `handle.go`, you can manage how __HTML__ partials or full pages are served, giving you complete control over the process.

## handle.go

There are two types of functions *__GoX__* looks for  `handle.go`:

1. #### [`func Handle`](#func-handle)
2. #### [`func SubRoute_`](#func-subroute-)

### Example

For better understanding, consider this directory hierarchy:
```bash
app
 ├─ index.html
 ├─ example
 │  ├─ page.md
 │  └─ handle.go
 └─ templates
    └─ article.html
```

## `func Handle`

When a request is made to your route, `func Handle` is directly invoked. This function supersedes all other user-defined handling provided through [Pages](#pages), [Data](#data-handling) & [Render](#custom-rendering).

- Below resolves to
__*/example*__

```go
// NOTE - The functionality below can be done through the Pages and Data

func Handle(w http.ResponseWriter, r *http.Request) {
	// Assuming you're already connected to a DB 
	vars := mux.Vars(r)
	articleID := vars["articleID"]

	var article Article
	err := db.Get(&article, "SELECT id, title FROM articles WHERE id = $1", articleID)
	// Handle if err

	tmpl := template.Must(template.ParseFiles("app/index.html", "templates/article.html"))
	tmpl.Execute(w, article)
}
```

## `func SubRoute_`

Appending an underscore to a function signals it as a sub-route. The function's name then defines the path, with camel-casing being split into hyphen-separated terms. This approach is especially useful for serving HTML partials.

- Below resolves to
__*/example/add-article*__


```go
func AddArticle_(w http.ResponseWriter, r *http.Request) {
	articleID := r.PostFormValue("articleID")
	articleName := r.PostFormValue("articleName")

	tmpl := template.Must(template.ParseFiles("templates/article.html"))
	tmpl.ExecuteTemplate(w, "article-element", Article{articleID, articleName})
}
```

### Directory Naming Conventions

It's important to note that any folder named with an ending underscore will utilize the parent directory in its route.

```bash
app
 ├─ index.html
 ├─ example_
 │  ├─ page.md
 │  └─ handle.go
 └─ templates
    └─ article.html
```

\
The above sub-route would be accessible at __*/add-article*__  