package books

import (
	"net/http"
	"strconv"

	"github.com/Alitindrawan24/go-books-api/internal/entity"
	"github.com/Alitindrawan24/go-books-api/internal/schema"
	"github.com/gin-gonic/gin"
)

func (handler *Handler) HandleGetBooks(c *gin.Context) {

	var querySchema schema.QueryParams
	querySchema.Author = c.Query("author")
	page := c.Query("page")
	limit := c.Query("limit")

	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
			return
		}

		querySchema.Page = pageInt
	}

	if limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit number"})
			return
		}

		querySchema.Limit = limitInt
	}

	books, err := handler.books.GetBooks(c.Request.Context(), querySchema)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
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
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
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

	c.JSON(http.StatusCreated, gin.H{"id": book.ID})
}

func (handler *Handler) HandleUpdateBook(c *gin.Context) {
	bookID := c.Param("id")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}
	var book entity.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err = handler.books.UpdateBook(c.Request.Context(), bookIDInt, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (handler *Handler) HandleDeleteBook(c *gin.Context) {
	bookID := c.Param("id")
	bookIDInt, err := strconv.Atoi(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid book ID"})
		return
	}

	err = handler.books.DeleteBook(c.Request.Context(), bookIDInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
