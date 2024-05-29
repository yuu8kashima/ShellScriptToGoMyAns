package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * ファイルの単語数を数える
 */

func main() {
	filename := os.Args[1]
	f, _ := os.Open(filename)
	defer f.Close()

	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	wc := 0
	for sc.Scan() {
		// fmt.Println(sc.Text())
		wc++
	}
	fmt.Println(wc)
}
