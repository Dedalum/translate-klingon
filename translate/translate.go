package translate

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type (
	// Translator holds the structure for converting from Latin alphabet to
	// Klingon alphabet
	Translator struct {
		// alphabetMap's keys are in lower case to the exception of 'Q' which
		// would conflict with the Klingon letter 'q'
		alphabetMap map[string]string
	}

	// TranslatorConfig holds the configuration for a translator object
	TranslatorConfig struct {
		AlphabetMapFilePath string `json:"alphabet_map_file_path"`
	}
)

// NewTranslator returns a new Translator object and loads the alphabetMap
// from the given file path
func NewTranslator(config *TranslatorConfig) (*Translator, error) {

	alphabetMapFile, err := os.Open(config.AlphabetMapFilePath)
	if err != nil {
		return &Translator{}, err
	}
	defer alphabetMapFile.Close()

	var alphabetMap map[string]string
	jsonParser := json.NewDecoder(alphabetMapFile)
	if err := jsonParser.Decode(&alphabetMap); err != nil {
		return &Translator{}, err
	}

	t := &Translator{
		alphabetMap: make(map[string]string),
	}

	// put the alphabet map to lower case for simpler usage, except for 'Q'
	// which is the only letter with its lower case 'q' as another letter in
	// the Klingon alphabet
	for key, val := range alphabetMap {
		if key == "Q" {
			t.alphabetMap[key] = val
		} else {
			t.alphabetMap[strings.ToLower(key)] = val
		}
	}

	return t, nil
}

// Convert converts a given string in Latin alphabet to Klingon alphabet and
// returns the result as a slice of strings (hex values for the Klingon letters)
func (t *Translator) Convert(word string) ([]string, error) {
	var translated []string

	i := 0
	for i < len(word) {
		if val, ok := t.checkMap(string(word[i])); ok {
			translated = append(translated, val)
			i++
		} else if val, ok := t.checkMap(word[i : i+1]); ok {
			translated = append(translated, val)
			i += 2
		} else if val, ok := t.checkMap(word[i : i+2]); ok {
			translated = append(translated, val)
			i += 3
		} else {
			return translated, fmt.Errorf("cannot translate '%s' into Klingo", word)
		}
	}

	return translated, nil
}

func (t *Translator) checkMap(word string) (string, bool) {

	if word == "q" || word == "Q" {
		return t.alphabetMap[word], true
	} else {
		word = strings.ToLower(word)
		if val, ok := t.alphabetMap[word]; ok {
			return val, true
		} else {
			return "", false
		}
	}

}
