package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type CountInfo struct {
	Ext   string `json:"ext"`
	Count int    `json:"count"`
}

/*
 * 問題15と同じ機能を実装し、各ファイルの拡張子ごとのファイル数をJSONで出力してください。
 */
func main() {
	start_path := os.Args[1]
	count_map := map[string]int{}

	filepath.Walk(start_path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			ext := filepath.Ext(info.Name())
			count_map[ext]++
		}

		return nil
	})
	// {"拡張子" : カウント} のjsonならこれでOK
	json_bytes, _ := json.Marshal(count_map)
	fmt.Println(string(json_bytes))

	// [{ext: "拡張子", count: カウント},...]のjsonならこう。
	countInfoObj := []CountInfo{}
	for k, v := range count_map {
		ci := CountInfo{
			Ext:   k,
			Count: v,
		}
		countInfoObj = append(countInfoObj, ci)
	}
	fmt.Println(countInfoObj)
	json_bytes, _ = json.Marshal(countInfoObj)
	fmt.Println(string(json_bytes))
}
