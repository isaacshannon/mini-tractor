package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
)

var port = false
var starboard = false

func main() {
	http.HandleFunc("/direction", direction)
	http.HandleFunc("/control", control)

	http.ListenAndServe(":8080", nil)
}

type directionResponse struct {
	PortWheel bool `json:"portWheel"`
	StarboardWheel bool `json:"starboardWheel"`
}

func control(w http.ResponseWriter, req *http.Request){
	if err := req.ParseForm(); err != nil {
		port = false
		starboard = false
	}

	portValue := req.FormValue("port")
	if portValue != "" {
		fmt.Println("toggling port")
		port = !port
	}
	starboardValue := req.FormValue("starboard")
	if starboardValue != "" {
		fmt.Println("toggling starboard")
		starboard = !starboard
	}
	w.Header().Set("Cache-Control", "no-store")
	http.ServeFile(w, req, filepath.Join("templates", "control.html"))
}

func direction(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := directionResponse{
		PortWheel:      port,
		StarboardWheel: starboard,
	}
	json.NewEncoder(w).Encode(resp)
}


