package services

import "io"

type YouTubeClient interface {
	GetStreamReader() (io.ReadCloser, error)
}
