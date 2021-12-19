package main

import "fmt"

func main() {

	type NoKtp string
	type Married bool

	var noKtpGhalib NoKtp = "3275020409980031"
	var MarriedStatus Married = false

	fmt.Println(noKtpGhalib)
	fmt.Println((MarriedStatus))

}
