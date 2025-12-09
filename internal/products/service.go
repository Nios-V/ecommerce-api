package products

import "context"

type Service interface {
	GetAllProducts(ctx context.Context) error
}

type svc struct {
	//repository
}

func NewService() Service {
	return &svc{}
}

func (s *svc) GetAllProducts(ctx context.Context) error {
	return nil
}
