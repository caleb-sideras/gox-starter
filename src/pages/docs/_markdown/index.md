## index.html

A index.html is UI that is shared between multiple pages. On navigation, layouts preserve state, remain interactive, and do not re-render. Layouts can also be nested.

You can define a layout by creating a index.html file and including a Go html/template body '{{template}' or '{{block}}' tag to be populated later by a route.

GoX supports dynamic updating of the metadata in the head of your html based on the route you navigate to.

Additionally, you can use your own '{{template}' or '{{block}}' tag and include the relavent file in your respective data.go file.
 
```html
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  {{ block "metadata" . }}{{end}}
</head>

<body class="flex flex-col">
  <header class="lg:sticky lg:top-0 lg:left-0">
    {{ block "nav" . }}{{end}}
  </header>
  <main>
    {{ block "body" . }}{{end}}
  </main>
</body>

</html>
```
