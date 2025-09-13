package main

import "fmt"

func main() {
	person := map[string]string{
		"name":   "yusuf",
		"alamat": "mangunan",
	}
	fmt.Println(person)
	fmt.Println("Nama: ", person["name"], "\nAlamat: ", person["alamat"])

	// cara buat map yg lain
	var book map[string]string = make(map[string]string)
	book["title"] = "Fisika dasar"
	book["penulis"] = "zami"
	book["terbit"] = "2019"
	fmt.Println(book)
	delete(book, "title")
	fmt.Println(book)
}
