package main

import (
	"flag"
	"github.com/Lineblaze/GwentGallery/app/internal/app"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	readEnv := flag.String("env-file", ".env", "Path to application env file.")
	configPath := flag.String("config-path", "config.yml", "Path to application config file.")
	flag.Parse()
	println(*configPath)
	if *readEnv != "" {
		if err := godotenv.Load(*readEnv); err != nil {
			log.Fatal(err)
		}
	}

	application, err := app.New(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := application.Start(); err != nil {
		log.Fatal(err)
	}
}
