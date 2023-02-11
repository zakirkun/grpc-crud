package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/zakirkun/grpc-crud/app/services"
	"github.com/zakirkun/grpc-crud/database"
	pb "github.com/zakirkun/grpc-crud/proto"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "gRPC server port")
)

func main() {
	fmt.Println("gRPC server running ...")

	_, _ = database.OpenDB()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMovieServiceServer(s, &services.ServerContext{})

	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve : %v", err)
	}
}
