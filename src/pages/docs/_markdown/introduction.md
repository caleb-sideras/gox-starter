# GoX Introduction

#### __*GoX*__ is a framework designed to make working with [__HTMX__](https://htmx.org/) and [__Go HTML/Templates__](https://pkg.go.dev/html/template) easier. __*GoX*__ achieves this by implementing the following:

- #### [__File Based Routing__](#routing-fundamentals)

File based routing provides an intuitive way to visualize your routes and isolate their respective functionality - removing the need for various config files

- #### [__Page Routing__]()

__*GoX*__ separates your html and data for static pages, allowing you to create reusable components. Additionally, __*GoX*__ will create body partials with unique __ETags__ accessible via __HTMX__

- #### [__Partial Rendering__]()

Within your desired route, you can define various functions that return your desired components and data. These functions are executed at build time rendering your html partials. At runtime, all routing to these partials are handled by __*GoX*__

- #### [__Partial Handling__]()

Some partials need to be dynamically rendered. __*GoX*__ allows you to define functions within your desired routes to accomplish this. Once again, __*GoX*__ handles all the routing 
