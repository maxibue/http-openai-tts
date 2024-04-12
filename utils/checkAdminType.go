package utils

func CheckAdminType(format string) bool {
	switch format {
	case "add":
		return true
	case "remove":
		return true
	default:
		return false
	}
}
