// Coding Exercise #1

// Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.

package main

import "fmt"

func main() {

	for i := 1; i <= 50; i++ {
		if i%7 == 0 { // if i is divisible by 7
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")
}

// package main

// import "fmt"

// func main() {

// 	a := 1

// 	for a < 50 {
// 		if a%7 == 0 {
// 			fmt.Printf("%d is divisable by 7\n", a)
// 		}
// 		a++

// 	}

// }
