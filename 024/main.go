package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

/*
 * 指定されたファイルからデータを読み込み、各行の長さを表示するプログラムを作成してください。
 * コメント: utf8の文字数でカウントするようにしてみた
 */
func main() {
	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("ファイルを開けません: %v\n", err)
		return
	}
	defer file.Close() // 関数終了時にファイルを閉じる

	// bufio.Scannerを使用して一行ずつ読み込む
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		count := utf8.RuneCountInString(line)
		fmt.Println(count)
	}
}
