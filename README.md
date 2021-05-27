## Environment Variables Unmarshal Utility for Golang

A convenient way to read environment variable.

### Usage:

1. `go get github.com/woohaha/envUnmarshal-go`

1. Define struct

```go
type Param struct {
First  string `env:"first"`
Second string `env:"second"`
Third  string `env:"third"`
}
```
> Note: Type of Fields must be `string`

1. Declare an object of the struct

```go
var param *Param
```

1. Inject environment variables into object

```go
import env "github.com/woohaha/envUnmarshal-go"
env.Unmarshal(param)
```

1. Enjoy

```go
fmt.Println(param.First)
```