# Handled Pre-Rendered Pages

Using GoX's `page.html` and `data.go` routing pattern will result in your page being rendered to html at build time and served based on the file path.

```
app
 ├─ index.html
 └─ home_
    ├─ page.html
    └─ data.go
```

## page.html

A page is UI that is unique to a route. Use nested folders to define a route and a page.html file encapsulated in a define template to make the route parsible by your index.html. 

```html
{{`{{ define body }}`}}

<div>
	<h1>Hello World!</h1>
	{{`{{ range .HomeCard }}`}}
	  <span>{{`{{ .Title }}`}}</span>
	  <span>{{`{{ .Description }}`}}</span>
	{{`{{ end }}`}}
</div>

{{`{{ end }}`}}
```

## data.go

`data.go` is used to define the templates and content you want to be parsed and populated with your page at build time.   
\
Define an exported variable called Data with the following type outlined below:

```go
package home

import "path-to-project/.gox/utils"

type HomeCard struct {
	Title       string
	Description string
}

var content HomeCard = HomeCard{
	Title: "Hello World",
	Description: "This is a test!"
}

var templates []string = []string{
	"templates/components/nav.html",
}

var Data utils.PageData = utils.PageData{
	Content:   content,
	Templates: templates,
}
```
\
`PageData` type defined in .gox/utils

```go
type PageData struct {
	Content   interface{}
	Templates []string
}
```

## HTMX

As GoX is a HTMX centered framework, any page that follows this pattern will have an additional `page-body.html` rendered at build-time to allow client-side routing using `HTMX`.

```
└─ public
   └─ html
      └─ home
         ├─ page.html
         └─ page-body.html
```
