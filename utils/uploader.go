package utils

import (
	post "piepay/controllers/POST"
	"time"
)

var timer *time.Ticker //initalise ticker to tick at 1000 s

func Uploader() {
	if timer == nil {
		timer = time.NewTicker(10 * time.Second) //initialise ticker
		go func() {
			for range timer.C {
				post.UploadVideoMetaData() //after every 1000s upload data to db
			}
		}()
	}
}
