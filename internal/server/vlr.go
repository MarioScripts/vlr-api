package main

import (
	"context"

	pb "github.com/MarioScripts/vlr-api/proto/gen/vlr/v1"
)

func (s *Server) GetMatches(ctx context.Context, in *pb.MatchesRequest) (*pb.MatchesResponse, error) {
	matches, err := getMatches(in)
	return &pb.MatchesResponse{
		Matches: matches,
	}, err
}

func (s *Server) GetMatch(ctx context.Context, in *pb.IdRequest) (*pb.Match, error) {
	return getMatchFromId(in.Id)
}

func (s *Server) GetTeam(ctx context.Context, in *pb.IdRequest) (*pb.TeamResponse, error) {
	return getTeam(in)
}

func (s *Server) GetPlayer(ctx context.Context, in *pb.IdRequest) (*pb.Player, error) {
	return getPlayer(in.GetId(), true), nil
}
