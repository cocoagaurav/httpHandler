package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testproc/models"
)

func registerformHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, registerpage)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	err := Db.Ping()
	if (err != nil) {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	newUser := &models.User{}
	err = json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	cred, err := Db.Prepare("insert into user value (?,?,?)")
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
