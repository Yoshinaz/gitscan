db-up:
	docker-compose -f docker-compose.yaml up -d

migrate:
	go run cmd/cli/main.go db migrate