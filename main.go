package main

import (
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
)

// Book is defined by this file , and then is not used by anyfile
type Book struct {
	gorm.Model
	Title       string `gorm:"column:title"`
	ISBN        string `gorm:"column:isbn"`
	Description string `gorm:"column:description"`
}

func main() {
	var (
		connectionName = os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user           = os.Getenv("CLOUDSQL_USER")
		password       = os.Getenv("CLOUDSQL_PASSWORD")
		database       = os.Getenv("CLOUDSQL_DATABASE")
	)
	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s", user, password, connectionName, database))
	if err != nil {
		panic(err)
	}
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
	appengine.Main()
}

func updateBook(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	b := &Book{}
	id := c.Param("id")
	before := db.Find(&b, id)
	db.Save(&before)
	return c.NoContent(http.StatusOK)
}

func deleteBook(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := c.Param("id")
	var Book Book
	db.Delete(&Book, id)
	return c.JSON(http.StatusNotFound, 404)
}

func getBook(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := c.Param("id")
	var Book Book
	book := db.Find(&Book, id)
	return c.JSON(http.StatusOK, book)
}

func allBooks(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var Books []Book
	db.Find(&Books)
	return c.JSON(http.StatusOK, Books)
}

func createBook(c echo.Context) (err error) {
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	b := new(Book)
	if err = c.Bind(b); err != nil {
		return
	}
	db.Create(&b)
	return c.JSON(http.StatusOK, b)
}
