package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/database"
	"github.com/cocoagaurav/httpHandler/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"log"
)

var (
	Ch   *amqp.Channel
	Mssg <-chan amqp.Delivery
)
var Db *sql.DB

func main() {

	RabbitConn()
	ConsumeMssg()
	Db = database.Opendatabase()
	listen := make(chan bool)
	go func() {
		for msg := range Mssg {
			post := &model.Post{}
			data := bytes.NewReader(msg.Body)
			err := json.NewDecoder(data).Decode(post)
			fmt.Printf("%v", post)

			if err != nil {
				log.Fatal(err)
				return
			}
			q, err := Db.Prepare("insert into post values(?,?,?)")
			defer q.Close()
			if err != nil {
				log.Fatal(err)
				return
			}
			_, err = q.Exec(post.Id, post.Title, post.Discription)
			if err != nil {
				log.Fatal(err.Error())
				return
			}

		}
	}()

	fmt.Println("listening....")
	<-listen
}
