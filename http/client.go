package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type (
	// Client holds the http client to query webpages
	Client struct {
		apiURL string
		http.Client
	}

	// ClientConfig holds the configuration for the HTTP client
	ClientConfig struct {
		APIHost string `json:"api_host"`
	}
)

const (
	apiPath            = "api/v1/rest"
	apiCharacterSearch = "character/search"
	apiCharacter       = "character"
)

// NewClient returns a pointer to a new client
func NewClient(config *ClientConfig) *Client {
	return &Client{
		fmt.Sprintf("http://%s/%s", config.APIHost, apiPath),
		http.Client{
			Timeout: time.Second * 3,
		},
	}
}

// GetCharacterUIDs queries the API for a given character name. It returns a
// list of characters (e.g. "Uhura" returns 3 different UIDs).
func (c *Client) GetCharacterUIDs(name string) ([]string, error) {
	var characters []string

	var dataStr = []byte(fmt.Sprintf("name=%s", name))
	resp, err := c.Post(fmt.Sprintf("%s/%s", c.apiURL, apiCharacterSearch),
		"data", bytes.NewBuffer(dataStr))

	if err != nil {
		return characters, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	uids, err := parseUIDs(body)
	if err != nil {
		return characters, err
	}

	return uids, nil
}

// parseUIDs returns the list of UIDs for a character given the JSON
// response
func parseUIDs(jsonData []byte) ([]string, error) {
	var uids []string

	type character struct {
		UID string `json:"uid"`
	}
	var data struct {
		Characters []character `json:"characters"`
	}

	if err := json.Unmarshal(jsonData, &data); err != nil {
		return uids, err
	}

	for _, character := range data.Characters {
		uids = append(uids, character.UID)
	}

	return uids, nil
}

// GetCharacterSpeciesList returns the species for a given UID
func (c *Client) GetCharacterSpeciesList(characterUID string) ([]string, error) {

	var speciesList []string

	resp, err := c.Get(
		fmt.Sprintf("%s/%s?uid=%s", c.apiURL, apiCharacter, characterUID))
	if err != nil {
		return speciesList, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	speciesList, err = parseSpeciesList(body)
	if err != nil {
		return speciesList, err
	}

	return speciesList, nil

}

// parseSpecies returns the list of species for a character given the JSON
// response
func parseSpeciesList(jsonData []byte) ([]string, error) {
	var speciesList []string

	type character struct {
		SpeciesList []struct {
			Name string `json:"name"`
		} `json:"characterSpecies"`
	}

	var data struct {
		Character character `json:"character"`
	}

	if err := json.Unmarshal(jsonData, &data); err != nil {
		return speciesList, err
	}

	for _, species := range data.Character.SpeciesList {
		speciesList = append(speciesList, species.Name)
	}

	return speciesList, nil
}
