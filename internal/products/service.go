package products

import "context"

type Service interface {
	GetAllProducts(ctx context.Context) ([]Product, error)
}

type svc struct {
	repository Repository
}

func NewService(repo Repository) Service {
	return &svc{
		repository: repo,
	}
}

func (s *svc) GetAllProducts(ctx context.Context) ([]Product, error) {
	return s.repository.GetAll(ctx)
}
