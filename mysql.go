package main

import (
	"database/sql"
 _ "github.com/go-sql-driver/mysql")

var DB *sql.DB

 func insert(matchs []Match){
	db, err := sql.Open("mysql", "root:q1w2e3r46@tcp(106.12.10.77:3306)/football")
	check(err)
	DB = db
}

func InserMatch(matchs []Match){

}

func check(err error) {
    if err != nil{
        fmt.Println(err)
        panic(err)
    }





