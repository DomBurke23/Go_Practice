package main

import (
// must be on new lines otherwise wont build. 
	"fmt"
	"math"
	"runtime"
	"time"
)

func main() {

	/*
	Go has only one looping construct, the for loop. 
	The basic for loop has three components separated by semicolons:
	the init statement: executed before the first iteration
	the condition expression: evaluated before every iteration
	the post statement: executed at the end of every iteration	
	*/
	sum1 := 0
	for i := 0; i < 10; i++ {
		sum1 += i
	}
	fmt.Println(sum1)
	
	// The init and post statements are optional. 
	// That is equivalent to a while loop but still called FOR 
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum2
	}
	fmt.Println(sum2)
	
	// if statement call 
	fmt.Println(sqrt(2), sqrt(-4))
	
	// if statement with short statement 
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
	
	// if else 	
	fmt.Println(
		pow2(3, 2, 10),
		pow2(3, 3, 20),
	)
	
	// newtons method 
	fmt.Println(Sqrt(2))
	
	// switch case 
	switchFunc()
	
	//switch case 2
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switchFunc2(today)
	
	//switch no condition
	switchFunc3()
	
	// Defer 
	// A defer statement defers the execution of a function until the surrounding function returns. 
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns. 
	defer fmt.Println("world")
	fmt.Println("hello")
	
	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order. 
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// if statement 
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	// The fmt.Sprint() function in Go is used to format and concatenate values into a string.
	return fmt.Sprint(math.Sqrt(x))
}

// the if statement can start with a short statement to execute before the condition.
// Variables declared by the statement are only in scope until the end of the if.
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

// if else 
func pow2(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

// let's implement a square root function: given a number x, 
// we want to find the number z for which z² is most nearly x.
// newtons method 
// in Go, if there are two methods with the same name but one starts with a capital letter and the other with a lowercase letter, 
// Go can differentiate between them based on the capitalization.
// Methods starting with a lowercase letter are private or package-private. 
// Methods starting with a capital letter are public or exported and can be accessed from other packages. 
const precision = 0.0000001 // desired precision
func Sqrt(x float64) float64 {
	z := x // initial guess
	for {
		newZ := z - (z*z-x)/(2*z) // Newton's method
		if math.Abs(newZ-z) < precision {
			return newZ // close enough
		}
		z = newZ
	}
}

// switch case , shorter than an if-else 
// the break keyword is not needed, it is implied by go 
func switchFunc(){
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// switch case passing in weekday 
func switchFunc2(today time.Weekday){
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// switch with no condition is the same as switch true 
// clean way to write long if-then-else chains 
func switchFunc3(){
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}