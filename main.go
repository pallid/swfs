package main

import (
	"net/http"
	"os"
	"swfs/handlers"
	log "swfs/loger"
	"swfs/weatherforecast"
	"swfs/weatherforecast/api"

	"github.com/joho/godotenv"
)

var port = "8888"

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Warning("Отсутствует конфигурационный файл .env")
	}
}

// Run the app
func main() {
	owm, err := api.NewOwmAPI()
	if err != nil {
		log.Error(err)
		panic(err)
	}
	if v, e := os.LookupEnv("PORT"); e {
		port = v
	}
	var impl weatherforecast.WF = owm
	weatherforecast.SetWeatherForecast(impl)
	log.Infof("Запуск приложения: порт %s", port)
	http.HandleFunc("/", handlers.StartHandler)
	http.ListenAndServe(":"+port, nil)
}
