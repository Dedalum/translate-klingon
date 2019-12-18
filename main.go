package main

import (
	"flag"
	"log"

	"github.com/Dedalum/translate-klingon/app"
	"github.com/Dedalum/translate-klingon/config"
)

var (
	klingonInput string
	configPath   string
)

func init() {
	flag.StringVar(&klingonInput, "name", "", "Klingon character in Latin alphabet to translate")
	flag.StringVar(&configPath, "config-path", "config.json", "Klingon word in Latin alphabet to translate")
}

func main() {
	flag.Parse()

	config, err := config.LoadAppConfig(configPath)
	if err != nil {
		log.Printf("%s\n", err)
		return
	}

	// Run app
	app.Run(config, klingonInput)
}
