// File: mymath/go.mod
module github.com/l4dRUSEN/mymath

go 1.16

// File: mymath/addition.go
package mymath

func Add(a, b int) int {
return a + b
}

// File: main/go.mod
module main

go 1.16

require (
github.com/l4dRUSEN/mymath latest
)

// File: main/main.go
