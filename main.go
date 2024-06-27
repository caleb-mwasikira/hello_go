package main

import (
	"bytes"
	"compress/gzip"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

func hashData(data []byte) string {
	b_hash := md5.New().Sum(data)[:20]
	return hex.EncodeToString(b_hash)
}

func printSummary(data []byte) {
	fmt.Println("Length: ", len(data))
	fmt.Println("Hash: ", hashData(data))
}

func main() {
	// Data to be compressed
	data := make([]byte, 1024)
	_, err := rand.Read(data)
	if err != nil {
		fmt.Printf("Error generating random bytes; %v\n", err)
		return
	}

	fmt.Println("Before compression: ")
	printSummary(data)

	// Create a buffer to store compressed data
	var buf bytes.Buffer

	// Any data written to gzip writer is automatically compressed
	gw := gzip.NewWriter(&buf)
	_, err = gw.Write(data)
	if err != nil {
		fmt.Println("Error during compression:", err)
		return
	}

	// Close the writer to flush any remaining data
	if err := gw.Close(); err != nil {
		fmt.Println("Error closing gzip writer:", err)
		return
	}

	fmt.Println("After compression: ")
	printSummary(buf.Bytes())

	// Decompress the data (optional)
	reader, err := gzip.NewReader(&buf)
	if err != nil {
		fmt.Println("error initialising reader:", err)
		return
	}

	decompressed, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error during decompression:", err)
		return
	}

	fmt.Println("After decompresion: ")
	printSummary(decompressed)
}
