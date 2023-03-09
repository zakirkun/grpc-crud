package contracts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/zakirkun/grpc-crud/app/domain/model"
	pb "github.com/zakirkun/grpc-crud/proto"
)

type DeliverClientServices interface {
	CreateMovie(ctx *gin.Context)
	GetMovie(ctx *gin.Context)
	GetMovies(ctx *gin.Context)
	UpdateMovie(ctx *gin.Context)
	DeleteMovie(ctx *gin.Context)
}

type ServiceGrpcServer interface {
	CreateMovie(ctx context.Context, req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error)
	GetMovie(ctx context.Context, req *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error)
	GetMovies(ctx context.Context, req *pb.ReadMoviesRequest) (*pb.ReadMoviesResponse, error)
	UpdateMovie(ctx context.Context, req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error)
	DeleteMovie(ctx context.Context, req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error)
}

type RepositoryGrpcServer interface {
	CreateMovie(data model.Movie, status chan error)
	FindMovie(id string, status chan error, data chan *model.Movie)
	FindAllMovie(status chan error, data chan []*pb.Movie)
	UpdateMovie(id string, req *pb.UpdateMovieRequest, status chan error)
	DeleteMovie(id string, status chan error)
}
