package models

type Book struct {
	Id      string   `json:"id"`
	Title   string   `json:"title"`
	Author  Author   `json:"author"`
	Release string   `json:"release"`
	Tags    []string `json:"tags"`
}

type Author struct {
	Id        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Birthday  string `json:"birthday"`
	Country   string `json:"country"`
}
