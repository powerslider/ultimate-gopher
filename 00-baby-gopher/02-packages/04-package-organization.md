# Package Organization

- [Overview](#overview)
- [Project Layout Approaches](#project-layout-approaches)
    - [Flat Structure](#flat-structure)
    - [Anti Pattern: Group by Function (Layered Architecture)](#anti-pattern-group-by-function-layered-architecture)
    - [Anti Pattern: Group by Module](#anti-pattern-group-by-module)
    - [Group by Context (Domain Driven Design)](#group-by-context-domain-driven-design)
- [Conclusion](#conclusion)

## Overview

- Unlike other languages, Go does not allow _circular package imports_.
- Projects require additional planning when grouping code into packages to ensure that dependencies do not import each
  other.
- Inevitably, every developer in Go asks the following question: _How do I organize my code?_
- There are a number of articles and approaches, and while some work well for some, they may not work well for others.

## Project Layout Approaches

- No approach is perfect, but there are a few that have gained some widespread adoption. Please check the following
  resources:
    - [Ben Johnson's](https://twitter.com/benbjohnson) article on
      [Standard Package Layout](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1).
    - [Golang Standards Project Layout](https://github.com/golang-standards/project-layout).
    - [Kat Zien's](https://twitter.com/kasiazien) excellent talk, presentation and code samples found
      [here](https://github.com/katzien/go-structure-examples).

### Flat Structure

- Rather than spending time trying to figure out how to break code into packages, an app with a flat structure would
  just place all of the `.go` files in a single package.
- At first this sounds awful, because we do not use packages which separate concerns while making it easier to navigate
  to the correct source files quickly.

#### Recommendations

- When using a flat structure you should still try to adhere to coding best practices. Here are some helpful tips:
    - You want to separate different parts of your application using different `.go` files:
    ```
    reservationapp/
      customer.go
      data.go
      handlers.go
      main.go
      reservation.go
      server.go
      storage.go
      storage_json.go
      storage_mem.go
    ```
    - Globals can still become problematic, so you should consider using types with methods to keep them out of your
      code:
    ```go
    package main 
  
    import(
        "net/http"
        "some"
        "someapi"
    )

    type Server struct {
      apiClient *someapi.Client
      router *some.Router
    }

    func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        s.router.ServeHTTP(w, r)
    }
    ```
    - And your `main()` function should probably still be stripped of most logic outside of setting up the application.
    - A possible improvement to the flat structure is to again put all of your code in a single package, but separate
      the `main` package. where you define the entrypoint of the application. This would allow you to use the common
      `cmd` subdirectory pattern:
    ```
    reservationapp/
      cmd/
        web/
          # package main
          main.go
        cli/
          # package main
          main.go
      # package reservationapp
      server.go
      customer_handler.go
      reservation_handler.go
      customer_store.go
      ...
    ```

### Anti Pattern: Group by Function (Layered Architecture)

- Layered architecture patterns are n-tiered patterns where the components are organized in _layers_.
- This is the traditional method for designing most software and is meant to be _self-independent_.
- This means that all the components are interconnected but do NOT depend on each other.
- We have all heard about the famous 3-tier MVC (Model-View-Controller) architecture where we split our application in
  the following 3 distinct layers:
    - _Presentation / User Interface (View)_
    - _Business Logic (Controller)_
    - _Storage / External Dependencies (Model)_
- This architecture translated into our example project layout would look like this:

```
reservationapp/
  # package main
  data.go
  handlers/
    # package handlers
    customers.go
    reservations.go
  # package main
  main.go
  models/
    # package models
    customer.go
    reservation.go
    storage.go
  storage/
    # package storage
    json.go
    memory.go
  ...
```

- Actually this type of layout should NOT compile at all due to circular dependencies.
- The `storage` package uses the `models` package to get the definitions for a `Customer` and a `Resevation` and the
  `models` package uses the `storage` package to make calls to the database.
- Another disadvantage of this structure is that it does not guide us about what the application actually does (at least
  not more than the flat structure).
- This type of layout is strongly NOT recommended when writing applications in Go so try to avoid it.

### Anti Pattern: Group by Module

- Grouping by Module offers us a slight improvement over the layered approach:

```
reservationapp/
  customers/
    # package customers
    customer.go
    handler.go
  # package main
  main.go
  reservations/
    # package reservations
    reservation.go
    handler.go
  storage/
    # package storage
    data.go
    json.go
    memory.go
    storage.go
  ...
```

- Now our application is structured logically, but that is probably the only advantage of this approach.
- It is still hard to decide, e.g., if `reservation` should go to the `customers` package because they are customer
  reservations or are they suited for having their own package.
- Naming is worse because we now have `reservations.Reservation` and `customers.Customer` which introduces stutter.
- Worst of all is the possibility for circular dependencies again if the `reservations` package needs to reference the
`customers` package and vice versa.

### Group by Context (Domain Driven Design)

## Conclusion

- No single right answer (sorry...)!
- "As simple as possible, but not simpler!"
- Maintain consistency.
- Flat and simple is ok for small projects.
- Two top-level directories:
    - `cmd` (for your binaries)
    - `pkg` (for your packages)
- Group by context, not generic functionality. Try `DDD`/`hex`.
- Dependencies: own packages.
- Mocks: shared subpackage.
- All other project files (fixtures, resources, docs, Docker, etc.): root dir of your project.
- `main` package initializes and ties everything together.
- Avoid global scope and `init()`.