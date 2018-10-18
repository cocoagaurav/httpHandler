package main

import (
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/database"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"net/http"
	"github.com/cocoagaurav/httpHandler/model"

)

func registerformHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Registerpage)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	err := database.Db.Ping()
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	newUser := &model.User{}
	err = json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	cred, err := database.Db.Prepare("insert into user value (?,?,?)")
	defer cred.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	_, err = cred.Exec(newUser.Name, newUser.Id, newUser.Age)
	if (err) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusOK)
}
