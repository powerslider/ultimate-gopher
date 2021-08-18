# Package Organization

- [Overview](#overview)
- [Project Layout Approaches](#project-layout-approaches)
    - [Flat Structure](#flat-structure)
- [Conclusion](#conclusion)

## Overview

- Unlike other languages, Go does not allow _circular package imports_.
- Projects require additional planning when grouping code into packages to ensure that dependencies do not import each
  other.
- Inevitably, every developer in Go asks the following question: _How do I organize my code?_
- There are a number of articles and approaches, and while some work well for some, they may not work well for others.

## Project Layout Approaches

- No approach is perfect, but there are a few that have gained some widespread adoption. Please check the following 
  resources which have (but not only) inspired
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
    myapp/
      main.go
      server.go
      user.go
      lesson.go
      course.go
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
    myapp/
      cmd/
        web/
          # package main
          main.go
        cli/
          # package main
          main.go
      # package myapp
      server.go
      user_handler.go
      user_store.go
      ...
    ```

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