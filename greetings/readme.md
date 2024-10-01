# Createa a module

## Create a module

```bash
go mod init [module_name]
```

## define a function in the module

define a function in the module that starts with a capital letter, so it can be exported and used by other modules.

```go
package greetings

import "fmt"


func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
```

now, the module is ready to be used by other modules.
