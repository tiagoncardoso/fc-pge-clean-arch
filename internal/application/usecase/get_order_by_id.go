package usecase

import (
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/dto"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/entity"
)

type GetOrderByIdUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetOrderByIdUseCase(orderRepository entity.OrderRepositoryInterface) *GetOrderByIdUseCase {
	return &GetOrderByIdUseCase{
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
