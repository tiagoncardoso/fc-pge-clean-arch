FROM golang:latest as builder
WORKDIR /app
COPY . .
RUN cd cmd/ordersystem  \
    && GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o app ./main.go ./wire_gen.go

FROM scratch
COPY --from=builder /app/cmd/ordersystem/app .
COPY --from=builder /app/cmd/ordersystem/.env.build .env
CMD ["./app"]