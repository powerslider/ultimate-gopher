# Iteration

- [Overview](#overview)
- [Iterating over Arrays](#iterating-over-arrays)
- [Continuing a Loop](#continuing-a-loop)
- [Breaking a Loop](#breaking-a-loop)
- [Do-While Loop](#do-while-loop)
- [The `range` Keyword](#the-range-keyword)

## Overview

- In Go there is only one looping construct - the `for` loop.
- It is used to fill in the semantics for `for`, `while`, `do while`, `do until`, etc.

```go
for i := 0; i < N; i++ {
    // do work until i equals N
}
```

## Iterating over Arrays

- Iterating over arrays is done using a `for` loop:

```go
package main

import "fmt"

func main() {
	names := [4]string{"Tim", "James", "George", "Lewis"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
```

[Run Code](https://play.golang.org/p/z5BXjMRUkYZ)

Output:

```
Tim
James
George
Lewis
```

- The `len` function returns the length of the array.

## Continuing a Loop

- The `continue` keyword allows us to go back to the start of the loop and stop executing the rest of the code in the
  `for` block:

```go
for {
    if i == 5 {
        // go to the start of the loop
        continue
    }
    // do work
}
```

- This does not stop the loop from executing, but rather ends that particular run of the loop.

## Breaking a Loop

- To stop execution of a loop we can use the `break` keyword:

```go
for {
    if i == 5 {
        // stop looping
        break
    }
    // do work
}
```

- The `for` loop is now stopped and will no longer run.

## Do-While Loop

- A `do while` loop is used when you want the loop to run at least 1 iteration, regardless of the condition.
- A C/Java-style example would look something like this:

```
do {
	task();
} while (condition);
```

- To create a `do while` style loop in Go a combination of an infinite loop and the `break` keyword can be used:

```go
var i int
for {
    fmt.Println(i)
    i += 2
    if i >= 3 {
        break
    }
}
```

## The `range` keyword

- Looping over arrays, and other collection types, is so common that Go created the `range` keyword to simplify this
  code:

```go
package main

import "fmt"

func main() {
	names := [4]string{"Tim", "James", "George", "Lewis"}

	for i, n := range names {
		fmt.Printf("%d - %s\n", i, n)
	}
}
```

[Run Code](https://play.golang.org/p/Ey3XUUQkPDj)

Output:

```
0 - Tim
1 - James
2 - George
3 - Lewis
```

- `range` returns the _index_ and the _value_ of the array.

[Next Chapter](../04-slices/README.md)

[Previous Section](01-arrays.md)

[Chapter Overview](README.md)

