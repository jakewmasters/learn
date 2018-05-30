// $ go run example.go
package main

import (
	"fmt"
	"math"
	"math/rand"
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

func print_math(){
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
	fmt.Println(math.Pi)
}

func add(x int, y int) int {
	return x + y
}

func main() {
	print_time()
	print_types()
	print_math()
	fmt.Println(add(42, 13))
}
