package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/kkdai/youtube/v2"
)

func main() {
	videoClient := youtube.Client{
		HTTPClient: &http.Client{},
	}
	video, err := videoClient.GetVideo("https://www.youtube.com/watch?v=shLUsd7kQCI")
	if err != nil {
		panic(err)
	}

	// Typically youtube only provides separate streams for video and audio.
	// If you want audio and video combined, take a look a the downloader package.
	formats := video.Formats.Quality("medium")
	reader, _, err := videoClient.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	p := make([]byte, 1024)
	for {
		_, err = reader.Read(p)
		if err == io.EOF {
			fmt.Println("Completed reading file.")
			break
		}

		if err != nil {
			log.Fatal("Error reading file: ", err)
		}
		fmt.Println("Reading...")
	}

}
