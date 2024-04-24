package main

import (
	"log"

	my_io "example.com/hello_go/io"
)

func main() {
	// src := "/home/netrunner/College"
	// err := my_io.ZipArchive(src, "")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("done archiving files")

	archive := "/home/netrunner/College.zip"
	err := my_io.Unzip(archive)
	if err != nil {
		log.Fatal(err)
	}
}
