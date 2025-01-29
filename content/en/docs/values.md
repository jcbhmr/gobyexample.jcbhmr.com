---
weight: 2
---

# Values

Go has various value types including strings,
integers, floats, booleans, etc. Here are a few
basic examples.

Strings, which can be added together with `+`.

{{% goplay %}}

```go
fmt.Println("go" + "lang")
// Output: golang
```

{{% /goplay %}}

Integers and floats.

{{% goplay %}}

```go
fmt.Println("1+1 =", 1+1)
fmt.Println("7.0/3.0 =", 7.0/3.0)
// Output:
// 1+1 = 2
// 7.0/3.0 = 2.3333333333333335
```

{{% /goplay %}}

Booleans, with boolean operators as you'd expect.

{{% goplay %}}

```go
fmt.Println(true && false)
fmt.Println(true || false)
fmt.Println(!true)
// Output:
// false
// true
// false
```

{{% /goplay %}}
