package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"time"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Formpage)
}
func loginhandler(w http.ResponseWriter, r *http.Request) {

	var (
		name string
		age  int
	)
	//ctx := context.Background()
	ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
	defer cancel()

	loginUser := &model.User{}
	err := json.NewDecoder(r.Body).Decode(loginUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("loginUser :[%+v]", loginUser)

	redirect := "/"
	cred := Db.QueryRow("select name,age from user where UID=?", loginUser.Id)

	err = cred.Scan(&name, &age)
	fmt.Println("database values are:", name, age)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	select {
	case <-ctx.Done():
		//	w.Write([]byte("request timeout"))
		log.Printf("contexterr is:%v/n", ctx.Err())
		http.Redirect(w, r, redirect, http.StatusRequestTimeout)

	}
	//fmt.Printf("context is:", <-ctx.Done())
	time.Sleep(4 * time.Second)
	status := http.StatusFound

	if loginUser.Name == name && loginUser.Age == age {
		redirect = "/success"
		token := uuid.NewV4()

		UserToken = token.String()
		UserCache[UserToken] = loginUser

		http.SetCookie(w, &http.Cookie{
			Name:    "sessiontoken",
			Value:   UserToken,
			Expires: time.Now().Add(24 * time.Hour),
		})

	} else {
		status = http.StatusNotFound
	}
	http.Redirect(w, r, redirect, status)
}
