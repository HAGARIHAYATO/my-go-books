package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Book struct {
	gorm.Model
	Title       string `gorm:"column:title"`
	ISBN        string `gorm:"column:isbn"`
	Description string `gorm:"column:description"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()
	db.AutoMigrate(&Book{})

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	g := e.Group("/api/v1")
	g.GET("/books", allBooks)
	g.POST("/books", createBook)
	g.GET("/books/:id", getBook)
	g.PUT("/books/:id", updateBook)
	g.DELETE("/books/:id", deleteBook)
	e.Logger.Fatal(e.Start(":8000"))
}

func dbOpen() {
	db, _ := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()
}

func updateBook(c echo.Context) error {
	db, _ := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()

	b := &Book{}
	id := c.Param("id")
	before :=  db.Find(&b, id)
	db.Save(&before)
	return c.NoContent(http.StatusOK)
}

func deleteBook(c echo.Context) error {
	db, _ := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()

	id := c.Param("id")
	var Book Book
	db.Delete(&Book, id)
	return c.JSON(http.StatusNotFound, 404)
}

func getBook(c echo.Context) error {
	db, _ := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()

	id := c.Param("id")
	var Book Book
	book := db.Find(&Book, id)
	return c.JSON(http.StatusOK, book)
}

func allBooks(c echo.Context) error {
	db, _ := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()

	var Books []Book
	db.Find(&Books)
	return c.JSON(http.StatusOK, Books)
}

func createBook(c echo.Context) (err error) {
	db, _ := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()
	
	b := new(Book)
	if err = c.Bind(b); err != nil {
		return
	}
	db.Create(&b)
	return c.JSON(http.StatusOK, b)
}
