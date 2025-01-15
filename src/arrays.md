# Arrays

In Go, an _array_ is a numbered sequence of elements of a
specific length. In typical Go code, [slices](slices.md) are
much more common; arrays are useful in some special
scenarios.

Here we create an array `a` that will hold exactly
5 `int`s. The type of elements and length are both
part of the array's type. By default an array is
zero-valued, which for `int`s means `0`s.

```go,editable
var a [5]int
fmt.Println("emp:", a)
```

We can set a value at an index using the
`array[index] = value` syntax, and get a value with
`array[index]`.

```go,editable
var a [5]int
a[4] = 100
fmt.Println("set:", a)
fmt.Println("get:", a[4])
```

The builtin `len` returns the length of an array.

```go,editable
var a [5]int
fmt.Println("len:", len(a))
```

Use this syntax to declare and initialize an array
in one line.

```go,editable
b := [5]int{1, 2, 3, 4, 5}
fmt.Println("dcl:", b)
```

You can also have the compiler count the number of
elements for you with `...`

```go,editable
b := [...]int{1, 2, 3, 4, 5}
fmt.Println("dcl:", b)
```

If you specify the index with `:`, the elements in
between will be zeroed.

```go,editable
b = [...]int{100, 3: 400, 500}
fmt.Println("idx:", b)
```

Array types are one-dimensional, but you can
compose types to build multi-dimensional data
structures.

```go,editable
var twoD [2][3]int
for i := 0; i < 2; i++ {
	for j := 0; j < 3; j++ {
		twoD[i][j] = i + j
	}
}
fmt.Println("2d: ", twoD)
```

You can create and initialize multi-dimensional
arrays at once too.

```go,editable
twoD = [2][3]int{
	{1, 2, 3},
	{1, 2, 3},
}
fmt.Println("2d: ", twoD)
```
