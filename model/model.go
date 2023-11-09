package model

type WebResponse[T any] struct {
	Data T `json:"data,omitempty"`
}

type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
