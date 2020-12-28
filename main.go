package main

import (
	"fmt"
	"go-compression/pkg/compress"
	"go-compression/pkg/decompress"
)

func main() {
	fmt.Println("Enter a string :")
	var exStr string
	fmt.Scanln(&exStr)

	// compress
	compressed := compress.Compress(exStr)
	fmt.Println("Post Compression : ", compressed)

	// decompress
	decompressed := decompress.Decompress(compressed)
	fmt.Println("Post decompression : ", decompressed)
}
