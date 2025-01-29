---
weight: 3
---

# Variables

In Go, _variables_ are explicitly declared and used by
the compiler to e.g. check type-correctness of function
calls.

`var` declares 1 or more variables.

```go
var a = "initial"
fmt.Println(a)
// Output: initial
```

You can declare multiple variables at once.

```go
var b, c int = 1, 2
fmt.Println(b, c)
// Output: 1 2
```

Go will infer the type of initialized variables.

```go
var d = true
fmt.Println(d)
// Output: true
```

Variables declared without a corresponding
initialization are _zero-valued_. For example, the
zero value for an `int` is `0`.

```go
var e int
fmt.Println(e)
// Output: 0
```

The `:=` syntax is shorthand for declaring and
initializing a variable, e.g. for
`var f string = "apple"` in this case.
This syntax is only available inside functions.

```go
f := "apple"
fmt.Println(f)
// Output: apple
```
