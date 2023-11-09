package model

type WebResponse[T any] struct {
	Data           T               `json:"data,omitempty"`
	PagingResponse *PagingResponse `json:"paging,omitempty"`
}

type PagingResponse struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
