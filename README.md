# Matrix Go Library

This is a simple and easy-to-use library for matrix calculations implemented in go. 
It provides a set of common operations such as addition, multiplication, transposition, inversion, determinant and others requisite for matrix manipulation in most applications.

## Installation

```
go get -u github.com/DQNEO/matrix
```

## Usage

To use this library, import it in your Go application:

```
import "github.com/DQNEO/matrix"
```

Here's an example that demonstrates how to create a new matrix and do some simple operations:

```go
package main

import (
    "fmt"
    "github.com/DQNEO/matrix"
)

func main() {
// Create a new 2x2 matrix
m := matrix.NewMatrix(2, 2, []float64{1, 2, 3, 4})

	// Print the matrix
	fmt.Println(m.String())

	// Transpose the matrix
	t := m.Tr()

	// Print the transposed matrix
	fmt.Println(t.String())
}
```

Please refer to the full API documentation for further details on available functions and methods.

## License

MIT license

## Authors

@DQNEO
