# Scope and Visibility

- [Overview](#overview)
- [Security](#security)

## Overview

- Go does not have the concept of `public`, `private`, or `protected` modifiers like other languages do.
- External visibility is controlled by capitalization according to the following rules:
    - __Rule 1:__ Types, Variables, Functions, etc. that start with a capital letter are available, publicly, outside
      the current package.
    - __Rule 2:__ A symbol that is visible outside its package is _exported_.
    - __Rule 3:__ Types, Variables, Functions, etc. that start with a lower case letter are unexported and are not
      available outside the current package.
    - __Rule 4:__ All variables and types declared inside a package are visible to everything else in the same package.

```go
// visible outside of the package:
func Foo() {}

// available only inside the package:
var bar string
```

---
__NOTE__
> While it is possible to return a non-exported type from an exported function, it is considered bad practice.
>```go
>// available only inside the package:
>type bar struct {}
>
>// "legal", but not encouraged
>func Foo() bar {
>   return bar{}
>}
>```
>Consider returning an exported (capitalized) struct or interface instead.
---

## Security

- It is important to understand that although you cannot directly access or change unexported fields in a struct, you
  can still get access to view the contents of them.
- One example of this is the `fmt` package which makes use of the `reflect` package:

```go
package main

import (
	"fmt"
	"github.com/powerslider/ultimate-gopher/00-baby-gopher/02-packages/example/foo"
)

func main() {
	user := foo.NewUser("Homer", "Simpson", "s3cr37")
	// You can see the contents of the private information...
	fmt.Printf("%+v\n", user)
	// output:  {FirstName:Homer LastName:Simpson password:s3cr37}

	// You cannot access or change it directly....

	//fmt.Println(user.password)
	//user.password = "new"

	// output:
	//./main.go:16:18: user.password undefined (cannot refer to unexported field or method password)
	//./main.go:17:6: user.password undefined (cannot refer to unexported field or method password)
}
```

[Run Code](example/main.go)

[Next Section](03-workspaces.md)

[Previous Section](01-definition.md)

[Chapter Overview](README.md)
