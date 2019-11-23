package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"projecttravel"
	"time"
)

//Seconds is UNIX time at startup

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
	serverstart := int(time.Now().Unix()) //logging unix time of server start as
	//a global variable, used as parameter above
	projecttravel.Seconds = serverstart
	port := "7979"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	http.HandleFunc("/project/v1/", defaultHandler)
	http.HandleFunc("/project/v1/weather/", projecttravel.HandlerWeather)
	http.HandleFunc("/project/v1/status/", projecttravel.HandlerStatus)
	//http.HandleFunc("/project/v1/travel/", projecttravel.HandlerTravel)
	http.HandleFunc("/project/v1/places/", projecttravel.HandlerPOI)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
