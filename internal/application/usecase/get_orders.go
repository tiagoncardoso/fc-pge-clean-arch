package usecase

import (
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/dto"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/domain/repository"
)

type GetOrdersUseCase struct {
	OrderRepository repository.OrderRepositoryInterface
}

func NewGetOrdersUseCase(orderRepository repository.OrderRepositoryInterface) *GetOrdersUseCase {
	return &GetOrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (g *GetOrdersUseCase) Execute() ([]dto.OrderOutputDTO, error) {
	orders, err := g.OrderRepository.FindOrders()
	if err != nil {
		return []dto.OrderOutputDTO{}, err
	}

	var ordersDTO []dto.OrderOutputDTO
	for _, order := range orders {
		orderDTO := dto.OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}

		ordersDTO = append(ordersDTO, orderDTO)
	}

	return ordersDTO, nil
}
