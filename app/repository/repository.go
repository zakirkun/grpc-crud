package repository

import (
	"errors"

	"github.com/zakirkun/grpc-crud/app/domain/model"
	"github.com/zakirkun/grpc-crud/database"

	pb "github.com/zakirkun/grpc-crud/proto"
)

type RepositoryContext struct {
}

func (r RepositoryContext) CreateMovie(data model.Movie) error {
	db, err := database.OpenDB()
	if err != nil {
		return err
	}

	res := db.Create(&data)

	if res.RowsAffected == 0 {
		return errors.New("failed create movie")
	}

	return nil
}

func (r RepositoryContext) FindMovie(id string) (error, *model.Movie) {
	db, err := database.OpenDB()
	if err != nil {
		return err, nil
	}

	var movie model.Movie

	res := db.Find(&movie, "id = ?", id)

	if res.RowsAffected == 0 {
		return errors.New("movie not found"), nil
	}

	return nil, &movie
}

func (r RepositoryContext) FindAllMovie() (error, []*pb.Movie) {
	db, err := database.OpenDB()
	if err != nil {
		return err, nil
	}

	var movies []*pb.Movie

	res := db.Find(&movies)
	if res.RowsAffected == 0 {
		return errors.New("movies not found"), nil
	}

	return nil, movies
}

func (r RepositoryContext) UpdateMovie(id string, req *pb.UpdateMovieRequest) error {
	db, err := database.OpenDB()
	if err != nil {
		return err
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
		return errors.New("movies not found")
	}

	return nil
}

func (r RepositoryContext) DeleteMovie(id string) error {
	db, err := database.OpenDB()
	if err != nil {
		return err
	}

	var movie model.Movie
	res := db.Where("id=?", id).Delete(&movie)
	if res.RowsAffected == 0 {
		return errors.New("movies not found")
	}

	return nil
}
