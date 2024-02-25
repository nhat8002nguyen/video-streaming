package services

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeService struct{}

func (s *YoutubeService) SearchVideos(text string, amount int64) ([]SearchedVideoMeta, error) {
	developerKey := os.Getenv("GOOGLE_DEV_KEY")

	query := flag.String("query", text, "Search term")
	maxResults := flag.Int64("max-results", amount, "Max YouTube results")

	flag.Parse()

	service, err := youtube.NewService(context.TODO(), option.WithAPIKey(developerKey))
	if err != nil {
		log.Printf("Error creating new YouTube client: %v", err)
		return nil, err
	}

	// Make the API call to YouTube.
	call := service.Search.List([]string{"id,snippet"}).
		Q(*query).
		MaxResults(*maxResults)
	response, err := call.Do()
	if err != nil {
		log.Printf("Error querying videos: %v", err)
		return nil, err
	}

	// Group video, channel, and playlist results in separate lists.
	videos := make(map[string]string)
	channels := make(map[string]string)
	playlists := make(map[string]string)

	// Iterate through each item and add it to the correct list.
	results := make([]SearchedVideoMeta, len(videos))
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			videos[item.Id.VideoId] = item.Snippet.Title
			results = append(results, SearchedVideoMeta{
				Id:      item.Id.VideoId,
				Title:   item.Snippet.Title,
				Website: "youtube",
			})
		case "youtube#channel":
			channels[item.Id.ChannelId] = item.Snippet.Title
		case "youtube#playlist":
			playlists[item.Id.PlaylistId] = item.Snippet.Title
		}
	}

	s.printIDs("Videos", videos)
	s.printIDs("Channels", channels)
	s.printIDs("Playlists", playlists)

	return results, nil
}

// Print the ID and title of each result in a list as well as a name that
// identifies the list. For example, print the word section name "Videos"
// above a list of video search results, followed by the video ID and title
// of each matching video.
func (s *YoutubeService) printIDs(sectionName string, matches map[string]string) {
	fmt.Printf("%v:\n", sectionName)
	for id, title := range matches {
		fmt.Printf("[%v] %v\n", id, title)
	}
	fmt.Printf("\n\n")
}
