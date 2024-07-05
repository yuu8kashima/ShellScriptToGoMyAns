package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dirCount := 0
	fileCount := 0
	start_path := os.Args[1]
	filepath.Walk(start_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() {
			dirCount++
		} else {
			fileCount++
		}
		return nil
	})
	fmt.Printf("dirCount: %d\n", dirCount)
	fmt.Printf("fileCount: %d\n", fileCount)
}
