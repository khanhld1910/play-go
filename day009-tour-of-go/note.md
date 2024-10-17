# Note

- some Go's basic types:

  - bool
  - string
  - int int8 int16 int32 int64
  - uint uint8 uint16 uint32 uint64 uintptr
  - byte // alias for uint8
  - rune // alias for int32
  - float32 float64
  - complex64 complex128

- bonus of 2 bitwise operators:

  - `<<` left shift: shift the bits of a number to the left. so m << n is equivalent to multiplying m by 2^n.
  - `>>` right shift: shift the bits of a number to the right. so m >> n is equivalent to dividing m by 2^n.

- variables declared without an explicit initial value are given their zero value:

  - `0` for numeric types,
  - `false` for the boolean type,
  - `""` for strings.

- type conversions:

  - the expression `T(v)` converts the value `v` to the type `T`.
  - some numeric conversions:

    - `int` to `float64`
    - `float64` to `int`
    - `uint` to `int`
    - `int` to `uint`

for example:

```go
var i int = 42
var f float64 = float64(i)
```

- constants:

  - constants are declared like variables, but with the `const` keyword.
  - constants can be character, string, boolean, or numeric values.
  - constants cannot be declared using the `:=` syntax.

For loop

- Go has only one looping construct, the `for` loop.

  - the basic for loop has three components separated by semicolons:

    - the init statement: executed before the first iteration.
    - the condition expression: evaluated before every iteration.
    - the post statement: executed at the end of every iteration.

  - the init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for loop.

```go

package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 10; i++ {
    sum += i
  }

  fmt.Println(sum)
}

// Output: 45
```

<!-- end at tour of GO > For loop -->
