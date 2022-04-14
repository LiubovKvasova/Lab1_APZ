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
}

func main() {
  http.ListenAndServe(":8795", nil)
} 
