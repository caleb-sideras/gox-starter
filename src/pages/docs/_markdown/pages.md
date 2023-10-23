# Pages: Route-Specific UI

Pages are a simple way to define and handle route specific UI

## page.html

Every route can have a distinct UI defined within a `page.html` file.

### Example

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
<html>

<body>
    {{`{{ block 'page' . }}`}} {{`{{end}}`}}
</body>

</html>	
```

- page.html (2)
```html
<h1>Hello World!</h1>
```

\
This resolves to:

1. URL Path: *__/home__*  
2. Render Hierarchy: `index.html` (1) &larr; `page.html` (2)

```html
<!DOCTYPE html>
<html>

<body>
	<h1>Hello World!</h1>
</body>

</html>	
```


## HTMX
As GoX is a __HTMX__ centered framework, any page that follows this pattern will have an additional `page-body.html` rendered at build-time.

\
Your output directory would look something like:

```bash
└─ public
   └─ html
      └─ home
         ├─ page.html
         └─ page-body.html
```
\
To delve deeper into how __GoX__ leverages this, read [GoX Router](/docs/router).
