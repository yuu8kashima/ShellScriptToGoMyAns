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

	sc := bufio.NewScanner(f)
	i := 0
	for sc.Scan() {
		i++
		fmt.Println(sc.Text())
	}
	fmt.Println(i)
}
