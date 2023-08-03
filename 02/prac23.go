package main

import "fmt"

func main() {

	sli := []int{1, 3, 66, 12}
	fmt.Println(sli)

	for i := 0; i < len(sli); i++ {
		fmt.Println(sli[i])
	}

}
