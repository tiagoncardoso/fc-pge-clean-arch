package main

import (
	"database/sql"
	"fmt"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/application/usecase"
	"net"
	"net/http"

	graphqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/streadway/amqp"
	"github.com/tiagoncardoso/fc/pge/clean-arch/configs"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/event/handler"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/infra/graph"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/infra/grpc/pb"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/infra/grpc/service"
	"github.com/tiagoncardoso/fc/pge/clean-arch/internal/infra/web/webserver"
	"github.com/tiagoncardoso/fc/pge/clean-arch/pkg/events"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db := initDB(cfg)
	rabbitMQChannel := initRabbitMQChannel(cfg)

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	getOrdersUseCase := NewGetOrdersUseCase(db)
	getOrderByIdUseCase := NewGetOrderByIdUseCase(db)
	initWebServer(cfg, db, eventDispatcher)
	initGrpcServer(cfg, createOrderUseCase, getOrdersUseCase, getOrderByIdUseCase)
	initGraphQLServer(cfg, createOrderUseCase, getOrdersUseCase, getOrderByIdUseCase)
}

func initDB(cfg configs.Conf) *sql.DB {
	db, err := sql.Open(cfg.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName))
	if err != nil {
		panic(err)
	}
	return db
}

func initRabbitMQChannel(cfg configs.Conf) *amqp.Channel {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/", cfg.RabbitMqUser, cfg.RabbitMqPassword, cfg.RabbitMqHost, cfg.RabbitMqPort))
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func initWebServer(cfg configs.Conf, db *sql.DB, eventDispatcher *events.EventDispatcher) {
	server := webserver.NewWebServer(cfg.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)

	server.AddHandler("/order", "POST", webOrderHandler.Create)
	server.AddHandler("/order", "GET", webOrderHandler.FindAll)
	server.AddHandler("/order/{id}", "GET", webOrderHandler.FindById)
	fmt.Println("Starting web server on port", cfg.WebServerPort)
	go server.Start()
}

func initGrpcServer(
	cfg configs.Conf,
	createOrderUseCase *usecase.CreateOrderUseCase,
	getOrdersUseCase *usecase.GetOrdersUseCase,
	getOrderByIdUseCase *usecase.GetOrderByIdUseCase,
) {
	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*createOrderUseCase, *getOrdersUseCase, *getOrderByIdUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", cfg.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", cfg.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)
}

func initGraphQLServer(
	cfg configs.Conf,
	createOrderUseCase *usecase.CreateOrderUseCase,
	getOrdersUseCase *usecase.GetOrdersUseCase,
	getOrderByIdUseCase *usecase.GetOrderByIdUseCase,
) {
	srv := graphqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase:  *createOrderUseCase,
		GetOrdersUseCase:    *getOrdersUseCase,
		GetOrderByIdUseCase: *getOrderByIdUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", cfg.GraphQLServerPort)
	http.ListenAndServe(":"+cfg.GraphQLServerPort, nil)
}
