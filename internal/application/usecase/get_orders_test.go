package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/dto"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/usecase/mocks"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/domain/entity"
	"testing"
)

var ordersFound = []*entity.Order{
	{
		ID:         "1",
		Price:      100.0,
		Tax:        10.0,
		FinalPrice: 110.0,
	},
	{
		ID:         "2",
		Price:      200.0,
		Tax:        20.0,
		FinalPrice: 240.0,
	},
}

var ordersFoundDto = []dto.OrderOutputDTO{
	{
		ID:         "1",
		Price:      100.0,
		Tax:        10.0,
		FinalPrice: 110.0,
	},
	{
		ID:         "2",
		Price:      200.0,
		Tax:        20.0,
		FinalPrice: 240.0,
	},
}

func Test_WhenICallGetOrdersUseCase_ThenIShouldReceiveOrders(t *testing.T) {
	mockOrderRepository := &mocks.OrderRepositoryMock{}
	usecase := NewGetOrdersUseCase(mockOrderRepository)

	mockOrderRepository.On("FindOrders").Return(ordersFound, nil)

	output, err := usecase.Execute()

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, output, ordersFoundDto)
}

func Test_WhenNoOrdersFound_ThenIShouldReceiveOrderNilAndError(t *testing.T) {
	mockOrderRepository := &mocks.OrderRepositoryMock{}
	usecase := NewGetOrdersUseCase(mockOrderRepository)

	mockOrderRepository.On("FindOrders").Return(ordersFound, errors.New("no orders found"))

	output, err := usecase.Execute()

	assert.NotNil(t, err)
	assert.Equal(t, output, []dto.OrderOutputDTO{})
	assert.Equal(t, "no orders found", err.Error())
}
