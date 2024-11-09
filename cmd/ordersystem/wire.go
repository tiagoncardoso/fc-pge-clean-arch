//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/usecase"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/domain/repository"

	"github.com/google/wire"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/event"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/infra/database"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/infra/web"
	"github.com/tiagoncardoso/fc/pge/clean-arch/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(repository.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewGetOrdersUseCase(db *sql.DB) *usecase.GetOrdersUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewGetOrdersUseCase,
	)
	return &usecase.GetOrdersUseCase{}
}

func NewGetOrderByIdUseCase(db *sql.DB) *usecase.GetOrderByIdUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		usecase.NewGetOrderByIdUseCase,
	)
	return &usecase.GetOrderByIdUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}
