package utils

func CheckFormat(format string) bool {
	switch format {
	case "mp3":
		return true
	case "opus":
		return true
	case "aac":
		return true
	case "flac":
		return true
	case "wav":
		return true
	case "pcm":
		return true
	default:
		return false
	}
}
