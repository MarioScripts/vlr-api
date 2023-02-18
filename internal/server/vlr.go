package main

import (
	"context"

	pb "github.com/MarioScripts/vlr-api/proto/gen/vlr/v1"
)

func (s *Server) GetMatches(ctx context.Context, in *pb.MatchesRequest) (*pb.MatchesResponse, error) {
	return &pb.MatchesResponse{
		Matches: getMatches(in),
	}, nil
}

func (s *Server) GetMatch(ctx context.Context, in *pb.MatchRequest) (*pb.Match, error) {
	return getMatchFromId(in.Id), nil
}
