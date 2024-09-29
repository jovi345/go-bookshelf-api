package repository

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jovi345/go-bookshelf-api/model/domain"
)

type bookRepositoryImplementation struct {
	books []domain.Book
}

// function ini ditambah belakangan
// setelah semua method pada interface
// diaplikasikan
func NewBookRepository() BookRepository {
	return &bookRepositoryImplementation{
		books: []domain.Book{},
	}
}

// r stands for repository
func (r *bookRepositoryImplementation) Save(book domain.Book) (domain.Book, error) {
	id := uuid.New()
	book.ID = id.String()
	book.InsertedAt = time.Now().Format(time.RFC3339)
	book.UpdatedAt = book.InsertedAt
	book.Finished = book.ReadPage == book.PageCount

	r.books = append(r.books, book)
	return book, nil
}

func (r *bookRepositoryImplementation) Update(id string, updatedBook domain.Book) (domain.Book, error) {
	for i, book := range r.books {
		if book.ID == id {
			updatedBook.ID = book.ID
			updatedBook.InsertedAt = book.InsertedAt
			updatedBook.UpdatedAt = time.Now().Format(time.RFC3339)
			updatedBook.Finished = updatedBook.ReadPage == updatedBook.PageCount
			r.books[i] = updatedBook
			return updatedBook, nil
		}
	}
	return domain.Book{}, errors.New("book not found")
}

func (r *bookRepositoryImplementation) Delete(id string) error {
	for i, book := range r.books {
		if book.ID == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

func (r *bookRepositoryImplementation) FindAll() ([]domain.Book, error) {
	return r.books, nil
}

func (r *bookRepositoryImplementation) FindByID(id string) (domain.Book, error) {
	for _, book := range r.books {
		if book.ID == id {
			return book, nil
		}
	}
	return domain.Book{}, errors.New("book not found")
}
