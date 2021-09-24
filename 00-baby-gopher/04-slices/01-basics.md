# Basics

- [Overview](#overview)
- [Iteration](#iteration)
- [Slice Internals](#slice-internals)
- [Appending to Slices](#appending-to-slices)
- [Growing a Slice](#growing-a-slice)
- [Making a Slice](#making-a-slice)
- [Two-Dimensional Slices](#two-dimensional-slices)

## Overview

- Slices are kind of similar to arrays, fixed type, dynamically sized and very flexible:

```go
package main

import "fmt"

func main() {
	// create an array of names
	namesArray := [4]string{"Tim", "James", "George", "Lewis"}
	// create a slice of names
	namesSlice := []string{"Tim", "James", "George", "Lewis"}

	fmt.Println(namesArray)
	fmt.Println(namesSlice)
}
```

[Run Code](https://play.golang.org/p/lfAiXotGCEK)

Output:

```
[Tim James George Lewis]
[Tim James George Lewis]
```

---
__NOTE__
> Unless you know that your list will contain a finite and fixed set of elements, you'll almost always use a slice when dealing with data.
---

- Other useful features of slices include:
    - Mutation without an allocation.
    - Ability to operate on subsections of the slice easily.
    - Dynamically grow the size of a slice.

- How can you tell the difference between a slice and an array?

```go
var array [4]string // array
var slice []string // slice
```

- Slices do not have a length in the declaration.

## Iteration

- Slices can be iterated over in the same ways that arrays can be:

```go
package main

import "fmt"

func main() {
	namesSlice := []string{"Tim", "James", "George", "Lewis"}

	for i, n := range namesSlice {
		fmt.Printf("%d - %s\n", i, n)
	}
}
```

[Run Code](https://play.golang.org/p/zNVEVP6A5nP)

Output:

```
0 - Tim
1 - James
2 - George
3 - Lewis
```

## Slice Internals

- A useful mental model of representing a slice is the following:

```go
type slice struct {
Length   int
Capacity int
Array    [10]array
}
```

---
__NOTE__
> This is not the actual definition, but a way you can think about a slice header.
---

## Appending to Slices

- The `append` built-in function allows us to add elements to a slice.

```go
package main

import "fmt"

func main() {
	names := []string{}
	names = append(names, "Tim")

	// Append multiple items at once
	names = append(names, "Maria", "Peter")

	// Append an entire slice to another slice
	moreNames := []string{"William", "Barney", "Jack"}
	names = append(names, moreNames...)

	fmt.Println(names)
}
```

[Run Code](https://play.golang.org/p/JSiSzpujAfo)

Output:

```
[Tim Maria Peter William Barney Jack]
```

---
__IMPORTANT__
> Should the slice not have enough space, Go will automatically reallocate the slice to have more capacity.
---

## Growing a Slice

- The `len` built-in function outputs how many elements the slice actually has.
- The `cap` built-in function outputs the capacity of the slice, or how many elements it can have.

```go
package main

import "fmt"

func main() {
	names := []string{}
	fmt.Println("len:", len(names)) // 0
	fmt.Println("cap:", cap(names)) // 0

	names = append(names, "Tim")
	fmt.Println("len:", len(names)) // 1
	fmt.Println("cap:", cap(names)) // 1

	names = append(names, "Barney")
	fmt.Println("len:", len(names)) // 2
	fmt.Println("cap:", cap(names)) // 2

	names = append(names, "Maria")
	fmt.Println("len:", len(names)) // 3
	fmt.Println("cap:", cap(names)) // 4

	names = append(names, "Jack")
	fmt.Println("len:", len(names)) // 4
	fmt.Println("cap:", cap(names)) // 4

	names = append(names, "Steven")
	fmt.Println("len:", len(names)) // 5
	fmt.Println("cap:", cap(names)) // 8
}
```

[Run code](https://play.golang.org/p/n6WSTRh76kP)

Output:

```
len: 0
cap: 0
len: 1
cap: 1
len: 2
cap: 2
len: 3
cap: 4
len: 4
cap: 4
len: 5
cap: 8
```

- In the next example we are going to visualize exactly how a slice resizes itself when in has reached its maximum
  capacity.
- Given a slice that contains 4 elements (`A`, `B`, `C`, `D`) that has reached its max capacity:

```
┌────────────────────────────────────────────────┐
│                  Slice Header                  │
├────────────────────────────────────────────────┤
│ Len -> 4                                       │
│                                                │
│ Cap -> 4                                       │
│                          ┌─────────────────┐   │
│ Arr -> Array Pointer ->  │Underlying Array │   │
│                          │                 │   │
│                          │┌─┬─┬─┬─┐        │   │
│                          ││A│B│C│D│        │   │
│                          │└─┴─┴─┴─┘        │   │
│                          └─────────────────┘   │
└────────────────────────────────────────────────┘
```

- When we append the values `E`, `F`, and `G`, it will force the slice to expand, as it currently has no capacity for
  the new values.
- It will create a new underlying array, copy the original values into the new one, and add the new values as well.

```
┌────────────────────────────────────────────────┐
│          append(slice, "E", "F", "G")          │
├────────────────────────────────────────────────┤
│                                                │
│ ┌─────────────────┐        ┌─────────────────┐ │
│ │Original Array   │        │New Array        │ │
│ │                 │        │                 │ │
│ │┌─┬─┬─┬─┐        │───────>│┌─┬─┬─┬─┬─┬─┬─┬─┐│ │
│ ││A│B│C│D│        │        ││A│B│C│D│E│F│G│ ││ │
│ │└─┴─┴─┴─┘        │        │└─┴─┴─┴─┴─┴─┴─┴─┘│ │
│ └─────────────────┘        └─────────────────┘ │
│                                                │
└────────────────────────────────────────────────┘
```

- This is what the new slice will look like:

```
┌────────────────────────────────────────────────┐
│                  Slice Header                  │
├────────────────────────────────────────────────┤
│ Len -> 7                                       │
│                                                │
│ Cap -> 8                                       │
│                          ┌─────────────────┐   │
│ Arr -> Array Pointer ->  │Underlying Array │   │
│                          │                 │   │
│                          │┌─┬─┬─┬─┬─┬─┬─┬─┐│   │
│                          ││A│B│C│D│E│F│G│ ││   │
│                          │└─┴─┴─┴─┴─┴─┴─┴─┘│   │
│                          └─────────────────┘   │
└────────────────────────────────────────────────┘
```

---
__IMPORTANT__
> If the original underlying array is not referenced by any other part of the program, it will be marked for garbage
> collection.
---

## Making a Slice

- Slices can also be created using the `make` built-in function:

```go
package main

import "fmt"

func main() {
	a := []string{}
	b := make([]string, 0)
	var c []string

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
```

[Run Code](https://play.golang.org/p/MdFjvpkCM9m)

Output:

```
[]
[]
[]
```

---
__NOTE__
> All the aforementioned example slices are functionally equivalent.
---

- The `make` built-in function allows us to define the starting _length_ of the slice, and optionally, the starting
  _capacity_ of the slice:

```go
package main

import "fmt"

func main() {
	a := make([]int, 1, 3)

	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
```

[Run Code](https://play.golang.org/p/yPYcSbcfTsQ)

Output:

```
[0]
1
3
```

- Be careful when using both `make` and `append`, as you may inadvertently create zero values in your slice:

```go
package main

import "fmt"

func main() {
	a := make([]string, 2)
	a = append(a, "foo", "bar")
	fmt.Printf("%q", a)
}
```

[Run Code](https://play.golang.org/p/lgrVcmgnFkG)

Output:

```
["" "" "foo" "bar"]
```

## Two-Dimensional Slices

- Go by default has only one-dimensional slices.
- To create a 2D version of a slice, it is necessary to create a slice of slices.
- Because slices have variable length, it is possible to have each inner slice be a different length.

```go
package main

import "fmt"

type Modules [][]string

func main() {
	modules := Modules{
		[]string{"Module 1"},
		[]string{"Module 2", "Module 3"},
		[]string{"Module 4", "Module 5", "Module 6"},
	}

	fmt.Println(modules)
}
```

[Run Code](https://play.golang.org/p/h7rEsSNvJM1)

Output:

```
[[Module 1] [Module 2 Module 3] [Module 4 Module 5 Module 6]]
```

[Next Section](02-slice-operations.md)

[Chapter Overview](README.md)
