## Definition

- [Overview](#overview)
- [Package Names](#package-names)
- [Executable Packages](#executable-packages)
- [Package Resolution](#package-resolution)

## Overview

- Go code is organized in packages.
- A package represents all the files in a single directory on disk.
- One directory can contain only files from the same package.
- You've seen this already several times. Our examples so far have all used package `main` declared at the top of the
  file.

## Package Names

- A package can only have one name, but it is not required to be the same as the folder it is in.
- However, it is _strongly encouraged_ to _match the folder name_, e.g., code in the folder `bar`
  should declare `package bar`.

```go
package bar

// put your code here
```

---
__IMPORTANT__
> All source files (`.go`) must declare the package name at the top of the file. No exceptions!
---

## Executable Packages

- Executable programs must have a `main` package that declares a `main()` function:

```go
package main

func main() {
	// program entrypoint code
}
```

---
__IMPORTANT__
> The `main` function can only be declared ONCE and receives NO ARGUMENTS, nor does it RETURN any values.
---

## Package Resolution

- If your code lives at `$GOPATH/src/foo/bar` then it's package name should be `bar`.
- The import statement for the `bar` package would be:

```go
import "foo/bar"
```

- Packages that live in source code repositories, like GitHub and GitLab, have the full location of the repository as
  part of their import path.
- For example, the source code at [https://github.com/gorilla/mux](https://github.com/gorilla/mux) would be imported
  using the following `import` path:

```go
import "github.com/gorilla/mux"
```

- The source code would be in the following location on disk:

```
$GOPATH/src/github.com/gorilla/mux
```

[Next Section](02-scope-and-visibility.md)
