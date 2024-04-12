package utils

func CheckModel(model string) bool {
	if model == "tts-1" || model == "tts-1-hd" {
		return true
	}
	return false
}
