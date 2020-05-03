package server

import (
	"context"
	"fmt"
	"net"

	"github.com/dmirou/otusgo/calendar/pkg/config"
	cevent "github.com/dmirou/otusgo/calendar/pkg/contracts/event"
	"github.com/dmirou/otusgo/calendar/pkg/contracts/request"
	"github.com/dmirou/otusgo/calendar/pkg/event"
	"github.com/dmirou/otusgo/calendar/pkg/helper"
	"github.com/golang/protobuf/ptypes/empty"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type CoreServer struct {
	cevent.EventServiceServer
	euc    event.UseCase
	server *grpc.Server
	cfg    *config.Config
	logger *zap.Logger
}

func NewCoreServer(euc event.UseCase, cfg *config.Config, logger *zap.Logger) *CoreServer {
	return &CoreServer{
		euc:    euc,
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

	authInterCeptor := func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (
		resp interface{},
		err error,
	) {
		headers, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.Unknown, "can not get request headers")
		}

		cs.logger.Debug("headers", zap.Any("headers", headers))

		userID, err := cs.getUserIDFromContext(ctx)
		if err != nil || userID == "" {
			return nil, status.Errorf(codes.Unauthenticated, "can not get user id from request")
		}

		cs.logger.Debug("auth info", zap.Any("userID", userID))

		return handler(ctx, req)
	}

	cs.server = grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(authInterCeptor),
		),
	)

	cevent.RegisterEventServiceServer(cs.server, cs)

	// Register reflection service on gRPC server.
	reflection.Register(cs.server)

	return cs.server.Serve(l)
}

// getUserIDFromContext returns user id from context metadata
func (cs *CoreServer) getUserIDFromContext(ctx context.Context) (string, error) {
	headers, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("can not get metadata from context")
	}

	userID, ok := headers["user-id"]
	if ok && len(userID) > 0 {
		return userID[0], nil
	}

	return "", nil
}

func (cs *CoreServer) CreateEvent(ctx context.Context, e *cevent.Event) (*cevent.Event, error) {
	userID, err := cs.getUserIDFromContext(ctx)
	if err != nil || userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "can not get user id from request")
	}

	e.UserId = userID

	ev, err := helper.EventFromProtobuf(e)
	if err != nil {
		return nil, err
	}

	err = cs.euc.CreateEvent(ctx, ev)
	if err != nil {
		return nil, err
	}

	e.Id = ev.ID

	return e, nil
}

func (cs *CoreServer) GetEventByID(ctx context.Context, req *request.ByID) (*cevent.Event, error) {
	userID, err := cs.getUserIDFromContext(ctx)
	if err != nil || userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "can not get user id from request")
	}

	ev, err := cs.euc.GetEventByID(ctx, userID, req.Id)
	if err != nil {
		return nil, err
	}

	return helper.EventToProtobuf(ev), nil
}

func (cs *CoreServer) UpdateEvent(ctx context.Context, e *cevent.Event) (*cevent.Event, error) {
	userID, err := cs.getUserIDFromContext(ctx)
	if err != nil || userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "can not get user id from request")
	}

	e.UserId = userID

	ev, err := helper.EventFromProtobuf(e)
	if err != nil {
		return nil, err
	}

	err = cs.euc.UpdateEvent(ctx, ev)
	if err != nil {
		return nil, err
	}

	return e, nil
}

func (cs *CoreServer) DeleteEvent(ctx context.Context, req *request.ByID) (*empty.Empty, error) {
	userID, err := cs.getUserIDFromContext(ctx)
	if err != nil || userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "can not get user id from request")
	}

	err = cs.euc.DeleteEvent(ctx, userID, req.Id)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (cs *CoreServer) ListEventsPerDate(ctx context.Context, req *request.ByDate) (
	*cevent.ListEventsResponse, error,
) {
	userID, err := cs.getUserIDFromContext(ctx)
	if err != nil || userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "can not get user id from request")
	}

	t, err := helper.ProtobufToTime(*req.Date)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "can not parse date")
	}

	evs, err := cs.euc.ListEventsPerDate(ctx, userID, t)
	if err != nil {
		return nil, err
	}

	resp := &cevent.ListEventsResponse{}
	resp.Events = make([]*cevent.Event, len(evs))

	for idx, e := range evs {
		resp.Events[idx] = helper.EventToProtobuf(e)
	}

	return &cevent.ListEventsResponse{}, nil
}

func (cs *CoreServer) ListEventsPerWeek(ctx context.Context, req *request.ByDate) (
	*cevent.ListEventsResponse, error,
) {
	userID, err := cs.getUserIDFromContext(ctx)
	if err != nil || userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "can not get user id from request")
	}

	t, err := helper.ProtobufToTime(*req.Date)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "can not parse date")
	}

	evs, err := cs.euc.ListEventsPerWeek(ctx, userID, t)
	if err != nil {
		return nil, err
	}

	resp := &cevent.ListEventsResponse{}
	resp.Events = make([]*cevent.Event, len(evs))

	for idx, e := range evs {
		resp.Events[idx] = helper.EventToProtobuf(e)
	}

	return &cevent.ListEventsResponse{}, nil
}

func (cs *CoreServer) ListEventsPerMonth(ctx context.Context, req *request.ByDate) (
	*cevent.ListEventsResponse, error,
) {
	userID, err := cs.getUserIDFromContext(ctx)
	if err != nil || userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "can not get user id from request")
	}

	t, err := helper.ProtobufToTime(*req.Date)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "can not parse date")
	}

	evs, err := cs.euc.ListEventsPerMonth(ctx, userID, t)
	if err != nil {
		return nil, err
	}

	resp := &cevent.ListEventsResponse{}
	resp.Events = make([]*cevent.Event, len(evs))

	for idx, e := range evs {
		resp.Events[idx] = helper.EventToProtobuf(e)
	}

	return &cevent.ListEventsResponse{}, nil
}

func (cs *CoreServer) Shutdown() {
	cs.server.GracefulStop()
	cs.logger.Info("core server was gracefully shutdown")
}
