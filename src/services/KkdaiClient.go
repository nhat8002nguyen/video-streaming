package services

import (
	"io"
	"net/http"

	"github.com/kkdai/youtube/v2"
)

type KkdaiClient struct{}

func (c KkdaiClient) GetStreamReader() (io.ReadCloser, error){
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
	
	return reader, err
}