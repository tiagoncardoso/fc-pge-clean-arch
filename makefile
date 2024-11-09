# Define a variable
PROJECT_DIR := ./cmd/ordersystem
RETRY_COUNT := 10

create-docker-images:
	$(MAKE) log msg="===========> Begin docker images creation (MySQL / RabbitMQ)"
	docker-compose down --remove-orphans
	docker-compose up -d --build

run-go-app:
	$(MAKE) log msg="===========> Running Go app"
	cd $(PROJECT_DIR) && go run ./main.go ./wire_gen.go

retry-run-go-app:
	@for i in $$(seq 1 $(RETRY_COUNT)); do \
		$(MAKE) run-go-app && break || echo "Waiting for MySQL and RabbitMQ. Retrying $$i/$(RETRY_COUNT) ..."; \
		sleep 5; \
	done

log:
	@echo "===========> $1"

run:
	$(MAKE) create-docker-images
	$(MAKE) retry-run-go-app