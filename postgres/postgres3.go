package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type triangle struct {
	fkatet float64
	skatet float64
	plosh  float64
	per    float64
	height float64
}

func perim(t *triangle) float64 {
	t.per = t.skatet / 2 * t.height
	x := t.per
	return x
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

	t := triangle{
		fkatet: 5,
		skatet: 3,
		height: 4,
	}

	t.per = perim(&t)

	t.plosh = 0.5 * t.fkatet * t.skatet

	fkatet := t.fkatet
	skatet := t.skatet
	plosh := t.plosh
	per := t.per
	height := t.height
	sqlStatement := `INSERT INTO triangle (fkatet, skatet, plosh, per, height ) VALUES ($1, $2,$3,$4,$5)`
	_, err = db.Exec(sqlStatement, fkatet, skatet, plosh, per, height)
	if err != nil {
		panic(err)

	}
	fmt.Println("Данные успешно добавлены!")

}
