package main

import "fmt"

func main() {

	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan counter ke : ", counter)
		counter++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println("Perulangan i ke : ", i)
	}

	nameSlice := []string{"Kazuma", "Kiryu", "Goro", "Mijima", "Ghalib", "Sasmito"}

	for i := 0; i < len(nameSlice); i++ {
		fmt.Println(nameSlice[i])
	}

	for i, value := range nameSlice {
		fmt.Println("Name index : ", i, " Value : ", value)
	}

	person := make(map[string]string)
	person["name"] = "Kazuma Kiryu"
	person["Job"] = "Real Estate Enterpreneur"

	for key, value := range person {
		fmt.Println(key, " = ", value)
	}

}
