package book

import (
	"math/rand"
	"book/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"go.opencensus.io/trace"
)

type Book struct {
	DB *gorm.DB
}

type book struct {
	BookName string `json:"book_name"`
	Author string `json:"author"`
	Qty uint32 `json:"qty"`	
}

func (b *Book) GetBooks(c *gin.Context) {
	db := b.DB
	var books []model.Book

	db.Find(&books)

	c.JSON(200, gin.H{
		"data": books,
	})

	GetBookTracing(c)
}

func (b *Book) GetBookById(c *gin.Context) {
	var book model.Book
	db := b.DB
	id := c.Param("id")

	db.Where("id = ?", id).Find(&book)

	c.JSON(200, gin.H{
		"book_name": book.BookName,
		"author": book.Author,
		"qty": book.Qty,
	})
}

func (b *Book) CreateBook(c *gin.Context) {
	var request book

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	b.DB.Create(&request)
	c.JSON(200, gin.H{
		"message": "success",
		"data": request,
	})
}

func GetBookTracing(c *gin.Context) {
	_, span := trace.StartSpan(c, "/books")
	defer span.End()
	time.Sleep(time.Duration(rand.Intn(800)+200) * time.Millisecond)
}