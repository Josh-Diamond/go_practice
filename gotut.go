// package main

// import ("fmt" // format - how we print
// 		"math"
// 		"math/rand") // import multiple packages
		
// func main() {
// 	fmt.Println("The square root of 4 is", math.Sqrt(4))
// 	fmt.Println("A number from 1-100", rand.Intn(100))
// }

package main

import ("fmt")

func main() {
	x := 15
	a := &x // points to x (memory address)
	fmt.Println(a)
	fmt.Println(*a) // points to x (read-through/find value)
	*a = 5 // re-assigns value of 5 to x
	fmt.Println(x)
	*a = *a**a // reassign value of a to value of a times value of a ( 5 * 5 = 25)
	fmt.Println(x)
}