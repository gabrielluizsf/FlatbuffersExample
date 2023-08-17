package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Greeting struct {
	Message string `json:"message"`
}

func startServer(port string) error{
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/greet", greetHandler)

	log.Printf("Server started on %s",port)
	return http.ListenAndServe(port, nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	response := Greeting{Message: "Hello, World!"}
	sendJSONResponse(w, response, http.StatusOK)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		sendErrorResponse(w, "Missing 'name' parameter", http.StatusBadRequest)
		return
	}
	request_name, _ := createPerson(name, 21)

	response := Greeting{Message: "Hello, " + request_name + "!"}
	sendJSONResponse(w, response, http.StatusOK)
}

func sendJSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]string{"error": message}
	sendJSONResponse(w, response, statusCode)
}
