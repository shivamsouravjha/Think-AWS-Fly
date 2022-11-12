package requests

type GetVideo struct {
	Page int `json:"page"  `
	Size int `json:"size"`
}

type SearchVideo struct {
	Title       string `json:"title"  `
	Description string `json:"description"`
	Page        int    `json:"page"  `
	Size        int    `json:"size"`
}
