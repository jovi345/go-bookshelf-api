package service

import (
	"github.com/jovi345/go-bookshelf-api/model/domain"
	"github.com/jovi345/go-bookshelf-api/model/web"
)

type BookService interface {
	Save(book domain.Book) (domain.Book, error)
	Update(id string, updatedBook domain.Book) error
	Delete(id string) error
	FindAll() ([]web.BookResponse, error)
	FindByID(id string) (domain.Book, error)
}
