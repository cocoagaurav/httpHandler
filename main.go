package main

import (
	"database/sql"
	"github.com/cocoagaurav/httpHandler/migration"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/olivere/elastic"
	"github.com/rubenv/sql-migrate"
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
	var err error
	ElasticConn()
	Conn = RabbitConn()
	//Db = Opendatabase()
	Db, err = sql.Open("mysql", "root:password123@tcp(mysql:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	migration1 := migration.Getmigration()
	_, err = migrate.Exec(Db, "mysql", migration1, migrate.Up)
	if err != nil {
		log.Printf("error is in migration:%v", err)
		return
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
