package server

import (
	"context"
	goerrors "errors"
	"fmt"
	"net"

	"github.com/dmirou/otusgo/calendar/pkg/config"
	cevent "github.com/dmirou/otusgo/calendar/pkg/contracts/event"
	"github.com/dmirou/otusgo/calendar/pkg/contracts/request"
	errors "github.com/dmirou/otusgo/calendar/pkg/error"
	"github.com/dmirou/otusgo/calendar/pkg/event"
	"github.com/dmirou/otusgo/calendar/pkg/helper"
	"github.com/golang/protobuf/ptypes/empty"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.uber.org/zap"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
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
	cfg    *config.Server
	logger *zap.Logger
}

func NewCoreServer(euc event.UseCase, cfg *config.Server, logger *zap.Logger) *CoreServer {
	return &CoreServer{
		euc:    euc,
		cfg:    cfg,
		logger: logger,
	}
}

func (cs *CoreServer) Run() error {
	addr := fmt.Sprintf("%s:%d", cs.cfg.IP, cs.cfg.Port)

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	fmt.Printf("core server listens on %s\n", addr)

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

	var (
		arg  *errors.InvalidArgError
		busy *errors.DateBusyError
	)

	switch {
	case err != nil && goerrors.As(err, &arg):
		return nil, cs.invalidArg(arg.Name, arg.Desc)
	case err != nil && goerrors.As(err, &busy):
		return nil, cs.invalidArg("start date", busy.Error())
	case err != nil:
		return nil, err
	default:
	}

	e.Id = ev.ID

	return e, nil
}

// invalidArg returns an error with InvalidArgument code
// and the field in the error details
func (cs *CoreServer) invalidArg(field, desc string) error {
	st := status.New(codes.InvalidArgument, "invalid "+field)
	v := &errdetails.BadRequest_FieldViolation{
		Field:       field,
		Description: desc,
	}
	br := &errdetails.BadRequest{}
	br.FieldViolations = append(br.FieldViolations, v)

	st, err := st.WithDetails(br)
	if err != nil {
		// If this errored, it will always error
		// here, so better panic so we can figure
		// out why than have this silently passing.
		panic(fmt.Sprintf("Unexpected error attaching metadata: %v", err))
	}

	return st.Err()
}

func (cs *CoreServer) GetEventByID(ctx context.Context, req *request.ByID) (*cevent.Event, error) {
	userID, err := cs.getUserIDFromContext(ctx)
	if err != nil || userID == "" {
		return nil, status.Errorf(codes.Unauthenticated, "can not get user id from request")
	}

	ev, err := cs.euc.GetEventByID(ctx, userID, req.Id)

	var notFound *errors.EventNotFoundError

	switch {
	case err != nil && goerrors.As(err, &notFound):
		return nil, status.Errorf(codes.NotFound, "event not found")
	case err != nil:
		return nil, err
	default:
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

	return resp, nil
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

	return resp, nil
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

	return resp, nil
}

func (cs *CoreServer) Shutdown() {
	cs.server.GracefulStop()
	cs.logger.Info("core server was gracefully shutdown")
}
