package usecase

import (
	"errors"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/usecase/mocks"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/domain/entity"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/dto"
)

func TestGivenAnOrderInput_WhenICallCreateOrderUseCase_ThenIShouldReceiveOrderOutput(t *testing.T) {
	mockOrderRepository := &mocks.OrderRepositoryMock{}
	mockOrderCreatedEvent := &mocks.OrderCreatedMock{}
	mockEventDispatcher := &mocks.EventDispatcherMock{}

	usecase := NewCreateOrderUseCase(mockOrderRepository, mockOrderCreatedEvent, mockEventDispatcher)

	input := dto.OrderInputDTO{
		ID:    "1",
		Price: 100.0,
		Tax:   10.0,
	}

	order := &entity.Order{
		ID:         input.ID,
		Price:      input.Price,
		Tax:        input.Tax,
		FinalPrice: 110.0,
	}

	mockOrderRepository.On("Save", order).Return(nil)
	mockOrderCreatedEvent.On("SetPayload", mock.Anything).Return()
	mockOrderCreatedEvent.On("GetName").Return("OrderCreated")
	mockOrderCreatedEvent.On("GetDateTime").Return(time.Now())
	mockOrderCreatedEvent.On("GetPayload").Return(order)
	mockEventDispatcher.On("Dispatch", mock.Anything).Return(nil)

	output, err := usecase.Execute(input)

	assert.Nil(t, err)
	assert.Equal(t, input.ID, output.ID)
	assert.Equal(t, input.Price, output.Price)
	assert.Equal(t, input.Tax, output.Tax)
	assert.Equal(t, 110.0, output.FinalPrice)

	mockOrderRepository.AssertExpectations(t)
	mockEventDispatcher.AssertExpectations(t)
	mockEventDispatcher.AssertExpectations(t)
}

func TestGivenAnOrderInput_WhenErrorOcoursInCreateOrderUseCase_ThenIShouldReceiveError(t *testing.T) {
	mockOrderRepository := &mocks.OrderRepositoryMock{}
	mockOrderCreatedEvent := &mocks.OrderCreatedMock{}
	mockEventDispatcher := &mocks.EventDispatcherMock{}

	usecase := NewCreateOrderUseCase(mockOrderRepository, mockOrderCreatedEvent, mockEventDispatcher)

	input := dto.OrderInputDTO{
		ID:    "1",
		Price: 100.0,
		Tax:   10.0,
	}

	order := &entity.Order{
		ID:         input.ID,
		Price:      input.Price,
		Tax:        input.Tax,
		FinalPrice: 110.0,
	}

	mockOrderRepository.On("Save", order).Return(errors.New("error saving order"))

	_, err := usecase.Execute(input)

	assert.NotNil(t, err)
	assert.Equal(t, "error saving order", err.Error())

	mockOrderRepository.AssertExpectations(t)
	mockEventDispatcher.AssertExpectations(t)
	mockEventDispatcher.AssertExpectations(t)
}
