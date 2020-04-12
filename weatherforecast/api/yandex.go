package api

import (
	"errors"
	"os"
)

// Yandex ...
type Yandex struct {
	GeoCode string
	Weather string
}

// NewYandexAPI ...
func NewYandexAPI() (*Yandex, error) {
	geocode := os.Getenv("YA_GEOCODE")
	if geocode == "" {
		return nil, errors.New("YA_GEOCODE - переменная окружение не установлена")
	}
	weather := os.Getenv("YA_WEATHER")
	if weather == "" {
		return nil, errors.New("YA_WEATHER - переменная окружение не установлена")
	}
	return &Yandex{GeoCode: geocode, Weather: weather}, nil

}

// GetCurrent ...
func (ya Yandex) GetCurrent(s string) error {
	return nil
}
