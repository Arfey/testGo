package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"./handlers"
	"./consts"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route


var routes = Routes{
	Route{
		"CreateUser",
		"POST",
		"/players/",
		handlers.CreateUser,
	},
	Route{
		"DeleteUsers",
		"DELETE",
		"/players/",
		handlers.DeleteUsers,
	},
	Route{
		"UpdateUser",
		"PUT",
		"/players/{id:[0-9]+}/",
		handlers.UpdateUser,
	},
	Route{
		"GetUser",
		"GET",
		"/players/{id:[0-9]+}/",
		handlers.GetUser,
	},
	Route{
		"DeleteUser",
		"DELETE",
		"/players/{id:[0-9]+}/",
		handlers.DeleteUser,
	},
	Route{
		"GetUser",
		"PUT",
		"/tournaments/{id:[0-9]+}/",
		handlers.SetWinner,
	},
	Route{
		"AddMember",
		"POST",
		"/tournaments/{id:[0-9]+}/",
		handlers.AddMember,
	},
	Route{
		"CreateTournament",
		"POST",
		"/tournaments/",
		handlers.CreateTournament,
	},
	Route{
		"Reset",
		"POST",
		"/reset/",
		handlers.Reset,
	},
}


func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
		Path("/" + consts.Version + route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
	}

	return router
}
