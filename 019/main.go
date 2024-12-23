package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

/**
 * DDLとDMLの実行をおこなってください。
 */
func main() {
	db_url := "host=127.0.0.1 port=5432 user=postgres password=hogehoge dbname=postgres"
	db, err := sql.Open("postgres", db_url)
	if err != nil {
		log.Fatalf("データベースの接続に失敗しました: %v", err)
	}
	defer db.Close()

	// DDL
	db.Exec("drop table if exists employees")
	result, err := db.Exec("create table if not exists employees (id integer, name varchar(10))")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())

	result, err = db.Exec("insert into employees (id, name) values (1, 'Mr.A'), (2, 'Mr.B')")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.RowsAffected())

	var row string
	db.QueryRow("select name from employees where id = 1").Scan(&row)
	fmt.Println(row)
}
