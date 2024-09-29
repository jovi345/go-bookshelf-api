package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jovi345/go-bookshelf-api/model/domain"
	"github.com/jovi345/go-bookshelf-api/service"
)

type BookHandler struct {
	service service.BookService
}

func NewBookHandler(service service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) AddBookHandler(c *gin.Context) {
	var book domain.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "invalid requeset",
		})
		return
	}

	if book.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal menambahkan buku. Mohon isi nama buku",
		})
		return
	}

	if book.ReadPage > book.PageCount {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal menambahkan buku. readPage tidak boleh lebih besar dari pageCount",
		})
		return
	}

	newBook, err := h.service.Save(book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "fail",
			"message": "Buku gagal ditambahkan",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Buku berhasil ditambahkan",
		"data": gin.H{
			"bookId": newBook.ID,
		},
	})
}

func (h *BookHandler) GetAllBookHandler(c *gin.Context) {
	books, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "Buku tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"books": books,
		},
	})
}

func (h *BookHandler) GetBookByIDHandler(c *gin.Context) {
	id := c.Param("id")
	book, err := h.service.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "Buku tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"book": book,
		},
	})
}

func (h *BookHandler) EditBookByIDHandler(c *gin.Context) {
	id := c.Param("id")

	var book domain.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "invalid requeset",
		})
		return
	}

	if book.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal memperbarui buku. Mohon isi nama buku",
		})
		return
	}

	if book.ReadPage > book.PageCount {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Gagal memperbarui buku. readPage tidak boleh lebih besar dari pageCount",
		})
		return
	}

	err = h.service.Update(id, book)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "fail",
			"message": "Gagal memperbarui buku. Id tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Buku berhasil diperbarui",
	})
}
