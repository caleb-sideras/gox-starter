# Data Handling

*__GoX__* provides a way to separate data from your __HTML,__ powering re-usuable components with dynamic data fetching 

## data.go

`data.go` is used to define the templates and content you want to be parsed and populated with your [Page](#pages) at build or request time. There are 2 things *__GoX__* looks for in your `data.go` files.

1. #### [`var Data`](#var-data) 
2. #### [`func Data`](#func-data)

## `var Data`

On build, *__GoX__* searches for an exported variable called `Data` with the following type outlined below.

```go
// .gox/utils

type PageData struct {
	Content   interface{}
	Templates []string
}
```

*__GoX__* will populate your page with the data from __Content__ and parse your page with your defined __Templates__. This will result in a static __HTML__ page being rendered.

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
var Data utils.PageData = utils.PageData{
  Content:  Card{
    Title: "Hello World",
    Description: "This is a test!"
  },
  Templates: []string{
    "templates/card.html",
  },
}
```

### Resolves to:
1. Path: *__/home__*  
2. File: `index.html` (1) &larr; `page.html` (2) &larr; `data.go` (2)
```html
<!DOCTYPE html>
<html lang="en">

<body>
  <h1>Hello World</h1>
  <p>This is a test!</p>
</body>

</html>	
```

## `func Data`

If your page has dynamic data simply define `func Data`. This function runs per request and is used with your [Pages](#pages), with its return value being used to populate your page.

```go
type Article struct {
	ID    int    `db:"id" json:"id"`
	Title string `db:"title" json:"title"`
}

func Data(w http.ResponseWriter, r *http.Request) any {
	// Assuming you're already connected to a DB 
	vars := mux.Vars(r)
	articleID := vars["articleID"]

	var article Article
	err := db.Get(&article, "SELECT id, title FROM articles WHERE id = $1", articleID)
	// Handle if err

	// Must return data
	return article
}
```