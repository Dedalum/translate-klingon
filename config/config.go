package config

import (
	"encoding/json"
	"os"

	"github.com/Dedalum/translate-klingon/http"
	"github.com/Dedalum/translate-klingon/translate"
)

type (
	// AppConfig configuration structure
	AppConfig struct {
		HTTPClient http.ClientConfig          `json:"http_client"`
		Translator translate.TranslatorConfig `json:"translator"`
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
