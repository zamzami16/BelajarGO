package main

import (
	"belajar-go-rest/helper"
	"fmt"
)

func main() {
	server := InitializeServer()
	fmt.Println("Berjalan di local host:3000")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
