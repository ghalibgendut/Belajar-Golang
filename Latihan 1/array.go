package main

import "fmt"

func main() {
	var names [3]string

	names[0] = "Ghalib"
	names[1] = "Sasmito"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2]) // ini berisi string kosong

	var val = [3]int{
		4,
		5,
		6,
	}

	fmt.Println(val)
	fmt.Println(val[0])
	fmt.Println(val[1])
	fmt.Println(val[2])

	fmt.Println(len(names)) //Cek panjang array

}
