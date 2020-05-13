package main

import ("fmt" // format - how we print
		"math"
		"math/rand") // import multiple packages
		
func main() {
	fmt.Println("The square root of 4 is", math.Sqrt(4))
	fmt.Println("A number from 1-100", rand.Intn(100))
}