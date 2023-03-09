package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zakirkun/grpc-crud/app/domain/contracts"
	"github.com/zakirkun/grpc-crud/app/domain/types"
	"github.com/zakirkun/grpc-crud/app/helper"
	pb "github.com/zakirkun/grpc-crud/proto"
)

type deliverClientContext struct {
	client pb.MovieServiceClient
}

func NewDelivery(client pb.MovieServiceClient) contracts.DeliverClientServices {
	return deliverClientContext{client: client}
}

func (d deliverClientContext) CreateMovie(ctx *gin.Context) {

	var movie types.Movie
	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.RestMessage(http.StatusBadRequest, "failed create movie", err.Error()))
	}

	data := &pb.Movie{
		Title:       movie.Title,
		Genre:       movie.Genre,
		Description: movie.Description,
		Thumbnail:   movie.Thumbnail,
	}

	res, err := d.client.CreateMovie(ctx, &pb.CreateMovieRequest{
		Movie: data,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.RestMessage(http.StatusBadRequest, "failed create movie", err.Error()))
	}

	ctx.JSON(http.StatusCreated, helper.RestMessage(http.StatusCreated, "Success", gin.H{
		"movie": res.Movie,
	}))
}

func (d deliverClientContext) GetMovie(ctx *gin.Context) {

	id := ctx.Param("id")
	res, err := d.client.GetMovie(ctx, &pb.ReadMovieRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.RestMessage(http.StatusNotFound, "failed get movie", err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.RestMessage(http.StatusOK, "Success", gin.H{
		"movie": res.Movie,
	}))
}

func (d deliverClientContext) GetMovies(ctx *gin.Context) {

	res, err := d.client.GetMovies(ctx, &pb.ReadMoviesRequest{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.RestMessage(http.StatusFound, "failed get movie", err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.RestMessage(http.StatusOK, "Success", gin.H{
		"movie": res.Movies,
	}))
}

func (d deliverClientContext) UpdateMovie(ctx *gin.Context) {

	var movie types.Movie

	if err := ctx.ShouldBindJSON(&movie); err != nil {
		ctx.JSON(http.StatusBadRequest, helper.RestMessage(http.StatusBadRequest, "failed update movie", err.Error()))
	}

	id := ctx.Param("id")

	res, err := d.client.UpdateMovie(ctx, &pb.UpdateMovieRequest{
		Movie: &pb.Movie{
			Id:          id,
			Title:       movie.ID,
			Genre:       movie.Genre,
			Description: movie.Description,
			Thumbnail:   movie.Thumbnail,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.RestMessage(http.StatusFound, "failed update movie", err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.RestMessage(http.StatusOK, "Success", gin.H{
		"movie": res.Movie,
	}))
}

func (d deliverClientContext) DeleteMovie(ctx *gin.Context) {

	id := ctx.Param("id")
	res, err := d.client.DeleteMovie(ctx, &pb.DeleteMovieRequest{Id: id})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, helper.RestMessage(http.StatusFound, "failed delete movie", err.Error()))
	}

	ctx.JSON(http.StatusOK, helper.RestMessage(http.StatusOK, "Success", gin.H{
		"deleted": res.Success,
	}))
}
