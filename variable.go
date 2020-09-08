package main

import "fmt"

func main(){
	//var name tipedata
	var name string

	name = "Leonanta Pramudya"
	fmt.Println(name)

	name = "leopra_nanta"
	fmt.Println(name)

	//var name = "String"
	var friendName = "Sukri"
	fmt.Println(friendName)

	var age = 25
	fmt.Println(age)

	//var tidak wajib,tapi pakai:= untuk deklarasi pertama
	country := "Indonesia"
	country = "Mongolia"
	fmt.Println(country)

	//atau pakai seperti ini

	var(
		firstName = "Leonanta Pramudya"
		lastName = "Kusuma"
	)

	fmt.Println(firstName, lastName)
	
}
