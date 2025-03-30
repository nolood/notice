package notification

import (
	notificationv1 "github.com/nolood/notice/protos/gen/go/notification"
)

type serverAPI struct {
	notificationv1.UnimplementedNotificationServiceServer
}

func NewServerAPI() *serverAPI {
	return &serverAPI{}
}
