package handler

import (
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/firebase"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"net/http"
)

func RegisterformHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Registerpage)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	err := Db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	newUser := &model.User{}
	err = json.NewDecoder(r.Body).Decode(newUser)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	email := "testuser@test.com"
	user := firebase.CreateFireBaseUser(newUser, email)
	cred, err := Db.Prepare("insert into user value (?,?,?,?)")
	defer cred.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}
	_, err = cred.Exec(newUser.Name, newUser.Id, newUser.Age, user.UID)
	if (err) != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusOK)
}
