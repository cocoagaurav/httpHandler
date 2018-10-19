package main

import (
	"github.com/cocoagaurav/httpHandler/database"
	"net/http"
)

func simpleMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		//tempUser := &models.User{}
		c, err := r.Cookie("sessiontoken")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token := c.Value
		_, ok := database.UserCache[token]
		if ok == true {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}

	}
}
