package main

import "fmt"

func main() {
	var months = [...]string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)      // [May, June, July] masih dalam bentuk array
	fmt.Println(len(slice1)) // Panjang dari slice1 hasilnya adalah 3
	fmt.Println(cap(slice1)) // Kapasitas dari slice1 hasilnya adalah 8

	// months[5] = "Berubah"
	// fmt.Println(slice1) // [May, Berubah, July]

	// slice1[0] = "Data di array ikut berubah"
	// fmt.Println(months) //data di array ikut beruah walau data slice yang diubah, karena slice reference ke array

	var slice2 = months[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "Kiryu Kazuma") //append akan membuat array baru apabila jumlah kapasitas array tidak cukup lagi, namun apabila masih cukup maka akan merubah array utama
	fmt.Println(slice3)
	slice3[1] = "Majima Goro"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 3, 5)
	newSlice[0] = "Ghalib Sasmito"
	newSlice[1] = "Kiryu Kazuma"
	newSlice[2] = "Majima Goro"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice) //[Ghalib Sasmito Kiryu Kazuma Majima Goro]

	iniArray := [5]int{1, 2, 3, 4, 5} // atau [...]int{1,2,3,4,5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
