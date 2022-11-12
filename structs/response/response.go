package response

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseFile struct {
	Status      string   `json:"status"`
	BucketNames []string `json:"buckets"`
}
