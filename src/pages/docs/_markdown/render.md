# Custom Rendering with GoX

GoX supports bespoke rendering processes for routes, ensuring flexibility and control over your application's responses.

## render.go

In `render.go` *__GoX__* specifically searches for two function types and executes them at build time:

1. #### [`func Render`](#func-render)
2. #### [`func SubRoute_`](#func-subroute-)

For both functions, GoX expects specific return types:

### Dynamic Rendering
Renders __HTML__ and integrates it with the [GoX router](#gox-router). Your page will be grouped and parsed with the closest `index.html` in your directory.

```go
type DynamicF struct {
	Templates []string            // Paths to the .html files intended for execution
	Content   interface{}         // Data to be executed within your template
}

type DynamicT struct {
	Template *template.Template   // A single template object
	Content  interface{}          // Data to be executed within your template
}
```

### Static Rendering 
This approach renders __HTML__ which will be served consistently upon every request.

```go
type StaticF struct {
	Templates []string            // Paths to the .html files intended for execution
	Content   interface{}         // Data to be used within your template
	Name      string              // Name of the desired template. If no execution is needed, use ""
}

type StaticT struct {
	Template *template.Template   // A single template object
	Content  interface{}          // Data to be used within your template
	Name     string               // Name of the desired template. If no execution is needed, use ""
}
```



### Example

For context, consider the folder hierarchy below:

```bash
app
 ├─ index.html
 ├─ example
 │  └─ render.go
 └─ templates
    └─ article.html
```

## `func Render`

It acts as your route handler, superseding any existing `page.html` within your route.

- Below resolves to
__*/example*__

```go
func Render() render.StaticF {
	// Assuming you're already connected to a DB 
	vars := mux.Vars(r)
	articleID := vars["articleID"]

	var article Article
	err := db.Get(&article, "SELECT id, title FROM articles WHERE id = $1", articleID)

  files := []string{"app/index.html", "templates/article.html"}
	return render.StaticF{files..., article, ""}
}
```

## `func SubRoute_`

Appending an underscore to a function declares it as a sub-route. The function's name is then treated as the path segment, with camel-casing being translated into hyphen-separated words. This can be harnessed for both full pages and __HTML__ partials.

- Below resolves to
__*/example/add-article*__

```go
func AddArticle_() render.StaticF {
	return render.StaticF{[]string{"templates/article.html"}, Article{"Article", "Description"}, "article"}
}
```


### Directory Naming Conventions

It's important to note that any folder named with an ending underscore will utilize the parent directory in its route.

```bash
app
 ├─ index.html
 └─ example_
    └─ render.go
```

\
The above sub-route would be accessible at __*/add-article*__  
