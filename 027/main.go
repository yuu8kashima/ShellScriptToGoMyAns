package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// 指定されたディレクトリのサイズを再帰的に計算し、結果を表示するプログラムを作成してください。
func main() {
	start_path := os.Args[1]
	sum_of_file_size := int64(0)
	filepath.Walk(start_path, func(path string, info os.FileInfo, err error) error {
		sum_of_file_size += info.Size()
		return nil
	})
	fmt.Println(sum_of_file_size)
}
