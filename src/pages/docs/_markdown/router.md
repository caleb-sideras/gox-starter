# GoX Router

The purpose of the GoX router is to preserve state on the client by controlling what is returned. This is established through [Index](/docs/index) groups. 


## Navigation

To access the *__GoX Router__* when navigating, use the __HTMX__ `hx-boost` attribute. *__GoX__* hijacks this, preventing full page swaps if index groups (i.e. the same `index.html`) are shared between routes.


### Example

In the directory structure below:

```bash
app               (1)
 ├─ index.html
 ├─ home          (2)
 │  └─ page.html
 └─ docs          (3)
    └─ page.html
```

The routes resolves to:

1. *__/home__* &larr; `index.html` (1) &larr; `page.html` (2)
2. *__/docs__* &larr; `index.html` (1) &larr; `page.html` (3)


Both *__/home__* and *__/docs__* use the same `index.html`. Navigation between these routes with `hx-boost` will always return the needed __HTML__ partials, replacing the body and thus maintaining state. 

\
However, when the routes have distinct index files:

```bash
app               (1)
 ├─ index.html
 ├─ home          (2)
 │  └─ page.html
 └─ docs          (3)
    ├─ index.html
    └─ page.html
```

Which resolves to:

1. *__/home__* &larr;  `index.html` (1) &larr; `page.html` (2)
2. *__/docs__* &larr;  `index.html` (3) &larr; `page.html` (3)

Therefore navigating between *__/home__* and *__/docs__* with `hx-boost` will always return the whole page. 

## Caching

To optimize performance and reduce redundant data transfers, the GoX Router employs various browser caching headers. This ensures that duplicate __HTML__ content isn't resent if it's already stored on the client's side.

## Overiding

You can overide the the *__GoX Router__* and handle the request based on how you see fit.

### Directly access HTML partials

Sometimes, you might want to fetch just the __HTML__ partial, skipping the surrounding layout. *__GoX__* facilitates this by allowing you to set the `index` query parameter to `false` in your request. By default, all GET requests have this parameter set to `true`.

- For instance:

```bash
https://example.com/example?index=false
```

### Custom handlers

When you want more control over the returned content, routes equipped with `func Handle` or those utilizing `func Render` with __StaticF__ or __StaticT__ return types will always dispatch the entire page. This gives you flexibility when you want to bypass the default behavior of the GoX Router.