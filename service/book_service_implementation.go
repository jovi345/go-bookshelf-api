package service

import (
	"github.com/jovi345/go-bookshelf-api/helper"
	"github.com/jovi345/go-bookshelf-api/model/domain"
	"github.com/jovi345/go-bookshelf-api/model/web"
	"github.com/jovi345/go-bookshelf-api/repository"
)

type bookServiceImplementation struct {
	repo repository.BookRepository
}

// function ini ditambah belakangan
// setelah semua method pada interface
// diaplikasikan
func NewBookService(repo repository.BookRepository) BookService {
	return &bookServiceImplementation{repo: repo}
}

// s stands for service
func (s *bookServiceImplementation) Save(book domain.Book) (domain.Book, error) {
	return s.repo.Save(book)
}

func (s *bookServiceImplementation) Update(id string, updatedBook domain.Book) error {
	return s.repo.Update(id, updatedBook)
}

func (s *bookServiceImplementation) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *bookServiceImplementation) FindAll() ([]web.BookResponse, error) {
	books, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	return helper.FormatToBookResponses(books), nil
}

func (s *bookServiceImplementation) FindByID(id string) (domain.Book, error) {
	return s.repo.FindByID(id)
}
