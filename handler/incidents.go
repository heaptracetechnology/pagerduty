package handler

import (
	"encoding/json"
	"github.com/PagerDuty/go-pagerduty"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
)

type Message struct {
    Success string `json:"success"`
    Result string `json:"result"`
}

func GetIncidents(w http.ResponseWriter, r *http.Request) {

	client := pagerduty.NewClient(os.Getenv("ACCESS_TOKEN"))
	
	vars := mux.Vars(r)
	options := pagerduty.ListIncidentsOptions{TimeZone: vars["timezone"], Since: vars["since"], Until: vars["until"], DateRange: vars["date_range"], IncidentKey: vars["incident_key"]}

	pagerincidents, err := client.ListIncidents(options)
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}

	fmt.Println(pagerincidents)
	bytes, err := json.Marshal(pagerincidents)
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
	}
	
	writeJsonResponse(w, bytes)
}

func GetIncidentsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id = vars["id"]

	client := pagerduty.NewClient(os.Getenv("ACCESS_TOKEN"))
	
	pagerincidents, err := client.GetIncident(id)
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}
	
	bytes, err := json.Marshal(pagerincidents)
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}
 
	writeJsonResponse(w, bytes)
}

func CreateIncidents(w http.ResponseWriter, r *http.Request) {

	from := r.FormValue("from")

	client := pagerduty.NewClient(os.Getenv("ACCESS_TOKEN"))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	createIncident := new(pagerduty.CreateIncident)
	err = json.Unmarshal([]byte(body), createIncident)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	pagerincidents, err := client.CreateIncident(from, createIncident)
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}
	bytes, err := json.Marshal(pagerincidents)
	w.WriteHeader(http.StatusCreated)

	writeJsonResponse(w, bytes)
}

func GetIncidentsNotesById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var id = vars["id"]
	
	client := pagerduty.NewClient(os.Getenv("ACCESS_TOKEN"))

	incidentNotes, err := client.ListIncidentNotes(id)
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}

	bytes, err := json.Marshal(incidentNotes)
	if err != nil {
		msg := Message{"false", err.Error()}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}

	if bytes != nil {
		msg := Message{"false", "Incident note are not found (404)"}
		msgbytes, err := json.Marshal(msg)
		fmt.Println(err);
		writeJsonResponse(w, msgbytes)
		return
	}

	writeJsonResponse(w, bytes)
}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}