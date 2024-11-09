package usecase

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/dto"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/usecase/mocks"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/domain/entity"
	"testing"
)

func TestGivenAnOrderId_WhenICallGetOrderByIdUseCase_ThenIShouldReceiveOrder(t *testing.T) {
	orderId := "1"
	orderFound := &entity.Order{
		ID:         "1",
		Price:      100.0,
		Tax:        10.0,
		FinalPrice: 110.0,
	}
	mockOrderRepository := &mocks.OrderRepositoryMock{}
	usecase := NewGetOrderByIdUseCase(mockOrderRepository)

	mockOrderRepository.On("FindOrderById", orderId).Return(orderFound, nil)

	output, err := usecase.Execute(orderId)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	assert.Equal(t, orderFound.ID, output.ID)
}

func TestGivenAnOrderId_WhenOrderNotFound_ThenIShouldReceiveOrderNilAndError(t *testing.T) {
	orderId := "1"
	orderRepositoryOutput := &entity.Order{}
	orderOutputDto := dto.OrderOutputDTO{}
	mockOrderRepository := &mocks.OrderRepositoryMock{}
	usecase := NewGetOrderByIdUseCase(mockOrderRepository)

	mockOrderRepository.On("FindOrderById", orderId).Return(orderRepositoryOutput, errors.New("order not found"))

	output, err := usecase.Execute(orderId)

	assert.NotNil(t, err)
	assert.Equal(t, output, orderOutputDto)
	assert.Equal(t, "order not found", err.Error())
}
