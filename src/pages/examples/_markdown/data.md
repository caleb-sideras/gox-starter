### Folder hierarchy

```bash
app
 ├─ index.html
 └─ data
    ├─ page.html
    └─ data.go
```

### page.html

```html
<div class="text-9xl font-bold m-auto px-8 ">
  {{`{{.Number}}`}}
</div>
```

### data.go

```go
package todo

import (
	"net/http"
	"github.com/caleb-sideras/goxstack/gox/data"
)

func Data(w http.ResponseWriter, r *http.Request) data.PageReturn {
	randomNumber, err := fetchRandomNumber()
	if err != nil {
		return data.PageReturn{data.Page{}, err}
	}

	return data.PageReturn{
		Page: data.Page{
			Content:   Number{randomNumber},
			Templates: []string{},
		},
		Error: nil,
	}
}
```