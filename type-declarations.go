package main

import "fmt"

func main(){
	//type NoKTP untuk string
	type NoKTP string
	type Married bool

	var noKtpKu NoKTP = "1234567890"
	var marriedStatus Married = false

	fmt.Println(noKtpKu)
	fmt.Println(marriedStatus)
}