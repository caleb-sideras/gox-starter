# Pages

*__GoX__* Pages are a simple way to handle routing to either dynamic or static routes 

## page.html

A page is UI that is unique to a route. Use nested folders to define a route and a `page.html` file surrounded in a __page__ block to make the route parsible by your `index.html`. Using a single `page.html` in your route will result in your page being rendered to __HTML__ at build time and served based on the file path.   



### Example:

- Filepath
```bash
app               (1)
 ├─ index.html
 └─ home          (2)
    └─ page.html
```

- index.html (1)
```html
<!DOCTYPE html>
<html lang="en">

<body>
    {{`{{ block 'page' . }}`}} {{`{{end}}`}}
</body>

</html>	
```

- page.html (2)
```html
{{`{{ define 'page' }}`}}

<h1>Hello World!</h1>

{{`{{ end }}`}}
```

- Resolves to:
1. Path: *__/home__*  
2. File: `index.html` (1) &larr; `page.html` (2)

```html
<!DOCTYPE html>
<html lang="en">

<body>
	<h1>Hello World!</h1>
</body>

</html>	
```


## HTMX
As GoX is a __HTMX__ centered framework, any page that follows this pattern will have an additional `page-body.html` rendered at build-time. All page based routes will have an __ETag__ in their header to prevent browser caching issues.

```
└─ public
   └─ html
      └─ home
         ├─ page.html
         └─ page-body.html
```
\
To access to `page-body.html` simply use __HTMX__ for the request to the same route
