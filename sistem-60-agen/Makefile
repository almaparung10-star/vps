.PHONY: run build compose-up compose-down migrate
run:
	go run ./cmd/api
build:
	go build -o bin/api ./cmd/api
compose-up:
	docker compose up -d --build
compose-down:
	docker compose down
