# Custom Handling

GoX provides an easy way to define route & sub-route handlers used for html partials or full page reloads


## handle.go

There are 3 types of functions *__GoX__* looks for  `handle.go`:

1. #### [Init](#func-init) 
2. #### [Handle](#func-handle)
3. #### [SubRoute](#func-subroute-)


- Consider the following folder structure for the documentation below:
```bash
app
 ├─ index.html
 ├─ example_
 │  ├─ page.md
 │  └─ handle.go
 └─ templates
    └─ article.html
```

## `func Init`

Is called intially per request to your route. Typically used for route specific middleware.

```go
func Init(w http.ResponseWriter, r *http.Request){
	// Connecting to a Database
	var err error
	connStr := "user=username dbname=mydb sslmode=disable password=password"
	db, err = sqlx.Connect("postgres", connStr)
	// Handle err
}
```

## `func Handle`

Upon request to your route, the [__GoX Router__](gox-router) will just simply call `func Handle`. This function overides all other user-defined handling through [Pages](#pages), [Data](#data-handling) & [Render](#custom-rendering).

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

- Below resolves to
__*/example/addarticle*__

`'Subroute'` is a placeholder for your own function name. Postfixing a function with an underscore defines a dynamic sub-route, with the name of the function resolving to the path. These should primarily be used for serving html partials.

```go
func AddArticle_(w http.ResponseWriter, r *http.Request) {
	articleID := r.PostFormValue("articleID")
	articleName := r.PostFormValue("articleName")

	tmpl := template.Must(template.ParseFiles("templates/article.html"))
	tmpl.ExecuteTemplate(w, "article-element", Article{articleID, articleName})
}
```

\
Subsequently, any folder post-fixed with an underscore will use the parent folder in the route
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
This sub-route will resolve to __*/addarticle*__  