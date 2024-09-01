package models

type Pokemon struct {
	ID       int    `json:"id"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
	Type     string `json:"type"`
}
