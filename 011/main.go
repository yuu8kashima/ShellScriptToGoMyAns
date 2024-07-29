package main

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

/**
 * 毎日午前3時にバックアップスクリプトを実行するcronジョブを作成する。
 */
func main() {
	c := cron.New()
	c.AddFunc("* * * * *", backupFunc)
	c.Start()
	defer c.Stop()

	select {}
}

func backupFunc() {
	fmt.Println("x")
}
