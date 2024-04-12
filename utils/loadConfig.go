package utils

import (
	"encoding/json"
	"os"

	"github.com/maximierung/http-openai-tts/structs"
)

func LoadConfig() (structs.Config, error) {
	configFile, err := os.Open("./config/config.json")
	if err != nil {
		return structs.Config{}, err
	}
	defer configFile.Close()

	var config structs.Config
	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(&config); err != nil {
		return structs.Config{}, err
	}

	return config, nil
}
