package service

import "go-book-inventory/repository"

type Controller struct {
	repo repository.BookRepository
}

func NewController(br repository.BookRepository) *Controller {
	return &Controller{repo: br}
}


