package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func timeRequest(rw http.ResponseWriter, req *http.Request){

	rw.Header().Set("Content-Type", "application/json")
	timeBroadcast := time.Now().Format(time.RFC3339)

	response := make(map[string]string)

    response["time"] = timeBroadcast

    error := json.NewEncoder(rw).Encode(response)
	if error != nil {
		log.Fatalf("Error occurred. Try again:) Error: %s", error)
	}
}

func main() {
  http.ListenAndServe(":8795", nil)
} 
