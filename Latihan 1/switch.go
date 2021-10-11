package main

import "fmt"

func main() {

	name := "Kazuma Kiryu"

	switch name {
	case "Kazuma Kiryu":
		fmt.Println("Kazuma Kiryu is an Ex-Yakuza")

	case "Mijima Goro":
		fmt.Println("Mad Dog Mijima Goro")

	default:
		fmt.Println("No Name")
	}

	// switch length := len(name); length >= 5 {
	// case true:
	// 	fmt.Println("Name is to long")
	// case false:
	// 	fmt.Println("Name is right")
	// }

	length := len(name)

	switch {
	case length > 50:
		fmt.Println("Name is to long!")
	case length < 20:
		fmt.Println("Name is quite long")
	default:
		fmt.Println("No Name Registered!")
	}

}
