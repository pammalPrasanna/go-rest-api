package main

import (
	"net/http"
	"os"
	"strconv"
	"strings"

	"errors"

	docs "go-books-api/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type book struct {
	ID     string `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type book2 struct {
	ID     int    `json:"id" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	ISBN   string `json:"isbn" binding:"required"`
}

type Headersv1 struct {
	AuthToken       string `header:"AuthToken" binding:"required"`
	MandateHeader1  string `header:"MandateHeader1" binding:"required"`
	OptionalHeader1 string `header:"OptionalHeader1" binding:"-"`
}

type Headersv2 struct {
	AuthToken       string `header:"AuthToken" binding:"required"`
	MandateHeader1  string `header:"MandateHeader1" binding:"required"`
	MandateHeader2  string `header:"MandateHeader2" binding:"required"`
	MandateHeader3  string `header:"MandateHeader3" binding:"required"`
	OptionalHeader1 string `header:"OptionalHeader1" binding:"-"`
	OptionalHeader2 string `header:"OptionalHeader2" binding:"-"`
	OptionalHeader3 string `header:"OptionalHeader3" binding:"-"`
}

type QueryParamsv1 struct {
	MandateQuery1  string `header:"MandateQuery1" binding:"required"`
	OptionalQuery1 string `header:"OptionalQuery1" binding:"-"`
}

type QueryParamsv2 struct {
	MandateQuery1  string `header:"MandateQuery1" binding:"required"`
	MandateQuery2  string `header:"MandateQuery2" binding:"required"`
	MandateQuery3  string `header:"MandateQuery3" binding:"required"`
	OptionalQuery1 string `header:"OptionalQuery1" binding:"-"`
	OptionalQuery2 string `header:"OptionalQuery2" binding:"-"`
	OptionalQuery3 string `header:"OptionalQuery3" binding:"-"`
}

type IError struct {
	Field string
	Tag   string
	Value string
}

var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust"},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald"},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy"},
}

