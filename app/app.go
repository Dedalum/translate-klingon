package app

import (
	"fmt"
	"log"

	"github.com/Dedalum/translate-klingon/config"
	"github.com/Dedalum/translate-klingon/http"
	"github.com/Dedalum/translate-klingon/translate"
)

func getCharacterSpecies(clientConfig *http.ClientConfig,
	characterName string) (string, error) {

	client := http.NewClient(clientConfig)
	characterUIDs, err := client.GetCharacterUIDs(characterName)
	if err != nil {
		return "", err
	}
	if len(characterUIDs) == 0 {
		return "", fmt.Errorf("No UID found for character %s", characterName)
	}

	characterSpeciesList, err := client.GetCharacterSpeciesList(characterUIDs[0])
	if err != nil {
		return "", err
	}
	if len(characterSpeciesList) == 0 {
		return "", fmt.Errorf("No species found for character %s", characterName)
	}

	return characterSpeciesList[0], nil
}

// Run runs the main app
func Run(config *config.AppConfig, characterName string) {
	species, err := getCharacterSpecies(&config.HTTPClient, characterName)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	translator, err := translate.NewTranslator(&config.Translator)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}
	klingon, err := translator.Convert(characterName)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// --- printing results
	for _, val := range klingon {
		fmt.Printf("%s ", val)
	}

	fmt.Printf("\n\n")
	fmt.Printf("%s\n", species)
}
