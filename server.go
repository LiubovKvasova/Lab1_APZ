package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func timeRequest(rw http.ResponseWriter, req *http.Request){}

func main() {
  http.ListenAndServe(":8795", nil)
} 
