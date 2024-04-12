package utils

func CheckSpeed(speed float64) bool {
	if speed >= 0.25 && speed <= 4.0 {
		return true
	}
	return false
}
