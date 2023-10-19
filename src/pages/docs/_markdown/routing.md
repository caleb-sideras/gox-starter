# Routing Fundamentals

The skeleton of every application is routing. This page will introduce you to the fundamental concepts of routing in __GoX__.

## Folder Structure

```bash
.
├─ .gox
│  ├─ main.go
│  ├─ structs.go
│  └─ generated.go
│     └─ utils
│        └─ utils.go
├─ public
│  ├─ favicon.ico
│  ├─ html
│  │  └─ home
│  │     ├─ page.html
│  │     └─ page-body.html
│  ├─ css
│  │  └─ output.css 
│  └─ js
│     └─ bundle.js
├─ src
│  ├─ main.go
│  ├─ app
│  │  ├─ index.html
│  │  └─ home_
│  │     ├─ page.html
│  │     ├─ data.go
│  │     ├─ handle.go
│  │     └─ render.go
│  ├─ server
│  │  └─ middleware.go
│  ├─ styles
│  │  ├─ app.css
│  │  └─ tailwind-material-colors.js
│  └─ templates
│     └─ components
│        └─ nav.html
├─ .env
├─ tailwind.config.js
├─ rollup.config.mjs
├─ index.js
├─ package.json
└─ README.md

```
## Roles of Folders and Files

GoX uses a file-system based router where:

- **Folders** are used to define routes. A route is a single path of nested folders, following the file-system hierarchy from the **root folder** down to a final **leaf folder**
- **Files** are used to create UI and drive the behavior associated with a particular route segment.

## Nested Routes

To create a nested route, you can nest folders inside each other. For example, you can add a new `/dashboard/settings` route by nesting two new folders in the `app` directory.  
\
The `/dashboard/settings` route is composed of three segments:

- `/` (Root segment)
- `dashboard` (Segment)
- `settings` (Leaf segment)

## File Conventions

GoX provides a set of special files to create UI with specific behavior in nested routes:

- `index.html` Shared UI for a segment and its children
- `page.html`	Unique UI of a route and make routes publicly accessible
- `data.go` Templates and page content needed to populate the UI
- `handle.go` Handlers unique to the route  
- `render.go` Render partials unique to the route

## Folder Conventions


### Underscore Postfix

Folders ending with an underscore will utilize the parent directory in its route. For example:

```bash
app
 ├─ index.html
 └─ home_
    └─ page.html
```

\
The above route would be accessible at the root &larr; __*/*__  

### Underscore Prefix

Folders starting with an underscore will simply be ignored by *__GoX__*
