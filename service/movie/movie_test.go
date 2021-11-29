package movie_test

import (
	"context"
	"log"
	"net"
	"testing"

	"github.com/TendiF/grpc-stockbit/proto"
	"github.com/TendiF/grpc-stockbit/service/movie"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

type testServer struct {
	proto.UnimplementedMovieServiceServer
}

var tcpPort = ":9999"

func DialTestServer() {

	// gRPC server section
	lis, err := net.Listen("tcp", tcpPort)

	if err != nil {
		glog.Errorf("fail to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	s := movie.Server{}

	proto.RegisterMovieServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		glog.Errorf("Failed to serve gRPC server over port 9000: %v", err)
	}
}

func TestGetMovie(t *testing.T) {
	go DialTestServer()
	t.Log("get list movie")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(tcpPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewMovieServiceClient(conn)

	response, err := c.GetMovies(context.Background(), &proto.GetMovieParams{
		Page:   1,
		Search: "batman",
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("movie : %d", response.GetTotalResults())

	if response.GetTotalResults() <= 0 {
		t.Error("empty data")
	}
}

func TestDetailMovie(t *testing.T) {
	t.Log("get detail movie")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(tcpPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := proto.NewMovieServiceClient(conn)

	response, err := c.GetDetailMovie(context.Background(), &proto.GetDetailMovieParams{
		Id: "tt2975590",
	})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("movie : %s", response.GetTitle())

	if response.GetTitle() == "" {
		t.Error("empty data")
	}

}
