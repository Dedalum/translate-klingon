package config

import (
	"encoding/json"
	"os"

	"github.com/Dedalum/translate-klingon/http"
)

type (
	// AppConfig configuration structure
	AppConfig struct {
		HTTPClient http.ClientConfig `json:"http_client"`
	}
)

// LoadAppConfig loads the application configuration
func LoadAppConfig(filePath string) (*AppConfig, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var appConfig *AppConfig
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(&appConfig); err != nil {
		return nil, err
	}

	return appConfig, nil
}
