package main

import ("net/http"
	"encoding/json"
	"strings"
)

const API_KEY = "8e22f31e4f12b135ecdcec6af24aaaee"

type weatherData struct{
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `jason:"temp"`
	} `json:""main`
}

func main()  {

	http.HandleFunc("/hi", hi)
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]

		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	})

	http.HandleFunc("/", hello)
	http.ListenAndServe(":8080", nil)

}
func hello(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hello!"))
}
func hi(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("hi Andela!"))
}
func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?qAPPID=8e22f31e4f12b135ecdcec6af24aaaee" + city)
	if err != nil {
		return weatherData{}, err
	}

	defer resp.Body.Close()

	var d weatherData

	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}

	return d, nil
}