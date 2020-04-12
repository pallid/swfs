package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	log "swfs/loger"
	"swfs/models"
	"swfs/weatherforecast"
)

// StartHandler стартовый хендлер приложения
func StartHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Error(err)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}
	t, err := template.ParseFiles("templates/start.html")
	if err != nil {
		log.Error(http.StatusInternalServerError)
		fmt.Fprint(w, http.StatusInternalServerError)
		return
	}
	wd := &models.WeatherData{}
	location := r.FormValue("location")
	if location != "" {
		log.Infof("requested location: %s", location)
		wd, err = weatherforecast.GetCurrent(location)
		if err != nil {
			log.Error(err)
		}
	}
	t.Execute(w, wd)
}
