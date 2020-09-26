package model

// Size type
type Size struct {
	Width int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

// Data type
type Data struct {
	Size Size `json:"size,omitempty"`
	Path string `json:"path,omitempty"`
}

// ImageObject Structure
type ImageObject struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data Data `json:"data,omitempty"`
}

