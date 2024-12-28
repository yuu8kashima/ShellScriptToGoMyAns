package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

// 実行ログを保存するディレクトリと拡張子
const save_dir = "./tmp"
const save_ext = "execlog"

/*
 * 指定されたコマンドを実行し、そのコマンドの出力をタイムスタンプ付きのファイル名で保存するプログラムを作成してください。
 * また、保存されたファイルの一覧を表示してください。
 */
func main() {
	cmd := os.Args // コマンドライン引数を外部コマンドと見なす
	out_file := makeExecLogFilename()

	// 実行
	out, err := exec.Command(cmd[1], cmd[2:]...).Output()
	if err != nil {
		log.Fatal(err)
	}
	// ファイル出力
	os.WriteFile(out_file, out, 0660)

	// 実行ログファイルをリストアップ
	listLogFilenames()
}

func makeExecLogFilename() string {
	now := time.Now()
	now_str := now.Format("2006-01-02-15-04-05")
	out_file := save_dir + "/" + now_str + "." + save_ext
	return out_file
}

func listLogFilenames() {
	pattern := save_dir + "/*." + save_ext
	files, err := filepath.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file)
	}
}
