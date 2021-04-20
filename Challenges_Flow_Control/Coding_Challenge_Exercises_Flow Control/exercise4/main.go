// Coding Exercise #4

// Using a for loop, an if statement and the logical and operator print out all the numbers between 1 and 500 that divisible both by 7 and 5.

package main

import "fmt"

func main() {

	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 { // if i is divisible both by 7 and 5
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}

// package main

// import "fmt"

// func main() {
// 	for a := 1; a < 500; a++ {
// 		if a%7 == 0 && a%5 == 0 {
// 			fmt.Printf("%d is divisable by 7 and 5\n", a)
// 		}
// 	}
// }
