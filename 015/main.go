package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
 * 指定したディレクトリ内のファイルをリストアップし、各ファイルの拡張子ごとにファイル数を表示する。
 */
func main() {
	start_path := os.Args[1]
	count_map := map[string]int{}

	filepath.Walk(start_path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			count_map[ext]++
		}

		return nil
	})
	for k, v := range count_map {
		fmt.Printf("%s: %d\n", k, v)
	}
}
