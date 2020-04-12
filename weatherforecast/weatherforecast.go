package weatherforecast

import "swfs/models"

// WF ...
type WF interface {
	GetCurrent(string) (*models.WeatherData, error)
}

var impl WF

// SetWeatherForecast ...
func SetWeatherForecast(wf WF) {
	impl = wf
}

// GetCurrent ...
func GetCurrent(s string) (*models.WeatherData, error) {
	return impl.GetCurrent(s)
}
