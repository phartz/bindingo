package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

type Result struct {
	TimeStamp string `json:"timestamp,omitempty"`
	Value     string `json:"value,omitempty"`
	Message   string `json:"string_value",omitempty"`
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	log.Printf("GetStatus DataService [%s] ID [%s] Credentials [%s].\n", params["dataService"], params["id"], params["credentials"])

	dataService, err := GetDataService(params)
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	credentials, err := GetCredentials(params["id"], params["credentials"])
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	result, err := dataService.GetStatus(credentials)
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	if result == 0 {
		WriteResponse(w, "0", "Database seems to work")
	} else {
		WriteResponse(w, "1", "Database doesn't work")
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	log.Printf("Insert DataService [%s] ID [%s] Credentials [%s].\n", params["dataService"], params["id"], params["credentials"])

	dataService, err := GetDataService(params)
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	credentials, err := GetCredentials(params["id"], params["credentials"])
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	err = dataService.Insert(credentials, params["name"])
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	log.Printf("Delete DataService [%s] ID [%s] Credentials [%s].\n", params["dataService"], params["id"], params["credentials"])

	dataService, err := GetDataService(params)
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	credentials, err := GetCredentials(params["id"], params["credentials"])
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	err = dataService.Delete(credentials, params["name"])
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	WriteResponse(w, "0", fmt.Sprintf("Entry deleted: id: #{id}, name: #{name}", params["id"], params["name"]))
}

func Exists(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	log.Printf("Exists DataService [%s] ID [%s] Credentials [%s].\n", params["dataService"], params["id"], params["credentials"])

	dataService, err := GetDataService(params)
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	credentials, err := GetCredentials(params["id"], params["credentials"])
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	exists, err := dataService.Exists(credentials, params["name"])
	if err != nil {
		WriteResponse(w, "2", err.Error())
		return
	}

	if exists {
		WriteResponse(w, "0", fmt.Sprintf("Entry exists: id: #{id}, name: #{name}", params["id"], params["name"]))
	} else {
		WriteResponse(w, "1", fmt.Sprintf("Entry not found: id: #{id}, name: #{name}", params["id"], params["name"]))
	}

}

func WriteResponse(w http.ResponseWriter, value string, message string) {
	t := time.Now()

	result := Result{
		TimeStamp: t.Format("20060102150405"),
		Value:     value,
		Message:   message}

	json.NewEncoder(w).Encode(result)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/service_instance/{dataService}/{id}/status", GetStatus).Methods("GET")
	router.HandleFunc("/service_instance/{dataService}/{id}/insert", GetStatus).Methods("PUT")
	router.HandleFunc("/service_instance/{dataService}/{id}/delete", GetStatus).Methods("DELETE")
	router.HandleFunc("/service_instance/{dataService}/{id}/exists", GetStatus).Methods("GET")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), router))
}
