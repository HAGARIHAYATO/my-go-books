package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/cast"
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
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook?parseTime=true")
	if appengine.IsDevAppServer() {
		db, err = gorm.Open("mysql", "root:roothagari@tcp(127.0.0.1:3306)/mybook?charset=utf8&parseTime=True&loc=Local")
	}
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
	// e.Logger.Fatal(e.Start(":8000"))
	http.Handle("/", e)
	appengine.Main()
}

func updateBook(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook?parseTime=true")
	if appengine.IsDevAppServer() {
		db, err = gorm.Open("mysql", "root:roothagari@tcp(127.0.0.1:3306)/mybook?charset=utf8&parseTime=True&loc=Local")
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	b := new(Book)
	if err = c.Bind(b); err != nil {
		return echo.ErrBadRequest
	}

	id := c.Param("id")
	b.ID = cast.ToUint(id)

	db.Save(&b)
	return c.NoContent(http.StatusOK)
}

func deleteBook(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook?parseTime=true")
	if appengine.IsDevAppServer() {
		db, err = gorm.Open("mysql", "root:roothagari@tcp(127.0.0.1:3306)/mybook?charset=utf8&parseTime=True&loc=Local")
	}
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
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook?parseTime=true")
	if appengine.IsDevAppServer() {
		db, err = gorm.Open("mysql", "root:roothagari@tcp(127.0.0.1:3306)/mybook?charset=utf8&parseTime=True&loc=Local")
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	id := c.Param("id")
	var b Book
	db.First(&b, id)
	return c.JSON(http.StatusOK, b)
}

func allBooks(c echo.Context) error {
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook?parseTime=true")
	if appengine.IsDevAppServer() {
		db, err = gorm.Open("mysql", "root:roothagari@tcp(127.0.0.1:3306)/mybook?charset=utf8&parseTime=True&loc=Local")
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var Books []Book
	db.Find(&Books)
	return c.JSON(http.StatusOK, Books)
}

func createBook(c echo.Context) (err error) {
	db, err := gorm.Open("mysql", "root:roothagari@unix(/cloudsql/my-go-app-259011:asia-northeast1:myinstance)/mybook?parseTime=true")
	if appengine.IsDevAppServer() {
		db, err = gorm.Open("mysql", "root:roothagari@tcp(127.0.0.1:3306)/mybook?charset=utf8&parseTime=True&loc=Local")
	}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	b := new(Book)
	if err = c.Bind(b); err != nil {
		return echo.ErrBadRequest
	}
	db.Create(&b)
	return c.JSON(http.StatusOK, b)
}
