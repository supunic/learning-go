package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer func(DbConnection *sql.DB) {
		if err := DbConnection.Close(); err != nil {
			log.Fatalln(err)
		}
	}(DbConnection)

	cmd := `CREATE TABLE IF NOT EXISTS person(
    			name STRING,
    			age  INT)`
	_, err := DbConnection.Exec(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	// インサート
	//cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	//_, err = DbConnection.Exec(cmd, "Mike", 24)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// アップデート
	//cmd = "UPDATE person SET age = ? WHERE name = ?"
	//_, err = DbConnection.Exec(cmd, 25, "Mike")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// マルチセレクト
	//cmd = "SELECT * FROM person"
	//rows, _ := DbConnection.Query(cmd)
	//defer func(rows *sql.Rows) {
	//	if err := rows.Close(); err != nil {
	//		log.Fatalln(err)
	//	}
	//}(rows)
	//var pp []Person
	//for rows.Next() {
	//	var p Person
	//	if err := rows.Scan(&p.Name, &p.Age); err != nil {
	//		log.Fatalln(err)
	//	}
	//	pp = append(pp, p)
	//}
	//if err = rows.Err(); err != nil {
	//	log.Fatalln(err)
	//}
	//for _, p := range pp {
	//	fmt.Println(p.Name, p.Age)
	//}

	// シングルセレクト
	//cmd = "SELECT * FROM person where age = ?"
	//row := DbConnection.QueryRow(cmd, 2500)
	//var p Person
	//if err = row.Scan(&p.Name, &p.Age); err != nil {
	//	if err == sql.ErrNoRows {
	//		log.Println("No row")
	//	} else {
	//		log.Fatalln(err)
	//	}
	//}
	//fmt.Println(p.Name, p.Age)

	// デリート
	//cmd = "DELETE FROM person WHERE name = ?"
	//_, err = DbConnection.Exec(cmd, "Mike")
	//if err != nil {
	//	log.Fatalln(err)
	//}

	tableName := "person"

	// sqlインジェクションに注意
	cmd = fmt.Sprintf("SELECT * FROM %s", tableName)
	rows, _ := DbConnection.Query(cmd)
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			log.Fatalln(err)
		}
	}(rows)
	var pp []Person
	for rows.Next() {
		var p Person
		if err := rows.Scan(&p.Name, &p.Age); err != nil {
			log.Fatalln(err)
		}
		pp = append(pp, p)
	}
	if err = rows.Err(); err != nil {
		log.Fatalln(err)
	}
	for _, p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}
