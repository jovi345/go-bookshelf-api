package repository

import "github.com/jovi345/go-bookshelf-api/model/domain"

type BookRepository interface {
	Save(book domain.Book) (domain.Book, error)
	Update(id string, updatedBook domain.Book) (domain.Book, error)
	Delete(id string) error
	FindAll() ([]domain.Book, error)
	FindByID(id string) (domain.Book, error)
}
