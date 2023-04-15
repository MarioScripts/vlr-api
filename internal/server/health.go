package main

import (
	"context"

	health "google.golang.org/grpc/health/grpc_health_v1"
)

func (s *Server) Check(ctx context.Context, in *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	return &health.HealthCheckResponse{Status: health.HealthCheckResponse_SERVING}, nil
}
