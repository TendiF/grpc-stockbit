package movie

import (
	"context"
	"log"

	"github.com/TendiF/grpc-stockbit/proto"
)

type Server struct {
	proto.UnimplementedMovieServiceServer
}

func (s *Server) GetMovies(ctx context.Context, params *proto.GetMovieParams) (*proto.GetMovieParams, error) {

	log.Print("Params: ", params)

	return params, nil
}
