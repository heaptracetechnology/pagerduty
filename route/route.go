package route

import (
	"github.com/gorilla/mux"
	"github.com/microservice/pagerduty/handler"
	"log"
	"net/http"
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
		"GetIncidents",
		"GET",
		"/listincidents",
		handler.GetIncidents,
	},

	Route{
		"GetIncidentsById",
		"GET",
		"/incidents/{id}",
		handler.GetIncidentsById,
	},

	Route{
		"CreateIncidents",
		"POST",
		"/incidents",
		handler.CreateIncidents,
	},

	Route{
		"GetIncidentsNotesById",
		"GET",
		"/incidents/{id}/notes",
		handler.GetIncidentsNotesById,
	},

	Route{
		"GetServicesList",
		"POST",
		"/services",
		handler.GetServicesList,
	},

	Route{
		"GetService",
		"POST",
		"/services/{id}",
		handler.GetService,
	},

	Route{
		"CreateService",
		"POST",
		"/services",
		handler.CreateService,
	}}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}