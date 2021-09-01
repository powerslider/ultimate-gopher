# Strings

- [Overview](#overview)
- [Raw String Literals](#raw-string-literals)
- [Interpreted String Literals](#interpreted-string-literals)

## Overview

A string is a sequence of one or more characters (letters, numbers, symbols) that can be either a constant or a
variable.

- Strings exist within either back quotes ` or double quotes " in Go and have different characteristics depending on
  which quotes you use.
- If you use the back quotes, you are creating a _raw string literal_. If you use the double quotes, you are creating
  an _interpreted string literal_.

## Raw String Literals

- Raw string literals are character sequences between back quotes, often called back ticks. Within the quotes, any
  character may appear except back quote.

```go
s := `Say "hello" to Go!`
```

- Backslashes have no special meaning inside of raw string literals.
- Raw string literals may also be used to create multiline strings:

```go
s := `Lorem Ipsum is simply dummy text of the printing and
      typesetting industry. Lorem Ipsum has been the industry's 
      standard dummy text ever since the 1500s, when an unknown 
      printer took a galley of type and scrambled it to make a 
      type specimen book. It has survived not only five centuries, 
      but also the leap into electronic typesetting, remaining essentially 
      unchanged. It was popularised in the 1960s with the release of Letraset 
      sheets containing Lorem Ipsum passages, and more recently with desktop 
      publishing software like Aldus PageMaker including versions of Lorem Ipsum.`
```

## Interpreted String Literals

- Interpreted string literals are character sequences between double quotes, as in `"bar"`. Within the quotes, any
  character may appear except newline and unescaped double quote.

```go
s := "Say \"hello\" to Go!"
```

- You will almost always use interpreted string literals because they allow for escape characters within them.

[Next Section](07-printing.md)

[Previous Section](05-booleans.md)

[Chapter Overview](README.md)