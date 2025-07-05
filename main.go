package main

import (
	"my-app/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func connect(c echo.Context) error {

	db, _ := model.DB.DB()
	defer db.Close()
	err := db.Ping()
	if err != nil {
		return c.String(http.StatusInternalServerError, "DB接続に失敗しました。")
	} else {
		return c.String(http.StatusOK, "DB接続に成功しました。")
	}
}

func main() {

	e := echo.New()
	e.GET("/", connect)
	e.Logger.Fatal(e.Start(":8080"))

	// e := echo.New()
	// db, _ := model.DB.DB()

	// defer db.Close()

	// e.GET("/users", controller.GetUsers)
	// e.GET("/users/:id", controller.GetUser)
	// e.POST("/users", controller.CreateUser)
	// e.PUT("/users/:id", controller.UpdateUser)
	// e.DELETE("/users/id,", controller.DeleteUser)
	// e.Logger.Fatal(e.Start(":8080"))
}
