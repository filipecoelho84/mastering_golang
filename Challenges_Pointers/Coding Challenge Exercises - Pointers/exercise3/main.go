// Coding Exercise #3

// Consider the following variable declaration:x, y := 5.5, 8.8

// Write a function that swaps the values of x and y. After calling the function x will be 8.8 and y will 5.5

package main

import "fmt"

func swap(a, b *float64) {
	*a, *b = *b, *a
}

func main() {
	x, y := 5.5, 8.8
	swap(&x, &y)

	fmt.Printf("x is %v and y is %v\n", x, y)

}

// package main

// import "fmt"

// func swapTheValues(x *float64, y *float64) {
// 	z := *x

// 	*x = *y

// 	*y = z

// }

// func main() {
// 	x, y := 5.5, 8.8
// 	swapTheValues(&x, &y)

// 	fmt.Printf("The value of x is %v and the value of y is %v!\n", x, y)
// }
