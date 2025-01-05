package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// 2つのディレクトリを同期するプログラムを作成してください。
// ソースディレクトリにあるファイルをターゲットディレクトリにコピーし、ターゲットディレクトリにのみ存在するファイルを削除します。
// メモ: 要するにrsyncの--delete付き
func main() {
	// 実用するならos.Args等で指定できるようにすべきだが、誤爆が怖いので対象は固定にする
	src_dir := "sync_src"
	dst_dir := "sync_dst"

	// srcを見てdstにファイルをコピー
	filepath.Walk(src_dir, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("walk: %s\n", path)
		in_dir_path := getInDirPath(path)                                 // path - src_dirのパス
		mirror_path := dst_dir + string(filepath.Separator) + in_dir_path //  ミラーリングでコピーする先のパス。
		mirror_path_exists := isExists(mirror_path)                       // ミラーリング先の確認

		// ディレクトリ階層を潜る際には、まずディレクトリが対象になり、次にそのディレクトリの下のファイルが来ることを仮定している
		if info.IsDir() {
			if !mirror_path_exists {
				fmt.Printf("%sにディレクトリ作成\n", mirror_path)
				os.Mkdir(mirror_path, info.Mode().Perm())
			} else {
				if !isDir(mirror_path) {
					// エラーにしてみたが、ファイル消去して処理続行する方が要求には合ってるかも
					log.Fatalf("エラー: %s にファイルがあるためディレクトリを作成出来ない", mirror_path)
				}
			}
		} else {
			fmt.Printf("%sから%sにファイルコピー\n", path, mirror_path)
			copyFile(path, mirror_path)
		}

		return nil
	})

	// dstを見て、srcに無いファイルを消す
	filepath.Walk(dst_dir, func(path string, info os.FileInfo, err error) error {
		fmt.Printf("delete walk: %s\n", path)
		in_dir_path := getInDirPath(path)                                 // path - dst_dirのパス
		mirror_path := src_dir + string(filepath.Separator) + in_dir_path // 被ミラーリング対象のパス。
		mirror_path_exists := isExists(mirror_path)                       // ミラーリング元の有無を確認

		if !mirror_path_exists {
			fmt.Printf("%sはsrcに無いため消去\n", path)
			os.Remove(path)
		}

		return nil
	})
}

// filewalkのpathが指定ディレクトリから始まるので、一段階ディレクトリを潜ったパスを返す
func getInDirPath(path string) string {
	sep := string(filepath.Separator)
	return strings.Join(strings.Split(path, sep)[1:], sep)
}

// ファイルとディレクトリの存在確認
func isExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

// ディレクトリ存在確認
func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// ファイルコピー
func copyFile(src, dst string) {
	c, err := os.Create(dst)
	if err != nil {
		log.Fatal(err)
	}

	r, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	_, err = io.Copy(c, r)
	if err != nil {
		log.Fatal(err)
	}
}
