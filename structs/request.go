package structs

type Request struct {
	Model          string  `json:"model"`
	Input          string  `json:"input"`
	Voice          string  `json:"voice"`           // Note: Make values optional asap
	ResponseFormat string  `json:"response_format"` // Optional, will use API's default if not provided
	Speed          float64 `json:"speed"`           // Optional, will use API's default if not provided
}
