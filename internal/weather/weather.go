package weather

// Forecast answers a fixed forecast; the sample never calls the network.
func Forecast(city, key string) map[string]any {
	source := "fixture"
	if key != "" {
		source = "openweather"
	}
	return map[string]any{"city": city, "high": 21, "low": 12, "source": source}
}
