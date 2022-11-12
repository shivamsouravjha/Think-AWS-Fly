package structs

type Video struct {
	ID          VideoID
	PublishedAt string `json:"publishedAt"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ChannelID   string `json:"channelId"`
	ChannelName string `json:"channelTitle"`
}
type VideoID struct {
	ID string `json:"videoID"`
}
