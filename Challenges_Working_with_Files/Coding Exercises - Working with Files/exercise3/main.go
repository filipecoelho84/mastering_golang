// Coding Exercise #3

// Remove the file created at exercise #1. Take care that the file is now called data.txt (it has been renamed at exercise #2).

// Perform error handling.

package main

import (
	"log"
	"os"
)

func main() {
	// removing the file
	err := os.Remove("data.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	newFile, err := os.Create("info.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	newFile.Close()

// 	oldFile := "info.txt"
// 	renamedFile := "data.txt"
// 	if renamedFile == oldFile {
// 		fmt.Println("File already exists!")
// 	} else {
// 		err = os.Rename(oldFile, renamedFile)
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = os.Remove("data.txt")
// 	// error handling
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// }
