// Coding Exercise #6

// Consider the following slice declaration: friends := []string{"Marry", "John", "Paul", "Diana"}

// Using copy() function create a copy of the slice. Prove that the slices are not connected by modifying one slice and notice that the other slice is not modified.

package main

import "fmt"

func main() {
	friends := []string{"Marry", "John", "Paul", "Diana"}
	yourFriends := make([]string, len(friends))
	copy(yourFriends, friends)

	yourFriends[0] = "Dan"

	fmt.Println(friends, yourFriends)

}

// package main

// import "fmt"

// func main() {
// 	friends := []string{"Marry", "John", "Paul", "Diana"}

// 	friends2 := make([]string, len(friends))

// 	friends3 := copy(friends2, friends)

// 	friends2 [0] = "Filipe"

// 	_ = friends3

// 	fmt.Println(friends, friends2)
// }
