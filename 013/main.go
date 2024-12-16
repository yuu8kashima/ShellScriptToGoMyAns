package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"
)

/*
 * 指定した複数のURLからデータをダウンロードする。なお、並列処理を利用してください。
 */
func main() {
	var wg sync.WaitGroup
	download_folder_path := os.Args[1]
	args_urls := os.Args[2:]

	for i, v := range args_urls {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			downloaded_path := download_folder_path + "\\" + strconv.Itoa(i+1) + ".html"

			fmt.Printf("Download %s to %s\n", url, downloaded_path)
			err := downloadFile(url, downloaded_path)
			if err != nil {
				fmt.Println(err)
			}
		}(i, v)
	}
	wg.Wait()
}

/**
 * ダウンロード関数。結局ChatGPTの例通りになってしまった
 */
func downloadFile(url, downloadedPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}

	outFile, _ := os.Create(downloadedPath)
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	return nil
}
