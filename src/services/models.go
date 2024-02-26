package services

type SearchedVideoMeta struct {
	Id        string          `json:"id"`
	Title     string          `json:"title"`
	Website   string          `json:"website"`
	URL       string          `json:"url"`
	Thumbnail *videoThumbnail `json:"thumbnail"`
}

type videoThumbnail struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
}
