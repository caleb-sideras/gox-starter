# Single Page Application (SPA) vs Multiple Page Application (MPA)

The **_"GoX Stack"_** uses a hybrid approach utilizing both `SPA` and `MPA` techniques, with the main distinction being in how navigation and page updates are handled.

## Single Page Applications (SPA)

An SPA dynamically updates the current page in response to user interaction, rather than loading entire new pages from the server.

The GoX stack uses this to:

1. Fetch and inject new HTML into the current page with HTMX, avoiding a full page reload.
2. Client-side routing to manage navigation within the app without causing a full page reload.

```javascript
// Element calls function
<md-navigation-tab id="home" label="Home" @click="handleTabClick"/>

handleTabClick(event) {
  // Checks if tab is active
  $dispatch(`navigationtabclicked${event.target.id}`);
}

// HTMX replaces body and changes url
<a hx-get="/get-home" hx-target="#body" hx-swap="innerHTML" x-data
  @navigationtabclickedhome="$el.click(); changeUrl('/');"/>
```

### Benefits of SPA in GoX stack:

- **Smooth User Experience:** SPA techniques in GoX can provide a seamless user experience, similar to a desktop application. Transitions between pages can be made smoother and faster with HTMX.
- **Less Bandwidth Usage:** Only necessary content is updated which reduces the data that needs to be transferred, beneficial for users on slow networks.
- **State Management:** SPAs can keep state in memory which can simplify certain programming tasks (Client-side routing).


## Multiple Page Applications (MPA)

An MPA involves multiple server routes where each route corresponds to a separate HTML page.

The GoX stack uses this to:

1. Server-side render the HTML of each page using Go's http/template.
2. Server-side routing to handle navigation between different server-rendered pages.

### Benefits of MPA in GoX stack:
- **SEO**: MPAs are traditionally better for SEO, as each page has its own unique URL, and the content for each page is fully rendered on the server.
- **Scalability**: MPAs can be more straightforward to scale horizontally, as each page can be served by different server instances.
- **Fallback**: MPAs have a better fallback when JavaScript is disabled or fails to load.

## Go Handlers

As we will have duality of using both SPA and MPA techniques, each handler should check the header for a "HX-request", and then only render the body.
