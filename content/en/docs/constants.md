---
weight: 4
---

# Constants

Go supports _constants_ of character, string, boolean,
and numeric values.

`const` declares a constant value.

```go
package main

import "fmt"

const s string = "constant"

func main() {
	fmt.Println(s)
	// Output: constant
}
```

A `const` statement can appear anywhere a `var`
statement can.

```go
package main

func main() {
	const n = 500000000
	fmt.Println(n)
	// Output: 500000000
}
```

Constant expressions perform arithmetic with
arbitrary precision.

```go
const n = 500000000
const d = 3e20 / n
fmt.Println(d)
// Output: 6e+11
```

A numeric constant has no type until it's given
one, such as by an explicit conversion.

```go
const d = 6e+11
fmt.Println(int64(d))
// Output: 600000000000
```

A number can be given a type by using it in a
context that requires one, such as a variable
assignment or function call. For example, here
`math.Sin` expects a `float64`.

```go
const n = 500000000
fmt.Println(math.Sin(n))
// Output: -0.28470407323754404
```
