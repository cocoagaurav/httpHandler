package main

import (
	"github.com/gorilla/mux"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/rubenv/sql-migrate"
	"github.com/shaleapps/agnus-server/migrations"
	"log"
	"net/http"
)
import _ "github.com/go-sql-driver/mysql"


func init(){
	UserCache = make(map[string]*model.User)
}

var UserToken string
var route =mux.NewRouter()

func main() {

	Db=opendatabase()
	_,err:=migrate.Exec(Db,"mysql",migrations)
	if(err!=nil){
		log.Fatal(err.Error())
		return
	}

	route.HandleFunc("/",formHandler)
	route.HandleFunc("/success",simpleMiddleware(afterLoginHandler))
	route.HandleFunc("/registerform",registerformHandler)
	route.HandleFunc("/register",registerHandler)
	route.HandleFunc("/post",posthandler)
	route.HandleFunc("/login",loginhandler)
	route.HandleFunc("/logout",logoutHandler)
	route.HandleFunc("/fetchformhandler",fetchformhandler)
	route.HandleFunc("/fetch",fetchHandler).Methods("POST")
	http.Handle("/",route)

	log.Printf("Starting Sever :%v",8081)
	http.ListenAndServe(":8081",nil)

}


