# Printing

- [Overview](#overview)
- [Print Verbs](#print-verbs)
    - [Using Wrong Verbs](#using-wrong-verbs)
- [Padding](#padding)
    - [Pad Right](#pad-right)
    - [Pad Left](#pad-left)
    - [Pad with Zeroes](#pad-with-zeroes)
    - [Pad with Arbitrary Length](#pad-with-arbitrary-length)
- [More Formatting Directives](#more-formatting-directives)

## Overview

Printing output is a mandatory feature for any programming language.

- Go uses the `fmt` package to provide us with such functionality.
- We will cover the `fmt` package more in depth later, but here is a quick introduction:

```go
fmt.Println("Printing a statement followed by a line return")
fmt.Printf("Hello %s\n", "gopher")
```

Output:

```
Printing a statement followed by a line return
Hello gopher
```

## Print Verbs

- All `Print` statements can use _verbs_ which create different styles of formatting.
- They are usually proceeded by a `%` or `\` character:

```
// Use %v to print the value
// Use %s to print a string
// Use %q to quote a string
// Use %d to print a decimal (int, int32, etc)
// Use %T to print out the data type of the variable
// Use \ to escape a character, like a quote:  \"
// Use \n to print a new line (line return)
// Use \t to insert a tab
// Use %+v to print the name and value
```

- Here are some examples:

```go
a := 1
fmt.Println("This will join all arguments and print them.  a =  ", a)

fmt.Printf("This is the `format` string. Escape \" and print value of a: %v\n", a)

type User struct {
Name string
}
u := User{Name: "Peter"}
fmt.Printf("user: %+v", u)
```

[Run Code](https://play.golang.org/p/GIv8uX2XfkU)

Output:

```
This will join all arguments and print them.  a =   1
This is the `format` string. Escape " and print value of a: 1
user: {Name:Peter}
```

### Using Wrong Verbs

- Occasionally you may accidentally use the wrong verb which will result in an invalid formatting output:

```go
package main

import "fmt"

func main() {
	fmt.Printf("This is an int: %s\n", 42)
	fmt.Printf("This is a string: %d\n", "hello")
}
```

[Run Code](https://play.golang.org/p/B5lHfvqUmz3)

- Desired output would be:

```
This is an int: 42
This is a string: hello
```

- However, the actual output is the following, due to the invalid formatting directives:

```
This is an int: %!s(int=42)
This is a string: %!d(string=hello)
```

---
__NOTE__

> While this will not hurt the performance of your program, it is certainly not desirable. To avoid it, always use
> `go vet` on your code:
> ```
> $ go vet main.go
> # command-line-arguments
> ./main.go:6:2: Printf format %s has arg 42 of wrong type int
> ./main.go:7:2: Printf format %d has arg "hello" of wrong type string
> ```
---

## Padding

- String padding in Go is referred to as the operation of prepending or appending spaces or characters to a string such
  that the total length of the final string is fixed, regardless of the input string’s length.
- You may have encountered a scenario where you had to display or format data in such a way, that it is aligned like in
  a table. Let’s check an example:

```
athletes distances
john          10km
marylin      131km
joe          0.5km
arthur         1km
```

- `athletes` are aligned to the left, `distances` are aligned to the right.
- Fortunately in Go padding can be done using just the standard library without having to write the padding logic by
  yourself or having to import third party libraries.
- The `fmt` package provides padding by utilizing the `width` option:

```
Width is specified by an optional decimal number immediately preceding the verb. If absent

%f default width, default precision
%9f width 9, default precision
%.2f default width, precision 2
%9.2f width 9, precision 2
%9.f width 9, precision 0
```

### Pad Right

```go
fmt.Printf("|%-10s|\n", "john")
```

[Run Code](https://play.golang.org/p/ln1wJ0gkuKA)

Output:

```
|john      |
```

### Pad Left

```go
fmt.Printf("|%10s|\n", "john")
```

[Run Code](https://play.golang.org/p/0-_DkCuw8pm)

Output:

```
|      john|
```

### Pad with Zeroes

```go
fmt.Printf("|%06d|%6d|\n", 12, 345)
```

[Run Code](https://play.golang.org/p/v0HxF7v-Art)

Output:

```
|000012|   345|
```

---
__NOTE__
> Notice the `0` in `%06d`, that will make it a width of 6 and pad it with zeros. The second one will pad with spaces.
---

### Pad with Arbitrary Length

- You can also define the padding _width_ using an _asterisk_ and specifying a parameter representing the _length_ of
  the padding:

```go
fmt.Printf("|%*dkm|\n", 10, 2)
```

[Run Code](https://play.golang.org/p/fOP36RHLnop)

Output:

```
|         2km|
```

## More Formatting Directives

- You can check more formatting directives that can be applied when using the `fmt` package in the official
  documentation page [here](https://pkg.go.dev/fmt#pkg-overview).

```
+	always print a sign for numeric values;
	guarantee ASCII-only output for %q (%+q)
-	pad with spaces on the right rather than the left (left-justify the field)
#	alternate format: add leading 0b for binary (%#b), 0 for octal (%#o),
	0x or 0X for hex (%#x or %#X); suppress 0x for %p (%#p);
	for %q, print a raw (backquoted) string if strconv.CanBackquote
	returns true;
	always print a decimal point for %e, %E, %f, %F, %g and %G;
	do not remove trailing zeros for %g and %G;
	write e.g. U+0078 'x' if the character is printable for %U (%#U).
' '	(space) leave a space for elided sign in numbers (% d);
	put spaces between bytes printing strings or slices in hex (% x, % X)
0	pad with leading zeros rather than spaces;
	for numbers, this moves the padding after the sign
```

[Next Section](08-utf-8.md)

[Previous Section](06-strings.md)

