# Index

*__GoX__* provides a way to have UI that is shared between multiple pages

## index.html

On navigation, if your desired route is within your current `index.html` group, it will preserve state, remain interactive, and not re-render. Layouts can also be nested.


* Consider the following structure:
```bash
app                  (1)
 ├─ index.html       
 └─ home             (2)
    ├─ page.html     
    └─ about         (3)
       ├─ index.html 
       ├─ page.html  
       └─ data.go    
```

- It resolves to:

1. *__/home__*: `index.html` (1) &larr; `page.html` (2)

2. *__/home/about__*: `index.html` (3) &larr; `page.html` (3) &larr; `data.go` (3)

You can define a index by creating an `index.html` file and including *template* or *block* tags that will be populated later by a route. __GoX__ by default supports dynamic updating of the __body__ and __metadata__. Simply add a `page.html` and `metadata.html` to your desired route and populate it with your data.

```html
<!DOCTYPE html>
<html lang="en">

<head>
  {{`{{ block 'metadata' . }}`}} {{`{{end}}`}}
</head>

<body>
  {{`{{ block 'page' . }}`}} {{`{{end}}`}}
</body>

</html>
```



When using your own *template* or *block* tags, simply include the relavent filepath in your respective `data.go` file. See below:


- index.html
```html
<header>
  {{`{{ block 'nav' . }}`}} {{`{{end}}`}}
</header>
```

- data.go
```go
var Data utils.PageData = utils.PageData{
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