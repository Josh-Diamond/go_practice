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

func add(x,y float64) float64 { //float64 refers to 64bit instead of 32
	return x + y
}

// const x int = 5
func multiple(a,b string) (string, string) {
	return a,b
}

func main() {
	// var num1 float64 = 0.1
	// var num2 float64 = 10.21
	// or
	// var num1, num2 float64 = 0.1, 10.21

	// fmt.Println(add(num1, num2))
	w1, w2 := "Hey", "Josh"
	fmt.Println(w1,w2 )
}