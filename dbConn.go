package main

import (
	"database/sql"
	"github.com/cocoagaurav/httpHandler/migration"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/labstack/gommon/log"
	"github.com/rubenv/sql-migrate"
	"time"
)

var DataBase *sql.DB

func init() {
	UserCache = make(map[string]*model.User)
}

var UserCache map[string]*model.User

func Opendatabase() *sql.DB {
	var err error
	DataBase, err = sql.Open("mysql", "root:password123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("not able to connect to database")
		time.Sleep(5 * time.Second)
		Opendatabase()
	}
	log.Printf("database is connected in func.......")

	return DataBase

}

func Migrate() {

	migration1 := migration.Getmigration()
	_, err := migrate.Exec(DataBase, "mysql", migration1, migrate.Up)
	if err != nil {
		log.Printf("error is in migration:%v", err)
		return
	}
}
