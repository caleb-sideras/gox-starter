# Data Handling with GoX

*__GoX__* provides a built-in way to separate data from your __HTML__. This enables you to build reusable components and dynamically fetch data.

## data.go

The `data.go` file serves as the bridge between your data and [pages](/docs/pages). When building or processing requests, *__GoX__* primarily looks for two components within the `data.go` files:

1. #### [`var Data`](#static-data-with-var-data) 
2. #### [`func Data`](#dynamic-data-with-func-data)

## Static Data with `var Data`

During the build process, *__GoX__* scouts for an exported variable named `Data` of a predefined type:

```go
// .gox/data

type Page struct {
	Content   interface{}
	Templates []string
}
```

With the data from `Content` and the templates listed in `Templates`, *__GoX__* populates and parses your page, producing a static __HTML__ output.

### Example:

- Filepath
```bash
app               (1)
 ├─ index.html
 ├─ home          (2)
 │  ├─ page.html
 │  └─ data.go
 └─ templates
    └─ card.html
```

- page.html
```html
{{`{{ define page }}`}}

{{`{{ range .HomeCard }}`}}

  {{`{{ template 'card' . }}`}}

{{`{{ end }}`}}

{{`{{ end }}`}}
```

- card.html
```html
{{`{{define 'card'}}`}}

<h1>{{`{{ .Title }}`}}</h1>
<p>{{`{{ .Description }}`}}</p>

{{`{{ end }}`}}
```

- data.go
```go
package home

import "path-to-project/.gox/utils"

// Example struct
type Card struct {
	Title       string
	Description string
}

// Exported variable used at build time
var Data data.Page = data.Page{
  Content:  Card{
    Title: "Hello World",
    Description: "This is a test!"
  },
  Templates: []string{
    "templates/card.html",
  },
}
```

### This resolves to the following:
1. URL Path: *__/home__*  
2. Render Hierarchy: `index.html` (1) &larr; `page.html` (2) &larr; `data.go` (2)
```html
<!DOCTYPE html>
<html>

<body>
  <h1>Hello World</h1>
  <p>This is a test!</p>
</body>

</html>	
```

## Dynamic Data with `func Data`

For content that is dynamic and fetched per-request, *__GoX__* offers `func Data` with a specific return type:

```go
// .gox/data

type PageReturn struct {
	Page
	Error error
}
```
\
As of now, if an `Error` arises, *__GoX__* communicates it directly to the client.


```go
type Article struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

func Data(w http.ResponseWriter, r *http.Request) PageReturn  {
	// Assuming you're already connected to a DB 
	vars := mux.Vars(r)
	articleID := vars["articleID"]

  var article Article
  _ := db.Get(&article, "SELECT id, title FROM articles WHERE id = $1", articleID)

  return data.PageReturn{data.Page{article, ""}, nil}
}
```