var books2 = []book2{
	{ID: 1, Title: "In Search of Lost Time", Author: "Marcel Proust", ISBN: "ABCD1"},
	{ID: 2, Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", ISBN: "ABCD2"},
	{ID: 3, Title: "War and Peace", Author: "Leo Tolstoy", ISBN: "ABCD3"},
}

// @Summary get all books
// @Success 200 {array} book
// @Failure 400 {string} error
// @Router /v1/books [get]
// @Tags books_v1
// @Produce json
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param OptionalHeader1 header string false "Optional header 1"
func getBooks(c *gin.Context) {
	if !headerValidation(c) {
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

// @Summary get all books
// @Success 200 {array} book2
// @Failure 400 {string} error
// @Router /v2/books [get]
// @Tags books_v2
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param MandateHeader2  header string true "Mandatory header 2"
// @Param MandateHeader3  header string true "Mandatory header 3"
// @Param OptionalHeader1 header string false "Optional header 1"
// @Param OptionalHeader2 header string false "Optinal header 2"
// @Param OptionalHeader3 header string false "Optinal header 3"
// @Produce json
func getBooksv2(c *gin.Context) {
	if !headerValidation(c) {
		return
	}
	c.IndentedJSON(http.StatusOK, books2)
}

// GetBookByID godoc
// @Summary queries API
// @Router /v1/queries [get]
// @Success 200 {object} QueryParamsv1
// @Failure 400 {string} error
// @Tags queries
// @Produce json
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param OptionalHeader1 header string false "Optional header 1"
// @Param MandateQuery1 query string true "Mandatory query 1"
// @Param OptionalQuery1 query string false "Optional query 1"
// @Accept json
func getQueriesv1(c *gin.Context) {
	if !headerValidation(c) {
		return
	}
	if !queryValidation(c) {
		return
	}
	c.IndentedJSON(http.StatusOK, c.Request.URL.Query())
}

// GetBookByID godoc
// @Summary queries API
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param MandateHeader2 header string true "Mandatory header 2"
// @Param MandateHeader3 header string true "Mandatory header 3"
// @Param MandateQuery1 query string true "Mandatory query 1"
// @Param MandateQuery2 query string true "Mandatory query 2"
// @Param MandateQuery3 query string true "Mandatory query 3"
// @Param OptionalQuery1 query string false "Optional query 1"
// @Param OptionalQuery2 query string false "Optional query 2"
// @Param OptionalQuery3 query string false "Optional query 3"
// @Router /v2/queries [get]
// @Success 200 {object} QueryParamsv2
// @Failure 400 {string} error
// @Tags queries
// @Produce json
// @Accept json
func getQueriesv2(c *gin.Context) {
	if !headerValidation(c) {
		return
	}
	if !queryValidation(c) {
		return
	}
	c.IndentedJSON(http.StatusOK, c.Request.URL.Query())
}

// GetBookByID godoc
// @Summary get a book by ID
// @Param id path string true "Book id"
// @Router /v1/books/{id} [get]
// @Success 200 {object} book
// @Failure 400 {string} error
// @Failure 404 {string} error
// @Tags books_v1
// @Produce json
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param OptionalHeader1 header string false "Optional header 1"
// @Accept json
func bookById(c *gin.Context) {
	if !headerValidation(c) {
		return
	}

	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// GetBookByID godoc
// @Summary get a book by ID
// @Param id path string true "Book id"
// @Router /v2/books/{id} [get]
// @Success 200 {object} book2
// @Failure 400 {string} error
// @Failure 404 {string} error
// @Tags books_v2
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param MandateHeader2  header string true "Mandatory header 2"
// @Param MandateHeader3  header string true "Mandatory header 3"
// @Param OptionalHeader1 header string false "Optional header 1"
// @Param OptionalHeader2 header string false "Optinal header 2"
// @Param OptionalHeader3 header string false "Optinal header 3"
// @Accept json
func bookByIdv2(c *gin.Context) {
	if !headerValidation(c) {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	book, err := getBookByIdv2(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with ID " + c.Param("id") + " not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func removeBookById(bookID string) (isRemoved bool) {
	bookToBeRemoved := -1
	for i := range books {
		if bookID == books[i].ID {
			bookToBeRemoved = i
			break
		}
	}
	if bookToBeRemoved != -1 {
		books = append(books[:bookToBeRemoved], books[bookToBeRemoved+1:]...)
		return true
	}
	return false
}

func removeBookByIdv2(bookID int) (isRemoved bool) {
	bookToBeRemoved := -1
	for i := range books2 {
		if bookID == books2[i].ID {
			bookToBeRemoved = i
			break
		}
	}
	if bookToBeRemoved != -1 {
		books2 = append(books2[:bookToBeRemoved], books2[bookToBeRemoved+1:]...)
		return true
	}
	return false
}

// DeleteBookByID godoc
// @Summary delete a book by ID
// @Param id path string true "Book id"
// @Router /v1/books/{id} [delete]
// @Success 200 {object} book
// @Failure 400 {string} error
// @Failure 404 {string} true "Book not found"
// @Failure 409 {string} true "Unable to delete book"
// @Tags books_v1
// @Produce json
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param OptionalHeader1 header string false "Optional header 1"
// @Accept json
func deleteBookById(c *gin.Context) {
	if !headerValidation(c) {
		return
	}

	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	status := removeBookById(book.ID)

	if !status {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Unable to delete book."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

// DeleteBookByID godoc
// @Summary delete a book by ID
// @Param id path string true "Book id"
// @Router /v2/books/{id} [delete]
// @Success 200 {object} book2
// @Failure 400 {string} error
// @Failure 404 {string} true "Book not found"
// @Failure 409 {string} true "Unable to delete book"
// @Tags books_v2
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param MandateHeader2  header string true "Mandatory header 2"
// @Param MandateHeader3  header string true "Mandatory header 3"
// @Param OptionalHeader1 header string false "Optional header 1"
// @Param OptionalHeader2 header string false "Optinal header 2"
// @Param OptionalHeader3 header string false "Optinal header 3"
// @Produce json
// @Accept json
func deleteBookByIdv2(c *gin.Context) {
	if !headerValidation(c) {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	book, err := getBookByIdv2(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book with ID " + c.Param("id") + " not found."})
		return
	}

	status := removeBookByIdv2(book.ID)

	if !status {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Unable to delete book."})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Book with ID " + strconv.Itoa(book.ID) + " was deleted successfully", "book": book})
}

func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func getBookByIdv2(id int) (*book2, error) {
	for i, b := range books2 {
		if b.ID == id {
			return &books2[i], nil
		}
	}

	return nil, errors.New("book not found")
}

// CreateBook godoc
// @Summary create a new book
// @Success 201 {object} book
// @Failure 400 {string} error
// @Param book body book true "Book data"
// @Router /v1/books [post]
// @Tags books_v1
// @Produce json
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param OptionalHeader1 header string false "Optional header 1"
// @Accept json
func createBook(c *gin.Context) {
	if !headerValidation(c) {
		return
	}

	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid datatype"})
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// CreateBook godoc
// @Summary create a new book
// @Success 201 {object} book2
// @Failure 400 {string} error
// @Param book body book2 true "Book data"
// @Router /v2/books [post]
// @Tags books_v2
// @Param AuthToken header string true "Authorization token"
// @Param MandateHeader1 header string true "Mandatory header 1"
// @Param MandateHeader2  header string true "Mandatory header 2"
// @Param MandateHeader3  header string true "Mandatory header 3"
// @Param OptionalHeader1 header string false "Optional header 1"
// @Param OptionalHeader2 header string false "Optinal header 2"
// @Param OptionalHeader3 header string false "Optinal header 3"
// @Accept json
func createBookv2(c *gin.Context) {
	if !headerValidation(c) {
		return
	}

	var newBook book2

	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	books2 = append(books2, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func headerValidation(c *gin.Context) (vstatus bool) {
	var headersv1 Headersv1
	var headersv2 Headersv2
	if strings.Contains(c.Request.URL.Path, "v1") {

		if err := c.ShouldBindHeader(&headersv1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return false
		}
	} else {
		if err := c.ShouldBindHeader(&headersv2); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return false
		}
	}

	return true
}

func queryValidation(c *gin.Context) (vstatus bool) {
	var queryParamsv1 QueryParamsv1
	var queryParamsv2 QueryParamsv2
	if strings.Contains(c.Request.URL.Path, "v1") {
		if err := c.ShouldBindQuery(&queryParamsv1); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return false
		}
	} else {
		if err := c.ShouldBindQuery(&queryParamsv2); err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return false
		}
	}
	return true
}

// @title Books API
// @description Books API documentation
func main() {
	args := os.Args
	port := "localhost:"
	if len(args) > 1 {
		port += args[1]
	} else {
		port += "4123"
	}
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/"

	router.GET("/api/v1/books", getBooks)
	router.GET("/api/v1/books/:id", bookById)
	router.POST("/api/v1/books", createBook)
	router.DELETE("/api/v1/books/:id", deleteBookById)

	router.GET("/api/v2/books", getBooksv2)
	router.GET("/api/v2/books/:id", bookByIdv2)
	router.POST("/api/v2/books", createBookv2)
	router.DELETE("/api/v2/books/:id", deleteBookByIdv2)

	router.GET("/api/v1/queries", getQueriesv1)
	router.GET("/api/v2/queries", getQueriesv2)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(port)
}
