package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zakirkun/grpc-crud/app/delivery"
	pb "github.com/zakirkun/grpc-crud/proto"
)

func RegisterRouter(client pb.MovieServiceClient) *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"192.168.72.1"})

	d := delivery.NewDelivery(client)

	v1 := "/v1"
	appv1 := r.Group(v1)
	{
		appv1.GET("/movies", d.GetMovies)
		appv1.GET("/movies/:id", d.GetMovie)
		appv1.POST("/movies", d.CreateMovie)
		appv1.PUT("/movies/:id", d.UpdateMovie)
		appv1.DELETE("/movies/:id", d.DeleteMovie)
	}

	return r
}
