# go-vector

This package is an example of a simple interface for vector manipulation in go. It provides operations like add, subtract, multiply, etc. It also has a test and benchmarking suite.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Installing

Get the package:
```
> go get github.com/sbarratt/go-vector
```

To use it in a project:
```go
import "github.com/sbarratt/go-vector/vector"
```

### Usage

Create and edit vectors:
```go
v = vector.MakeVector(2)
v.Set(0, 1)
v.Get(0) // 1
```

Familiar Operations:
```go
vector.Add(v1, v2)
vector.AddScalar(v1, 1)
v.Min()
v.Mean()
v.Var()
v.Std()
```

Print:
```
fmt.Println(v.ToString())
```

## Contributing

Submit a pull request.

## Author

[Shane Barratt](shanebarratt.com)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details