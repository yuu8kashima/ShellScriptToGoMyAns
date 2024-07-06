package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/**
 * 実行中のすべてのプロセスの一覧を表示し、特定のプロセス名をフィルタリングする。
 * （Windows版）
 */
func main() {
	// out, err := exec.Command("ps", "-o", "pid,command").Output()
	searchProcessName := os.Args[1]
	out, err := exec.Command("tasklist").Output()
	if err != nil {
		fmt.Println("psコマンドの実行に失敗しました:", err)
		return
	}

	processes := strings.Split(string(out), "\n")
	processes = processes[3:] // 最初の3行はヘッダみたいなものなので飛ばす
	for _, process := range processes {
		fields := strings.Fields(process)
		if len(fields) > 2 {
			command := fields[0]
			pid := fields[1]

			if strings.Contains(command, searchProcessName) {
				fmt.Println("PID:", pid, ", コマンド:", command)
			}
		}
	}
}
