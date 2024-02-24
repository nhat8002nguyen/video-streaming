package main

import (
	"log"
	"video-streaming/src/apis"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading envs.")
	}
}

func main() {
	apis.SearchVideos()
}
