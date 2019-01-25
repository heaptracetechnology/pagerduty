	package handler

	import (
		"github.com/gorilla/mux"
		"encoding/json"
		"io/ioutil"
		"github.com/PagerDuty/go-pagerduty"
		"net/http"
		"os"
	)

	func GetIncidents(w http.ResponseWriter, r *http.Request) {

		var access_token = os.Getenv("access_token")
		client := pagerduty.NewClient(access_token)


		vars := mux.Vars(r)
		options := pagerduty.ListIncidentsOptions{TimeZone:vars["timezone"],Since:vars["since"],Until:vars["until"],DateRange:vars["date_range"],IncidentKey:vars["incident_key"]}
		
		pagerincidents, err := client.ListIncidents(options)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		bytes, err := json.Marshal(pagerincidents)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		writeJsonResponse(w, bytes)
	}

	func GetIncidentsById(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		var id = vars["id"]

		var access_token = os.Getenv("access_token")
		client := pagerduty.NewClient(access_token)
		
		pagerincidents, err := client.GetIncident(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		bytes, err := json.Marshal(pagerincidents)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		writeJsonResponse(w, bytes)
	}

	func CreateIncidents(w http.ResponseWriter, r *http.Request) {

		from := r.FormValue("from")
		
		var access_token = os.Getenv("access_token")
		
		client := pagerduty.NewClient(access_token)
		
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
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return;
		}

		bytes, err := json.Marshal(pagerincidents)
		w.WriteHeader(http.StatusCreated)
		
		writeJsonResponse(w, bytes)
	}
	
	func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Write(bytes)
	}

