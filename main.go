package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jovi345/go-bookshelf-api/handler"
	"github.com/jovi345/go-bookshelf-api/repository"
	"github.com/jovi345/go-bookshelf-api/service"
)

func main() {
	bookRepo := repository.NewBookRepository()
	bookService := service.NewBookService(bookRepo)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	router.POST("/books", bookHandler.AddBookHandler)
	router.GET("/books", bookHandler.GetAllBookHandler)
	router.GET("/books/:id", bookHandler.GetBookByIDHandler)
	router.PUT("/books/:id", bookHandler.EditBookByIDHandler)

	router.Run("localhost:9000")
}
