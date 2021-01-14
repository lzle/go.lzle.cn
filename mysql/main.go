package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main()  {
	fmt.Println("123")

	db, err := sql.Open("mysql", "root:123456@/lzle")
	if err != nil {
		log.Print(err)
	}
	rows,err := db.Query("SELECT * FROM lzle.t where id = 1;")
	if err != nil {
		log.Print(err)
	}

	fmt.Println(rows)

}