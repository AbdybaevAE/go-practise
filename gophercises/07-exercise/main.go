package main

import (
	"database/sql"
	"fmt"
	"log"
	"unicode"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "postgres"
)
type Phone struct {
	Id int `db:"id"`
	Value string `db:"content"`
	FormattedPhone sql.NullString `db:"formatted"`
}
func main() {
	//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	//	"password=%s dbname=%s sslmode=disable",
	//	host, port, user, password, dbname)
	//db, err := sql.Open("postgres", psqlInfo)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close()
	//
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}
	//
	//fmt.Println("Successfully connected!")
	db, err := sqlx.Connect("postgres", "user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	allPhones := make([]Phone,0)
	sel1 := db.Select(&allPhones, "select * from phones")
	fmt.Println(sel1,allPhones)
	fmt.Println(len(allPhones))
	for _, phone := range allPhones {
		formatted := format(phone.Value)
		var count int
		db.Get(&count, "select count(*) from phones where phones.formatted=$1",formatted)
		if count == 0 {
			_, err := db.Exec(`update phones set formatted=$1 where id=$2`,formatted,phone.Id)
			if err != nil {
				panic(err)
			}
		} else {
			_, err := db.Exec(`delete from phones where id=$1`,phone.Id)
			if err != nil {
				panic(err)
			}
		}
	}

}

func format(num string) string  {
	str := []rune(num)
	out := make([]rune, 0)
	for _, v := range str {
		if unicode.IsDigit(v) {
			out = append(out, v)
		}
	}
	return string(out)

}

