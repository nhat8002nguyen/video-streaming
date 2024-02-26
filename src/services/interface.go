package services

import "io"

type YouTubeClient interface {
	GetStreamReader() (io.ReadCloser, error)
}

type VideoService interface {
	SearchVideos(text string, amount int64) ([]SearchedVideoMeta, error)
	GetStreamReader() (io.ReadCloser, error)
}
