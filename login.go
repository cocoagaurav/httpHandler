package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"net/http"
	"time"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Formpage)
}
func loginhandler(w http.ResponseWriter, r *http.Request) {

	var (
		name   string
		age    int
		authId string
	)
	loginUser := &model.User{}
	err := json.NewDecoder(r.Body).Decode(loginUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("loginUser :[%+v]", loginUser)

	//redirect := "/"
	cred := Db.QueryRow("select name,age,auth_id from user where UID=?", loginUser.Id)

	err = cred.Scan(&name, &age, &authId)
	fmt.Println("database values are:", name, age, authId)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	fmt.Println("label 2")

	status := http.StatusFound

	if loginUser.Name == name && loginUser.Age == age {
		fmt.Println("label 3")
		//	redirect = "/success"
		token := "token123"
		fmt.Println("label 4")
		http.SetCookie(w, &http.Cookie{
			Name:    "sessiontoken",
			Value:   token,
			Expires: time.Now().Add(24 * time.Hour),
		})
		fmt.Println("label 5")

	} else {
		status = http.StatusNotFound
	}
	fmt.Println("label 6")
	//http.Redirect(w, r, redirect, status)
	w.WriteHeader(status)
}
