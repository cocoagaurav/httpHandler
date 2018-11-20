package main

import "net/http"

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	clearSession(w)
	http.Redirect(w, r, "/", 302)
}
