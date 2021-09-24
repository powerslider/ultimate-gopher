# Slice Operations

- [Overview](#overview)
- [Slice Subsets](#slice-subsets)
    - [Defining Slice Subsets](#defining-slice-subsets)
    - [Mutating Slice Subsets](#mutating-slice-subsets)
- [Copying Slices](#copying-slices)
- [More Slice Operations](#more-slice-operations)

## Overview

- This section will focus on how to utilize slices in more non-trivial ways, e.g. how to build two-dimensional slices,
  how to mutate slice subsets and how to implement more complex slice operations only by using the built-in `append` and
  `copy` functions.

## Slice Subsets

### Defining Slice Subsets

- Subsets of a slice (or a slice of a slice) allow us to work with just section of a slice:

```go
package main

import (
	"fmt"
)

func main() {
	names := []string{"Tim", "Jack", "Maria", "Pamela"}

	fmt.Println(names)

	// Format of slice subsets:
	// slice[starting_index : (starting_index + length)]

	// Get 2 elements starting with the second element (index 1)
	fmt.Println(names[1:3]) // [Paul George] - names[1:1+2]

	// Get all elements starting from the third at index 2 until 
	// the end of the slice. Both statements are functionally equivalent.
	fmt.Println(names[2:len(names)])
	fmt.Println(names[2:])

	// Get all elements starting from the beginning of the slice (index 0) 
	// until the third element (index 2). Both statements are functionally 
	// equivalent.
	fmt.Println(names[0:2])
	fmt.Println(names[:2])
}
```

[Run Code](https://play.golang.org/p/3XPjGimVJIF)

Output:

```
[Tim Jack Maria Pamela]
[Jack Maria]
[Maria Pamela]
[Maria Pamela]
[Tim Jack]
[Tim Jack]
```

- Negative numbers are not allowed as slice indices. The following results in invalid syntax:

```go
names[:-2]
```

```
invalid slice index -2 (index must be non-negative)
```

### Mutating Slice Subsets

- It is important to remember that when grabbing a subset of a slice, you are just getting a "window" into that slice.
- When mutating the subset, the original slice gets mutated as well:

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	names := []string{"Tim", "Jack", "Maria", "Pamela"}

	fmt.Println(names)

	// Get the first three elements of the `names` slice
	subset := names[:3]

	fmt.Println(subset)

	for i, g := range subset {
		subset[i] = strings.ToUpper(g)
	}

	// Print out our original slice of names
	fmt.Println(names)
}
```

[Run Code](https://play.golang.org/p/8V-jS2GBm3z)

Output:

```
[Tim Jack Maria Pamela]
[Tim Jack Maria]
[TIM JACK MARIA Pamela]
```

## Copying Slices

- You can use the `copy` built-in function to make a copy without sharing reference to the original underlying array:

```go
package main

import "fmt"

func main() {
	original := []string{"banana", "melon", "kiwi", "cherry"}
	// Create a new reference to the existing slice
	ref := original

	// Initialize a variable named 'dup' to the
	// same size as the original slice
	dup := make([]string, len(original))

	// Copy the values from original to dup
	// NOTE: if the slices are not the same size, it will
	// only copy what it has space for (length)
	copy(dup, original)

	// Changing either `ref` or `original` will change either
	// of them as they still share the reference to the same
	// backing array
	ref[0] = "orange"
	original[1] = "grapes"

	fmt.Println("Original: ", original)
	fmt.Println("Ref:      ", ref)
	fmt.Println("Dup:      ", dup)
}
```

[Run Code](https://play.golang.org/p/WjlnLeVIbNw)

Output:

```
Original:  [orange grapes kiwi cherry]
Ref:       [orange grapes kiwi cherry]
Dup:       [banana melon kiwi cherry]
```

## More Slice Operations

- By utilizing built-in functions such as `make`, `append`, `copy`, `len` and `cap` and slice subsets using the `:`
  syntax, lots of tricks and operations become possible very easily and take advantage of evolving compiler
  optimizations.
- A good reference is the following wiki in Go's own Github
  repo: [Slice Tricks](https://github.com/golang/go/wiki/SliceTricks).
  
[Next Chapter](../05-maps/README.md)

[Previous Section](01-basics.md)

[Chapter Overview](README.md)
