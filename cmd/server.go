package main

import (
	"aegis_task/configs"
	"aegis_task/internal/database"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Server struct {
	configs *configs.Config
	db      *gorm.DB
}

func main() {
	// init server
	server := Server{}
	e := echo.New()

	// router
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "server is up and running!")
	})

	// start server
	server.configs = configs.InitConfig()
	database.InitializeDatabaseConnection()
	if database.DB == nil {
		panic("Failed to initialize database connection!")
	}
	fmt.Printf("Database connection established!")
	server.db = database.DB
	e.Logger.Fatal(e.Start(":1323"))
}
