# Custom Rendering

GoX provides an easy way to render & serve your page and html partials


## render.go

*__GoX__* looks for 2 types of functions in `render.go` then runs them at build-time.

1. #### [`func Render`](#func-render)
2. #### [`func SubRoute_`](#func-subroute-)

Both of these functions should either return filepaths or template objects for rendering. See return types below:

- FileRender: Render HTML from filepaths
```go
// - Data: A struct containing the data you want executed in your template
// - StrArr: A list of strings, where each string represents represents the path to a .html file you want executed
// - Str: A string that indicates the template you want executed. Use "" for no template execution
type RenderFilesType struct {
	Data  interface{}
	StrArr []string
	Str    string
}
```

- Template: Render HTML from a template object  
```go
// - Data: A struct containing the data you want executed in your template
// - Tmpl: A *template.Template object
// - Str: A string that indicates the template you want executed. Use "" for no template execution
type RenderTemplateType struct {
	Data interface{}
	Tmpl  *template.Template
	Str   string
}
```

\
Consider the following folder structure for the documentation below:
```bash
app
 ├─ index.html
 ├─ example
 │  ├─ page.md
 │  └─ render.go
 └─ templates
    └─ article.html
```

## `func Render`

Typically used when [Pages](#pages) isn't sufficent for your route rendering process, necessitating your own implementation. This does override [Pages](#pages), so the [*__GoX Router__*](#gox-router) will simply serve the output.

- Below resolves to
__*/example*__

```go
// NOTE - The functionality below can be done through Pages and Data

func Render() utils.RenderFilesType {
	// Assuming you're already connected to a DB 
	vars := mux.Vars(r)
	articleID := vars["articleID"]

	var article Article
	err := db.Get(&article, "SELECT id, title FROM articles WHERE id = $1", articleID)

  files := []string{"app/index.html", "templates/article.html"}
	return utils.RenderFileType{article, files..., ""}
}
```

## `func SubRoute_`

`'Subroute'` is a placeholder for your own function name. Postfixing a function with an underscore defines a static sub-route, with the name of the function resolving to the path. These should primarily be used for serving html partials.

- Below resolves to
__*/example/addarticle*__

```go
func AddArticle_() utils.RenderFilesType {
	return utils.RenderFilesType{ Article{"Article", "Description"}, []string{"templates/article.html"}, "article" }
}
```

\
Subsequently, any folder post-fixed with an underscore will use the parent folder in the route
```bash
app
 ├─ index.html
 ├─ example_
 │  ├─ page.md
 │  └─ render.go
 └─ templates
    └─ article.html
```

\
This sub-route will resolve to __*/addarticle*__  
