package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/cocoagaurav/httpHandler/firebase"
	"github.com/cocoagaurav/httpHandler/htmlPages"
	"github.com/cocoagaurav/httpHandler/model"
	"net/http"
	"time"
)

func FormHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPages.Formpage)
}
func Loginhandler(w http.ResponseWriter, r *http.Request) {

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
	userCred := firebase.GetUserCreds(authId)

	if loginUser.Name == userCred.DisplayName && loginUser.Age == age && authId == userCred.UID {
		token := firebase.GenerateToken(authId)
		http.SetCookie(w, &http.Cookie{
			Name:    "sessiontoken",
			Value:   token,
			Expires: time.Now().Add(24 * time.Hour),
		})
		json.NewEncoder(w).Encode(token)

	}
}
