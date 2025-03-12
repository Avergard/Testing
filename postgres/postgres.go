package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable port=5434"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)

	for i := range 30000 {
		_, err := db.Exec("insert into table1 (id, name) values($1, $2)", i, i)
		if err != nil {
			fmt.Println(err)
		}
		_, err = db.Exec("insert into table2 (id, name) values($1, $2)", i, i)
		if err != nil {
			fmt.Println(err)
		}
		_, err = db.Exec("insert into table3 (id, name) values($1, $2)", i, i)
		if err != nil {
			fmt.Println(err)
		}
		_, err = db.Exec("insert into table4 (id, name) values($1, $2)", i, i)
		if err != nil {
			fmt.Println(err)
		}
		_, err = db.Exec("insert into table5 (id, name) values($1, $2)", i, i)
		if err != nil {
			fmt.Println(err)
		}

	}
}
