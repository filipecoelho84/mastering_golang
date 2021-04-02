// Coding Exercise #3

// Consider the following map declaration: m := map[int]bool{"1": true, 2: false, 3: false}

// 1. The above map declaration has an error. Solve the error!

// 2. Delete a key:value pair from the map.

// 3. Iterate over the map and print out each key and value.

package main

import "fmt"

func main() {

	// 1.
	m := map[int]bool{1: true, 2: false, 3: false}

	// 2.
	delete(m, 2)

	// 3.
	for k, v := range m {
		fmt.Printf("k: %d, v: %t\n", k, v)
	}

}

// package main

// import "fmt"

// func main() {
// 	m := map[int]bool{1: true, 2: false, 3: false}

// 	_ = m

// 	fmt.Println(m)

// 	delete(m, 3)

// 	fmt.Println(m)

// 	for k, v := range m {
// 		fmt.Printf("%d -> %t\n", k, v)
// 	}
// }
