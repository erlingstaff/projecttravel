package projecttravel

import "net/http"

func handlerWeather(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
}
