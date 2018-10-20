package main

import (
	"github.com/streadway/amqp"
	"log"
)

var Conn *amqp.Connection

func RabbitConn() {
	var err error
	Conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	Ch, err = Conn.Channel()
}
