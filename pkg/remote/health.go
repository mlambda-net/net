package remote

import (
  "context"
  "google.golang.org/grpc/health/grpc_health_v1"
)

type HealthChecker struct{}

func (h HealthChecker) Check(_ context.Context, _ *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
  return &grpc_health_v1.HealthCheckResponse{
    Status: grpc_health_v1.HealthCheckResponse_SERVING,
  }, nil
}

func (h HealthChecker) Watch(_ *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
  return server.Send(&grpc_health_v1.HealthCheckResponse{
    Status: grpc_health_v1.HealthCheckResponse_SERVING,
  })
}

func NewHealthChecker() *HealthChecker {
  return &HealthChecker{}
}
