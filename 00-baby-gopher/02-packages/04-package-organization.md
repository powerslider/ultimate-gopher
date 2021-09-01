# Package Organization

- [Overview](#overview)
- [Project Layout Approaches](#project-layout-approaches)
    - [Flat Structure](#flat-structure)
    - [Anti Pattern: Group by Function (Layered Architecture)](#anti-pattern-group-by-function-layered-architecture)
    - [Anti Pattern: Group by Module](#anti-pattern-group-by-module)
    - [Group by Context (Domain Driven Design)](#group-by-context-domain-driven-design)
    - [Group by Context (Domain Driven Design + Hexagonal Architecture)](#group-by-context-domain-driven-design--hexagonal-architecture)
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
- Now we are going to explore what options are available for structuring your applications and distinguish between some
  good and bad practices.
- Additionally, we will explore how to build robust mental models for reasoning about your problem domains and how to
  represent that in your project layouts.

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
- The `storage` package uses the `models` package to get the definitions for a `Customer` and a `Reservation` and the
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

- This way of thinking about your applications is called _Domain Driven Design (DDD)_.
- In its essence it guides you to think about the domain you are dealing with and all the business logic without even
  writing a single line of code.
- There are three main components that need to be defined:
    - _Bounded contexts_
    - _Models within each context_
    - _Ubiquitous language_

#### Bounded Contexts

- A bounded context is a fancy term defining limits upon your models.
- An example would be, e.g., a `User` entity that might have different properties attached to it based on the context:
    - A `User` in a sales department context might have properties like `leadTime`, `costOfAquisition`, etc.
    - A `User` in a customer support context might have properties like `responseTime`, `numberOfTicketsHandled`, etc.
    - This showcases that a `User` means different things to different people, and the meaning depends heavily on the
      _context_.
- The bounded context also helps in deciding what has to stay consistent within a particular boundary and what can
  change independently.
    - If we decide later on to add a new property to the `User` from the sales department context, that would not affect
      the `User` model in the customer support context.

#### Ubiquitous Language

- _Ubiquitous Language_ is the term used in _Domain Driven Design_ for the practice of building up a common, rigorous
  language between developers and users.
- This language is based on the _Domain Model_ used in the software, and it evolves up to the point of being able to
  express complex ideas by combining simple elements of the _Domain Model_.

#### Categorising the building blocks

- Based on the DDD methodology we will now start to reason about our domain by constructing its building blocks.
- If we take our _Restaurant Reservation System_ example, we would have the following elements:
    - __Context__: Booking Reservations.
    - __Language__: reservation, customer, storage, ...
    - __Entities__: Reservation, Customer, ...
    - __Value Objects__: Restaurant, Host, ...
    - __Aggregates__: BookedReservation
    - __Service__: Reservation lister/listing, Reservation adder/adding, Customer adder/adding, Customer lister/listing,
      ...
    - __Events__: ReservationAdded, CustomerAdded, ReservationAlreadyExists, ReservationNotFound, ...
    - __Repository__: ReservationRepository, CustomerRepository, ...
- Now after defining those blocks we can translate them in our project layout:

```
reservationapp/
  adding/
    endpoint.go
    service.go
  customers/
    customer.go
    sample_customers.go
  listing/
    endpoint.go
    service.go
  main.go
  reservations/
    reservation.go
    sample_reservations.go
  storage/
    json.go
    memory.go
    type.go
```

- The main advantage here is that our packages now communicate what they PROVIDE and not what they CONTAIN.
- This makes it easier to avoid _circular dependencies_, because:
    - `adding` and `listing` talk to `storage`.
    - `storage` pulls from `customers` and `reservations`.
    - Model packages like `reservations` and `customers` do not care about `storage` directly.

### Group by Context (Domain Driven Design + Hexagonal Architecture)

- So far we managed to structure our application according to _DDD_, eliminated _circular dependencies_ and made it
  intuitive what each package does only by looking at the directory and file names.
- We still have some problems:
    - How can we start a version of our application that contains sample data?
    - In our current version sample data is bundled with the application's entrypoint in `main.go` and we have only one
      `main.go`.
    - We have no option to run a sample data version separately from the main version of the application.
    - Maybe we want a pure cli version of the application where instead of adding reservations through HTTP requests, we
      want the command line to prompt us for each reservation property?

#### Hexagonal Architecture

- This type of architecture distinguishes the parts of the system which form your _core domain_ and all the external
  dependencies are just implementation details.
- External dependencies could be databases, external APIs, mail clients, cloud services etc., anything that your
  application interacts with.
- The problem this solves is giving you the ability to change one part of the application without affecting the rest,
  e.g., swapping databases or transport protocols (HTTP to gRPC).
- This is not in any way similar to the _MVC (layered)_ model, because:
    - _MVC_ tends to look at inputs and outputs in a top to bottom way (input -> main logic -> output).
    - _Hex_ treats inputs and outputs on the same level. It does not care if something is an input or an output, it is
      just an external interface.
- The key rule in the _hex_ model is that dependencies only point __INWARDS__ (only __outer layers depend upon inner
  layers__ and not the other way around). This is called the _Dependency Inversion Principle_.
  Check [this excellent article](https://martinfowler.com/articles/dipInTheWild.html) by Martin Fowler to learn more.
- Based on this approach our project structure could look like this:

```
reservationapp/
  cmd/
    # HTTP server
    reservation-server/
      main.go
    # CLI app
    reservation-cli/
      main.go
    # HTTP server with seeded data
    reservation-sample-data/
      main.go
      sample_reservation.go
      sample_customers.go
  pkg/
    adding/
      reservation.go
      endpoint.go
      service.go
    listing/
      customer.go
      reservation.go
      endpoint.go
      service.go
    transport/
      http/
        server.go
    main.go
    storage/
      json/
        customer.go
        repository.go
        reservation.go
      memory/
        customer.go
        repository.go
        reservation.go
```

- To solve the multiple app version binaries problem we utilize the `cmd` subdirectory pattern which we mentioned as an
  improvement to the _flat structure_ layout.
- We are now able to produce 3 different binaries used to serve different purposes:
    - `reservation-server` - main version of the app deploying an HTTP server.
    - `reservation-cli` - a CLI version with a removed transport layer offering a CLI interface for interaction.
    - `reservation-sample-data` - a sample data seeded version used mainly for testing.
- We introduce the `pkg` package which separates our Go code from the `cmd` binaries and non-code resources, e.g. DB
  scripts, configs, documentation, etc. which should be found on the same level under project root.

---
__NOTE__
> Using the `cmd` and `pkg` directories has become somewhat of a trend in the Go community. It is not a standard
> by any means, but a good recommendation that should definitely should be considered.
---

- According to `DDD` we keep the `adding` and `listing` packages which represent our core domain.
- We remove `reservations` and `customers` packages and instead introduce models in each of the core domain packages,
  e.g., `adding.Reservation`, `adding.Customer`, `listing.Reservation`, etc.
- The advantage here is that we have separate representations per model according to the _bounded context_
  (`adding` or `listing`). This allows decoupled model modification and avoids _circular dependencies_.
- We introduce a `transport` package which contains all transport protocol implementations, e.g. HTTP or maybe gRPC in
  their own respective subpackages.
- The `storage` package is another _bounded context_ which features its own model representations on a storage level and
  subpackages for storage implementations, e.g. `json`, `memory`, etc.
- Again, `main.go` ties everything together and should not contain any logic that would require testing.

---
__NOTE__
> Testing is another important aspect that we have not discussed yet and should be considered when structuring your projects, but we will skip it for
> now because we will have special chapters dedicated only to testing approaches, types and techniques.
---

## Conclusion

- No single right answer (sorry...)!
- "As simple as possible, but not simpler!"
- Maintain consistency.
- Flat and simple is ok for small projects.
- Two top-level directories:
    - `cmd` (for your binaries)
    - `pkg` (for your packages)
- Group by context, not generic functionality. Try _DDD/hex_.
- Dependencies: own packages.
- All other project files (fixtures, resources, docs, Docker, etc.) - root dir of your project.
- `main` package initializes and ties everything together.
- Avoid global scope and `init()`.

[Next Chapter](../03-arrays-and-iteration/README.md)

[Previous Section](03-workspaces.md)

[Chapter Overview](README.md)