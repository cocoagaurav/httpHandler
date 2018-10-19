package Post

import (
	"bytes"
	"github.com/streadway/amqp"
	"io"
	"log"
)

func Publish(post []byte) {
	err := Ch.Publish(
		"",
		Q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        post,
		})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
}

func ConsumeMssg() {
	var err error
	//Mssg = make(<-chan amqp.Delivery)
	Mssg, err = Ch.Consume(
		Q.Name,
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
}

func StreamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}
