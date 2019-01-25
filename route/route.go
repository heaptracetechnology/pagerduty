package route

import (
    "log"
    "net/http"
	"github.com/gorilla/mux"
	"github.com/microservice/pagerduty/handler"
)

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes {
    Route {
        "GetIncidents",
        "GET",
        "/listincidents",
        handler.GetIncidents,
    },
    
    Route {
        "GetIncidentsById",
        "GET",
        "/incidents/{id}",
        handler.GetIncidentsById,
	},
    
    Route {
        "CreateIncidents",
        "POST",
        "/incidents",
        handler.CreateIncidents,
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