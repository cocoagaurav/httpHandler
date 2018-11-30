package main

import (
	"fmt"
	"github.com/cocoagaurav/httpHandler/firebase"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/cocoagaurav/httpHandler/router"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {
	firebase.FirebaseStartAuth()

	//	ElasticClient := ElasticConn()

	Conn := RabbitConn()

	DataBase := Opendatabase()

	config := &model.Configs{
		Db:     DataBase,
		Rabbit: Conn,
	}

	MigrateUp()
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Setuproutes(config),
	}

	fmt.Printf("Starting Sever :%v", 8080)

	server.ListenAndServe()

}
