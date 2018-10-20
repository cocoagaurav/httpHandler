package main

import (
	"github.com/cocoagaurav/httpHandler/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/streadway/amqp"
	"log"
	"net/http"
)

//func init() {
//	UserCache = make(map[string]*model.User)
//}

var UserToken string
var route = mux.NewRouter()
var Conn *amqp.Connection

func main() {
	var err error
	database.Db = database.Opendatabase()
	Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	//migration1:=migration.Getmigration()
	//_,err:=migrate.Exec(database.Db,"mysql",migration1,migrate.Up)
	//if(err!=nil){
	//	log.Fatal(err.Error())
	//	return
	//}
	//Db , err  := openDatabase("root:password123@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	//defer Db.Close()
	//if err!=nil{
	//	log.Fatalf("Failed to initialize DB")
	//}
	route.HandleFunc("/", formHandler)
	route.HandleFunc("/success", simpleMiddleware(AfterLoginHandler))
	route.HandleFunc("/registerform", registerformHandler)
	route.HandleFunc("/register", registerHandler)
	route.HandleFunc("/post", Posthandler)
	route.HandleFunc("/login", loginhandler)
	route.HandleFunc("/logout", logoutHandler)
	route.HandleFunc("/fetchformhandler", Fetchformhandler)
	route.HandleFunc("/fetch", FetchHandler).Methods("POST")
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
