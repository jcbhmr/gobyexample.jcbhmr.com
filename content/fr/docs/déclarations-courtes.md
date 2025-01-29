---
weight: 5
---
# Déclarations courtes
```go
package main

import "fmt"

func main() {
    // `x := val` est la version courte pour
    // `var x type = val`.
    x := "Hello var"
    fmt.Println(x)
}
```
```sh
$ go run short-declarations.go
Hello var
```
