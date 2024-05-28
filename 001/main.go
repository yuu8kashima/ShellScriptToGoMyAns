package main

import (
	"fmt"
	"os"
)

func main() {
	// dir := "C:\\Users\\yuu8k\\Work\\my-blog"
	// ex, _ := os.Executable()
	// currentDir := filepath.Dir(ex)
	// fmt.Println("currentDir=" + currentDir)

	files, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		txt := ""
		if file.IsDir() {
			txt += "D "
		} else {
			txt += "  "
		}
		fmt.Println(file.Name())
	}
}
