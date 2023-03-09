package repository

import (
	"errors"

	"github.com/zakirkun/grpc-crud/app/domain/model"
	"github.com/zakirkun/grpc-crud/database"

	pb "github.com/zakirkun/grpc-crud/proto"
)

type RepositoryContext struct {
}

func (r RepositoryContext) CreateMovie(data model.Movie, status chan error) {
	db, err := database.OpenDB()
	if err != nil {
		status <- err
	}

	res := db.Create(&data)

	if res.RowsAffected == 0 {
		err = errors.New("failed create movie")
		status <- err
	}

	status <- nil
}

func (r RepositoryContext) FindMovie(id string, status chan error, data chan *model.Movie) {
	db, err := database.OpenDB()
	if err != nil {
		status <- err
		data <- nil
	}

	var movie model.Movie

	res := db.Find(&movie, "id = ?", id)

	if res.RowsAffected == 0 {
		status <- errors.New("movie not found")
		data <- nil
	}

	status <- nil
	data <- &movie
}

func (r RepositoryContext) FindAllMovie(status chan error, data chan []*pb.Movie) {
	db, err := database.OpenDB()
	if err != nil {
		status <- err
		data <- nil
	}

	var movies []*pb.Movie

	res := db.Find(&movies)
	if res.RowsAffected == 0 {
		status <- errors.New("record empty")
		data <- nil
	}

	status <- nil
	data <- movies
}

func (r RepositoryContext) UpdateMovie(id string, req *pb.UpdateMovieRequest, status chan error) {
	db, err := database.OpenDB()
	if err != nil {
		status <- err
	}

	var movie model.Movie
	request := req.GetMovie()

	res := db.Model(&movie).Where("id=?", id).Updates(model.Movie{
		Title:       request.GetTitle(),
		Genre:       request.GetGenre(),
		Description: request.GetDescription(),
		Thumbnail:   request.GetThumbnail(),
	})

	if res.RowsAffected == 0 {
		status <- errors.New("movies not found")
	}

	status <- nil
}

func (r RepositoryContext) DeleteMovie(id string, status chan error) {
	db, err := database.OpenDB()
	if err != nil {
		status <- err
	}

	var movie model.Movie
	res := db.Where("id=?", id).Delete(&movie)
	if res.RowsAffected == 0 {
		status <- errors.New("movies not found")
	}

	status <- nil
}
