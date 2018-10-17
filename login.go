package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/cocoagaurav/httpHandler/model"
	"net/http"
	"time"
)

var UserCache map[string]*model.User

func formHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, formpage)
}
func loginhandler(w http.ResponseWriter, r *http.Request) {

	var(
		name string
		age int
	)

	loginUser := &model.User{}
	err:= json.NewDecoder(r.Body).Decode(loginUser)
	if(err!=nil){
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Printf("loginUser :[%+v]",loginUser)

	redirect := "/"
	cred := Db.QueryRow("select name,age from user where UID=?", loginUser.Id)

	err = cred.Scan(&name, &age)
	fmt.Println("database values are:",name,age)
	if (err != nil){
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNonAuthoritativeInfo)
			return
		}else{
			w.WriteHeader(http.StatusInternalServerError)
			return
	}
	}
	status:=http.StatusFound
	if loginUser.Name == name && loginUser.Age == age {

		redirect = "/success"
		token:= uuid.NewV4()
		//if err!=nil{
		//	w.WriteHeader(http.StatusInternalServerError)
		//	return
		//}
		UserToken=token.String()
		UserCache[UserToken] = loginUser

		http.SetCookie(w,&http.Cookie{
			Name:  "sessiontoken",
			Value: UserToken,
			Expires:time.Now().Add(24*time.Hour),
		})

	}else{
			status=http.StatusNotFound
		}
	http.Redirect(w, r, redirect, status)
}



















