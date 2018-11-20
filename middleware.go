package main

import (
	"net/http"
)

func simpleMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("sessiontoken")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tk := VerifyToken(c.Value)
		if tk.UID != "" {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	})
}
