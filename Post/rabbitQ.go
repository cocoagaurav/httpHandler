package Post

import (
	"github.com/streadway/amqp"
	"log"
)

var (
	Ch   *amqp.Channel
	Q    amqp.Queue
	Mssg <-chan amqp.Delivery
)

func MakeRabbitQ() {
	var err error
	Ch, err = Conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
	Q, err = Ch.QueueDeclare(
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
}
