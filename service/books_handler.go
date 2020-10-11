package service

import (
	"encoding/json"
	"fmt"
	"go-book-inventory/models"
	"io/ioutil"
	"net/http"
)

func (c Controller) booksHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errStr := fmt.Sprintf("error reading body: %v", err.Error())
		w.Write([]byte(errStr))
		return
	}
	book := models.Book{}
	err = json.Unmarshal(bodyBytes, &book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errStr := fmt.Sprintf("error reading body: %v", err.Error())
		w.Write([]byte(errStr))
		return
	}
	createdBook:=c.repo.AddBook(book)
	var marshalledBook []byte
	marshalledBook, err = json.Marshal(createdBook)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errStr := fmt.Sprintf("error marshalling response: %v", err.Error())
		w.Write([]byte(errStr))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(marshalledBook)
}
