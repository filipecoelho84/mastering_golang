package main

import (
	"fmt"
	"time"
)

func main() {
	language := "golang"

	switch language {
	case "Python":
		fmt.Println("You are learning Python! You don't use curly braces but indentation!!9")

	case "Go", "golang":
		fmt.Println("Good, go for Go! You are using curly braces {}!")

	default:
		fmt.Println("Any other programming language is a good start!")
	}

	n := 5
	switch true {
	case n%2 == 0:
		fmt.Println("Even number!")
	case n%2 != 0:
		fmt.Println("Odd number!")

	default:
		fmt.Println("Never here!!")
	}

	hour := time.Now().Hour()
	// fmt.Println(hour)

	switch {
	case hour < 12:
		fmt.Println("Good Morning!")
	case hour < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening!")
	}

}
