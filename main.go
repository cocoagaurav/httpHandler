package main

import (
	"fmt"
	"github.com/cocoagaurav/httpHandler/firebase"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/cocoagaurav/httpHandler/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"log"
	"net/http"
)

func main() {

	//	ElasticClient := ElasticConn()
	var env model.Env
	err := envconfig.Process("myapi", &env)
	if err != nil {
		log.Printf("error while getting env variables:%s", err.Error())
		return
	}
	Conn := RabbitConn(env)
	DataBase := Opendatabase(env)
	firebase.FirebaseStartAuth(env)

	config := &model.Configs{
		Db:     DataBase,
		Rabbit: Conn,
		Env:    env,
	}

	MigrateUp(config.Db)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router.Setuproutes(config),
	}

	fmt.Printf("Starting Sever :%v", 8080)

	server.ListenAndServe()

}
