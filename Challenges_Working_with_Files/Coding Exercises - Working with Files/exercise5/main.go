// Coding Exercise #5

// Create a Go Program that reads the entire contents of a file called info.txt  using a scanner (bufio package) line by line.

// The file exists in the current working directory.

/////////////////////////////////
// Run the program: go run main.go
/////////////////////////////////

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// opening the file in read-only mode. The file must exist (in the current working directory)
	// use a valid path!
	file, err := os.Open("info.txt")
	// error handling
	if err != nil {
		log.Fatal(err)
	}
	// defer closing the file
	defer file.Close()

	// the file value returned by os.Open() is wrapped in a bufio.Scanner just like a buffered reader.
	scanner := bufio.NewScanner(file)

	// reading the whole file line by line:
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// checking for any possible errors:
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {
// 	file, err := os.Open("info.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	success := scanner.Scan()
// 	if success == false {
// 		err = scanner.Err()
// 		if err != nil {
// 			log.Println("Scan was completed and it reached the end of the file.")
// 		} else {
// 			log.Fatal(err)
// 		}
// 	}

// 	fmt.Println("First Line found:\n", scanner.Text())

// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// }
