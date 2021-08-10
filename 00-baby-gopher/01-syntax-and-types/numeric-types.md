# Numeric Types

- [Overview](#overview)
- [How to pick to correct numeric type?](#how-to-pick-the-correct-numeric-type)
- [Overflow vs. Wraparound](#overflow-vs-wraparound)
- [Saturation](#saturation)

## Overview

Go has two types of numeric types:

- The first type is an _architecture independent type_, i.e. regardless of the architecture you compile for, the type
  will have the correct size (bytes).
- The second type is a _implementation specific type_, i.e. the byte size of that numeric type can vary based on the
  architecture the program is built for.

Go has the following _architecture independent_ numeric types:

```go
uint8       unsigned  8-bit integers (0 to 255)
uint16      unsigned 16-bit integers (0 to 65535)
uint32      unsigned 32-bit integers (0 to 4294967295)
uint64      unsigned 64-bit integers (0 to 18446744073709551615)
int8        signed  8-bit integers (-128 to 127)
int16       signed 16-bit integers (-32768 to 32767)
int32       signed 32-bit integers (-2147483648 to 2147483647)
int64       signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
float32     IEEE-754 32-bit floating-point numbers (+- 1O-45 -> +- 3.4 * 1038 )
float64     IEEE-754 64-bit floating-point numbers (+- 5 * 10-324 -> 1.7 * 10308 )
complex64   complex numbers with float32 real and imaginary parts
complex128  complex numbers with float64 real and imaginary parts
byte        alias for uint8
rune        alias for int32
```

In addition, Go has the following _implementation specific_ types:

```go
uint     either 32 or 64 bits
int      same size as uint
uintptr  an unsigned integer large enough to store the uninterpreted bits of a pointer value
```

---
__NOTE__
> _Implementation specific types_ will have their size defined by the architecture the program is compiled for.
---

## How to pick the correct numeric type?

- Picking the correct type usually has more to do with performance for the target architecture you are programming for
  than the size of the data you are working with.
- However, without needing to know the specific ramifications of performance for your program, you can follow some of
  these basic guidelines when first starting out:
    - > __Option 1__: For integer data, it's common in Go to use the _implementation types_ like `int` or `uint`. This will typically result in the fastest processing speed for your target architecture.
    - > __Option 2__: If you know you will not exceed a specific size range, then picking an _architecture independent_
      type can both increase speed decrease memory usage. To understand integer ranges, we can look at the following examples:
      > ```go
      > int8 (-128 -> 127)
      > int16 (-32768 -> 32767)
      > int32 (− 2,147,483,648 -> 2,147,483,647)
      > int64 (− 9,223,372,036,854,775,808 -> 9,223,372,036,854,775,807)
      > ```
      > For unsigned integers, we have the following ranges:
      > ```go
      > uint8 (with alias byte, 0 -> 255)
      > uint16 (0 -> 65,535)
      > uint32 (0 -> 4,294,967,295)
      > uint64 (0 -> 18,446,744,073,709,551,615)
      > ```
      > And for floats:
      > ```go
      > float32 (+- 1O-45 -> +- 3.4 * 1038 )
      > (IEEE-754) float64 (+- 5 * 10-324 -> 1.7 * 10308 )
      > ```

## Overflow vs. Wraparound

- Go can overflow a number as well as wraparound a number.
- An overflow happens when you try to store a value larger than the data type allows.
- Which of those two options happens depends on if the value can be calculated at _compile time_ or at _runtime_.
    - At compile time, if the compiler can determine a value will be too large to hold in the data type specified, it
      will throw an overflow error. If we take the following example:
    ```go
    var maxUint8 uint8 = 255 // Max Uint8 size
    fmt.Println(maxUint8)
    ```
    - This outputs `255`, but if we add `1` to the value at runtime it will wraparound to `0`:
    ```go
    fmt.Println(maxUint8 + 1)
    ```
    - If we change the program to add `1` to the variable when we assign it, it will not compile:
    ```go
    var maxUint8 uint8 = 255 + 1
    fmt.Println(maxUint8)
    ```
    - Because the compiler can determine the overflow, it throws the following error:
    ```go
    constant 256 overflows uint8
    ```
  [Run code](https://play.golang.org/p/lWanW2rZj5i)

## Saturation

Go does not saturate variables during mathematical operations such as addition or multiplication.
- In languages that saturate, if you had a `uint8` with a max value of `255`, and added `1`, the value would still be 
  the max (saturated) value of `255`.
- In Go, however, it will always wrap around. There is no saturation in Go.

```go
var maxUint8 uint8 = 11
maxUint8 = maxUint8 * 25
fmt.Println("new value:", maxUint8)
```
- Output:
```go
new value: 19
```

[Next Section](booleans.md)

[Previous Section](declaring-variables.md)



