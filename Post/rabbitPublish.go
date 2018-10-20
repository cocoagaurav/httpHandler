package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
)

func ConsumeMssg() {
	var err error
	//Mssg = make(<-chan amqp.Delivery)
	Mssg, err = Ch.Consume(
		"PostQ",
		"",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("consuming....")
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
