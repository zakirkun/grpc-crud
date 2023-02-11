package main

import (
	"flag"
	"log"

	"github.com/zakirkun/grpc-crud/cmd/client/routes"
	pb "github.com/zakirkun/grpc-crud/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func init() {
	flag.Parse()
}

func main() {

	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewMovieServiceClient(conn)
	app := routes.RegisterRouter(client)

	app.Run(":5000")
}
