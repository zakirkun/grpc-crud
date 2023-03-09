package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/zakirkun/grpc-crud/app/domain/model"
	"github.com/zakirkun/grpc-crud/app/repository"
	pb "github.com/zakirkun/grpc-crud/proto"
)

type ServerContext struct {
	pb.UnimplementedMovieServiceServer
}

func (*ServerContext) CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {

	movie := req.GetMovie()
	movie.Id = uuid.New().String()

	data := model.Movie{
		ID:          movie.GetId(),
		Title:       movie.GetTitle(),
		Genre:       movie.GetGenre(),
		Description: movie.GetDescription(),
		Thumbnail:   movie.GetThumbnail(),
	}

	repo := new(repository.RepositoryContext)

	// make chan status
	status := make(chan error)

	go repo.CreateMovie(data, status)

	// recive channel
	err := <-status

	if err != nil {
		return nil, err
	}

	// close channel
	close(status)

	return &pb.CreateMovieResponse{
		Movie: &pb.Movie{
			Id:          movie.GetId(),
			Title:       movie.GetTitle(),
			Genre:       movie.GetGenre(),
			Description: movie.GetDescription(),
			Thumbnail:   movie.GetThumbnail(),
		},
	}, nil
}

func (*ServerContext) GetMovie(ctx context.Context, req *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {

	repo := new(repository.RepositoryContext)

	// make chan status
	status := make(chan error)

	// make chan movie data
	data := make(chan *model.Movie)

	go repo.FindMovie(req.GetId(), status, data)

	// recive channel
	err := <-status
	getMovie := <-data

	if err != nil {
		return nil, err
	}

	// close all channel
	close(status)
	close(data)

	return &pb.ReadMovieResponse{
		Movie: &pb.Movie{
			Id:          getMovie.ID,
			Title:       getMovie.Title,
			Genre:       getMovie.Genre,
			Description: getMovie.Description,
			Thumbnail:   getMovie.Thumbnail,
		},
	}, nil
}

func (*ServerContext) GetMovies(ctx context.Context, req *pb.ReadMoviesRequest) (*pb.ReadMoviesResponse, error) {

	repo := new(repository.RepositoryContext)

	// make chan status
	status := make(chan error)

	// make chat data
	data := make(chan []*pb.Movie)

	go repo.FindAllMovie(status, data)

	// recive channel
	err := <-status
	movies := <-data

	if err != nil {
		return nil, err
	}

	// close all open channel
	close(status)
	close(data)

	return &pb.ReadMoviesResponse{
		Movies: movies,
	}, nil
}

func (*ServerContext) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {

	repo := new(repository.RepositoryContext)

	// make chan status
	status := make(chan error)

	go repo.UpdateMovie(req.Movie.GetId(), req, status)

	// recive channel
	err := <-status

	if err != nil {
		return nil, err
	}

	// close channel
	close(status)

	return &pb.UpdateMovieResponse{
		Movie: &pb.Movie{
			Id:          req.Movie.GetId(),
			Title:       req.Movie.GetTitle(),
			Genre:       req.Movie.GetGenre(),
			Description: req.Movie.GetDescription(),
			Thumbnail:   req.Movie.GetThumbnail(),
		},
	}, nil
}

func (*ServerContext) DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {

	repo := new(repository.RepositoryContext)

	// make chan status
	status := make(chan error)

	go repo.DeleteMovie(req.GetId(), status)

	// recive channel
	err := <-status

	if err != nil {
		return nil, err
	}

	return &pb.DeleteMovieResponse{
		Success: true,
	}, nil
}
