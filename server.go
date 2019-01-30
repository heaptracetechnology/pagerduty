package main

import (
	"github.com/microservice/pagerduty/route"
	"log"
	"net/http"
)

func main() {
	router := route.NewRouter()
	log.Fatal(http.ListenAndServe(":5000", router))
}