package post

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"piepay/config"
	"piepay/services/es"
	"piepay/structs"
	"strings"
	"time"

	"github.com/olivere/elastic/v7"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

var (
	maxResults = flag.Int64("max-results", 25, "Max YouTube results") //maxmimum number of videos called per call
	videotype  = flag.String("type", "video", "Type of video")        //calling video data from youtube api
)

func UploadVideoMetaData() {
	var developerKey = config.Get().YoutubeKey //developer key

	startafter := time.Now().Add(-time.Second * 1000).Format(time.RFC3339) //getting videos uploaded in last 1000s

	flag.Parse()

	client := &http.Client{
		Transport: &transport.APIKey{Key: developerKey}, //setting key
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	arraya := []string{"id", "snippet"} //parts asked
	call := service.Search.List(arraya).
		MaxResults(*maxResults).PublishedAfter(startafter).Type(*videotype) //calling data
	response, err := call.Do()
	if err != nil {
		errorText := strings.Split(err.Error(), ", ")
		if errorText[1] == "quotaExceeded" {
			fmt.Println("Keys changed")
			config.UpdateKey() //if error is of quota exceeded , past new key
		}
	} else {
		videos := []structs.Video{}
		for _, item := range response.Items { //store all required data from data
			switch item.Id.Kind {
			case "youtube#video":
				videos = append(videos, structs.Video{
					ID:          structs.VideoID{item.Id.VideoId},
					PublishedAt: item.Snippet.PublishedAt,
					ChannelName: item.Snippet.ChannelTitle,
					Title:       item.Snippet.Title,
					Description: item.Snippet.Description,
					ChannelID:   item.Snippet.ChannelId,
				})
			}
		}
		index := config.Get().Index
		uploadOnEs(index, videos)
	}
}

func uploadOnEs(index string, matches []structs.Video) {
	bulk := es.Client().Bulk()

	for i := range matches { //adding each unit in bulk
		idStr := matches[i].ID.ID
		req := elastic.NewBulkIndexRequest().Id(idStr).Index(index).Doc(matches[i])
		bulk = bulk.Add(req)
	}

	resp, _ := bulk.Do(context.Background()) //bulk upload data
	fmt.Println(resp)
	fmt.Printf("\n\n")

}
