package myleetcode

func convertTemperature(celsius float64) []float64 {
	result := []float64{}
	result = append(result, celsius+float64(273.15))
	result = append(result, celsius*float64(1.8)+float64(32))
	return result
}
