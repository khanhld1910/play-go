# Call module from another module

## Call module from another module

import the module to the main module and call the function.

```go
package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Sneeze Cheese")
	fmt.Println(message)
}
```

see `example.com/greetings` is the module name that we created in the previous step, and is now being used in the main module.

## replace the module by the local module, since we have not published the module yet.

```bash
go mod edit -replace example.com/greetings=../greetings
```

after that, clean up the go.mod

```bash
go mod tidy
```

now, the `example.com/greetings` module is replaced by the local module `../greetings`.

## run the main module

```bash
go run .
```
