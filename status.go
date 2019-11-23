package projecttravel

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var Seconds int //global value to store starttime of the server in UNIX time

//diagnosticsStruct
type diagStruct struct {
	MapsAPI        int
	OpenWeatherAPI int
	Database       int
	Version        string
	Uptime         int
}

//three constants, test-URLS and version (to change version need to go into code, so i figured
//it would be good practise to keep it as a constant)
const exampleURLGMapsAPI = "https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=60.7917739,10.6819978&radius=1500&key=AIzaSyBw7w2Yy6edsg28xfUcIwQ-yPHcY3fPqx0&limt=4"
const versjon = "v1"
const exampleURLOpenWeatherAPI = "http://api.openweathermap.org/data/2.5/weather?APPID=bf6b17711f24e7dd4c19aaa3abb904c5&lat=10&lon=10"

//HandlerStatus is the handler from main.go
func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get(exampleURLGMapsAPI) //checking response/error from first URL
	ds := &diagStruct{}
	if err != nil {
		fmt.Println("Something wrong with Get exampleURLGMapsAPI")
		fmt.Fprintf(w, "520"+http.StatusText(520)) //520 Server error code
		return
	}
	statusCode1 := resp.StatusCode
	defer resp.Body.Close() //closing response.Body

	resp, err = http.Get(exampleURLOpenWeatherAPI) //changing test-URL to Database
	if err != nil {                                //checking if error or new URL
		fmt.Println("Something wrong with Get exampleURLOpenWatherAPI")
		fmt.Fprintf(w, "520"+http.StatusText(520))
		return
	}
	statusCode2 := resp.StatusCode
	defer resp.Body.Close()
	revisedSeconds := int(time.Now().Unix()) - Seconds //getting uptime from unix-parameter
	ds.MapsAPI = statusCode1                           //setting ds (diagnostics) struct to correct values manually
	ds.OpenWeatherAPI = statusCode2
	ds.Database = statusCode2
	ds.Uptime = revisedSeconds
	ds.Version = versjon
	err = json.NewEncoder(w).Encode(ds) //encoding and essentially printing the struct
	if err != nil {                     //checking for issues
		fmt.Println("Error with Encoder!")
		fmt.Fprintf(w, "500"+http.StatusText(500))
		return
	}

}
