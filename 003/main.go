package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := os.Args[1]
	f, _ := os.Open(filename)
	defer f.Close()

	// tacs := make([]string, 0)
	tacs := []string{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		tacs = append(tacs, sc.Text())
	}

	size := len(tacs)
	// 配列を逆順にしているが、流石に無駄（要件的には配列をお尻から走査するだけで良かった）
	for i := 0; i < size/2; i++ {
		tacs[i], tacs[size-i-1] = tacs[size-i-1], tacs[i]
	}
	for i := 0; i < size; i++ {
		fmt.Println(reverse(tacs[i]))
	}
}

func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
