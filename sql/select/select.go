package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type usuario struct {
	id   int
	nome string
}

func main() {
	db, err := sql.Open("mysql", "manager:1234@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var usuarios []usuario = []usuario{}
	rows, _ := db.Query("select id, nome from usuarios where id > ?", 0)
	defer rows.Close()

	for rows.Next() {
		var u usuario
		rows.Scan(&u.id, &u.nome)
		fmt.Println(u)
		usuarios = append(usuarios, u)
	}

	fmt.Println(usuarios)
}
