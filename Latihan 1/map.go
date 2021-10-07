package main

import "fmt"

func main() {

	persons := map[string]string{
		"name":    "Kiryu Kazuma",
		"address": "Kamurocho",
	}

	fmt.Println(persons)
	fmt.Println(persons["name"])
	fmt.Println(persons["address"])

	persons["title"] = "Ex-Yakuza" //Cara menambah data pada map

	fmt.Println(persons["title"])
	fmt.Println(persons)

	var cases map[string]string = make(map[string]string)

	cases["title"] = "Empty Lot Murder Case"
	cases["suspect"] = "Kiryu Kazuma"
	cases["investigator"] = "Takayumi Yagami"

	fmt.Println(cases)
	fmt.Println(len(cases)) //3

	delete(cases, "investigator") // ini untuk menghapus data pada map

	fmt.Println(cases)
	fmt.Println(len(cases)) //2

}
