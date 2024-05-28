// Tutorial from https://go.dev/tour/welcome/1 

package main

import (
	"fmt"
	"time"
	"math"
	"math/rand"
	"math/cmplx"
)

// variables 
var c, python, java bool
var i, j int = 1, 2

// variables can be factored into blocks 
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// standard print line 
	fmt.Println("Welcome to the playground!")
	
	// print with a line break 
	fmt.Printf("Dominique \nSmith.\n")
	
	// print current date time 
	fmt.Println("The time is", time.Now())
	
	// print random integer 
	fmt.Println("My favorite number is", rand.Intn(10))
	
	// print square root value using string interpolation 
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	
	// print value of Pi , must be a capital letter as its an exported name. 
	fmt.Println(math.Pi)
	
	// print sum passing params to another function
	fmt.Println(add_v1(42, 13))
	
	// print sum passing params to another function
	fmt.Println(add(42, 13))
	
	// set variables then call function swap 
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	
	// Naked Return function 
	fmt.Println(split(17))
	
	// print variables 
	var i int
	fmt.Println(i, c, python, java)
	
	// initialise variables 
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
	
	// Inside a function, the := short assignment statement can be used in place of a var declaration with implicit type.
	var ii, j int = 1, 2
	k := 3
	cc, python, java := true, false, "no!"
	fmt.Println(ii, j, k, cc, python, java)
	
	// print basic types 
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	
	// variables declared without an explicit initial value are given zero value
	var iii int
	var f float64
	var bb bool
	var s string
	fmt.Printf("%v %v %v %q\n", iii, f, bb, s)
	
	// Type Conversion 
	// The expression T(v) converts the value v to the type T.
	var x, y int = 3, 4
	var ff float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(ff)
	fmt.Println(x, y, z)
	
	// type inference 
	//When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax), the variable's type is inferred from the value on the right hand side.
	v := 42 // int
	fmt.Printf("v is of type %T\n", v)
	vv := 42.5 // int
	fmt.Printf("v is of type %T\n", vv)
}

// functions that the main function an call. 

// Notice the return type is AFTER the variable name and params. 
func add_v1(x int, y int) int {
	return x + y
}

// this version omits the 2nd type as it shares the same type. 
func add(x, y int) int {
	return x + y
}

// a function can return multiple results 
func swap(x, y string) (string, string) {
	return y, x
}

// Go's return values may be named. 
// If so, they are treated as variables defined at the top of the function.
// A return statement without arguments returns the named return values. This is known as a "naked" return.
// Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
