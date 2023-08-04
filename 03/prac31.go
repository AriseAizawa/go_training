package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.Intn(100)
	fmt.Println(r)

	if r%3 == 0 && r%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if r%3 == 0 {
		fmt.Println("Fizz")
	} else if r%5 == 0 {
		fmt.Println("Buzz")
	}

}
