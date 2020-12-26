package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type item struct {
	id       int    `json:"id"`
	name     string `json:"name"`
	category string `json:"category"`
}

func main() {
	// DB settings
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// API settings
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		rows, err := db.Query("SELECT * FROM daily_items")
		if err != nil {
			panic(err.Error())
		}
		var items []item
		for rows.Next() {
			item := item{}
			if err := rows.Scan(&item.id, &item.name, &item.category); err != nil {
				panic(err.Error())
			}
			items = append(items, item)
		}
		fmt.Println(items)
		return c.JSON(http.StatusOK, items)
	})

	// API port listen
	e.Logger.Fatal(e.Start(":8081"))
}
