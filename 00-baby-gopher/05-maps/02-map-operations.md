# Map Operations

- [Overview](#overview)
- [Iterating Maps](#iterating-maps)
- [Getting Map Keys](#getting-map-keys)
- [Deleting Map Keys](#deleting-map-keys)
- [Checking Map Values](#checking-map-values)
- [Sorting Maps](#sorting-maps)
    - [Sorting by Key](#sorting-by-key)
    - [Sorting by Value](#sorting-by-value)

## Overview

- In this section we will learn how to perform operations on maps such as: iteration, deletion of map keys, checking for
  map values, sorting maps by key and by value, etc.

## Iterating Maps

- Maps can be iterated over in the same ways as arrays and slices with the subtle difference that the `range` keyword
  returns the `key` and the `value` for each entry in the map:

```go
package main

import "fmt"

func main() {
	fruitColorMap := map[string]string{
		"banana":     "yellow",
		"strawberry": "red",
		"kiwi":       "green",
		"plum":       "purple",
	}

	for key, value := range fruitColorMap {
		fmt.Printf("A %s is %s.\n", key, value)
	}
}
```

[Run Code](https://play.golang.org/p/4f3Z_2GeD1y)

Output:

```
A banana is yellow.
A strawberry is red.
A kiwi is green.
A plum is purple.
```

---
__IMPORTANT__
> When iterating over a `map` with a `range` loop, the iteration order is NOT specified and is not guaranteed to be
> the same from one iteration to the next.
---

## Getting Map Keys

- Go does not provide a way to get a list of keys or values from a map.
- We must build that list ourselves:

```go
package main

import "fmt"

func main() {
	fruitColorMap := map[string]string{
		"banana":     "yellow",
		"strawberry": "red",
		"kiwi":       "green",
		"plum":       "purple",
	}

	keys := make([]string, 0, len(fruitColorMap))

	for k := range fruitColorMap {
		keys = append(keys, k)
	}
	fmt.Printf("%+v", keys)
}
```

Output:

```
[banana strawberry kiwi plum]
```

## Deleting Map Keys

- The `delete` built-in function can be used to remove a key and its value from a map:

```go
package main

import "fmt"

func main() {
	fruitColorMap := map[string]string{
		"banana":     "yellow",
		"strawberry": "red",
		"kiwi":       "green",
		"plum":       "purple",
	}

	delete(fruitColorMap, "banana")

	fmt.Println(fruitColorMap)
}
```

[Run Code](https://play.golang.org/p/o26mLljvXOM)

Output:

```
map[kiwi:green plum:purple strawberry:red]
```

## Checking Map Values

- Maps in Go return an optional second argument that will tell you if the key exists in the map:

```go
package main

import "fmt"

func main() {
	fruitColorMap := map[string]string{
		"banana":     "yellow",
		"strawberry": "red",
		"kiwi":       "green",
		"plum":       "purple",
	}

	key := "banana"
	value, ok := fruitColorMap[key]
	if ok {
		fmt.Printf("Found key %q: %q", key, value)
	} else {
		fmt.Printf("Key not found: %q", key)
	}
}
```

[Run Code](https://play.golang.org/p/0_jwui01d58)

Output:

```
Found key "banana": "yellow"
```

- We can simplify a bit the syntax by inlining the `ok` variable:

```go
package main

import "fmt"

func main() {
	fruitColorMap := map[string]string{
		"banana":     "yellow",
		"strawberry": "red",
		"kiwi":       "green",
		"plum":       "purple",
	}

	key := "banana"
	if value, ok := fruitColorMap[key]; ok {
		fmt.Printf("Found key %q: %q", key, value)
	} else {
		fmt.Printf("Key not found: %q", key)
	}
}
```

- In the following example, we show that even if we don't get an error, you still have a bug because we didn't check for
  the existence of the value:

```go
package main

import "fmt"

type person struct {
	ID   int
	Name string
}

func main() {
	data := map[int]person{}
	s1 := person{ID: 1, Name: "Tim"}
	data[1] = s1

	value := data[10]

	// Because of the way zero values in Go work, we still get a zero 
	// value representation of the struct which is certainly a bug.
	fmt.Printf("%+v", value)
}
```

[Run Code](https://play.golang.org/p/aHa6qBINVmp)

Output:

```
{ID:0 Name:}
```

Desired Output:

```
{ID:1 Name:Tim}
```

- The way to mitigate this bug in to check for existence of the key in your code to avoid a zero value being returned
  and creating a bug:

```go
package main

import "fmt"

type person struct {
	ID   int
	Name string
}

func main() {
	data := map[int]person{}
	s1 := person{ID: 1, Name: "Tim"}
	data[1] = s1

	if value, ok := data[10]; ok {
		fmt.Printf("%+v", value)
	}
}
```

[Run Code](https://play.golang.org/p/sFJNyX_0YwL)

- This will return us an empty output, because we checked that for key `10`, there is no available value, hence avoiding
  returning a zero value entry.

## Sorting Maps

- Maps are not sorted:

```go
package main

import "fmt"

func main() {
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	for k, v := range months {
		fmt.Println(k, v)
	}
}
```

[Run Code](https://play.golang.org/p/zvA_zBGbfrN)

Output:

```
4 April
5 May
9 September
8 August
10 October
11 November
1 January
2 February
3 March
6 June
7 July
12 December
```

---
__NOTE__

> This output will be different on each run and of course never sorted!
---

### Sorting by Key

- Sorting by key is achieved by gathering all the keys and using the `sort` package to sort them.
- By sorting a slice of all the keys, we can iterate them and access each map value in the sorted order:

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	keys := make([]int, 0, len(months))

	for k := range months {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Printf("keys: %+v\n", keys)
	for _, k := range keys {
		fmt.Println(k, months[k])
	}
}
```

[Run Code](https://play.golang.org/p/Y3QiclS2OaC)

Output:

```
keys: [1 2 3 4 5 6 7 8 9 10 11 12]
1 January
2 February
3 March
4 April
5 May
6 June
7 July
8 August
9 September
10 October
11 November
12 December
```

### Sorting by Value

- Sorting by value has more steps involved:
    - Invert the map so that you have the values as keys.
    - Perform sorting by key.
- Using our example with all months, we aim to sort them in an alphabetic order:

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	months := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	// create an inverted map where values are now the keys
	sorted := make(map[string]int)

	keys := make([]string, 0, len(months))

	for k, v := range months {
		keys = append(keys, v)
		sorted[v] = k
	}

	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println(sorted[k], k)
	}
}
```

Output:

```
4 April
8 August
12 December
2 February
1 January
7 July
6 June
3 March
5 May
11 November
10 October
9 September
```

[Next Section](03-map-tips.md)

[Previous Section](01-map-basics.md)

[Chapter Overview](README.md)