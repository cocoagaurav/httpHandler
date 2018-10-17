package model

type Post struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Discription string `json:"discription,omitempty"`
}
