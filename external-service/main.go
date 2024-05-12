package main

import (
	"context"
	"log"
	"net"

	musicv1 "github.com/Le0nar/find-music-protos/gen/go/music"
	"github.com/Le0nar/find-music/external-service/internal/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

type musicHandler interface {
    CreateMusic(ctx context.Context, singer, track string) (int64,  error)
    // TODO: mb rename to searchValue
    GetMusic(ctx context.Context, searchQuery string) (*models.Music,  error)
    UpdateMusic(ctx context.Context, singer, track string, id int64) (int64, error)
    DeleteMusic(ctx context.Context, id int64) (int64, error)
}

type server struct {
	musicv1.UnimplementedMusicServer
	music musicHandler
}

func (s *server) CreateMusic(ctx context.Context, in *musicv1.CreateMusicRequest) (*musicv1.CreateMusicResponse, error)  {
    if in.Singer == "" {
        return nil, status.Error(codes.InvalidArgument, "singer is required")
    }
    if in.Track == "" {
        return nil, status.Error(codes.InvalidArgument, "track is required")
    }

    // TODO: i dont know what is better: "in.Singer" or "in.GetSinger()"
    id, err := s.music.CreateMusic(ctx, in.GetSinger(), in.GetTrack())
    if err != nil {
        return nil, status.Error(codes.Internal, "failed to create the music")
    }

    return &musicv1.CreateMusicResponse{Id: id}, nil
}

func (s *server) GetMusic(ctx context.Context, in *musicv1.GetMusicRequest) (*musicv1.GetMusicResponse, error)  {
    if in.SearchValue == "" {
        return nil, status.Error(codes.InvalidArgument, "search value is required")
    }

    // TODO: rename
    musicInstance, err := s.music.GetMusic(ctx, in.GetSearchValue())
    if err != nil {
        // TODO: mb handle 2 cases: internal error and music not found
        return nil, status.Error(codes.Internal, "internal server error")
    }

    return &musicv1.GetMusicResponse{Singer: musicInstance.Singer, Track: musicInstance.Track, Id: musicInstance.Id}, nil
}

func (s *server) UpdateMusic(ctx context.Context, in *musicv1.UpdateMusicRequest) (*musicv1.UpdateMusicResponse, error)  {
    if in.Singer == "" {
        return nil, status.Error(codes.InvalidArgument, "singer is required")
    }
    if in.Track == "" {
        return nil, status.Error(codes.InvalidArgument, "track is required")
    }
    if in.Id == emptyValue {
        return nil, status.Error(codes.InvalidArgument, "id is required")
    }

    id, err := s.music.UpdateMusic(ctx, in.GetSinger(), in.GetTrack(), in.GetId())
    if err != nil {
        return nil, status.Error(codes.Internal, "failed to update the music")
    }

    return &musicv1.UpdateMusicResponse{Id: id}, nil
}

func (s *server) DeleteMusic(ctx context.Context, in *musicv1.DeleteMusicRequest) (*musicv1.DeleteMusicResponse, error)  {
    if in.Id == emptyValue {
        return nil, status.Error(codes.InvalidArgument, "id is required")
    }

    id, err := s.music.DeleteMusic(ctx, in.GetId())
    if err != nil {
        return nil, status.Error(codes.Internal, "failed to delete the music")
    }

    return &musicv1.DeleteMusicResponse{Id: id}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
	 	log.Fatalf("failed to listen on port 50051: %v", err)
	}
   
	s := grpc.NewServer()
	musicv1.RegisterMusicServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
	 	log.Fatalf("failed to serve: %v", err)
	}
}