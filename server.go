package main

import (
	"log"
	"net"

	"github.com/TendiF/grpc-stockbit/proto"
	"github.com/TendiF/grpc-stockbit/service/movie"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("fail to listen on port 9000: %v", err)
	}

	s := movie.Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterMovieServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
