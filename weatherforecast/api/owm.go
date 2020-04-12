package api

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"swfs/models"

	"github.com/briandowns/openweathermap"
)

// OWM ...
type OWM struct {
	Key string
}

// NewOwmAPI метод для создания экземпляра интерфейса weatherforecast.WF
func NewOwmAPI() (*OWM, error) {
	key := os.Getenv("OWM_KEY")
	if key == "" {
		return nil, errors.New("OWM_KEY - переменная окружение не установлена")
	}
	return &OWM{key}, nil

}

// GetCurrent метод для получения прогноза погоды
func (owm OWM) GetCurrent(location string) (*models.WeatherData, error) {
	cwd, err := owm.getCurrent(location)
	if err != nil {
		return nil, err
	}
	wd := &models.WeatherData{}
	if len(cwd.Name) == 0 { // некорректная локация
		wd.Location = location
		wd.Successful = false
		return wd, nil
	}
	wd.Successful = true
	wd.Location = cwd.Name
	var descriptions []string
	for _, item := range cwd.Weather {
		descriptions = append(descriptions, item.Description)
	}
	wd.WeatherDescription = strings.Join(descriptions, ",")
	wd.Temp = cwd.Main.Temp
	wd.TempMin = convertTempToString(cwd.Main.TempMin)
	wd.TempMax = convertTempToString(cwd.Main.TempMax)
	wd.RecommendationDescription = wd.GetRecommendationDescription()

	return wd, nil
}

// convertTempToString метод для конвертации температуры в строку
// делается округление температуры до целого числа
// и к положительной температуре добавляется префикс "+"
func convertTempToString(f float64) string {
	if f > 0 && f < 1 {
		return "0"
	}
	pref := ""
	if f > 0 {
		pref = "+"
	}
	s := strconv.FormatFloat(f, 'f', 0, 32)
	s = fmt.Sprintf("%s%s", pref, s)
	return s
}

// getCurrent внутренний метод
// для использования библиотеки openweathermap
func (owm OWM) getCurrent(location string) (*openweathermap.CurrentWeatherData, error) {
	w, err := openweathermap.NewCurrent("C", "ru", owm.Key)
	if err != nil {
		return nil, err
	}
	w.CurrentByName(location)
	return w, nil
}
