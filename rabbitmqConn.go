package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

var Connection *amqp.Connection

func RabbitConn() *amqp.Connection {
	var err error
	Connection, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("not able to connect to rabbitmq")
		time.Sleep(5 * time.Second)
		RabbitConn()
	}
	fmt.Printf("connected to rabbitmq...... \n\n")
	return Connection
}
