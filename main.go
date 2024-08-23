package main

import (
	"fmt"
	"os"

	"github.com/theluckiestsoul/meta-mancer/internal/file"
	"github.com/theluckiestsoul/meta-mancer/pkg/metadata"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file path")
		return
	}
	filePath := os.Args[1]
	fmt.Printf("Processing file: %s\n", filePath)

	fileType := file.GetFileType(filePath)
	fmt.Printf("File type: %s\n", fileType)

	switch fileType {
	case "image/jpeg":
		metadata.ExtractJPEGMetadata(filePath)
	default:
		fmt.Println("Unsupported file type")
	}
}
