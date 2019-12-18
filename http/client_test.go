package http

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseUIDs(t *testing.T) {
	expectedUIDs := []string{
		"CHMA0000115364",
		"CHMA0000023576",
		"CHMA0000068639",
	}
	jsonData := helperLoadBytes(t, "uhura.json")

	uids, err := parseUIDs(jsonData)
	assert.Nil(t, err)
	if assert.NotNil(t, uids) {
		assert.Equal(t, expectedUIDs, uids)
	}
}

func TestGetCharacterUIDs(t *testing.T) {

	config := &ClientConfig{
		APIHost: "stapi.co",
	}
	characterName := "Uhura"
	expectedUIDs := []string{
		"CHMA0000115364",
		"CHMA0000023576",
		"CHMA0000068639",
	}

	c := NewClient(config)
	uids, err := c.GetCharacterUIDs(characterName)
	assert.Nil(t, err)
	if assert.NotNil(t, uids) {
		assert.Equal(t, expectedUIDs, uids)
	}
}

func TestParseSpeciesList(t *testing.T) {
	expectedSpeciesList := []string{"Human"}
	jsonData := helperLoadBytes(t, "uhura_full.json")

	speciesList, err := parseSpeciesList(jsonData)
	assert.Nil(t, err)
	if assert.NotNil(t, speciesList) {
		assert.Equal(t, expectedSpeciesList, speciesList)
	}
}

func TestGetCharacterSpeciesList(t *testing.T) {

	config := &ClientConfig{
		APIHost: "stapi.co",
	}
	characterUID := "CHMA0000115364"
	expectedSpeciesList := []string{
		"Human",
	}

	c := NewClient(config)
	speciesList, err := c.GetCharacterSpeciesList(characterUID)
	assert.Nil(t, err)
	if assert.NotNil(t, speciesList) {
		assert.Equal(t, expectedSpeciesList, speciesList)
	}
}

func helperLoadBytes(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name) // relative path
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return bytes
}
