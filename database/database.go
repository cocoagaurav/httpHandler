package database

import (
	"database/sql"
	"github.com/cocoagaurav/httpHandler/model"
)

var Db, DataBase *sql.DB

func init() {
	UserCache = make(map[string]*model.User)
}

var UserCache map[string]*model.User

func Opendatabase() *sql.DB {
	var err error
	//DataBase, err = sql.Open("mysql", "root:password123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	DataBase, err = sql.Open("mysql", "root:password123@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	return DataBase

}
