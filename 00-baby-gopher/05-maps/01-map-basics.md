# Map Basics

- [Overview](#overview)
- [Defining Maps](#defining-maps)
- [Initializing Maps](#initializing-maps)
- [Length and Capacity](#length-and-capacity)
- [Map Values](#map-values)
- [Map Keys](#map-keys)

## Overview

- Maps are somewhat similar to what other languages call _dictionaries_ or _hashes_.
- A map is an _unordered set_ of values indexed by a _unique key_.
- The computer science term for this data structure is called a _hash table_.
- _Hash tables_ provide fast lookups, inserts and deletes.
- Go's built-in `map` type implements a _hash table_.

## Defining Maps

- A `map` definition has the following form:

```
map[<key_type>]<value_type>
```

- Putting entries into an empty map after declaration has the following form:

```
mapInstance[<key>] = <value>
```

- `<key>` and `<value>` should be of the corresponding types declared in the `map` definition.

## Initializing Maps

- Maps can have values assigned by initializing an empty map and putting entries into it:

```go
package main

import "fmt"

func main() {
	fruitColorMap := map[string]string{}

	fruitColorMap["banana"] = "yellow"
	fruitColorMap["strawberry"] = "red"
	fruitColorMap["kiwi"] = "green"
	fruitColorMap["plum"] = "purple"

	fmt.Println(fruitColorMap)
}
```

[Run Code](https://play.golang.org/p/-WqjH3an5GJ)

Output:

```
map[banana:yellow kiwi:green plum:purple strawberry:red]
```

- Maps can also have their values assigned at creation time, just like arrays:

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

	fmt.Println(fruitColorMap)
}
```

[Run Code](https://play.golang.org/p/whLvpgizKzL)

Output:

```
map[banana:yellow kiwi:green plum:purple strawberry:red]
```

- If you do not initialize a `map`, and try to access its values, you will receive a runtime error:

```go
package main

func main() {
	var fruitColorMap map[string]string
	fruitColorMap["banana"] = "yellow"
}
```

[Run Code](https://play.golang.org/p/N1TNMQ1idQK)

Output:

```
panic: assignment to entry in nil map
```

## Length and Capacity

- The `len` function can be used to find the length (the number of keys) in the map:

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

	fmt.Printf("Length: %d\n", len(fruitColorMap))
}
```

[Run Code](https://play.golang.org/p/DmuyPpeLHVS)

Output:

```
Length: 4
```

- Maps do not have a capacity, since they can grow as needed, so the `cap` function would raise an error:

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

	fmt.Printf("Capacity: %d\n", cap(fruitColorMap))
}
```

[Run Code](https://play.golang.org/p/vnJmW7Rz_PU)

Output:

```
invalid argument fruitColorMap (type map[string]string) for cap
```

## Map Values

- Map values can be set and retrieved using the `[]` syntax:

```go
package main

import "fmt"

func main() {
	fruitColorMap := map[string]string{}

	fruitColorMap["banana"] = "yellow"

	banana := fruitColorMap["banana"]
	fmt.Println(banana)
}
```

[Run Code](https://play.golang.org/p/THC-RYALC7W)

Output:

```
yellow
```

## Map Keys

- Map keys must be comparable.
- Functions, maps or slices cannot be used as key types in our maps.

```go
package main

import "fmt"

func main() {
	m := map[func()]string{}
	fmt.Println(m)
}
```

[Run Code](https://play.golang.org/p/Hmw12IHg2tG)

Output:

```
invalid map key type func()
```

- Structs are ok if they do not contain complex types:

```go
package main

import "fmt"

type simple struct {
	ID int
}

type complex struct {
	f func(id int) simple
}

func main() {
	m := map[simple]string{}

	fmt.Println(m)

	// invalid map key type complex
	//m1 := map[complex]string{}
}
```

[Run Code](https://play.golang.org/p/jJRrHzlgSj2)

Output:

```
map[]
```

- If we uncomment `m1`, again we would get:

```
invalid map key type complex
```

[Next Section](02-map-operations.md)

[Chapter Overview](README.md)



