package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"xsis-test/config"
	"xsis-test/models"
	"xsis-test/repository"

	"github.com/labstack/echo/v4"
)

type MovieServiceContract interface {
	GetMovies(c echo.Context) error
	UploadImage(c echo.Context) error
	CreateMovie(c echo.Context) error
	GetDetailMovie(c echo.Context) error
	UpdateMovie(c echo.Context) error
	DeleteMovie(c echo.Context) error
}

type MovieService struct {
	movieRepository repository.MovieRepositoryContract
}

func NewMovieService(repo repository.MovieRepositoryContract) MovieServiceContract {
	return &MovieService{
		movieRepository: repo,
	}
}

func (h *MovieService) GetMovies(c echo.Context) error {
	var response models.ListMovieResponse
	var params config.PaginationRequest
	var meta config.Meta

	err := c.Bind(&params)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	meta.SetLimit(params.Limit)
	meta.SetPage(params.Page)
	filters := config.GeneralFilters{}
	filters.Limit = meta.GetLimit()
	filters.Search = params.Search
	filters.Page = meta.GetPage()

	movies, total, err := h.movieRepository.GetMovies(filters)

	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Success = true
	response.Message = "Success get data"
	response.Data = movies
	meta.GetMeta(total, len(movies))
	response.Meta = meta

	return c.JSON(http.StatusOK, response)
}

func (h *MovieService) UploadImage(c echo.Context) error {
	var response models.DefaulResponse
	types := c.FormValue("type")
	if types != "movies" {
		response.Success = false
		response.Message = "wrong type"
		return c.JSON(http.StatusBadRequest, response)
	}

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filename := fmt.Sprintf("%d-%s", config.TimeMillisecond(), file.Filename)
	var path = fmt.Sprintf("assets/%s/%s", types, filename)
	dst, err := os.Create(path)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusOK, response)
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	response.Success = true
	response.Message = path
	return c.JSON(http.StatusOK, response)
}

func (h *MovieService) CreateMovie(c echo.Context) error {
	response := models.DefaulResponse{
		Success: true,
		Message: "Movie Successfully Created",
	}
	var params models.CreateMovieRequest
	err := c.Bind(&params)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(params); err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.movieRepository.CreateMovie(params)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *MovieService) GetDetailMovie(c echo.Context) error {
	response := models.DetailMovieResponse{
		Success: true,
		Message: "Success",
	}
	movie, err := h.movieRepository.GetDetailMovie(config.StringToInt64(c.Param("ID"), 64))
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	response.Data = movie
	return c.JSON(http.StatusOK, response)
}

func (h *MovieService) UpdateMovie(c echo.Context) error {
	response := models.DefaulResponse{
		Success: true,
		Message: "Success",
	}
	id := config.StringToInt64(c.Param("ID"), 64)

	var params models.CreateMovieRequest
	err := c.Bind(&params)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := c.Validate(params); err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	_, err = h.movieRepository.GetDetailMovie(id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.movieRepository.UpdateMovie(id, params)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *MovieService) DeleteMovie(c echo.Context) error {
	response := models.DefaulResponse{
		Success: true,
		Message: "Success",
	}
	id := config.StringToInt64(c.Param("ID"), 64)

	_, err := h.movieRepository.GetDetailMovie(id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	err = h.movieRepository.DeleteMovie(id)
	if err != nil {
		response.Success = false
		response.Message = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, response)
}
