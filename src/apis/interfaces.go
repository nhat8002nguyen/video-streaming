package apis

import "net/http"

type VideoAPI interface {
	SearchVideos(w http.ResponseWriter, r *http.Request)
}
