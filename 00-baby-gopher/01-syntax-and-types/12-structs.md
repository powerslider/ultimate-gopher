# Structs

- [Overview](#overview)
- [Defining Structs](#defining-structs)
- [Initializing Structs](#initializing-structs)
- [Struct Tags](#struct-tags)
    - [Encoding](#encoding)
- [Printing Structs](#printing-structs)

## Overview

- A `struct` is a collection of fields, often called members (or attributes).
- Structs are used to create custom complex types in Go.
- When trying to understand structs, it helps to think of them as a blueprint for what the new type will do.
- A `struct` definition, does NOT contain any data.

## Defining Structs

- Structs may have 0 or more fields:

```go
type User struct {
    Name  string
    Email string
}
```

## Initializing Structs

- Without initial values:

```go
u := User{}
```

- With initial values using the `field: value` syntax:

```go
u := User{
    Name:  "Homer Simpson",
    Email: "homer@example.com",
}
```

- With initial values using the `field: value` syntax in one line:

```go
u := User{Name:  "Homer Simpson", Email: "homer@example.com"}
```

- With initial values without mentioning field names:

```go
u := User{"Homer Simpson", "homer@example.com"}
```

---
__NOTE__
> Initializing without field names is considered bad practice as it can lead to future undesired refactoring. If we
> decide to change the `struct` definition:
> ```go
> type User struct {
>	    ID    int
>	    Name  string
>	    Email string
> }
> ```
> the code would no longer compile, and we would get this error:
> ```
> too few values in User literal
> ```
> As a result, you may have several areas of your code that now need to be changed because you didn't use the
> `field: value` syntax.
---

- Fields can be referenced using a period, and the field name:

```go
fmt.Println(u.Name)
```

- You can set as many (or as few) of the field values on a struct at initialization time as you want.

```go
u := User{Email: "marge@example.com"}
u.Name = "Marge Simpson"
```

## Struct Tags

- Struct tags are small pieces of metadata attached to fields of a struct that provide instructions to other Go code
  that works with the struct.
- A struct tag looks like this, with the tag offset with backtick ` characters:

```go
type User struct {
    Name string `example:"name"`
}
```

- Other Go code is then capable of examining these structs and extracting the values assigned to specific keys it
  requests.
- Struct tags have no effect on the operation of your code without some other code that examines them.

### Encoding

- A common use for struct tags is encoding the data of your struct to some type of other format, e.g. `JSON`, `XML`,
  `Protobuf`, etc.

```go
package main

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	ID       int
	Name     string
	Phone    string
	Password string
}

func main() {
	u := User{
		ID:       1,
		Name:     "Rob Pike",
		Password: "goIsAwesome",
	}

	err := json.NewEncoder(os.Stdout).Encode(&u)
	if err != nil {
		log.Fatal(err)
	}
}
```

[Run Code](https://play.golang.org/p/3W659GPKdgm)

- If we look at the code above, we can see that by using the built-in `encoding/json` package we get less than optimal
  encoding:

```
{"ID": 1,"Name": "Rob Pike","Phone": "","Password": "goIsAwesome"}
```

- While the encoder does emit `JSON`, it is not idiomatic. Here are the following problems:
    - Fields are not cased properly (they start with a capital letter).
    - Empty fields are still encoded.
    - Sensitive information such as the password was also encoded.
- If we give the proper struct tags, we can tell the encoder how we want the `JSON` to be serialized:

```go
package main

import (
	"encoding/json"
	"log"
	"os"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"-"`
}

func main() {
	u := User{
		ID:       1,
		Name:     "Rob Pike",
		Password: "goIsAwesome",
	}

	err := json.NewEncoder(os.Stdout).Encode(&u)
	if err != nil {
		log.Fatal(err)
	}
}
```

[Run Code](https://play.golang.org/p/4XvunCYfb4o)

Output:

```
{"id":1,"name":"Rob Pike"}
```

- The encoder now uses the proper cased names we provided in the struct tags.
- Additionally, we used the special case of `omitempty` to state that we do not want the _Phone_ field encoded if the
  value is empty, as well as the `-` to tell the encoder to skip encoding the _Password_ field.
- Each package that uses struct tags will have directions on how to specify tags for the proper behavior.
- For an in depth look at the options the `encoding/json` package uses for marshalling with struct tags, you can refer
  to its package [documentation](https://pkg.go.dev/encoding/json#Marshal).

## Printing Structs

- Using the `%+v` verb is helpful in showing the field names within a struct:

```go
package main

import "fmt"

type User struct {
	First string
	Last  string
	Email string
}

func main() {
	u := User{First: "Homer", Last: "Simpson", Email: "homer@example.com"}
	fmt.Printf("user: %v\n", u)

	// using %+v shows the field names in the struct when it prints out
	fmt.Printf("user: %+v\n", u)
}
```

[Run Code](https://play.golang.org/p/jM1BJgkPll4)

Output:

```
user: {Homer Simpson homer@example.com}
user: {First:Homer Last:Simpson Email:homer@example.com}
```

[Next Chapter](../02-packages/README.md)

[Previous Section](11-iota.md)

[Chapter Overview](README.md)

