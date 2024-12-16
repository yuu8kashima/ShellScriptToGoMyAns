package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
 * 指定したディレクトリ内の全ファイルサイズの合計を計算する
 */
func main() {
	var total_file_size int64 = 0
	start_path := os.Args[1]
	filepath.Walk(start_path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() {
			// NOP
		} else {
			fmt.Printf("name: %v, size: %v\n", info.Name(), info.Size())
			total_file_size += info.Size()
		}
		return nil
	})
	fmt.Printf("total: %d\n", total_file_size)
}
