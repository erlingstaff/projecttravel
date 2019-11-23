package projecttravel

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//HandlerWeather is the handler from main.go
func HandlerWeather(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	lon := ""
	lat := ""
	city := ""
	fs := FinishedStruct{}
	kelvin := -273.15
	keys, ok := r.URL.Query()["lon"]
	if ok {
		lon = keys[0]
	}
	keys, ok = r.URL.Query()["lat"]
	if ok {
		lat = keys[0]
	}
	if lon == "" && lat == "" {
		keys, ok = r.URL.Query()["city"]
		if ok {
			city = keys[0]
		} else {
			fmt.Fprintln(w, "500 "+http.StatusText(500)) //!!
			fmt.Fprintln(w, "\nPlease use correct path, syntax is /project/v1/places/?city={city_name} || ?lon={lon_coordinate}&lat={lat_coordinate}")
			return
		}
		printed, err := findWeatherCity(city)
		fs.Name = printed.Name
		fs.Degrees = printed.Main.Kelvin + kelvin
		fs.WeatherDescription = printed.Weather
		if err != nil {
			fmt.Fprintln(w, "500 "+http.StatusText(500)) //!!
		} else {
			err = json.NewEncoder(w).Encode(fs) //encodes it to the w parameter
			if err != nil {
				fmt.Println("Error Encoding")
				fmt.Fprintf(w, "500"+http.StatusText(500))
				return
			}
		}
	} else {
		printed, err := findWeatherXY(lon, lat)
		fs.Name = printed.Name
		fs.Degrees = printed.Main.Kelvin + kelvin
		fs.WeatherDescription = printed.Weather
		if err != nil {
			fmt.Fprintln(w, "500 "+http.StatusText(500)) //!!
		} else {
			err = json.NewEncoder(w).Encode(fs) //encodes it to the w parameter
			if err != nil {
				fmt.Println("Error Encoding")
				fmt.Fprintf(w, "500"+http.StatusText(500))
				return
			}
		}
	}

}

func findWeatherXY(lon string, lat string) (weatherData, error) {
	apiConfig, err := loadAPIConfig("APIConfig.json")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapAPIKey + "&lat=" + lat + "&lon=" + lon)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return weatherData{}, err
	}
	return d, nil
}

func findWeatherCity(city string) (weatherData, error) {
	apiConfig, err := loadAPIConfig("APIConfig.json")
	if err != nil {
		return weatherData{}, err
	}

	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=" + apiConfig.OpenWeatherMapAPIKey + "&q=" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	err = json.NewDecoder(resp.Body).Decode(&d)
	if err != nil {
		return weatherData{}, err
	}
	return d, nil
}

func loadAPIConfig(filename string) (apiConfigData, error) {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		return apiConfigData{}, err
	}

	var c apiConfigData
	err = json.Unmarshal(bytes, &c)
	if err != nil {
		return apiConfigData{}, err
	}

	return c, nil
}

type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type FinishedStruct struct {
	Name               string
	Degrees            float64
	WeatherDescription []struct {
		Description string `json:"description"`
	}
}

type apiConfigData struct {
	OpenWeatherMapAPIKey string `json:"OpenWeatherMapApiKey"`
	//GoogleApiKey         string `json:"GoogleApiKey"`
}
