package repository

import "github.com/tiagoncardoso/fc/pge/clean-arch/internal/domain/entity"

type OrderRepositoryInterface interface {
	Save(order *entity.Order) error
	FindOrderById(id string) (*entity.Order, error)
	FindOrders() ([]*entity.Order, error)
}
