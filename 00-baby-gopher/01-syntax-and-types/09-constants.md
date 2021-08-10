# Constants

- [Overview](#overview)
- [Untyped Constants](#untyped-constants)
- [Typed Constants](#typed-constants)
- [Type Inference](#type-inference)

## Overview

- Constants are like variables, except they cannot be modified once they have been declared, i.e they are _immutable_.
- They can be a character, `string`, `boolean`, or numeric values.
- They cannot be declared using the `:=` syntax.
- They can appear anywhere a `var` statement can.
- They also run faster, because the compiler can make specific optimizations.

```go
const gopher = "Pesho"
fmt.Println(gopher)
```

Output:

```
Pesho
```

- Constant modification after declaration is not permitted, so you will get a compile time error:

```go
const gopher = "Pesho"
gopher = "George"
fmt.Println(gopher)
```

Output:

```
cannot assign to gopher
```

## Untyped Constants

- Constants can be _untyped_.
- This can be useful when working with numbers such as integers.
- If the constant is _untyped_, it is explicitly converted, where _typed_ constants are not.

```go
package main

import "fmt"

const (
	year     = 365        // untyped
	leapYear = int32(366) // typed
)

func main() {
	hours := 24
	minutes := int32(60)
	fmt.Println(hours * year)       // multiplying an int and untyped
	fmt.Println(minutes * year)     // multiplying an int32 and untyped
	fmt.Println(minutes * leapYear) // multiplying both int32 types
}
```

[Run Code](https://play.golang.org/p/EwXnYvrdPY0)

Output:

```
8760
21900
21960
```

## Typed Constants

- If you declare a constant with a type, it will be that exact type.
- `leapYear` was defined as data type `int32`. This means it can only operate with `int32` data types.
- `year` is declared with no type, so it is considered _untyped_. Because of this, you can use it with any integer data
  type.


- The following code outputs a compile time error:

```go
package main

import "fmt"

const (
	leapYear = int32(366) // typed
)

func main() {
	hours := 24
	fmt.Println(hours * leapYear) // multiplying int and int32 types}
}
```

[Run Code](https://play.golang.org/p/N_hB588ETvm)

Output:

```
invalid operation: hours * leapYear (mismatched types int and int32)
```

---
__NOTE__
> If you try to use a _typed_ constant with anything other than it's type, Go will throw a compile time error.
---

## Type Inference

- An _untyped_ `const` or `var` will be converted to the type it is combined for any mathematical operation:

```go
package main

import (
	"fmt"
)

const (
	a = 2
	b = 2
	c = int32(2)
)

func main() {
	fmt.Printf("a = %[1]d (%[1]T)\n", a)
	fmt.Printf("b = %[1]d (%[1]T)\n", b)
	fmt.Printf("c = %[1]d (%[1]T)\n", c)

	fmt.Printf("a*b = %[1]d (%[1]T)\n", a*b)
	fmt.Printf("a*c = %[1]d (%[1]T)\n", a*c)

	d := 4
	e := int32(4)

	fmt.Printf("a*d = %[1]d (%[1]T)\n", a*d)
	fmt.Printf("a*e = %[1]d (%[1]T)\n", a*e)
}
```

[Run Code](https://play.golang.org/p/GYZmktyViex)

Output:

```
a = 2 (int)
b = 2 (int)
c = 2 (int32)
a*b = 4 (int)
a*c = 4 (int32)
a*d = 8 (int)
a*e = 8 (int32)
```

[Next Section](10-iota.md)

[Previous Section](08-utf-8.md)
