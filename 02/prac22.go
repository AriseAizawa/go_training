package main

import (
	"fmt"
	"strconv"
)

var var_str = "3.14"
var var_flo = 3.14

func main() {
	fmt.Println(var_str)
	fmt.Printf("%T\n", var_str)
	fmt.Println(var_flo)
	fmt.Printf("%T\n", var_flo)

	var_str2, _ := strconv.ParseFloat(var_str, 64)
	var_flo2 := strconv.FormatFloat(var_flo, 'f', -1, 64)

	fmt.Println(var_str2)
	fmt.Printf("%T\n", var_str2)
	fmt.Println(var_flo2)
	fmt.Printf("%T\n", var_flo2)

}
