package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type User struct {
	id   uint
	name string
}

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

	user1 := User{}

	id2 := user1.id
	name2 := user1.name

	fmt.Scan(&id2)
	fmt.Scan(&name2)

	id := id2
	name := name2

	sqlStatement := `INSERT INTO user1 (id, name) VALUES ($1, $2)`
	_, err = db.Exec(sqlStatement, id, name)
	if err != nil {
		panic(err)

	}
	fmt.Println("Данные успешно добавлены!")

}
