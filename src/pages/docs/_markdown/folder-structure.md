# Routing Fundamentals

The skeleton of every application is routing. This page will introduce you to the fundamental concepts of routing for the web and how to handle routing in Gox.

## Folder Structure

```
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
- **Files** are used to create UI and handle behavior for a route segment.

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
- `handle.go` Sub-route handlers unique to the route  
- `render.go` Sub-route pages unique to the route
