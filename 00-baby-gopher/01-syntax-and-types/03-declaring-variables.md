# Declaring Variables

- [Overview](#overview)
- [Zero Values](#zero-values)
- [Nil](#nil)
- [Naming Variables](#naming-variables)
- [Naming Style](#naming-style)
- [Multiple Assignment](#multiple-assignment)
- [Typing](#typing)

## Overview

There are several ways to declare a variable and more than one way to declare the exact same variable and value.

- Let's declare a variable called `a` of data type `int` without initialization. This means we will declare a space to
  put a value without assigning it an initial value:

```go
var a int
```

- Now we have a variable declared as `a` of data type `int`.

- Let's initialize the value by using the equal (`=`) operator:

```go
var a int = 1
```

- The above mentioned declaration forms are called __long variable declarations__ in Go.
- You can also use __short variable declaration__:

```go
a := 1
```

- In this case, you have a variable called `a`, and a data type of `int` which Go will infer automatically.

Based on the three ways to declare variables, the Go community has adopted the following idioms:

- __Rule 1__: Use __long form declaration__ when you are not initializing a variable:

```go
var a int
```

- __Rule 2__: Use ___short form declaration__ when declaring and initializing a variable:

```go
a := 1
```

- __Rule 3__: If you want to use __short form declaration__, but bypass the compiler's type inference, you can wrap your
  value in your desired type.

```go
i := int64(1)
```

---
__NOTE__
> The __long form declaration__ is seldom used and not considered idiomatic in Go when you are also initializing the
> value:

```go
var a int = 1
```

---

## Zero Values

All built-in types have a zero value. Any allocated variable is usable even if it never has an assigned value:

```go
package main

import "fmt"

func main() {
	var a int
	var b bool
	var c float64
	var d string

	fmt.Printf("var a %T = %d\n", a, a)
	fmt.Printf("var b %T = %t\n", b, b)
	fmt.Printf("var c %T = %f\n", c, c)
	fmt.Printf("var d %T = %q\n\n", d, d)
}

/* output
var a int = 0
var b bool = false
var c float64 = 0.000000
var d string = ""
*/
```

[Run Code](https://play.golang.org/p/1WqHi2SSntf)

- Because all values have a zero value in Go, you cannot have `undefined` values like in some other languages, e.g.a
  `boolean`could represent `undefined`, `true`, or `false`, thus allowing for three possible values to the variable. In
  Go the `bool` type allows just `true` and `false` as possible values following the zero value concept.

## Nil

Another type in Go is the `nil` type, which serves numerous purposes:

- One of them is serving as zero value for many common types:
    - maps
    - slices
    - functions
    - channels
    - interfaces
    - errors
- The other purposes we will cover in later sections, when we introduce more features and concepts of the language.

## Naming Variables

The naming of variables is quite flexible, but there are some rules you need to keep in mind:

- __Rule 1:__ Variable names must only be one word (as in no spaces).
- __Rule 2:__ Variable names must be made up of only letters, numbers and underscore (_)
- __Rule 3:__ Variable names cannot begin with a number.

Following the rules above, let’s look at both valid and invalid variable names:

| Valid     | Invalid    | Explanation                 |
|-----------|------------|-----------------------------|
| firstName | first-name | Hyphens are not permitted.  |
| person1   | 1person    | Cannot begin with a number. |
| user      | $user      | Symbols are not permitted.  |
| firstName | first name | Must be only one word.      |

- __Rule 4:__ Variable names are _case-sensitive_, meaning that `userName`, `USERNAME`, `UserName`, and `uSERnAME` are
  all completely different variables. You should avoid using similar variable names within a program in order to avoid
  confusion.

- __Rule 5:__ The first letter of a variable has a special meaning in Go.
    - If a variable starts with an uppercase letter, then that variable is accessible _outside the package it was
      declared in (or exported)_.
    - If a variable starts with a lowercase letter, then it is only available _within the package it is declared in_.

```go
var userName string
var Email string
```

- `Email` starts with an uppercase letter and can be accessed by other packages.
- `password` starts with a lowercase letter, and is only accessible inside the package it is declared in.

## Naming Style

- It is common in Go to use very terse (or short) variable names. Given the choice between using `userName` and `user`
  for a variable, it would be idiomatic to choose `user`.

- Scope also plays a role in the terseness of the variable name. The rule is that the smaller the scope the variable
  exists in, the shorter the variable name.

```go
names := []string{"Susan", "George", "Rob", "Anna"}
for i, n := range names {
fmt.Printf("index: %d = %q\n", i, n)
}
```

- The variable `names` is used in a larger scope, so it would be common to give it a more meaningful name to help
  remember what it means in the program.
- However, the variables `i` and `n` are used immediately in the next line of code, and are never used again. Because of
  this, someone reading the code will not be confused about where they are used, or what they mean.

---
__NOTE__

Some notes about style. Go adopts using `MixedCaps` or `mixedCaps` rather than underscores for multiword names.

|Good      | Bad      | Explanation                           |
|----------|----------|---------------------------------------|
|userName  |user_name |Underscores are not conventional.      |
|i           |index      |Prefer `i` over index as it is shorter.|
|serveHTTP |serveHttp |Acronyms should be capitalized.        |
|userID    |UserId      |Acronyms should be capitalized.        |

---

## Multiple Assignment

- Go allows you to assign several values to several variables within the same line.
- Each of these values can be of a different data type:

```go
a, b, c := "gosho", 3.07, 23
fmt.Println(j)
fmt.Println(k)
fmt.Println(l)

/*
output:
gosho
3.07
23
*/
```

- In the example above:
    - The variable `a` was assigned to the string `"gosho"`.
    - The variable `b` was assigned to the float `3.07`.
    - The variable `c` was assigned to the integer `15`.

- This approach to assigning multiple variables to multiple values in one line can keep your lines of code down, but
  make sure you are not compromising readability for fewer lines of code.

## Typing

- Go is a __statically typed__ language.
- __Statically typed__ means that each statement in the program is checked at compile time and that the data type is
  _bound to the variable_.
- In __dynamically linked__ languages, e.g. Javascript, PHP, Python, etc., the data type is _bound to the value_.

The type is declared when declaring a variable:

```go
var pi float64 = 3.14
var count int = 7
```

In languages like PHP, the data type is associated to the value:

```php
$s = "Gophers rock";        // $s is a string
$s = 123;                   // $s is now an integer
```

[Next Section](04-numeric-types.md)

[Previous Section](02-operators-and-delimiters.md)