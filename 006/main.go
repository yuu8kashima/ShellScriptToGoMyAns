package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	filename := os.Args[1]
	patternString := os.Args[2]

	pattern, err := regexp.Compile(patternString)
	if err != nil {
		fmt.Println("error")
		return
	}

	f, _ := os.Open(filename)
	defer f.Close()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		match := pattern.MatchString(line)
		if match {
			fmt.Println(line)
		}
	}
}
