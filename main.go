package main

import (
	"log"

	my_io "example.com/hello_go/io"
)

func main() {
	src := "/home/netrunner/College"

	err := my_io.ZipArchive(src, "")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("done archiving files")
}
