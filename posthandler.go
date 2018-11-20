package main

import (
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/labstack/gommon/log"
	"github.com/streadway/amqp"
	"net/http"
	"strconv"
)

func AfterLoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, htmlPages.InternalPage)
}
func Posthandler(w http.ResponseWriter, r *http.Request) {
	err := Db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	c, err := r.Cookie("sessiontoken")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newpost := &model.Post{}
	err = json.NewDecoder(r.Body).Decode(newpost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Printf("\n\npost id:%d \n post title:%s \n post disc:%s \n token is:%s", newpost.Id, newpost.Title, newpost.Discription, c.Value)
	fmt.Println("label 2")

	tok := VerifyToken(c.Value)
	//id := &tok.UID

	fmt.Println("label 3")

	fmt.Println(tok.UID)
	fmt.Println("label 4")
	fmt.Println("label 46")
	//	fmt.Println(tok.UID)
	var uid string
	_ = Db.QueryRow("select id from user where auth_id =?", tok).Scan(&uid)
	newpost.Id, _ = strconv.Atoi(uid)

	jsonpost, err := json.Marshal(newpost)

	if err != nil {
		log.Fatal(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Printf("json data is:%s", string(jsonpost))

	Ch, err := Conn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}
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
	err = Ch.Publish(
		"",
		Q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jsonpost,
		})
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	http.Redirect(w, r, "/success", 302)
}
