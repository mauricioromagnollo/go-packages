# go-packages

> A simple repository to learn how to publish reusable packages in golang.

## Get The Package

```sh
go get "github.com/mauricioromagnollo/go-packages"
```

## Usage Example

```go
import "github.com/mauricioromagnollo/go-packages/typevalidator"

typevalidator.GetTypeOf(1)
// "int"

typevalidator.GetTypeOf("John")
// "string"

...
```
