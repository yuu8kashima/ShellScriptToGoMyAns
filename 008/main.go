package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/**
 * ファイルの中で特定の文字列を検索し、その文字列が現れた行番号と行を表示する。
 */
func main() {
	filename := os.Args[1]
	searchString := os.Args[2]

	f, _ := os.Open(filename)
	defer f.Close()

	sc := bufio.NewScanner(f)
	lineCount := 0
	for sc.Scan() {
		lineCount++
		line := sc.Text()
		if strings.Contains(line, searchString) {
			fmt.Println(strconv.Itoa(lineCount) + ": " + line)
		}
	}
}
