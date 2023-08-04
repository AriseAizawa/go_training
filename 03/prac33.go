package main

import "fmt"

func judge_odd_even(n int) {
	if n%2 == 1 {
		fmt.Print(n)
		fmt.Println(":odd")
	} else {
		fmt.Print(n)
		fmt.Println(":even")
	}

}

func main() {
	sli := []int{1, 5, 214, 53, 6, 24, 72}
	fmt.Println(sli)

	for _, n := range sli {
		judge_odd_even(n)

	}

}
