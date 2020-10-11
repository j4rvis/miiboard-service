package repository

import (
	"container/list"
	"fmt"
	"math/rand"
	"miiboard-service/models"
	"strings"
	"time"
)

var hexChars = [16]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}

type BookRepository struct {
	bookList *list.List
}

func NewBookRepository() *BookRepository {
	return &BookRepository{bookList: list.New()}
}

func (br *BookRepository) AddBook(b models.Book) models.Book {
	b.Id = GenerateObjectId()
	br.bookList.PushBack(b)
	return b
}

func GenerateObjectId() string {
	oidBuilder := strings.Builder{}
	oidBuilder.WriteString(fmt.Sprintf("%12x", time.Now().UnixNano()/1000))
	len := oidBuilder.Len()
	for n := 0; n < 24-len; n++ {
		oidBuilder.WriteRune(hexChars[rand.Intn(15)])
	}
	return oidBuilder.String()
}
