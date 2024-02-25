package main

import (
	"log"
	"net/http"
	"video-streaming/src/apis"
	"video-streaming/src/services"

	"github.com/joho/godotenv"
)

var videoAPI apis.VideoAPI

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when loading envs.")
	}

	// Init services and dependencies
	youtubeService := services.YoutubeService{}
	videoAPI = &apis.VideoAPIImpl{
		Service: &youtubeService,
	}
}

func main() {
	http.HandleFunc("/search", videoAPI.SearchVideos)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
