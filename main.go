package main

import (
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
	"xsis-test/config"
	"xsis-test/repository"
	"xsis-test/services"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config.ConnectDB()
}

func MovieInjector() (service services.MovieServiceContract) {
	movieRepository := repository.NewMovieRepository(config.DB)
	service = services.NewMovieService(movieRepository)
	return
}

func main() {
	e := echo.New()
	// register validator
	val := validator.New()
	val.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	e.Validator = &config.Validator{Validator: val}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	movieService := MovieInjector()
	e.GET("/Movies", movieService.GetMovies)
	e.POST("/upload-image", movieService.UploadImage)
	e.POST("/Movies", movieService.CreateMovie)
	e.GET("/Movies/:ID", movieService.GetDetailMovie)
	e.PATCH("/Movies/:ID", movieService.UpdateMovie)
	e.DELETE("/Movies/:ID", movieService.DeleteMovie)

	e.Logger.Fatal(e.Start(":" + os.Getenv("APP_PORT")))
}
