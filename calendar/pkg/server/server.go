package server

import (
	"context"
	"fmt"
	"net"

	"github.com/dmirou/otusgo/calendar/pkg/config"
	"github.com/dmirou/otusgo/calendar/pkg/contracts/event"
	"github.com/dmirou/otusgo/calendar/pkg/contracts/request"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type CoreServer struct {
	event.EventServiceServer
	server *grpc.Server
	cfg    *config.Config
	logger *zap.Logger
}

func NewCoreServer(cfg *config.Config, logger *zap.Logger) *CoreServer {
	return &CoreServer{
		cfg:    cfg,
		logger: logger,
	}
}

func (cs *CoreServer) Run() error {
	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		return err
	}

	// TODO move to config
	fmt.Printf("core server listens on %v:%v\n", "127.0.0.1", "9000")

	cs.server = grpc.NewServer()

	event.RegisterEventServiceServer(cs.server, cs)

	// Register reflection service on gRPC server.
	reflection.Register(cs.server)

	return cs.server.Serve(l)
}

func (cs *CoreServer) CreateEvent(ctx context.Context, e *event.Event) (*event.Event, error) {
	cs.logger.Info("CreateEvent called", zap.String("event", e.String()))

	return &event.Event{}, nil
}

func (cs *CoreServer) GetEventByID(ctx context.Context, req *request.ByID) (*event.Event, error) {
	cs.logger.Info("GetEventByID called", zap.Any("request", req))

	return &event.Event{}, nil
}

func (cs *CoreServer) UpdateEvent(ctx context.Context, e *event.Event) (*event.Event, error) {
	cs.logger.Info("UpdateEvent called", zap.String("event", e.String()))

	return &event.Event{}, nil
}

func (cs *CoreServer) DeleteEvent(ctx context.Context, req *request.ByID) (*empty.Empty, error) {
	cs.logger.Info("DeleteEvent called", zap.Any("request", req))

	return &empty.Empty{}, nil
}

func (cs *CoreServer) ListEventsPerDate(ctx context.Context, req *request.ByDate) (
	*event.ListEventsResponse, error,
) {
	cs.logger.Info("ListEventsPerDate called", zap.Any("request", req))

	return &event.ListEventsResponse{}, nil
}

func (cs *CoreServer) ListEventsPerWeek(ctx context.Context, req *request.ByDate) (
	*event.ListEventsResponse, error,
) {
	cs.logger.Info("ListEventsPerWeek called", zap.Any("request", req))

	return &event.ListEventsResponse{}, nil
}

func (cs *CoreServer) ListEventsPerMonth(ctx context.Context, req *request.ByDate) (
	*event.ListEventsResponse, error,
) {
	cs.logger.Info("ListEventsPerMonth called", zap.Any("request", req))

	return &event.ListEventsResponse{}, nil
}

func (cs *CoreServer) Shutdown() {
	cs.server.GracefulStop()
	cs.logger.Info("core server was gracefully shutdown")
}
