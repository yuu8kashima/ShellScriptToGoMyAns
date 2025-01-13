package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
)

// 2つの文字列が与えられたとき、それらがアナグラムであるかどうかを判定するプログラムを作成してください。
func main() {
	str1, str2 := os.Args[1], os.Args[2]
	if str1 == str2 {
		// メモ: 同一文字列はアナグラムでないとしておく
		fmt.Println("eq")
		return
	}

	// 中身をソートしてeqならアナグラム
	r_str1 := []rune(str1)
	r_str2 := []rune(str2)

	sort.Slice(r_str1, func(i, j int) bool { return r_str1[i] < r_str1[j] })
	sort.Slice(r_str2, func(i, j int) bool { return r_str2[i] < r_str2[j] })
	fmt.Println(r_str1, r_str2)
	if reflect.DeepEqual(r_str1, r_str2) {
		fmt.Println("アナグラム")
	}
}
