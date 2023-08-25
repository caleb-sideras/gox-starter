# Client Side Routing

The **_GoX Stack_** uses client side routing to prevent a full page reload. This is powered by HTMX and Alpine.js, paired with an additional body template.

- Define all of your routes using HTMX and Alpine.js in the navigation tab. These routes must listen for an event dispatched from anywhere within the document.
```javascript
// These following code listens for a "navigationtabclickedhome" event, which triggers "hx-get" 
<a hx-get="/get-home" hx-target="#body" hx-swap="innerHTML" x-data
  @navigationtabclickedhome="$el.click(); changeUrl('/');"/>
```

- Have your button dispatch the relavent event
```javascript
// Element calls function or dispatches directly
<md-navigation-tab id="home" label="Home" @click="handleTabClick"/>

handleTabClick(event) {
  // Perform some logic here
  $dispatch(`navigationtabclicked${event.target.id}`);
}
```

- Each Go handler should check the header for the "HX-request", and then only render the body.
```go
r.HandleFunc("/home", GetHome)

func GetHome(w http.ResponseWriter, r *http.Request) {
  r.Header.Get("HX-Request")
}
```