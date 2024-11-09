package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/domain/entity"
)

type OrderRepositoryMock struct {
	mock.Mock
}

func (r *OrderRepositoryMock) FindOrderById(id string) (*entity.Order, error) {
	args := r.Called(id)
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (r *OrderRepositoryMock) FindOrders() ([]*entity.Order, error) {
	args := r.Called()
	return args.Get(0).([]*entity.Order), args.Error(1)
}

func (r *OrderRepositoryMock) Save(order *entity.Order) error {
	args := r.Called(order)
	return args.Error(0)
}
