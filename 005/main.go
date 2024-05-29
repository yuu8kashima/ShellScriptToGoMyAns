package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	start_path := os.Args[1]
	filepath.Walk(start_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(path)
		return nil
	})
}
