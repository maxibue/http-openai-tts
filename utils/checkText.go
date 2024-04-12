package utils

func CheckText(text int) bool {
	if text < 4097 && text > 0 {
		return true
	}
	return false
}
