package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func main() {
  http.ListenAndServe(":80", nil)
} 