package model

// Size type
type Size struct {
	Width int `json:"width"`
	Height int `json:"height"`
}

// Data type
type Data struct {
	Size Size `json:"size"`
	Path string `json:"path"`
}

// ImageObject Structure
type ImageObject struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data Data `json:"data"`
}

