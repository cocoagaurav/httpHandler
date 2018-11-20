package main

import (
	"database/sql"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

var (
	route         = chi.NewRouter()
	Conn          *amqp.Connection
	ElasticClient *elastic.Client
	Db            *sql.DB
)

func main() {
	FirebaseStartAuth()
	ElasticConn()
	Conn = RabbitConn()
	Db = Opendatabase()
	Migrate()

	route.Post("/", formHandler)
	route.Handle("/success", simpleMiddleware(http.HandlerFunc(AfterLoginHandler)))
	route.Post("/registerform", registerformHandler)
	route.Post("/register", registerHandler)
	route.Post("/post", Posthandler)
	route.Post("/login", loginhandler)
	route.Post("/logout", logoutHandler)
	route.Post("/fetchformhandler", Fetchformhandler)
	route.Post("/fetch", FetchHandler)
	route.Get("/quote/{date}", Getquote)

	http.Handle("/", route)

	log.Printf("Starting Sever :%v", 8080)
	http.ListenAndServe(":8080", nil)
}
