### Folder hierarchy

```bash
app
 ├─ index.html
 └─ render
    ├─ render.go
    └─ example.md
```

### render.go

```go
package render

import (
	"github.com/caleb-sideras/goxstack/gox/render"
	"github.com/caleb-sideras/goxstack/src/global"
	"html/template"
)

func Markdown_() render.StaticT {
	tmpl := template.Must(global.MarkdownToHTML("app/render/example.md"))

	return render.StaticT{tmpl, nil, ""}
}
```

### example.md

```md
## This is rendered Markdown!

reasons to use *__GoX__*

1. cool mascot 
2. htmx is good for the environment
3. gox sounds like c*cks
```
