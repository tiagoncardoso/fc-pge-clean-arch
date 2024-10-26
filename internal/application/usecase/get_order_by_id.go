package usecase

import (
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/dto"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/entity"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/infra/database"
)

type GetOrderByIdUseCase struct {
	OrderRepository database.OrderRepository
}

func NewGetOrderByIdUseCase(orderRepository entity.OrderRepositoryInterface) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (g *GetOrderByIdUseCase) Execute(orderId string) (dto.OrderOutputDTO, error) {
	order, err := g.OrderRepository.FindOrderById(orderId)
	if err != nil {
		return dto.OrderOutputDTO{}, err
	}

	return dto.OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
