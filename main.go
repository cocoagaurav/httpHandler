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

	//	ElasticClient := ElasticConn()
	Conn := RabbitConn()
	DataBase := Opendatabase()
	firebase.FirebaseStartAuth()

	config := &model.Configs{
		Db:     DataBase,
		Rabbit: Conn,
	}

	MigrateUp(config.Db)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Setuproutes(config),
	}

	fmt.Printf("Starting Sever :%v", 8080)

	server.ListenAndServe()

}
