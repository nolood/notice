package template

import (
	"context"

	templatev1 "github.com/nolood/notice-protos/gen/go/template"
	"google.golang.org/grpc"
)

type serverAPI struct {
	templatev1.UnimplementedTemplateServiceServer
}

func Register(gRPC *grpc.Server) {
	templatev1.RegisterTemplateServiceServer(gRPC, &serverAPI{})
}

func (s *serverAPI) AddTemplate(ctx context.Context, req *templatev1.AddTemplateRequest) (*templatev1.AddTemplateResponse, error) {
	panic("not implemented")
}

func (s *serverAPI) DeleteTemplate(ctx context.Context, req *templatev1.DeleteTemplateRequest) (*templatev1.DeleteTemplateResponse, error) {
	panic("not implemented")
}

func (s *serverAPI) GetTemplate(ctx context.Context, req *templatev1.GetTemplateRequest) (*templatev1.GetTemplateResponse, error) {
	panic("not implemented")
}

func (s *serverAPI) UpdateTemplate(ctx context.Context, req *templatev1.UpdateTemplateRequest) (*templatev1.UpdateTemplateResponse, error) {
	panic("not implemented")
}
