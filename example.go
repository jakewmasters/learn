// $ go run example.go
package main

import (
	"fmt"
	"math/cmplx"
	"time"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func print_time(){
	fmt.Println("Welcome. The time is", time.Now())
}

func print_types(){
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func main() {
	print_time()
	print_types()
}
