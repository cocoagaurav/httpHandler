package main

import (
	"net/http"
)

func simpleMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		c, err := r.Cookie("sessiontoken")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token := c.Value
		_, ok := UserCache[token]
		if ok == true {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}

	}
}
