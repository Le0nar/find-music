package main

import (
	"log"
	"net"

	musicv1 "github.com/Le0nar/find-music-protos/gen/go/music"
	"github.com/Le0nar/find-music/external-service/internal/handler"
	"github.com/Le0nar/find-music/external-service/internal/service"
	"google.golang.org/grpc"
)



func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
	 	log.Fatalf("failed to listen on port 50051: %v", err)
	}
   
	service := service.New()
	handler := handler.New(service)

	s := grpc.NewServer()
	musicv1.RegisterMusicServer(s, handler)
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
	 	log.Fatalf("failed to serve: %v", err)
	}
}