package model

type Post struct {
	Name        string `json:"name,omitempty"`
	Title       string `json:"title,omitempty"`
	Discription string `json:"discription,omitempty"`
}
