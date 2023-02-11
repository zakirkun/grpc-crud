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

	err := repo.CreateMovie(data)
	if err != nil {
		return nil, err
	}

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

	err, getMovie := repo.FindMovie(req.GetId())
	if err != nil {
		return nil, err
	}

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

	err, movies := repo.FindAllMovie()
	if err != nil {
		return nil, err
	}

	return &pb.ReadMoviesResponse{
		Movies: movies,
	}, nil
}

func (*ServerContext) UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {

	repo := new(repository.RepositoryContext)
	err := repo.UpdateMovie(req.Movie.GetId(), req)
	if err != nil {
		return nil, err
	}

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

	err := repo.DeleteMovie(req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.DeleteMovieResponse{
		Success: true,
	}, nil
}
