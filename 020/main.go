package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

/**
 * 全ての利用可能な書籍をタイトルと著者名で表示するSQL文を作成してください。
 * 指定されたユーザーが借りた書籍の一覧を表示するSQL文を作成してください。
 * 期限切れの貸出書籍をユーザー名とタイトルで表示するSQL文を作成してください。
 *
 * 自分のメモ:
 *  * 本来は問18からの続いている問題のようだ
 *  * 毎回ポスグレ立てるのも面倒になってきたし、勉強ついでにsqlite3を使う
 *  * sqlite testdb.sqlite3 < init.sql を実行後にこのプログラムを動かしていると考えること
 */
func main() {
	db, err := sql.Open("sqlite3", "./testdb.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 全ての利用可能な書籍をタイトルと著者名で表示するSQL文を作成してください。
	sqlStmt := `
		select b.title, a.name from books b join authors a on b.author_id = a.id
	`
	rows, err := db.Query(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var title, name string
		err = rows.Scan(&title, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("タイトル:%s, 著者: %s\n", title, name)
	}

	// 指定されたユーザーが借りた書籍の一覧を表示するSQL文を作成してください。
	sqlStmt = `
		select b.title, u.name from loans l join books b on l.book_id = b.id join users u on l.user_id = u.id where l.user_id = ?
	`
	sqlPre, err := db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlPre.Close()

	target_user_id := 1 // ユーザーID部分は本来はArgs等で入力する
	rows, err = sqlPre.Query(target_user_id)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var title, username string
		err = rows.Scan(&title, &username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("タイトル:%s, ユーザー: %s\n", title, username)
	}

	// 期限切れの貸出書籍をユーザー名とタイトルで表示するSQL文を作成してください。
	now := "2022-01-16 00:00:00"
	sqlStmt = `
		select b.title, u.name 
		from loans l join books b on l.book_id = b.id join users u on l.user_id = u.id 
		where l.due_date < ?
	`
	sqlPre, err = db.Prepare(sqlStmt)
	if err != nil {
		log.Fatal(err)
	}
	defer sqlPre.Close()
	rows, err = sqlPre.Query(now)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var title, username string
		err = rows.Scan(&title, &username)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("期限切れ: タイトル:%s, ユーザー: %s\n", title, username)
	}
}
