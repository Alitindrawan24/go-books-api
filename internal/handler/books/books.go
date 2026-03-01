package books

import (
	"net/http"
	"strconv"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) HandleGetBooks(c *gin.Context) {

	books, err := handler.books.GetBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Books retrieved successfully",
		"status":  true,
		"data":    books,
	})
}

func (handler *Handler) HandleGetBook(c *gin.Context) {
	bookID := c.Param("id")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	book, err := handler.books.GetBook(c.Request.Context(), bookIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book retrieved successfully",
		"status":  true,
		"data":    book,
	})
}

func (handler *Handler) HandleStoreBook(c *gin.Context) {
	var book entity.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := handler.books.CreateBook(c.Request.Context(), book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Book created successfully",
		"status":  true,
		"data":    book,
	})
}
