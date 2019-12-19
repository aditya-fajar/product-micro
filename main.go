package main

import (
	"book/book"
	"book/config"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
)

func main() {
	r := gin.Default()
	db := config.DBInit()
	book := book.Book{DB: db}

	p := ginprometheus.NewPrometheus("gin")
	p.Use(r)

	config.RegisterConsul()
	config.RegisterZipkin()

	r.GET("/books", book.GetBooks)
	r.GET("/book/:id", book.GetBookById)
	r.POST("/books", book.CreateBook)

	r.GET("/healthcheck", config.Healthcheck)

	r.Run()
}
