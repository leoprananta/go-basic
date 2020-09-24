package main

import "fmt"

func main() {
	var names[3]string

	names[0] = "Leonanta"
	names[1] = "Kurniawan"
	names[2] = "Kusuma"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//langsung deklarasi
	var values = [3]int{
		90,
		86,
		89,
	}

	fmt.Println(values)

	//func len
	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string

	fmt.Println(len(lagi))
}