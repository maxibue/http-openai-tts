package utils

import "fmt"

func AvailableRoutes(allowHosting bool, allowAdmin bool) {
	fmt.Println("\nList of all available API routes:\n/ping")
	if allowHosting {
		fmt.Println("/tts\n/raw\n/*")
	} else {
		fmt.Println("/tts")
	}
	if allowAdmin {
		fmt.Println("/admin/add\n/admin/remove")
	}
	fmt.Print("\n")
}
