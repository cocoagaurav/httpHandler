package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
	"net/http"
)

func Getquote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ch, err := Conn.Channel()
	if err != nil {
		log.Printf("error while creating channle")
		return
	}
	defer ch.Close()
	c, err := r.Cookie("sessiontoken")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	uid := VerifyToken(c.Value)
	_, err = ch.QueueDeclare(
		"response",
		false,
		false,
		false,
		false,
		nil)

	_, err = ch.QueueDeclare(
		"quoteQ",
		false,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Printf("error while creating a Q:%v", err)
	}
	err = ch.Publish(
		"",
		"quoteQ",
		false,
		false,
		amqp.Publishing{
			ContentType:   "text/plain",
			ReplyTo:       "response",
			CorrelationId: uid.UID,
			Body:          []byte(params["date"]),
		})

	msg, err := ch.Consume(
		"response",
		"",
		false,
		false,
		false,
		false,
		nil)

	finish := make(chan bool)

	go func() {
		for mssg := range msg {
			fmt.Printf("message is:%v", string(mssg.Body))

			if uid.UID == mssg.CorrelationId {
				fmt.Fprint(w, string(mssg.Body))
				mssg.Ack(false)
				break

			}
		}
		finish <- true

	}()
	<-finish
}
