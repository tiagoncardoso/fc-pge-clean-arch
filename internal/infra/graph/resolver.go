package graph

import (
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase  usecase.CreateOrderUseCase
	GetOrdersUseCase    usecase.GetOrdersUseCase
	GetOrderByIdUseCase usecase.GetOrderByIdUseCase
}
