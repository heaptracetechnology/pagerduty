package handler

import (
	"encoding/json"
	"fmt"
	"github.com/PagerDuty/go-pagerduty"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

func GetServicesList(w http.ResponseWriter, r *http.Request) {

	client := pagerduty.NewClient(os.Getenv("ACCESS_TOKEN"))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	listServiceOptions := new(pagerduty.ListServiceOptions)

	err = json.Unmarshal([]byte(body), listServiceOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	listServiceOptionsPtr := *listServiceOptions
	pagerServices, err := client.ListServices(listServiceOptionsPtr)
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}

	bytes, err := json.Marshal(pagerServices)
	writeJsonResponse(w, bytes)
}

func GetService(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id = vars["id"]

	client := pagerduty.NewClient(os.Getenv("ACCESS_TOKEN"))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	getServiceOptions := new(pagerduty.GetServiceOptions)

	err = json.Unmarshal([]byte(body), getServiceOptions)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	pagerServices, er := client.GetService(id, getServiceOptions)
	if er != nil {
		msg := Message{"false", er.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}

	bytes, err := json.Marshal(pagerServices)
	writeJsonResponse(w, bytes)
}

func CreateService(w http.ResponseWriter, r *http.Request) {

	client := pagerduty.NewClient(os.Getenv("ACCESS_TOKEN"))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	service := new(pagerduty.Service)
	err = json.Unmarshal([]byte(body), service)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Println(service)

	servicePtr := *service
	service, er := client.CreateService(servicePtr)
	if er != nil {
		msg := Message{"false", er.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}

	bytes, err := json.Marshal(service)
	w.WriteHeader(http.StatusCreated)
	writeJsonResponse(w, bytes)
}