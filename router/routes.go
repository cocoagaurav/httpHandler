package router

import (
	"github.com/cocoagaurav/httpHandler/handler"
	"github.com/go-chi/chi"
	"github.com/shaleapps/agnus-server/mware"
	"net/http"
)

func Setuproutes() {
	var route = chi.NewRouter()

	route.Post("/", handler.FormHandler)
	route.Post("/login", handler.Loginhandler)
	route.Mount("/", Authroutes())

}

func Authroutes() http.Handler {

	var route = chi.NewRouter()

	route.Use()
	route.Use(mware.database())

	route.Post("/registerform", handler.RegisterformHandler)
	route.Post("/register", handler.RegisterHandler)
	route.Post("/post", handler.Posthandler)
	route.Post("/logout", handler.LogoutHandler)
	route.Post("/fetchformhandler", handler.Fetchformhandler)
	route.Post("/fetch", handler.FetchHandler)
	route.Get("/quote/{date}", handler.Getquote)

	http.Handle("/", route)

}
