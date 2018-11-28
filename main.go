package main

import (
	"database/sql"
	"fmt"
	"github.com/cocoagaurav/httpHandler/firebase"
	"github.com/cocoagaurav/httpHandler/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic"
	"github.com/streadway/amqp"
	"net/http"
)

var (
	Conn          *amqp.Connection
	ElasticClient *elastic.Client
	Db            *sql.DB
)

func main() {
	firebase.FirebaseStartAuth()

	ElasticConn()

	Conn = RabbitConn()

	Db = Opendatabase()

	Migrate()

	router.Setuproutes()

	fmt.Printf("Starting Sever :%v", 8080)

	http.ListenAndServe(":8080", nil)
}
