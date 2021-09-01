# Definition

- [Overview](#overview)
- [Package Names](#package-names)
- [File Names](#file-names)
- [Executable Packages](#executable-packages)
- [Package Resolution](#package-resolution)

## Overview

- Go code is organized in packages.
- Packages collect related code.
- They can be big or small and may be spread across multiple files.
    - The `net/http` package exports more than 100 names (18 files).
    - The `errors` package exports just one (1 file).
- All the files in a package live in a single directory.
- Our examples so far have all used package `main` declared at the top of the file.

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

- Keep package names short and meaningful.
- Do not use underscores `_`, they make package names long.
    - `suffixarray` not `suffix_array`
- Do not overgeneralize. A `util` or `common` package could be anything, so it is highly discouraged to name packages
  this way.
- The name of the package is part of its type and function names.
- Avoid stutter (e.g., `strings.Reader` not `strings.StringReader`). 
- On its own, type `Buffer` is ambiguous. But users see:

```go
buf := new(bytes.Buffer)
```

## File Names

- Inside packages there are no requirements as to what the names of the files inside that package are to be named.
- However, it is common practice to name the _entrypoint_ file after the name of the package.
- For example, a package named `storage` would probably have a `storage.go` file inside as the main _entrypoint_ to that
  package. It is also the file where you would write your top level comments that will show up when you generate your
  code documentation.

```
storage/
  storage.go
  driver.go
  ...
```

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

- For very small programs, `main` is the only package you need to write.

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

[Previous Chapter](../01-syntax-and-types/README.md)
