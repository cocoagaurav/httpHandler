package main

import (
	"database/sql"
)
var Db,DataBase *sql.DB

func opendatabase() *sql.DB{
	var err error
	DataBase, err = sql.Open("mysql", "root:password123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	return DataBase

}
