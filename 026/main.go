package main

import (
	"fmt"
	"os"
)

/*
 * ユーザーが入力した文字列が回文であるかどうかを判定するプログラムを作成してください。
 */
func main() {
	sentence := os.Args[1]
	reverse_sencente := str_reverse(sentence)
	if sentence == reverse_sencente {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

/*
 * 文字列を逆にする関数 // https://seven-901.hatenablog.com/entry/2021/06/14/234000
 */
func str_reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
