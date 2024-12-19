package main

import (
	"fmt"
	"os"

	"github.com/go-ping/ping"
)

/*
 * 指定したホスト名またはIPアドレスへのpingが成功するかどうかを確認する。
 */
func main() {
	address := os.Args[1]
	// 以下はhttps://github.com/go-ping/pingのサンプルほぼコピペ
	pinger, err := ping.NewPinger(address)
	pinger.SetPrivileged(true)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Println(stats)
}
