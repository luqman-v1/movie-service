package grpc

import (
	"log"
	"net"
	"os"

	pb "movie-service/transport/grpc/proto/movie"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start() {
	lis, err := net.Listen("tcp", ":"+os.Getenv("grpc_port"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Println("Starting grpc server at port :" + os.Getenv("grpc_port"))
	}

	s := grpc.NewServer()

	pb.RegisterMovieServer(s, NewServer())
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
