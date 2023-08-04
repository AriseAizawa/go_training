package main

import (
	"fmt"
	"math/rand"
)

func main() {
	r := rand.Intn(100)
	fmt.Println(r)

	switch true {
	case r%3 == 0 && r%5 == 0:
		fmt.Println("FizzBuzz")

	case r%3 == 0:
		fmt.Println("Fizz")

	case r%5 == 0:
		fmt.Println("Buzz")
	}

}
