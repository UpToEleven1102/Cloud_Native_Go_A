package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
)

func main() {
	engine := gin.Default()

	//config routes
	routeConfig(engine)

	//config static files
	engine.LoadHTMLGlob("./templates/*.html")
	engine.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Home page",
		})
	})

	engine.Run(port())
}

func routeConfig(engine *gin.Engine) {
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	engine.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from gin framework"})
	})

	engine.GET("/api/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, AllBook())
	})

	engine.POST("/api/books", func(context *gin.Context) {
		var book Book
		if context.BindJSON(&book) == nil {
			isbn, created := CreateBook(book)
			if created {
				context.Header("Location", "/api/books/" + isbn)
				context.Status(http.StatusCreated)
			} else {
				context.Status(http.StatusConflict)
			}
		}
	})

	engine.GET("/api/books/:isbn", func(context *gin.Context) {
		isbn := context.Params.ByName("isbn")
		book, found := GetBook(isbn)

		if found {
			context.JSON(http.StatusOK, book)
		} else {
			context.AbortWithStatus(http.StatusNotFound)
		}
	})

	engine.PUT("/api/books/:isbn", func(context *gin.Context) {
		var book Book
		isbn := context.Params.ByName("isbn")
		if context.BindJSON(&book) == nil{
			book, updated := UpdateBook(isbn, book)
			if updated {
				context.JSON(http.StatusOK, book)
			} else {
				context.Status(http.StatusNotFound)
			}
		}

	})

	engine.DELETE("/api/books/:isbn", func(context *gin.Context) {
		isbn := context.Params.ByName("isbn")
		DeleteBook(isbn)

		context.Status(http.StatusOK)
	})
}

func port() string {
	port:= os.Getenv("PORT")

	if len(port) == 0 {
		port = "8080"
	}

	return ":" + port
}

