package models

// WeatherData основная структура приложения
// отображает информацию о погоде для локации
type WeatherData struct {
	Location                  string
	WeatherDescription        string
	Temp                      float64
	TempMin                   string
	TempMax                   string
	Successful                bool
	RecommendationDescription string
}

// GetRecommendationDescription метод для получения рекомендации
// на основе заполненных данных
func (wd WeatherData) GetRecommendationDescription() string {

	switch {
	case wd.Temp < -10:
		return "Одевайся потеплее, возьми теплую шапку..."
	case wd.Temp < 5:
		return "Куртка, шапка - без них никуда"
	case wd.Temp < 15:
		return "Можно одеть демисезонную куртку"
	case wd.Temp < 20:
		return "Одень легкую куртку"
	case wd.Temp < 25:
		return "Отличная погода, куртка ненужна"
	case wd.Temp < 30:
		return "Аккуратней на солнце, возьми головной убор"
	}

	return "#сидидома"
}
