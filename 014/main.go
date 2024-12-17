package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type MyFileInfo struct {
	name string
	size int64
}

/*
 * 指定したディレクトリ内のファイルをリストアップし、各ファイルのサイズを表示する。ファイル情報を型として定義してください。
 */
func main() {
	start_path := os.Args[1]
	filepath.Walk(start_path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			myFileInfo := MyFileInfo{
				name: path + string(filepath.Separator) + info.Name(), // filepath.Joinを使う方が良いが、今回は試しにSeparator使用
				size: info.Size(),
			}
			fmt.Println(myFileInfo)
		}

		return nil
	})
}
