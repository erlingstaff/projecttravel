package projecttravel

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=60.7917739,10.6819978&radius=1500&key=AIzaSyBw7w2Yy6edsg28xfUcIwQ-yPHcY3fPqx0
//https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=
//&radius=1500&key=AIzaSyBw7w2Yy6edsg28xfUcIwQ-yPHcY3fPqx0
//https://maps.googleapis.com/maps/api/geocode/json?address=
//&key=AIzaSyBw7w2Yy6edsg28xfUcIwQ-yPHcY3fPqx0

//HandlerPOI is the handler from main.go
func HandlerPOI(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	city := ""
	keys, ok := r.URL.Query()["city"]
	if ok != true {
		fmt.Fprintln(w, "500 "+http.StatusText(500)) //!!
		fmt.Fprintln(w, "\nPlease use correct path, syntax is /project/v1/places/?city={city_name}")
		return
	}
	city = keys[0]
	lonlatObj := findCityXY(city, w)
	llng := fmt.Sprintf("%f", lonlatObj.Results[0].Geometry.Location.Lng)
	llat := fmt.Sprintf("%f", lonlatObj.Results[0].Geometry.Location.Lat)
	xyStr := llat + "," + llng
	fmt.Println(xyStr)
	resp, err := http.Get("https://maps.googleapis.com/maps/api/place/nearbysearch/json?location=" + xyStr + "&radius=1500&key=AIzaSyBw7w2Yy6edsg28xfUcIwQ-yPHcY3fPqx0" + "&sensor=false")
	var d poi
	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		fmt.Fprintln(w, "500 "+http.StatusText(500))
		return
	}
	err = json.NewEncoder(w).Encode(d) //encodes it to the w parameter
	if err != nil {
		fmt.Println("Error Encoding")
		fmt.Fprintf(w, "500"+http.StatusText(500))
		return
	}

}

func findCityXY(city string, w http.ResponseWriter) lonlat {

	resp, err := http.Get("https://maps.googleapis.com/maps/api/geocode/json?address=" + city + "&key=AIzaSyBw7w2Yy6edsg28xfUcIwQ-yPHcY3fPqx0")
	if err != nil {
		fmt.Fprintln(w, "500 "+http.StatusText(500))
		return lonlat{}
	}
	defer resp.Body.Close()

	var d lonlat

	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		fmt.Fprintln(w, "500 "+http.StatusText(500))
		return lonlat{}
	}

	return d
}

type poi struct {
	Results []struct {
		Name         string `json:"name"`
		Vicinity     string `json:"vicinity"`
		OpeningHours struct {
			OpenNow bool `json:"open_now"`
		} `json:"opening_hours,omitempty"`
		Rating float64 `json:"rating,omitempty"`
	} `json:"results"`
}

type lonlat struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}
