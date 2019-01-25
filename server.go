package main

import (
	"log"
	"net/http"
	"github.com/microservice/pagerduty/route"
)

func main() {
	router := route.NewRouter() 
	log.Fatal(http.ListenAndServe(":5000", router))
}