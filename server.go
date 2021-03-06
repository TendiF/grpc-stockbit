package main

import (
	"context"
	"flag"
	"net"

	"github.com/TendiF/grpc-stockbit/gateway"
	"github.com/TendiF/grpc-stockbit/proto"
	"github.com/TendiF/grpc-stockbit/service/movie"
	"github.com/TendiF/grpc-stockbit/utils"
	"github.com/golang/glog"
	"google.golang.org/grpc"
)

var (
	endpoint = flag.String("endpoint", "localhost:9000", "endpoint of the gRPC service")
	network  = flag.String("network", "tcp", `one of "tcp" or "unix". Must be consistent to -endpoint`)
)

func serveHttp() {
	// gRPC gateway section
	flag.Parse()
	defer glog.Flush()

	ctx := context.Background()
	opts := gateway.Options{
		Addr: ":9090",
		GRPCServer: gateway.Endpoint{
			Network: *network,
			Addr:    *endpoint,
		},
	}
	if err := gateway.Run(ctx, opts); err != nil {
		glog.Fatal(err)
	}
}

func main() {
	go serveHttp()

	// gRPC server section
	lis, err := net.Listen("tcp", ":9000")

	if err != nil {
		glog.Errorf("fail to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer(
		utils.WithServerUnaryInterceptor(),
	)

	s := movie.Server{}

	proto.RegisterMovieServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		glog.Errorf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
