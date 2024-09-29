package helper

import (
	"github.com/jovi345/go-bookshelf-api/model/domain"
	"github.com/jovi345/go-bookshelf-api/model/web"
)

func FormatToBookResponse(book domain.Book) web.BookResponse {
	return web.BookResponse{
		ID:        book.ID,
		Name:      book.Name,
		Publisher: book.Publisher,
	}
}

func FormatToBookResponses(books []domain.Book) []web.BookResponse {
	var bookResponses []web.BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, FormatToBookResponse(book))
	}

	return bookResponses
}
