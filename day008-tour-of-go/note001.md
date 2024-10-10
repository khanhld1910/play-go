# Note 1

- every Go program is made up of packages
- programs start running in package main
- by convention, the package name is the same as the last element of the import path. For instance, the "math/rand" package comprises files that begin with the statement package rand.
- you can write multiple import statements, but it is good practice to use the factored import statement
- `fmt` is short for format, is a Go built-in package

- Export names:

  - In Go, a name is exported if it begins with a capital letter. For example, `Pizza` is an exported name, as is `Pi`, which is exported from the math package.
  - `pizza` and `pi` do not start with a capital letter, so they are not exported.

- Functions:

  - A function can take zero or more arguments.
  - The type comes after the variable name.

  ```go
  func add(x int, y int) int {
    return x + y
  }
  ```

- functions continue:

  - When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
  - In this example, x and y are both of type int.

  ```go
  func add(x, y int) int {
    return x + y
  }
  ```

  so we shortened from

  ```go
  x int, y int
  ```

  to

  ```go
  x, y int
  ```

- multiple results:

  - A function can return any number of results.
  - The swap function returns two strings.

  ```go
  func swap(x, y string) (string, string) {
    return y, x
  }
  ```

- named return values:

  - Go's return values may be named. If so, they are treated as variables defined at the top of the function.
  - These names should be used to document the meaning of the return values.
  - A return statement without arguments returns the named return values. This is known as a "naked" return.
  - Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

  ```go
  package main

  import "fmt"

  func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return // naked return => so x and y are returned
  }

  func main() {
    fmt.Println(split(17))
  }

  // Output: 7 10
  ```

  name return

  ```go
  package main

  import "fmt"

  func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return x, y
  }

  func main() {
    a, b := split(17)
    fmt.Println(a, b)
  }

  // Output: 7 10
  ```

- variables

- The var statement declares a list of variables; as in function argument lists, the type is last.
- The var statement can be used at package level or function level.

```go

package main

import "fmt"

var c, python, java bool

func main() {
  var i int
  fmt.Println(i, c, python, java)
}

// Output: 0 false false false
```

- variables with initializers

a var declaration can include initializers, one per variable.
if an initializer is present, the type can be omitted; the variable will take the type of the initializer.

```go

package main

import "fmt"

var i, j int = 1, 2

func main() {
  var c, python, java = true, false, "no!"
  fmt.Println(i, j, c, python, java)
}

// Output: 1 2 true false no!
```

- short variable declarations

- Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
- Outside a function, every statement begins with a keyword (var, func, and so on) and so the := construct is not available.

```go

package main

import "fmt"

func main() {
  var i, j int = 1, 2
  k := 3
  c, python, java := true, false, "no!"
  fmt.Println(i, j, k, c, python, java)
}

// Output: 1 2 3 true false no!
```

(10/17 of the tour of go)
