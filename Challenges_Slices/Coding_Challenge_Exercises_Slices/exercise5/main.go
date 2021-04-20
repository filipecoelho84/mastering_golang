// Coding Exercise #5

// Consider the following slice declaration: nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

// Using a slice expression and a for loop iterate over the slice ignoring the first and the last two elements.

// Print those elements and their sum.

package main

import "fmt"

func main() {
	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	sum := 0
	for _, v := range nums[2 : len(nums)-2] {
		fmt.Println(v)
		sum += v
	}
	fmt.Println("Sum:", sum)

}

// package main

// import (
// 	"fmt"
// 	"strconv")

// func main() {
// 	nums := []int{5, -1, 9, 10, 1100, 6, -1, 6}

// 	nums1 := nums[1:6]

// 	sum := 0

// 	for _, v := range nums1 {
// 		num, err := strconv.ParseInt(v, 0, 64)
// 		if err != nil {
// 			continue
// 		}
// 		sum += num
// 	}

// 	fmt.Printf("Sum: %d\n", sum)
// }
