package database

import (
	"database/sql"
	"github.com/cocoagaurav/httpHandler/model"
)
var Db,DataBase *sql.DB
var UserCache map[string]*model.User

func Opendatabase() *sql.DB{
	var err error
	DataBase, err = sql.Open("mysql", "root:password123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return DataBase

}
