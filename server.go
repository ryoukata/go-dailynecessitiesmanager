package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type dailyItem struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

func gormConnect() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "mysql"
	PROTOCOL := "tcp(172.18.0.2:3306)"
	DBNAME := "mysql"

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	// DB settings
	db := gormConnect()
	defer db.Close()

	// API settings
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/", func(c echo.Context) error {
		// create instance of item struct
		itemsEx := []dailyItem{}

		db.Find(&itemsEx)

		return c.JSON(http.StatusOK, itemsEx)
	})

	// API port listen
	e.Logger.Fatal(e.Start(":8081"))
}
