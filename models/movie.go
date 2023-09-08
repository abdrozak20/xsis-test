package models

import (
	"xsis-test/config"

	"gorm.io/gorm"
)

const MOVIE_TABLE_NAME = "movies"

type Movie struct {
	gorm.Model
	ID          int64
	Title       string
	Description string
	Rating      float32
	Image       string
}

type DefaulResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ListMovieResponse struct {
	Success bool            `json:"success"`
	Message string          `json:"message"`
	Data    []MovieResponse `json:"data"`
	Meta    config.Meta     `json:"meta"`
}

type MovieResponse struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
}

type CreateMovieRequest struct {
	Title       string  `json:"title" validate:"required"`
	Description string  `json:"description"`
	Rating      float32 `json:"rating"`
	Image       string  `json:"image"`
}

type DetailMovieResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    MovieResponse
}
