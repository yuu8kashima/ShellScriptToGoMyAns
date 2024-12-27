package main

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

/**
 * 指定されたディレクトリを監視し、新しいファイルが作成されたらそのファイル名を表示する。
 * fsnotifyパッケージを使えばできそう… だったのだが、Windowsだと変化があったディレクトリの監視が限界らしい。
 * WSLでは一応動いた。
 */
func main() {
	dirname := os.Args[1]

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				// Writeの例:
				if event.Has(fsnotify.Write) {
					log.Println("modified file:", event.Name)
				}
				// 今回のお題はCreate
				if event.Has(fsnotify.Create) {
					log.Println("Create file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(dirname) // (5)
	if err != nil {
		log.Fatal(err)
	}

	// これでmain関数を永遠にブロックできるらしい
	<-make(chan struct{})
}
