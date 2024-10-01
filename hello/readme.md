# Getting start

## prerequisites

- make sure to install Go

## Notes

### to enable dependency tracking for the code

```bash
go mod init [module name]
```

for example:

```bash
go mod init example/hello
```

### run the code

```bash
go run .
```

it will run the main package in the current directory

### see list of commands

```bash
go help
```

### add new module requirement and sums

```bash
go mod tidy
```

this command helps to check the dependencies and update the go.mod file, clean up unused dependencies and add missing dependencies that are required by the code but not listed in the go.mod file.
