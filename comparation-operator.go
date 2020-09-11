package main

import "fmt"

func main() {
	//comparation for string
	var name1 = "Leo"
	var name2 = "Nanta"

	var result1 bool = name1 == name2
	var result2 bool = name1 > name2
	fmt.Println(result1)
	fmt.Println(result2)

	//comparation for number
	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}