package apis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"video-streaming/src/services"

	"github.com/gorilla/websocket"
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Replace this with your origin validation logic
		// (e.g., check allowed origins or perform authentication)
		return true
	},
}

func (impl *VideoAPIImpl) HandleStreamWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	streamReader, err := impl.Service.GetStreamReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		// Process the received message
		fmt.Printf("Received message: %s\n", message)

		chunk := make([]byte, 1024)
		_, err = streamReader.Read(chunk)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				break
			}
		}

		// Optionally, send a response message
		err = conn.WriteMessage(messageType, chunk)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
}
