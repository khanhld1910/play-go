# Getting started with multi-module workspaces

## create a module

create a module using

```bash
go mod init [module-name]
```

add a dependency to the module, using `go get`

```bash
go get [module-name]
```

## create the work space

- initialize the work space

```bash
go work init ./[existing-module-name]
```

this will create a `go.work` file to manage the workspace, containing the modules in the `existing-module-name` directory.
the `go.work` file has similar syntax to the `go.mod` file.

- run the program in the workspace directory

```bash
go run ./[existing-directory-name]
```

The Go command includes all the modules in the workspace as main modules

## download and modify a dependency

- clone the repository

```bash
git clone [repository-url]
```

- add the module to the workspace

```bash
go work use ./[repository-name]
```

the `go.work` file will be updated to include the new module.

- modify the program to use the function
