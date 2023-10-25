package model

type WebResponse[T any] struct {
	Data           T              `json:"data"`
	PagingResponse PagingResponse `json:"paging"`
}

type PagingResponse struct {
	Page  int `json:"page"`
	Size  int `json:"size"`
	Total int `json:"total"`
}
