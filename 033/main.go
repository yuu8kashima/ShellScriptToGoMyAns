package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 指定されたディレクトリ内のすべてのファイルを再帰的に検索し、
// 指定された拡張子のファイルのみをリストアップするプログラムを作成してください。
func main() {
	ext := os.Args[1]
	start_path := os.Args[2]

	filepath.Walk(start_path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if filepath.Ext(path) == ext {
				fmt.Println(path)
			}
		}

		return nil
	})
}
