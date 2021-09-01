# Comments

- [Overview](#overview)
- [Block Comments](#block-comments)
- [Inline Comments](#inline-comments)
- [Commenting Out Code For Testing](#commenting-out-code-for-testing)
- [GoDoc](#godoc)

## Overview

- Comments in Go begin with a set of forward slashes `//` and continue to the end of the line.
- It is idiomatic to have a white space after the set of forward slashes.

```go
// This a an exceptional comment.
```

---
__NOTE__
> When commenting code, you should be looking to answer the _why_ behind the code as opposed to the _what_ or _how_.
> Unless the code is particularly tricky, looking at it can generally tell what the code is doing or how it is doing it.
---

## Block Comments

You can create block comments two ways in Go:

- The first is by using a set of double forward slashes and repeating them for every line.

```go
// First line of a block comment
// Second line of a block comment
```

- The second is to use opening tags `/*` and closing tags `*/`. For documenting code, it is considered idiomatic to
  always use `//` syntax.
- You would only use the `/* ... */` syntax for debugging, which we will cover later.

```go
/*
Everything here
will be considered
a block comment
*/
```

- Here is an example of a block comment that defines what is happening in the `MustGet()` function defined below:

```go
// MustGet will retrieve a url and return the body of the page.
// If Get encounters any errors, it will panic.
func MustGet(url string) string {
...
}
```

___

__Useful Recommendations__

- It is common to see block comments at the beginning of exported functions in Go (these comments are also what generate
  your code documentation).
- Block comments are also being used when operations are less straightforward and are therefore demanding of a thorough
  explanation.
- Except for documenting functions, you should try to avoid over-commenting the code and should tend to trust other
  programmers to understand Go unless you are writing for a particular audience.

---

## Inline Comments

- Inline comments occur on the same line of a statement, following the code itself.
- Like other comments, they begin with a set of forward slashes.
- Not required, but considered idiomatic is to have a whitespace after the forward slashes.
- The general format of an inline comment looks like this:

```
[code]  // Inline comment about the code.
```

- Inline comments should be used sparingly, but can be effective in some scenarios:
    - Explaining tricky parts of the code.
    - Not remembering a line of the code you are writing in the future.
    - Collaborating with someone who may not be familiar with all aspects of the code.
- A good example would be the following (if your collaborators may not know that the following statement creates a
  complex number):

```go
z := x % 2 // Get the modulus of x.
```

- Another good use case in explaining the reason behind doing something, or some extra information, etc.:

```go
x := 8 // Initialize x with an arbitrary number
```

---
__NOTE__
> Comments that are made inline should be used only when necessary and when they can provide helpful guidance for the
> person reading the program.
---

## Commenting Out Code For Testing

- In addition to using comments as a way to document code, you can also use for the following reasons:
    - Comment out code that you do not want to execute while you are testing or debugging a program you are currently
      creating.
    - Troubleshooting the precise issue when experiencing errors after implementing new lines of code.
    - Trying alternatives while you are determining how to set up your code.
    - Disable failing code while continuing to work on other parts of the code.

- Using block comments:

```go
func main() {
/*
	In this example, we're commenting out the addTwoNumbers
	function because it is failing, therefore preventing it from executing.
	Only the multiplyTwoNumbers function will run

	a := addTwoNumbers(3, 5)
	fmt.Println(a)

*/

    m := multiplyTwoNumbers(5, 9)
    fmt.Println(m)
}
```

- Using basic comments (this option is usually only viable when using an IDE or an editor that prepends `//` on each
  line of the code automatically by marking the code snippet and using a keyboard shortcut):

```go
func main() {
//In this example, we're commenting out the addTwoNumbers
//function because it is failing, therefore preventing it from executing.
//Only the multiplyTwoNumbers function will run
//
//a := addTwoNumbers(3, 5)
//fmt.Println(a)

    m := multiplyTwoNumbers(5, 9)
    fmt.Println(m)
}
```

## GoDoc

- Documenting your code is important.
- Go is created with this statement as part of the design consideration.
- `GoDoc` is a tool that generates code documentation from your comments.
- For an in depth look at formatting your comments to create good project documentation, read the following article:
  [Godoc: documenting Go code](https://blog.golang.org/godoc).

[Next Section](10-constants.md)

[Previous Section](08-utf-8.md)

[Chapter Overview](README.md)