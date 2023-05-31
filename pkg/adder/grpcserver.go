package adder

import (
	"api/pkg/api"
	"context"
)

type GRPCServer struct {
	api.UnimplementedAdderServer // Вбудований нереалізований сервер
}

func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {
}

func (s *GRPCServer) Add(cnt context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}

func (s *GRPCServer) SetName(cnt context.Context, req *api.Person) (*api.PersonFullName, error) {
	return &api.PersonFullName{FullName: req.GetName() + req.GetSubName()}, nil
}
