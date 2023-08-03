package main

import "fmt"

func main() {

	fruit := make(map[string]int, 2)
	fruit["apple"] = 100
	fruit["banana"] = 200
	fruit["orange"] = 50

	fru_list := []string{"apple", "banana", "orange", "lemon"}

	for _, v := range fru_list {

		if fruit[v] == 0 {
			fmt.Println(v + " はありません。")
		} else {
			fmt.Print(v + " の値段は、")
			fmt.Print(fruit[v])
			fmt.Println("円です。")
		}

	}

}
