package main

import (
	"aegis_task/configs"
	"aegis_task/internal/broker"
	"aegis_task/internal/database"
	usermodel "aegis_task/internal/user_service/models"
	"aegis_task/internal/user_service/repositories"
	"aegis_task/internal/user_service/service"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	// initial config
	server.configs = configs.InitConfig()
	database.InitializeDatabaseConnection()
	if database.DB == nil {
		panic("Failed to initialize database connection!")
	}
	fmt.Printf("Database connection established!")
	server.db = database.DB

	// init repo and service
	userRepo := repositories.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepo)

	// router
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "server is up and running!")
	})
	e.GET("/produceMessage", func(c echo.Context) error {
		return broker.ProducerMessage(c, "hello kafka again and again!")
	})
	e.POST("/user", func(ctx echo.Context) error {
		request := usermodel.User{
			Name:         ctx.FormValue("name"),
			Username:     ctx.FormValue("username"),
			Password:     ctx.FormValue("password"),
			Role:         ctx.FormValue("role"),
			Organization: ctx.FormValue("organization"),
		}

		return userService.CreateUser(ctx, request)
	})

	// start server
	go func() {
		if err := e.Start(":1323"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server!")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancle := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancle()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
