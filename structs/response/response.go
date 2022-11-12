package response

import "piepay/structs"

type VideoResponse struct {
	Status  string   `json:"status"`
	Message string   `json:"message"`
	Data    GetVideo `json:"data"`
}

type GetVideo struct {
	VideoDetails []structs.Video
	Page         int `json:"page"`
	Size         int `json:"size"`
}
