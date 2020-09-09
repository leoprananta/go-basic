package main

import "fmt"

func main(){
	// const firstName string = "Leonanta"
	// const lastName = "Pramudya"

	//atau 
	const(
		firstName string = "Leonanta"
		lastName = "Pramudya"
	)

	// firstName = "Leo"
	fmt.Println(firstName, lastName)
}

//const tidak bisa diubah nilainya, seperti val di kotlin