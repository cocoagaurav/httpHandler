package main

import (
	"github.com/cocoagaurav/httpHandler/database"
	"github.com/cocoagaurav/httpHandler/migration"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/olivere/elastic"
	"github.com/rubenv/sql-migrate"
	"github.com/streadway/amqp"
	"log"
	"net/http"
	"time"
)

var (
	UserToken     string
	route         = mux.NewRouter()
	Conn          *amqp.Connection
	ElasticClient *elastic.Client
)

func main() {
	var err error
	ElasticConn()
here:
	Conn, err = amqp.Dial("amqp://guest:guest@rabbitmq-server:5672/")
	if err != nil {
		log.Printf("not able to connect to rabbitmq")
		time.Sleep(5 * time.Second)
		goto here
	}
	log.Printf("rabbitmq is connected/.................")
	database.Db = database.Opendatabase()
	log.Printf("database is connected/.................")
	migration1 := migration.Getmigration()
	_, err = migrate.Exec(database.Db, "mysql", migration1, migrate.Up)
	if err != nil {
		log.Printf("error is in migration")
	}
	route.HandleFunc("/", formHandler)
	route.HandleFunc("/success", simpleMiddleware(AfterLoginHandler))
	route.HandleFunc("/registerform", registerformHandler)
	route.HandleFunc("/register", registerHandler)
	route.HandleFunc("/post", Posthandler)
	route.HandleFunc("/login", loginhandler)
	route.HandleFunc("/logout", logoutHandler)
	route.HandleFunc("/fetchformhandler", Fetchformhandler)
	route.HandleFunc("/fetch", FetchHandler).Methods("Post")
	http.Handle("/", route)

	log.Printf("Starting Sever :%v", 8080)
	http.ListenAndServe(":8080", nil)

}

//func openDatabase(url string) (*gorm.DB ,error){
//
//	db, err := gorm.Open("mysql", url)
//	if err!=nil{
//		return nil,err
//	}
//
//	err = db.AutoMigrate(
//		&model.User{},
//		&model.Post{},
//		).Error
//	if err!=nil{
//		fmt.Errorf("Failed to AutoMigrate as err = [%v]",err)
//	}
//	return db,nil
//}
//Db , err  := openDatabase("root:password123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
//defer Db.Close()
//if err!=nil{
//	log.Fatalf("Failed to initialize DB")
//}
