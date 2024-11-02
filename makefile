# Define a variable
PROJECT_DIR := ./cmd/ordersystem

create-docker-images:
	$(MAKE) log msg="===========> Begin docker images creation (MySQL / RabbitMQ)"
	docker-compose down --remove-orphans
	docker-compose up -d --build

wait-for-rabbitmq:
	$(MAKE) log msg="===========> Waiting for RabbitMQ to be ready"
	sleep 15  # Wait for 15 seconds

run-go-app:
	$(MAKE) log msg="===========> Running Go app"
	cd $(PROJECT_DIR) && go run ./main.go ./wire_gen.go

log:
	@echo "===========> $1"

run:
	$(MAKE) create-docker-images
	$(MAKE) wait-for-rabbitmq
	$(MAKE) run-go-app