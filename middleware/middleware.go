package middleware

import (
	"github.com/cocoagaurav/httpHandler/firebase"
	"net/http"
)

func SimpleMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("sessiontoken")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		tk := firebase.VerifyToken(c.Value)
		if tk.UID != "" {
			handler.ServeHTTP(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}

	})
}
