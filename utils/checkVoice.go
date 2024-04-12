package utils

func CheckVoice(voice string) bool {
	switch voice {
	case "alloy":
		return true
	case "echo":
		return true
	case "fable":
		return true
	case "onyx":
		return true
	case "nova":
		return true
	case "shimmer":
		return true
	default:
		return false
	}
}
