package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

/**
 * 特定のプロセスIDを持つプロセスを終了させる。
 */
func main() {
	pid := os.Args[1] // 終了したいプロセスのPID
	_, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println("プロセスIDが数値ではない:", err)
		return
	}

	err = exec.Command("taskkill", "/PID", pid, "/T").Run()
	if err != nil {
		fmt.Println("プロセス終了に失敗しました:", err)
		return
	}

	fmt.Println("PID", pid, "のプロセスを終了しました")
}
