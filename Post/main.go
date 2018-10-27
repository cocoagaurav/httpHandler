package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/olivere/elastic"
	"github.com/streadway/amqp"
	"log"
)

var (
	Ch            *amqp.Channel
	Mssg          <-chan amqp.Delivery
	ElasticClient *elastic.Client
	Db            *sql.DB
	Bulk          *elastic.BulkService
)

func main() {
	ElasticOpen()
	RabbitConn()
	Db = Opendatabase()
	listen := make(chan bool)
	Ch, _ = Conn.Channel()
	Q, err := Ch.QueueDeclare(
		"PostQ",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	mssg, err := Ch.Consume(
		Q.Name,
		"",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	go func() {
		for msg := range mssg {
			go func(msg amqp.Delivery) {
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
				Datainsert(string(msg.Body))
				msg.Ack(false)
			}(msg)
		}
	}()

	fmt.Println("listening....")
	<-listen
}
