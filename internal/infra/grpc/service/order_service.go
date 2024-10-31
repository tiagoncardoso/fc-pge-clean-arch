package service

import (
	"context"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/dto"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/usecase"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/infra/grpc/pb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase  usecase.CreateOrderUseCase
	GetOrdersUseCase    usecase.GetOrdersUseCase
	GetOrderByIdUseCase usecase.GetOrderByIdUseCase
}

func NewOrderService(
	createOrderUseCase usecase.CreateOrderUseCase,
	getOrdersUseCase usecase.GetOrdersUseCase,
	getOrderByIdUseCase usecase.GetOrderByIdUseCase,
) *OrderService {
	return &OrderService{
		CreateOrderUseCase:  createOrderUseCase,
		GetOrdersUseCase:    getOrdersUseCase,
		GetOrderByIdUseCase: getOrderByIdUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	dto := dto.OrderInputDTO{
		ID:    in.Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         output.ID,
		Price:      float32(output.Price),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrderById(ctx context.Context, in *pb.ListOrderByIdRequest) (*pb.Order, error) {
	order, err := s.GetOrderByIdUseCase.Execute(in.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Order{
		Id:         order.ID,
		Price:      float32(order.Price),
		Tax:        float32(order.Tax),
		FinalPrice: float32(order.FinalPrice),
	}, nil
}

func (s *OrderService) ListOrders(ctx context.Context, in *pb.Blank) (*pb.OrdersList, error) {
	var ordersList pb.OrdersList
	orders, err := s.GetOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		ordersList.Orders = append(ordersList.Orders, &pb.Order{
			Id:         order.ID,
			Price:      float32(order.Price),
			Tax:        float32(order.Tax),
			FinalPrice: float32(order.FinalPrice),
		})
	}

	return &ordersList, nil
}
