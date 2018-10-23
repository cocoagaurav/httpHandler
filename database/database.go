package database

import (
	"database/sql"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/labstack/gommon/log"
	"time"
)

var Db, DataBase *sql.DB

func init() {
	UserCache = make(map[string]*model.User)
}

var UserCache map[string]*model.User

func Opendatabase() *sql.DB {
	var err error
	DataBase, err = sql.Open("mysql", "root:password123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Printf("not able to connect to database")
		time.Sleep(5 * time.Second)
		Opendatabase()
	}
	log.Printf("database is connected.......")
	return DataBase

}

//DataBase, err = sql.Open("mysql", "root:password123@tcp(127.0.0.1:3306)/test")
