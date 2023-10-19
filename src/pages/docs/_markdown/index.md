# Index: Shared UI and State

*__GoX__* provides a way to have UI that is shared between multiple pages

## index.html

On navigation, routes that have the same `index.html` will preserve state, remain interactive, and not refresh the browser.

\
Consider this directory structure:
```bash
app                  (1)
 ├─ index.html       
 └─ home             (2)
    ├─ page.html     
    └─ about         (3)
       ├─ index.html 
       └─ page.html  
```

\
Which resolves to:

1. *__/home__* &rarr; `index.html` (1) &larr; `page.html` (2)

2. *__/home/about__* &rarr; `index.html` (3) &larr; `page.html` (3)

To establish a shared UI, create an `index.html` file and add a "page" *template* or *block* tag. Depending on the route navigated to, appropriate [pages](/docs/pages) will be populated. Additionally, custom *template* or *block* tags can be added and later populated by your routes [`data.go`](/docs/data) file. 

```html
<!DOCTYPE html>
<html>

<body>
  {{`{{ block 'page' . }}`}} {{`{{end}}`}}
</body>

</html>
```

When incorporating custom *template* or *block* tags, ensure you link the relevant filepath in the associated `data.go` file. For instance:

- index.html
```html
<header>
  {{`{{ block 'nav' . }}`}} {{`{{end}}`}}
</header>
```

- data.go
```go
var Data data.Page = data.Page{
	Templates: []string{
  	"templates/nav.html",
  }
}
```

- nav.html
```html
{{`{{define 'nav'}}`}}

<div>navigation</div>

{{`{{end}}`}}
```