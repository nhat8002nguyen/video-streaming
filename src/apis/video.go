package apis

import (
	"encoding/json"
	"net/http"
	"strconv"
	"video-streaming/src/services"
)

type VideoAPIImpl struct {
	Service services.VideoService
}

func (impl *VideoAPIImpl) SearchVideos(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	search := query.Get("text")

	amount, err := strconv.ParseInt(query.Get("amount"), 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	searchedMetas, err := impl.Service.SearchVideos(search, int64(amount))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	data := make(map[string]any)
	data["total"] = len(searchedMetas)
	data["data"] = searchedMetas

	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}
}
