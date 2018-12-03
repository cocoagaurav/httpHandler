package main

import (
	"database/sql"
	"fmt"
	"github.com/cocoagaurav/httpHandler/migration"
	"github.com/labstack/gommon/log"
	"github.com/rubenv/sql-migrate"
	"time"
)

func Opendatabase() *sql.DB {
	DataBase, err := sql.Open("mysql", "root:password123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("not able to connect to database")
		time.Sleep(5 * time.Second)
		Opendatabase()
	}
	fmt.Printf("Starting database... \n\n")

	return DataBase
}

func MigrateUp(dataBase *sql.DB) {

	migration1 := migration.Getmigration()
	_, err := migrate.Exec(dataBase, "mysql", migration1, migrate.Up)
	if err != nil {
		log.Printf("error is in migration:%v", err)
		time.Sleep(5 * time.Second)
		MigrateUp(dataBase)
	}
	fmt.Println("migration is ready")

}
