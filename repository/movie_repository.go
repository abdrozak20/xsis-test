package repository

import (
	"errors"
	"xsis-test/config"
	"xsis-test/models"

	"gorm.io/gorm"
)

type MovieRepositoryContract interface {
	GetMovies(filter config.GeneralFilters) ([]models.MovieResponse, int64, error)
	CreateMovie(param models.CreateMovieRequest) error
	GetDetailMovie(ID int64) (movie models.MovieResponse, err error)
	UpdateMovie(ID int64, param models.CreateMovieRequest) error
	DeleteMovie(ID int64) error
}

type MovieRepository struct {
	DB *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepositoryContract {
	return &MovieRepository{
		DB: db,
	}
}

func (repo MovieRepository) GetMovies(filter config.GeneralFilters) ([]models.MovieResponse, int64, error) {
	// var movie []models.Movie
	var response []models.MovieResponse
	var total int64

	offset := (filter.Page - 1) * filter.Limit
	query := repo.DB.
		Table(models.MOVIE_TABLE_NAME).Where("deleted_at IS NULL")
	queryTotal := repo.DB.
		Table(models.MOVIE_TABLE_NAME).Where("deleted_at IS NULL")

	if filter.Search != "" {
		query.Where("title LIKE '%" + filter.Search + "%'")
		queryTotal.Where("title LIKE '%" + filter.Search + "%'")

	}

	err := queryTotal.
		Select(`count(id) as TotalRow`).
		Take(&total).Error

	if err != nil {
		return nil, 0, err
	}

	if total == 0 {
		return nil, 0, errors.New("data not found")
	}

	err = query.
		Limit(filter.Limit).
		Offset(offset).
		Find(&response).Error

	return response, total, err
}

func (repo MovieRepository) CreateMovie(param models.CreateMovieRequest) error {
	movie := models.Movie{
		Title:       param.Title,
		Description: param.Description,
		Rating:      param.Rating,
		Image:       param.Image,
	}
	err := repo.DB.Create(&movie).Error
	return err
}

func (repo MovieRepository) GetDetailMovie(ID int64) (movie models.MovieResponse, err error) {
	err = repo.DB.Table(models.MOVIE_TABLE_NAME).
		Where("deleted_at IS NULL").First(&movie, "id = ?", ID).Error
	return movie, err
}

func (repo MovieRepository) UpdateMovie(id int64, param models.CreateMovieRequest) error {
	movie := models.Movie{
		ID:          id,
		Title:       param.Title,
		Description: param.Description,
		Rating:      param.Rating,
		Image:       param.Image,
	}
	err := repo.DB.Save(&movie).Error

	return err
}

func (repo MovieRepository) DeleteMovie(id int64) error {
	movie := models.Movie{}
	err := repo.DB.Delete(&movie, id).Error
	return err
}
