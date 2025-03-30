package notification

import (
	"context"

	notificationv1 "github.com/nolood/notice-protos/gen/go/notification"
	"google.golang.org/grpc"
)

type serverAPI struct {
	notificationv1.UnimplementedNotificationServiceServer
}

func Register(gRPC *grpc.Server) {
	notificationv1.RegisterNotificationServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) SendNotification(ctx context.Context, req *notificationv1.SendNotificationRequest) (*notificationv1.SendNotificationResponse, error) {
	panic("not implemented")
}

func (s *serverAPI) CheckNotificationStatus(ctx context.Context, req *notificationv1.CheckNotificationStatusRequest) (*notificationv1.CheckNotificationStatusResponse, error) {
	panic("not implemented")
}
