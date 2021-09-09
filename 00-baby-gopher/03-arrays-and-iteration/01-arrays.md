# Arrays

- [Overview](#overview)
- [Defining Arrays](#defining-arrays)
- [Initializing Arrays](#initializing-arrays)
- [Indexing Arrays](#indexing-arrays)
- [Array Type](#array-type)
- [Setting Array Values](#setting-array-values)
- [Two Dimensional Arrays](#two-dimensional-arrays)

## Overview

- Arrays in Go are useful when planning for a detailed layout of memory.
- Using arrays can sometimes help avoid allocation. However, their primary use is for the building blocks of _slices_.

## Defining Arrays

- Arrays are _fixed length_, _fixed type_ and _zero based_:

```go
fruits := [4]string{}
fruits[0] = "banana"
fruits[1] = "strawberry"
fruits[2] = "pineapple"
fruits[3] = "orange"
```

---
__NOTE__
> The _capacity_ of an array is defined at creation time. Once an array has allocated it's size, the size can no longer
> be changed.
---

## Initializing Arrays

- Arrays can have their values set at initialization time.

```go
package main

import "fmt"

func main() {
	fruits := [4]string{"banana", "strawberry", "pineapple", "orange"}

	fmt.Println(fruits)
}
```

## Array Zero Value

- The zero value of each element in an array is the zero value for the type of elements in the array:

```go
names := [4]string{}
digits := [4]int{}

fmt.Printf("%q\n", names)
fmt.Println(digits)
```

Output:

```
["" "" "" ""]
[0 0 0 0]
```

- For integers that is `0`, for strings `""`.

## Indexing Arrays

- You will receive an error (either compile time or a panic) when trying to access an index of the array beyond its
  size.

```go
fruits := [4]string{}
fruits[4] = "lemon"
```

```
invalid array index 4 (out of bounds for 4-element array)
```

## Array Type

- An array can only be of the type it is declared. The following will result in compiler errors:

```go
names := [4]string{"John", "Ryan", "Ana", "Maria"}
digits := [4]int{1, 2, 3, 4}

names[0] = 5 // cannot put an int in a string array
digits[0] = "five" // cannot put a string in an int array
```

[Run Code](https://play.golang.org/p/Ee0_j3d8FTH)

Output:

```
./prog.go:8:11: cannot use 5 (type untyped int) as type string in assignment
./prog.go:9:12: cannot use "five" (type untyped string) as type int in assignment
```

- The length is actually part of the type that is defined for arrays:

```go
package main

import "fmt"

func main() {
	a1 := [2]string{"one", "two"}
	a2 := [2]string{}

	a2 = a1

	fmt.Println(a2)
	a3 := [3]string{}

	// This cannot be done, as it is not of the same type
	a3 = a2

	fmt.Println(a3)
}
```

[Run Code](https://play.golang.org/p/RhNHN7yd8WW)

Output:

```
./prog.go:15:5: cannot use a2 (type [2]string) as type [3]string in assignment
```

## Setting Array Values

---
__IMPORTANT__

When creating two arrays, and then setting the value of one array to the other, they still continue to have their own
memory space.
---

```go
package main

import "fmt"

func main() {
	a1 := [2]string{"one", "two"}
	a2 := [2]string{}

	a2 = a1

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)

	a1[0] = "bob"

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
}
```

[Run Code](https://play.golang.org/p/crjtMBBsZ3O)

Output:

```
a1: [one two]
a2: [one two]
a1: [bob two]
a2: [one two]
```

## Two Dimensional Arrays

- Go's arrays are one-dimensional. To create an equivalent of a 2D array, it is necessary to define an array-of-arrays:

```go
package main

import "fmt"

// Matrix - a 3x3 array, really an array of arrays.
type Matrix [3][3]int

func main() {
	m := Matrix{
		{0, 0, 0},
		{1, 1, 1},
		{2, 2, 2},
	}

	fmt.Println(m)
}
```

[Run Code](https://play.golang.org/p/6PycSKrie8d)

Output:

```
[[0 0 0] [1 1 1] [2 2 2]]
```

[Next Section](02-iteration.md)

[Chapter Overview](README.md)