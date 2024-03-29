package translate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTranslator(t *testing.T) {

	expectedAlphabetMap := map[string]string{
		"a":   "0xF8D0",
		"b":   "0xF8D1",
		"ch":  "0xF8D2",
		"d":   "0xF8D3",
		"e":   "0xF8D4",
		"gh":  "0xF8D5",
		"h":   "0xF8D6",
		"i":   "0xF8D7",
		"j":   "0xF8D8",
		"l":   "0xF8D9",
		"m":   "0xF8DA",
		"n":   "0xF8DB",
		"ng":  "0xF8DC",
		"o":   "0xF8DD",
		"p":   "0xF8DE",
		"q":   "0xF8DF",
		"Q":   "0xF8E0",
		"r":   "0xF8E1",
		"s":   "0xF8E2",
		"t":   "0xF8E3",
		"tlh": "0xF8E4",
		"u":   "0xF8E5",
		"v":   "0xF8E6",
		"w":   "0xF8E7",
		"y":   "0xF8E8",
		"'":   "0xF8E9",
		"0":   "0xF8F0",
		"1":   "0xF8F1",
		"2":   "0xF8F2",
		"3":   "0xF8F3",
		"4":   "0xF8F4",
		"5":   "0xF8F5",
		"6":   "0xF8F6",
		"7":   "0xF8F7",
		"8":   "0xF8F8",
		"9":   "0xF8F9",
		",":   "0xF8FD",
		".":   "0xF8FE",
		" ":   "0x0020",
	}
	config := &TranslatorConfig{
		AlphabetMapFilePath: "../en_to_klingon.json",
	}
	translator, err := NewTranslator(config)

	assert.Nil(t, err)
	if assert.NotNil(t, translator.alphabetMap) {
		assert.Equal(t, expectedAlphabetMap, translator.alphabetMap)
	}
}

func TestConvert(t *testing.T) {
	expectedKlingon := []string{
		"0xF8E5",
		"0xF8D6",
		"0xF8E5",
		"0xF8E1",
		"0xF8D0",
	}
	config := &TranslatorConfig{
		AlphabetMapFilePath: "../en_to_klingon.json",
	}
	translator, err := NewTranslator(config)

	klingon, err := translator.Convert("Uhura")
	assert.Nil(t, err)
	if assert.NotEmpty(t, klingon) {
		assert.Equal(t, expectedKlingon, klingon)
	}
}

func TestConvertSpecific(t *testing.T) {
	expectedKlingon := []string{
		"0xF8E4",
		"0xF8D5",
		"0xF8E9",
		"0xF8D0",
		"0x0020",
		"0xF8D2",
	}
	config := &TranslatorConfig{
		AlphabetMapFilePath: "../en_to_klingon.json",
	}
	translator, err := NewTranslator(config)

	klingon, err := translator.Convert("Tlhgh'a ch")
	assert.Nil(t, err)
	if assert.NotEmpty(t, klingon) {
		assert.Equal(t, expectedKlingon, klingon)
	}
}

func TestConvertUntranslatable(t *testing.T) {
	config := &TranslatorConfig{
		AlphabetMapFilePath: "../en_to_klingon.json",
	}
	translator, err := NewTranslator(config)

	word := "fUhura"
	klingon, err := translator.Convert(word)
	assert.Error(t, err, "cannot translate '%s' into Klingon", word)
	assert.Empty(t, klingon)
}
