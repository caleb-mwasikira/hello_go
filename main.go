package main

import (
	io "example.com/hello_go/io"
)

func main() {
	filename := "./data/file.md"
	new_filename := "./data/silent_planet.md"
	config_file := "./data/test.conf"

	io.ReadFileStats(filename)
	io.ReadWholeFile(filename)
	io.ReadByLine(new_filename)
	io.ReadByWord(config_file)
}
