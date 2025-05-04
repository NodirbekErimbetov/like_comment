package models

type Like struct {
	PostId string `json:"post_id"`
	Count  int64  `json:"count"`
}


