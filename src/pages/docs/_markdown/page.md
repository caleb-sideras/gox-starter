## page.html

A page is UI that is unique to a route. Use nested folders to define a route and a page.html file encapsulated in {{define "template"}}{{end}} to make the route publicly accessible.

Create your first page by adding a page.html file inside the app directory:

```html
{{define "body"}}

<div>
<h1>Hello World!</h1>
</div>

{{end}}
```