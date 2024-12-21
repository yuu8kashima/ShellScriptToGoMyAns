package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"database/sql"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
)

// yamlのパースに構造体が欲しい
type Config struct {
	User, Password string
}

/**
 * 指定したYAMLファイルからPostgreSQLサーバーの接続情報を取得し、接続が成功するかどうかを確認する。
 * DBの設定関連については、localhostにpostgresユーザーでパスワードがhogehogeで建っていることにする
 */
func main() {
	data, err := os.ReadFile("./database.yml")
	if err != nil {
		log.Fatalf("ファイルの読み込みに失敗しました: %v", err)
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("YAMLの解析に失敗しました: %v", err)
	}
	fmt.Println(config)
	// yaml読み込みここまで

	// 以降はpostgresへの接続
	hostname := "127.0.0.1"
	port := 5432
	user := config.User
	password := config.Password
	dbname := "postgres"
	db_url := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", hostname, strconv.Itoa(port), user, password, dbname)
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました: %v", err)
	}
	defer db.Close()
	rows, err := db.Query("select * from user")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var u string
		rows.Scan(&u)
		fmt.Println(u)
	}
}
