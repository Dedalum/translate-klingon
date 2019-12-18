package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoConfigFile(t *testing.T) {
	appConfig, err := LoadAppConfig("")
	assert.NotNil(t, err)
	assert.Nil(t, appConfig)
}

func TestLoadAppConfig(t *testing.T) {
	const configTest = `
{
	"http_client": {
		"api_host": "exemple.fr"
	}
}
`
	file, err := ioutil.TempFile(os.TempDir(), "prefix")
	if err != nil {
		fmt.Println("couldn't create temp file")
	}
	defer os.Remove(file.Name())

	file.WriteString(configTest)
	appConfig, err := LoadAppConfig(file.Name())

	assert.Nil(t, err)
	if assert.NotNil(t, appConfig) {
		assert.Equal(t, appConfig.HTTPClient.APIHost, "exemple.fr")
	}
}
