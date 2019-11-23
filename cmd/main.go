package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"projecttravel"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Handler not implemented, bad method", http.StatusNotImplemented)
		return
	}
	_, err := fmt.Fprintf(w, "IMT2681 Project - Travel")
	if err != nil {
		log.Printf("Error, bad endpoint")
	}
}

func main() {
	port := "7979"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	http.HandleFunc("/project/v1/", defaultHandler)
	http.HandleFunc("/project/v1/weather/", projecttravel.HandlerWeather)
	//http.HandleFunc("/project/v1/status/", projecttravel.HandlerStatus)
	//http.HandleFunc("/project/v1/travel/", projecttravel.HandlerTravel)
	http.HandleFunc("/project/v1/places/", projecttravel.handlerPOI)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
