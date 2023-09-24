# GoX Router

The *__GoX Router__* controls what __HTML__ is returned to the client based on a variety of factors.


## Conditions

Because *__GoX__* allows you to define various types of routes, certain checks need to run server-side to select the most appropriate response based on the request. These factors include:

- Index reload or insertion

- Body or full page response

- ETags

- Page vs Data vs Render vs Handler
 

## Overiding

You can overide the the *__GoX Router__* and handle the request based on how you see fit. This can be done in the following ways:

- [`func Handle()`]() in `handle.go`
- Define your own handlers in `main.go`
