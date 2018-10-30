package main

import (
	"context"
	"fmt"
	"net/http"
)

func simpleMiddleware(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		fmt.Printf("request context value is :%v", r.Context().Value("UserId"))
		c, err := r.Cookie("sessiontoken")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		token := c.Value
		val, ok := UserCache[token]
		ctx = context.WithValue(ctx, "UserId", val.Id)

		if ok == true {
			next(w, r.WithContext(ctx))
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}

	}
}
