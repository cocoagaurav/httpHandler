package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/olivere/elastic"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

var (
	UserToken     string
	route         = mux.NewRouter()
	Conn          *amqp.Connection
	ElasticClient *elastic.Client
	Db            *sql.DB
)

func main() {
	ElasticConn()
	Conn = RabbitConn()
	Db = Opendatabase()
	Migrate()

	route.HandleFunc("/", formHandler)
	route.HandleFunc("/success", simpleMiddleware(AfterLoginHandler))
	route.HandleFunc("/registerform", registerformHandler)
	route.HandleFunc("/register", registerHandler)
	route.HandleFunc("/post", Posthandler)
	route.HandleFunc("/login", loginhandler)
	route.HandleFunc("/logout", logoutHandler)
	route.HandleFunc("/fetchformhandler", Fetchformhandler)
	route.HandleFunc("/fetch", FetchHandler).Methods("Post")
	route.HandleFunc("/quote/{date}", Getquote).Methods("Get")
	http.Handle("/", route)

	log.Printf("Starting Sever :%v", 8080)
	http.ListenAndServe(":8080", nil)
}
