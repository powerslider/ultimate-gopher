# Iota

- [Overview](#overview)
- [Constant Order](#constant-order)
- [Bit Shifting](#bit-shifting)
- [Powers](#powers)
- [Advice Using Iota](#advice-using-iota)

## Overview

- The `iota` keyword represents successive integer constants 0, 1, 2, etc., to simplify definitions of incrementing
  numbers.
- Because it can be used in expressions, it provides a generality beyond that of simple enumerations.
- It resets to 0 whenever the keyword `const` appears in the source code and increments after each `const`
  specification.

```go
package main

import "fmt"

const (
	Apple int = iota
	Orange
	Banana
)

func main() {
	fmt.Printf("The value of Apple is %v\n", Apple)
	fmt.Printf("The value of Orange is %v\n", Orange)
	fmt.Printf("The value of Banana is %v\n", Banana)
}
```

[Run Code](https://play.golang.org/p/n-4xBc9RNRj)

Output:

```
The value of Apple is 0 
The value of Orange is 1 
The value of Banana is 2
```

- You can use constant shorthand (leaving out everything after the constant name) to declare consecutive numbered
  constants.

## Constant Order

- Changing the order of the defined constants in an `iota` block will change their values. Therefore, it is very
  important that you understand how this will affect your code before you change the order.

```go
package main

import "fmt"

const (
	Apple int = iota
	Banana
	Orange
)

func main() {
	fmt.Printf("The value of Apple is %v\n", Apple)
	fmt.Printf("The value of Banana is %v\n", Banana)
	fmt.Printf("The value of Orange is %v\n", Orange)
}
```

[Run Code](https://play.golang.org/p/ZE0X6bdO3MY)

Output:

```
The value of Apple is 0
The value of Banana is 1
The value of Orange is 2
```

## Bit Shifting

- `iota` combined with the bitshift operator `<<` can be used to create bitmasks as well:

```go
package main

import "fmt"

const (
	read   = 1 << iota // 00000001 = 1
	write              // 00000010 = 2
	remove             // 00000100 = 4

	// admin will have all of the permissions
	admin = read | write | remove
)

func main() {
	fmt.Printf("read =  %v\n", read)
	fmt.Printf("write =  %v\n", write)
	fmt.Printf("remove =  %v\n", remove)
	fmt.Printf("admin =  %v\n", admin)
}
```

[Run Code](https://play.golang.org/p/Wohy0W1FMDj)

Output:

```
read =  1
write =  2
remove =  4
admin =  7
```

## Powers

- `iota` can also help in defining patterns with `powers` calculations:

```go
package main

import "fmt"

const (
	_  = 1 << (iota * 10) // ignore the first value
	KB                    // decimal:       1024 -> binary 00000000000000000000010000000000
	MB                    // decimal:    1048576 -> binary 00000000000100000000000000000000
	GB                    // decimal: 1073741824 -> binary 01000000000000000000000000000000
)

func main() {
	fmt.Printf("KB =  %v\n", KB)
	fmt.Printf("MB =  %v\n", MB)
	fmt.Printf("GB =  %v\n", GB)
}
```

[Run Code](https://play.golang.org/p/ObLEE1SOVbp)

Output:

```
KB =  1024
MB =  1048576
GB =  1073741824
```

## Advice Using `iota`

- `iota` does provide some powerful features (as we saw in the previous examples) in a very easy to achieve manner.
- In general given that the order of the constants matters, does make your code quite brittle.
- It is considered best practice to avoid using `iota` when possible.
- Even with something as complex as the shifting multiplier example we saw, we could re-write it as just constants and
  ensure that the code is not brittle:

```go
package main

import "fmt"

const (
	KB = 1024       // binary 00000000000000000000010000000000
	MB = 1048576    // binary 00000000000100000000000000000000
	GB = 1073741824 // binary 01000000000000000000000000000000
)

func main() {
	fmt.Printf("KB =  %v\n", KB)
	fmt.Printf("MB =  %v\n", MB)
	fmt.Printf("GB =  %v\n", GB)
}
```

[Run Code](https://play.golang.org/p/p0Qsfajtwxh)

Output:

```
KB =  1024
MB =  1048576
GB =  1073741824
```

[Next Section]()

[Previous Section](10-constants.md)



